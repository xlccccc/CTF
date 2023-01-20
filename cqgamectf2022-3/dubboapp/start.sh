#!/bin/sh

/dubbo/java/apache-zookeeper-3.5.9-bin/bin/zkServer.sh start
sleep 10
python3 -c 'import os;import pty;pty.spawn(["/bin/bash","-c","/dubbo/java/jdk1.8.0_202/bin/java -jar /dubbo/java/target/mydubbo-server.jar"])'
