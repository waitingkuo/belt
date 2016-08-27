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
  sudo curl -o /usr/local/bin/belt https://storage.googleapis.com/waitingkuo-belt/belt-linux
  sudo chmod +x /usr/local/bin/belt
elif [[ $platform == 'Darwin' ]]; then
  curl -o /usr/local/bin/belt https://storage.googleapis.com/waitingkuo-belt/belt-darwin
  chmod +x /usr/local/bin/belt
fi
