#! /bin/sh

kill -9 $(pgrep webserver)

cd ~/webServer/

git pull https://github.com/caoyu5779/Cloud.git

cd server

./webServer &