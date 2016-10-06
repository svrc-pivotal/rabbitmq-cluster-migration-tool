#!/bin/bash -eu

if ! which lpass > /dev/null 2>&1
then
  echo "must have lastpass CLI installed" >&2
  exit 1
fi

ssh_key=$(lpass show "Shared-London Services"/london-ci/git-ssh-key --notes)

fly -t london set-pipeline \
    --pipeline rabbitmq-migration-tool \
    --config pipeline.yml \
    --var git-ssh-key="${ssh_key}"
