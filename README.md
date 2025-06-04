# yaksha-opaque
Final Year Project

# Opaque Benchmarking Application

## Overview

This project is a benchmarking application designed to evaluate the efficiency of different OPAQUE protocols by simulating the key exchange process between servers and clients. The main focus is to measure and compare key performance metrics such as latency, loading speed, and calculation speed across various OPAQUE-based authentication protocols including OPRF-based OPAQUE and 3HDH.

## Purpose

The application aims to provide a comprehensive and systematic way to benchmark different secure password authentication protocols to help researchers and developers understand their performance characteristics and trade-offs. This is particularly useful for selecting the most efficient protocol for specific application scenarios involving secure authentication.

## Overall Architecture

The system is designed with a modular and extensible architecture consisting of the following components:

- **Protocol Modules**  
  Each OPAQUE protocol implementation resides in its own module handling cryptographic operations and key exchange logic.

- **Server-Client Architecture**  
  The server hosts the protocol implementations and manages requests from multiple clients. The clients simulate authentication attempts, initiating the key exchange process.

- **Benchmarking Engine**  
  This component orchestrates benchmarking by measuring latency, loading time, and computational effort involved in each protocol run. Metrics are collected and aggregated to produce comparative results.

The project uses Go for its strong concurrency support, performance, and rich standard libraries, making it well-suited for networked cryptographic benchmarking tasks.

The folder structure reflects this design, organizing code by responsibilities such as protocol implementations, benchmarking logic, server and client applications, utility functions, and configuration.

---

