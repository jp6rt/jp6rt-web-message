#!/bin/bash

set -e
aws cloudformation deploy --stack-name ${STACK_NAME} --template-file ./infra/resources-template.yml

if [ $? -eq 255 ]
then 
  echo "No updates needed to deploy"
  exit 0
else
  echo "Failure $?"
fi