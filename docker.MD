# burst server

docker run --name burst-server -p 10086:10086 -p 39399:39399 -d fzdwx/burst-server:1.4

netstat -lntp |grep 39399