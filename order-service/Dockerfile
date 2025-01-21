# Step 1: Use your custom image
FROM kingstonduy/custom-image:latest

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the executable and environment file into the container
COPY main .
# Copy the file to the container
COPY resources/local.env ./resources/


# Step 5: Define the command to run your Go executable
CMD ["./main"]
