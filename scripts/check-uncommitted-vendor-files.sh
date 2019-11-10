#!/bin/bash

set -euo pipefail

GOPROXY="https://proxy.golang.org,direct" GOSUMDB=off GO111MODULE=on go mod tidy

untracked_files=$(git ls-files --others --exclude-standard | wc -l)

diff_stat=$(git diff --shortstat)

if [[ "${untracked_files}" -ne 0 || -n "${diff_stat}" ]]; then
  echo 'Untracked or diff in tracked vendor files found. Please run "go mod tidy" and commit the changes'
  exit 1
fi
