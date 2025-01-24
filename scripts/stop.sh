#!/bin/bash

# Go back one directory
cd ..

# Stop Docker containers
docker-compose stop

# Check if Docker containers were stopped successfully
if [ $? -eq 0 ]; then
    echo "Docker containers stopped successfully!"
else
    echo "Failed to stop Docker containers."
    exit 1
fi
