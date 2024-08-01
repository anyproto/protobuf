regenerate:
	go install github.com/anyproto/protobuf/protoc-min-version
	go install github.com/anyproto/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/anyproto/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/anyproto/protobuf/types:. timefail.proto
