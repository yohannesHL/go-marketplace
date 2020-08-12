#!/bin/bash

if [[ "$1" =~ ^v[0-9]{1,2}.[0-9]{1,3}.[0-9]{1,4}$ ]]; then
    echo "Bumped version to $1"
    echo "$1" > ./mod.version
else
    echo "Version string not in expected semver format (ex: v1.1.0). "
    exit 1
fi

# git commit -m "Bumped version to $1"
# git tag $1
# git push --tags 