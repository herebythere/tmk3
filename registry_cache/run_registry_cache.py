# run_registry_cache
# brian taylor vann

import subprocess

def build_and_run_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    build_and_run_podman()
