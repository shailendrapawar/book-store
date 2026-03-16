#!/bin/bash

echo "Generating swagger docs..."
swag init -g cmd/api/main.go
echo "Done."