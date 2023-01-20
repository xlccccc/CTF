#!/bin/bash

echo 'flag'  > /flag
chmod 700 /flag
su - dotnet -c "cd /app && dotnet WebApplication1.dll urls='http://*:8888'"
tail -f /dev/null
