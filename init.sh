#!/bin/bash

echo 'add local development domain'
if [[ ! $(cat /etc/hosts | grep cryptoapp.com) ]]; then
  echo "127.0.0.1 cryptoapp.com" | sudo sh -c 'cat - >> /etc/hosts'
fi