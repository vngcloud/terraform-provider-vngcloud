#!/bin/bash

#cd terraform/examples
cd ./examples || exit

# Set the log level to debug
export TF_LOG=debug

#verbose logging
export TF_LOG_PROVIDERS=debug

# Run the terraform command with the specified arguments
terraform "$@"

#running ./terraform-debug.sh apply -auto-approve


