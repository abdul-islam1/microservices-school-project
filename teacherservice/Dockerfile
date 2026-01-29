# Stage 1: Build the JAR
FROM maven:3.9.5-amazoncorretto-17 AS builder

WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline -B
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Runtime image (Alpine)
FROM amazoncorretto:17-alpine-jdk AS runtime

WORKDIR /app
COPY --from=builder /app/target/teacherservice-1.0.0.jar ./teacherservice.jar

EXPOSE 5002
CMD ["java", "-jar", "teacherservice.jar"]
