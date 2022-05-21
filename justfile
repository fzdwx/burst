#!/usr/bin/env just --justfile
mvn_setting_path := "/mnt/c/Users/98065/.m2/settings.xml"

target_path := "/root/burst-server.jar"
# maven build without tests
build:
   mvn -s {{mvn_setting_path}} -DskipTests clean package -Pprod

kill:
    ssh  root@114.132.249.192 "sh /root/killbrust.sh"
    ssh  root@114.132.249.192 "rm -rf {{target_path}}"

run:
    ssh  root@114.132.249.192 "java -jar {{target_path}}"

pub: run
    scp -r burst-server/target/burst-server-1.0-SNAPSHOT.jar root@114.132.249.192:{{target_path}}

# dependencies tree for compile
dependencies:
  mvn dependency:tree -Dscope=compile > dependencies.txt

# display updates
updates:
  mvn versions:display-dependency-updates > updates.txt    