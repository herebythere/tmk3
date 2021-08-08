import json
import os
import subprocess
from string import Template


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
    available_services_conf = config["available_services"]
    config_conf = config["config"]
    credentials_conf = config["credentials"]
    skeleton_keys_conf = config["skeleton_keys"]

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml", {
                        "available_services_filepath_test": available_services_conf["filepath_test"],
                        "available_services_filepath": available_services_conf["filepath"],
                        "config_filepath_test": config_conf["filepath_test"],
                        "config_filepath": config_conf["filepath"],
                        "credentials_filepath_test": credentials_conf["filepath_test"],
                        "credentials_filepath": credentials_conf["filepath"],
                        "https_port": server_conf["https_port"],
                        "local_cache_address": server_conf["local_cache_address"],
                        "session_cookie_label": server_conf["session_cookie_label"],
                        "skeleton_keys_filepath_test": skeleton_keys_conf["filepath_test"],
                        "skeleton_keys_filepath": skeleton_keys_conf["filepath"],
                        "service_name": config["service_name"]})


def build_registry_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])


if __name__ == "__main__":
    config = get_config("config/config.json")
    create_required_templates(config)
    # build_registry_with_podman()
