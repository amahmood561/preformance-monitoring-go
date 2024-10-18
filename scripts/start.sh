#!/bin/bash

# Build the application
echo "Building the Go Performance Monitoring Toolkit..."
go build -o monitor ./cmd/monitor

# Run the application
echo "Starting the monitor..."
./monitor
