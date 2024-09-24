#!/bin/sh

# Mounts the Docker socket & binary to allow for spawning containers from the container
docker run --rm -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker -v ./traces:/app/traces rust-test:latest