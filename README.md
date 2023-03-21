# proto2cli

This utility generates cli binaries from a set of grpc protoc to enable both direct access to the grpc server as well as to a grpc gateway server (either a default one or using the spec in the grpc-gateway annotations).

The CLI generated has a few artifacts:
1. A python library for the main cli module
2. A library that can be incorporated as extended cli options in a larger cli binary

## Development

`make install` installs the plugin locally and you can use it in buf or via the protoc binary as a plugin when processing protobuf files.

## Future Work

* Generate artifacts for languages other than python so host CLIs can be other languages.
