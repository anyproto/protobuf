regenerate:
	go install github.com/anyproto/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/anyproto/protobuf/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/anyproto/protobuf/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
