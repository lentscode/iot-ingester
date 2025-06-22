# Microservices Example for an IoT ingester system

This project simulate (or at least tries to 😅) a distributed data ingestion system for IoT devices connected to the cloud, built with Go, MySQL and Apache Kafka.
The idea behind this project consists in few points:

- The single device sends data from real-world events and sends it to the `ingester` service via HTTP request.
- The `ingester` then creates a Kafka event notifying that some data needs to be processed.
- This event is captured by the `processor`, which saves the data into its DB.