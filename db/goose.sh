#! /bin/bash

set -a
. ../.env
set +a

goose -dir ./migrations "$@"