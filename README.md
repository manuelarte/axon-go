# Axon-Go

[![Go](https://github.com/manuelarte/axonserver-connector-go/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/axon-go/actions/workflows/go.yml)

This repository contains an [Axon Server][axon-server] Integration with Go, based on the [open api][axon-server-integration-swagger-ui] generated code.
Some demo projects can be checked in the [./examples](./examples) directory.

[AxonServer Integration Docs](https://docs.axoniq.io/axon-server-reference/v2025.0/axon-server/administration/integration/)

## Bugs?

+ Register endpoint needs a field `contentType`.
+ The Endpoint schema type and wrappingType can be an enum, not a string.

## Examples

Check the [examples](./examples) folder with different ways to use AxonServer with Go.


[axon-server]: https://axoniq.io/product-overview/axon-server
[axon-server-integration-swagger-ui]: http://localhost:8024/swagger-ui
