#!/bin/sh

echo $FLAG1 > /flag1
echo $FLAG2 > /flag2
rm -f ./start.sh
node main.js