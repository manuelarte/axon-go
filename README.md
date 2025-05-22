# Axon-Go

[![Go](https://github.com/manuelarte/axonserver-connector-go/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/axon-go/actions/workflows/go.yml)

This repository contains an [Axon Server][1] Integration with Go, based on the open api generated code.
Some demo projects can be checked in the [./examples](./examples) directory.

[AxonServer Integration Docs](https://docs.axoniq.io/axon-server-reference/v2025.0/axon-server/administration/integration/)

## Bugs?

+ Register endpoint needs a field `contentType`.
+ The Endpoint schema type and wrappingType can be an enum, not a string.


[1]: https://axoniq.io/product-overview/axon-server
