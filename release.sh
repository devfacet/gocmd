#!/bin/bash

set -e

# Init vars
GIT_TAG=$1

if [ -z "$GIT_TAG" ]; then
  echo "invalid tag"
  exit 1
fi;

# Push the tag
git commit -m "$GIT_TAG"
git tag -a $GIT_TAG -m "$GIT_TAG"
git push --follow-tags
