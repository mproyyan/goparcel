# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/cargos internal/common

# Copy cargos and common to internal folder
COPY internal/cargos internal/cargos
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/cargos internal/common
RUN go work sync

# Build application
RUN go build -o cargo-service /app/internal/cargos

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/cargo-service .

# Run application
CMD [ "./cargo-service" ]