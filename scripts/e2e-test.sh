#!/bin/bash

echo "Running E2E test...."
docker run -it --rm --name gx-test -v "$PWD/scripts":"/tmp" gx:latest gx -u https://github.com/darox/gx -s README.md -t /tmp/ -b main
# check if file exists
if [ -f "$PWD/scripts/README.md" ]; then
    echo "E2E test passed✅"
else
    echo "E2E test failed❌"
    exit 1
fi
# remove file
rm -f "$PWD/scripts/README.md"