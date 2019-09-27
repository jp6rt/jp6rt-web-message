#!/bin/bash

set -e
aws cloudformation deploy --stack-name ${STACK_NAME} --template-file ./infra/resources-template.yml || EXIT_CODE=$?

if [ "$EXIT_CODE" -eq 255 ]
then 
  echo "No updates needed to deploy"
  exit 0
else
  echo "Failure $?"
fi
