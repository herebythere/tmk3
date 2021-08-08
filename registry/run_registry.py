import subprocess

def run_registry_with_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    run_registry_with_podman()
