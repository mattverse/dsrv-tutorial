#! /bin/bash -x
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# Build the docker image first.
sudo docker build --tag arbiter/protogen -f "$DIR"/Dockerfile "$DIR"/../..

# Make sure that previous container not exist.
sudo docker rm --force arbiter_protogen 1>/dev/null 2>&1

sudo docker run -v "$DIR"/../..:/app --name arbiter_protogen arbiter/protogen sh dev/proto/protogen.sh

# Clear the container
sudo docker rm --force arbiter_protogen 1>/dev/null 2>&1
