#!/bin/bash

# Go back one directory
cd ..

# Stop and remove Docker containers
docker-compose down

# Check if Docker containers were stopped successfully
if [ $? -eq 0 ]; then
    echo "Docker containers stopped and removed successfully!"
else
    echo "Failed to stop Docker containers."
    exit 1
fi
