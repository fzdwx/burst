#!/usr/bin/env just --justfile
mvn_setting_path := "/mnt/c/Users/98065/.m2/settings.xml"

target_path := "/root/burst-server.jar"
log_path := "/root/burst-server.log"

# maven build without tests
l:
   @just -l

# build burst-server
build:
   mvn -s {{mvn_setting_path}} -DskipTests clean package -Pprod

runServer:
   java -jar burst-server/target/burst-server-1.0-SNAPSHOT.jar

# compile proto files
proto:
   protoc --java_out=./burst-server/src/main/java ./protocol/burst.proto
   protoc --go_out=./burst-client/ ./protocol/burst.proto

kill:
    ssh  root@114.132.249.192 "sh /root/killbrust.sh"
    ssh  root@114.132.249.192 "rm -rf {{target_path}}"

run:
    ssh  root@114.132.249.192 "nohup java -jar {{target_path}} > /root/burst-server.log 2>&1 &"

log action:
    ssh  root@114.132.249.192 "{{action}} {{log_path}}"
pub:
    scp -r burst-server/target/burst-server-1.0-SNAPSHOT.jar root@114.132.249.192:{{target_path}}

# dependencies tree for compile
dependencies:
  mvn dependency:tree -Dscope=compile > dependencies.txt

# display updates
updates:
  mvn versions:display-dependency-updates > updates.txt


#set GOARCH=amd64
#set GOOS=linux
# nohup java -jar /root/burst-server.jar > /root/burst-server.log 2>&1 &