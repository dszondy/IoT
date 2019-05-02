# Select openjdk container for base
FROM openjdk:11.0.2-jdk-slim-stretch as final
RUN mkdir src
COPY ./src/ /src
WORKDIR src
RUN ls
RUN mkdir build; javac -cp postgresql-42.2.5.jar -d build -sourcepath . ./*/*.java
WORKDIR build
ENTRYPOINT java -cp .:../postgresql-42.2.5.jar main.Main