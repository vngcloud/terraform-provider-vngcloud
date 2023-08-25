#!/bin/bash

#cd terraform/examples
cd ./examples || exit

# Run the terraform command with the specified arguments
terraform "$@"

