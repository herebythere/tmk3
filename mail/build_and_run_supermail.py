# build_and_run_gateway
# brian taylor vann

import json
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
    config_conf = config["config"]
    cert_conf = config["certs"]

    create_template("templates/main.cf.template",
                    "webapi/main.cf", {"domain_tld": config["tld"],
                                       "cert_filepath": cert_conf["cert_filepath"],
                                       "private_key_filepath": cert_conf["private_key_filepath"]})

    create_template("templates/webapi.dockerfile.template",
                    "webapi/dockerfile", server_conf)

    create_template("templates/docker-compose.yml.template",
                    "docker-compose.yml", {"service_name": config["service_name"],
                                           "http_port": server_conf["http_port"],
                                           "mx_port": server_conf["mx_port"],
                                           "filepath": config_conf["filepath"],
                                           "filepath_test": config_conf["filepath_test"]})


def build_and_run_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    config = get_config("config/config.json")
    create_required_templates(config)
    build_and_run_podman()
