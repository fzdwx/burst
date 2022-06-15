FROM openjdk:17
ARG JAR_FILE=target/burst-server-1.3.jar
COPY ${JAR_FILE} burst.jar
ENTRYPOINT ["java","-jar","/burst.jar"]