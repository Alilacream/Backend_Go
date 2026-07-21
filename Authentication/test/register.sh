#!/bin/bash
if [ $# -eq 0 ]; then
  echo "no first comment and second argument"
  exit 1
elif [$# -eq 1]; then
  echo "Need second argument"
  exit 1
fi

curl -X POST http://localhost:8080/register \
  -d "username:$1&password:$2"
