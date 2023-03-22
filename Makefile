
VERSION=v0.0.1

install:
	cd cmd/protoc-gen-cli && go install

publish:
	git push && buf push
