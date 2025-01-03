# Use the official Ubuntu base image
FROM ubuntu:20.04

# Set environment variables to avoid prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Install prerequisites
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    gnupg \
    unzip \
    software-properties-common \
    openjdk-11-jdk \
    && apt-get clean

# Add Confluent's APT repository and install Confluent Platform
RUN curl -O https://packages.confluent.io/archive/7.8/confluent-7.8.0.zip \
    && unzip confluent-7.8.0.zip -d /opt \
    && rm confluent-7.8.0.zip

# Set Confluent environment variables
ENV CONFLUENT_HOME=/opt/confluent-7.8.0
ENV PATH="$CONFLUENT_HOME/bin:$PATH"

# Download and install Confluent Hub CLI from tarball
RUN wget -q http://client.hub.confluent.io/confluent-hub-client-latest.tar.gz \
    && tar -xzf confluent-hub-client-latest.tar.gz -C /usr/local \
    && rm confluent-hub-client-latest.tar.gz

# Add Confluent Hub CLI to PATH
ENV PATH="/usr/local/confluent-hub-client/bin:$PATH"

# Verify Confluent Hub CLI installation
RUN which confluent-hub && confluent-hub version

# Install Debezium PostgreSQL Connector explicitly specifying the component directory
RUN confluent-hub install --no-prompt debezium/debezium-connector-postgresql:latest \
    --component-dir $CONFLUENT_HOME/share/java

# Expose default Confluent ports
EXPOSE 9092 9021 8083 8082 2181

# Set the working directory
WORKDIR /opt/confluent

# Command to start the Confluent Platform
CMD ["confluent", "local", "services", "start"]
