regenerate:
	go install github.com/anyproto/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/anyproto/protobuf/protobuf/ \
	--gogo_out=. events.proto 
