FROM openjdk:17
ARG JAR_FILE=burst-server-1.3.jar
COPY ${JAR_FILE} burst.jar
EXPOSE 10086
ENTRYPOINT ["java","-jar","/burst.jar"]