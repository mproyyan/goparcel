#!/bin/bash

# Initialize MongoDB server as a replica set (unauthenticated version)

if [[ $(mongosh --quiet --eval "rs.status().ok" 2>/dev/null) -ne 1 ]]; then
    echo "Replica set not initialized, attempting to initiate..."
    if ! mongosh --quiet --eval "rs.initiate()" >/dev/null; then
        echo "❌ Failed to initiate replica set"
        exit 1
    else
        echo "✅ Replica set initiated successfully."
    fi
else
    echo "ℹ️ Replica set is already initialized."
fi

# Final health check
if [[ $(mongosh --quiet --eval "rs.status().ok" 2>/dev/null) -ne 1 ]]; then
    echo "❌ Health check failed: Replica set is not ok."
    exit 1
else
    echo "✅ Health check passed: Replica set is ok."
fi
