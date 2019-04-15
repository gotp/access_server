#! /bin/bash

pkill -f "access_server"
./access_server -alsologtostderr -v=2 > ./stdout 2>&1 &
