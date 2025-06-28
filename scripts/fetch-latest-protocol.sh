#!/bin/bash

UPSTREAM_REPO="https://github.com/LIFX/public-protocol.git"
CLONE_DIR="$(mktemp -d)"
SRC_DIR=$(git rev-parse --show-toplevel)/src

echo "Cloning LIFX/public-protocol repo"

# Clone the upstream repo into a temp dir
git clone --depth 1 "$UPSTREAM_REPO" "$CLONE_DIR" --quiet

# Get the latest commit hash from the clone (main branch assumed)
cd "$CLONE_DIR"
LATEST_COMMIT=$(git rev-parse HEAD)

echo "Using upstream protocol commit: $LATEST_COMMIT"

# Copy protocol.yml into src/
cp protocol.yml "$SRC_DIR/protocol.yml"

# Store commit hash in a file alongside it
echo "$LATEST_COMMIT" > "$SRC_DIR/protocol_commit.txt"

# Clean up
rm -rf "$CLONE_DIR"
