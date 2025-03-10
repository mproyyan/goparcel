<h1 align="center">
    <strong>
        <i>GoParcel</i>
    </strong>
</h1>

## About GoParcel

GoParcel is a microservice-based logistics management system designed to streamline parcel delivery operations. Built using Golang, it follows Domain-Driven Design (DDD) principles and implements a clean architecture to structure its components efficiently. The system leverages gRPC for synchronous communication between services, enabling seamless data exchange across different modules.

## Core Feature
- Receive parcels from customers for shipment
- Transfer parcels between locations during transit
- Route parcels to their final destinations
- Ship parcels via cargo services
- Prevent parcel loss by scanning at every transfer point
- Deliver parcels to recipients using couriers

## Services

- Cargo service
- Courier service
- User service
- Location service
- Shipment service
- API gateway

## Tech Stack

- Golang
- MongoDB
- GRPC
- GraphQL

## Installation

- Clone this repo `git@github.com:mproyyan/goparcel.git`
- Copy `.env` from `.env.example` and fill the required data
- Run `docker compose up -d`