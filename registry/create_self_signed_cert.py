# brian taylor vann
# create_self_signed_cert.py

import subprocess

def build_and_run_podman():
    subprocess.run(["openssl", "req", "-x509", "-nodes", "-newkey", "rsa:4096", "-keyout", "./config/privkey.pem", "-out",
                   "./config/fullchain.pem", "-days", "365", "-subj", "/C=US/ST=California/L=San Francisco/O=tmk3/OU=Org/CN=registry.tmk3.com"])

if __name__ == "__main__":
    build_and_run_podman()
