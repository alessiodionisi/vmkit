#!/usr/bin/env bash

for i in $(find * -iname "*.go" -o -iname "*.m" -o -iname "*.h")
do
  if ! grep -q Copyright $i
  then
    cat ./scripts/license.txt $i >$i.new && mv $i.new $i
  fi
done
