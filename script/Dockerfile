FROM openjdk:17
ARG JAR_FILE=burst-server-1.5.jar
COPY ${JAR_FILE} burst.jar
EXPOSE 10086
EXPOSE 39399
ENTRYPOINT ["java","-jar","/burst.jar"]