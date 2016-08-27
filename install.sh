#!/bin/env bash

platform='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
  platform='Linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
  platform='Darwin'
fi

echo "Downloading belt ..."
if [[ $platform == 'Linux' ]]; then
  curl -o /usr/local/bin/belt https://storage.googleapis.com/waitingkuo-belt/belt-linux
elif [[ $platform == 'Darwin' ]]; then
  curl -o /usr/local/bin/belt https://storage.googleapis.com/waitingkuo-belt/belt-darwin
fi
