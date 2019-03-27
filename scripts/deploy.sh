#!/bin/bash

# Setting -e and -v as pert https://docs.travis-ci.com/user/customizing-the-build/#Implementing-Complex-Build-Steps
# -e: immediately exit if any command has a non-zero exit status
# -v: print all lines in the script before executing them
# -o: prevents errors in a pipeline from being masked
set -euo pipefail

# IFS new value is less likely to cause confusing bugs when looping arrays or arguments (e.g. $@)
IFS=$'\n\t'

# Initialize digital ocean token
doctl auth init --access-token $DO_TOKEN

# Get the kubectl config
doctl kubernetes cluster kubeconfig show $DO_CLUSTER_ID > ~/.kube/do.yaml

# Set the configs
export KUBECONFIG=~/.kube/config:~/.kube/do.yaml

echo "Initialising helm..."
helm init --client-only

echo "Installing release"
helm upgrade --install $RELEASE_NAME $CHART \
    --set "image.repository=$IMAGE" \
    --set "image.tag=$TAG" \
    --set "image.command=$COMMAND" \
    --set "image.args={-t,$DISCORD_TOKEN}"