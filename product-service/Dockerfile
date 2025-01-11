# Stage 1: Set up Wine to run the Windows executable
FROM golang:1.22.4 AS go-build

# Set environment variables to prevent interactive prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Install Wine, curl, and other necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    wine64 \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Stage 2: Create the runtime image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Windows executable and any required environment files into the container
COPY ./main /app/
COPY ./resources/dev.env /app/resources/dev.env

# Install Wine in this stage as well
RUN apt-get update && \
    apt-get install -y wine64 && \
    rm -rf /var/lib/apt/lists/*

# Expose the port that the service will run on (optional, based on your app)
EXPOSE 7002

# Command to run the executable using Wine
CMD ["wine", "/app/main"]
