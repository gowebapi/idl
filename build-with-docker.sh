#!/bin/bash

set -ex

# docker build -t gowebapi-idl-compile .
id=$(docker run -d gowebapi-idl-compile echo hello world)

echo "ID ${id}"
docker wait "${id}"

docker cp "${id}:/go/output" "output"

