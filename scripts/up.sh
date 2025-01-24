#!/bin/bash

# Go back one directory
cd ..

# Run docker-compose up -d
docker-compose up -d

# Check if Docker started successfully
if [ $? -eq 0 ]; then
    echo "Docker containers started successfully!"
else
    echo "Failed to start Docker containers."
    exit 1
fi

# Run the Go program
go run main.go

# Check if the Go program ran successfully
if [ $? -eq 0 ]; then
    echo "Go program executed successfully!"
else
    echo "Failed to execute the Go program."
    exit 1
fi
