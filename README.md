# TTL Cache Implementation

### Introduction

This project implements a custom cache system in Go, which supports time-to-live (TTL) functionality with an active expiration strategy. Unlike traditional passive (or lazy) expiration where items are only checked and potentially purged on access, this cache actively manages and removes expired items in the background.

### Key Features

- Active TTL management using background routines
- Uses a custom linked list to store cache items
- Supports string-like byte slice keys and byte slice values
- Minimal memory allocation (at most one per operation)
- Thread-safe implementation

## Installation

This project requires Go 1.15 or higher. [Click here](https://go.dev/dl/) to download Go.

To run the benchmarks, execute the following command:

```bash
cd cache
go test -bench=.
```

## Usage

For sample usage, run the following command:

```bash
cd main
go run main.go
```

## project structure

```bash
/cache
├── cache.go     # Implementation of the TTL cache
└── cache_test.go     # Benchmark for the cache
```

## Author
Parsa Mazaheri