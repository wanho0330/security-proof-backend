# Ethereum 2.0 Custody Proof Storage System Using Smart Contracts

This project is a custody proof storage system that utilizes Ethereum 2.0 smart contracts.

## Overview
- The package structure is inspired by project structures commonly used in Google projects.
- Interfaces are implemented to support mock testing.
- Latest libraries, including `goverter` and `gojet`, are used for code generation.
- RPC and web communication using `Protobuf` and `Connect RPC` to optimize data transfer.
- Token-based authentication is used, with `JWT` tokens stored in `Redis` as access tokens and refresh tokens.
- JWT tokens contain user index and role information, allowing for an authorization mechanism.
- For dashboard statistics, this app communicate with `Elasticsearch`, utilizing a type-safe, `typed API` approach.
- All packages are maintained according to consistent coding standards using `golangci-lint`.

## Package Structure
- `cmd` : Contains applications that can be executed.
- `internal` : Contains packages that are not accessible externally.
- `pkg` : Contains packages accessible by external references.
- `config` : Contains various configuration files, such as SSL certificates.
- `script` : Contains a collection of shell scripts.
