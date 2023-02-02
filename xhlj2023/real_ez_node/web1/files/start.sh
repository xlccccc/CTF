#!/bin/bash
echo $DASFLAG > /flag.txt
while true
do 
node /app/bin/www&
sleep 120
echo `ps | grep "node" | grep -v "grep" | awk '{print $1}'`
kill -9 `ps | grep "node" | grep -v "grep" | awk '{print $1}'`
done

