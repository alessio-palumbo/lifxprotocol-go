#!/bin/bash

UPSTREAM_REPO="https://github.com/LIFX/public-protocol.git"
CLONE_DIR="$(mktemp -d)"
SRC_DIR=$(git rev-parse --show-toplevel)/cmd/protocol-gen/src
COMMIT_FILE="$SRC_DIR/protocol_commit.txt"

echo "Cloning LIFX/public-protocol repo"

# Clone the upstream repo into a temp dir
git clone --depth 1 "$UPSTREAM_REPO" "$CLONE_DIR" --quiet

# Get the latest commit hash from the clone (main branch assumed)
cd "$CLONE_DIR"
LATEST_COMMIT=$(git rev-parse HEAD)

# Compare to currently recorded commit
if [ -f "$COMMIT_FILE" ]; then
    CURRENT_COMMIT=$(cat "$COMMIT_FILE")
    if [ "$CURRENT_COMMIT" = "$LATEST_COMMIT" ]; then
        echo "Protocol commit is up to date ($CURRENT_COMMIT). No regeneration needed."
        rm -rf "$CLONE_DIR"
        exit 1  # ⚠️ exit non-zero to signal "no update needed"
    fi
fi

echo "Protocol updated: $LATEST_COMMIT"
# Copy new protocol and commit hash
cp protocol.yml "$SRC_DIR/protocol.yml"
echo "$LATEST_COMMIT" > "$SRC_DIR/protocol_commit.txt"

# Clean up
rm -rf "$CLONE_DIR"
