#!/bin/sh

TAG=""

# Parse command-line arguments
while getopts ":t:" opt; do
  case ${opt} in
    t ) TAG="$OPTARG";;
    \? ) echo "Usage: $0 -t <tag-version>"; exit 0;;
    : ) echo "Invalid option: -$OPTARG requires an argument"; exit 1;;
  esac
done

VERSION="${TAG:-latest}"
REPO_URL="github.com/KainNhantumbo/go-notes-cli"

echo "Installing the todo cli app from ${REPO_URL}"

go install ${REPO_URL}@${VERSION}

echo "App installed successfuly!"

go-notes-cli -h

echo "If you want to improve your ux, set an alias for the command in your shell."