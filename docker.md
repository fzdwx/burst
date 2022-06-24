# burst server

docker run --name burst-server -p 10086:10086 -p 20000-58000:20000-58000 -d fzdwx/burst-server:1.5

netstat -lntp |grep 39399

docker tag fzdwx/burst-server:1.5 likelovec/burst-server:1.5
docker push likelovec/burst-server:1.5
docker run --name burst-server -p 10086:10086 -p 39399:39399 -d likelovec/burst-server:1.5