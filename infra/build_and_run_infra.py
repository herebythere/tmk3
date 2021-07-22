# build_and_run_gateway
# brian taylor vann

import json
import os
import subprocess
from string import Template


def create_required_directories():
    if not os.path.exists("cache_store/conf"):
        os.makedirs("cache_store/conf")
    if not os.path.exists("cache_store/data"):
        os.makedirs("cache_store/data")

def get_config(source):
    config_file = open(source, 'r')
    config = json.load(config_file)
    config_file.close()

    return config


def create_template(source, target, keywords):
    source_file = open(source, 'r')
    source_file_template = Template(source_file.read())
    source_file.close()
    updated_source_file_template = source_file_template.substitute(**keywords)

    target_file = open(target, "w+")
    target_file.write(updated_source_file_template)
    target_file.close()


def create_required_templates(config):
    server_conf = config["server"]
    config_conf = config["config"]
    cache_conf = config["cache"]
    skeleton_keys_conf = config["skeleton_keys"]

    create_template("templates/cache.dockerfile.template",
                    "cache_store/dockerfile", cache_conf)

    create_template("templates/redis.conf.template",
                    "cache_store/conf/redis.conf", cache_conf)

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml", {"service_name": config["service_name"],
                                           "https_port": server_conf["https_port"],
                                           "config_filepath": config_conf["filepath"],
                                           "config_filepath_test": config_conf["filepath_test"],
                                           "skeleton_keys_filepath": skeleton_keys_conf["filepath"],
                                           "skeleton_keys_filepath_test": skeleton_keys_conf["filepath_test"]})


def build_and_run_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    create_required_directories()
    config = get_config("config/config.json")
    create_required_templates(config)
    build_and_run_podman()
