
VERSION=v0.0.1

docker:
	docker build -f Dockerfile.proto2cli -t plugins.buf.build/panyam/proto2cli:${VERSION}-1 .

install:
	cd cmd/protoc-gen-cli && go install
