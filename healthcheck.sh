#!/bin/bash

# Initialize MongoDB server as a replica set
# NOTE: This is a workaround since initializing the MongoDB server as
# a replica set cannot be done in the script placed in the
# /docker-entrypoint-initdb.d/ directory.

if [[ $(mongosh --quiet -u "$MONGODB_INITDB_ROOT_USERNAME" -p "$MONGODB_INITDB_ROOT_PASSWORD" --authenticationDatabase admin --eval "rs.status().ok") -ne 1 ]]; then
    echo "Replica set not initialized, attempting to initiate..."
    if ! mongosh --quiet -u "$MONGODB_INITDB_ROOT_USERNAME" -p "$MONGODB_INITDB_ROOT_PASSWORD" --authenticationDatabase admin --eval "rs.initiate()"; then
        echo "Failed to initiate replica set"
        exit 1
    else
        echo "Replica set initiated successfully."
    fi
else
    echo "Replica set is already initialized."
fi

# Final health check
if [[ $(mongosh --quiet -u "$MONGODB_INITDB_ROOT_USERNAME" -p "$MONGODB_INITDB_ROOT_PASSWORD" --authenticationDatabase admin --eval "rs.status().ok") -ne 1 ]]; then
    echo "Health check failed: Replica set is not ok."
    exit 1
else
    echo "Health check passed: Replica set is ok."
fi