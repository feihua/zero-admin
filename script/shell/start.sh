#!/bin/sh

app=$1
if [ -z "$app" ]; then
  echo "Please input app name"
  exit 1
fi

chmod +x ./"$app"
pkill -f "$app"
nohup ./"$app" -f ./"$app".yaml  > /dev/null 2>&1 &