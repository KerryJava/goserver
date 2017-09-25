#!/bin/bash

COMMIT=$(git log | head -n 1 | cut -d ' ' -f 2)
echo $COMMIT

target="server-*-${COMMIT:0:7}"
echo $target
#rm server*

./build.sh  && \
./$target --log_dir=./log
