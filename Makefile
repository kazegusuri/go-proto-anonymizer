REPOSITORY=github.com/kazegusuri/go-proto-anonymizer

regenerate:
	protoc -I. -I${GOPATH}/src --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/gogoproto/descriptor.proto:. anonymizer.proto

	protoc -I. --go_out=. examples/*.proto
	protoc -I. --goanonymizer_out=. examples/*.proto
	protoc -I. --go_out=. examples/base/*.proto
	protoc -I. --goanonymizer_out=. examples/base/*.proto

install:
	go install $(REPOSITORY)/protoc-gen-goanonymizer

test:
	go test -v $(REPOSITORY)/examples
