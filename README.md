# go-proto-anonymizer

A protoc plugin which generates `Anonymize()` methods for each generated go structs.

## Anonymization rule

```
syntax = "proto3";
package kazegusuri.example;
import "anonymizer.proto";

message SimpleMessage {
  string string_value = 1 [(kazegusuri.anonymizer.field) = {method: NULL}];
  bool bool_value = 2 [(kazegusuri.anonymizer.field) = {method: NULL}];
}

message NestedMessage {
  SimpleMessage nested = 1;
  SimpleMessage masked_nested = 2  [(kazegusuri.anonymizer.field) = {method: NULL}];
}
```

```golang
package main

import (
	fmt
	anonymizer "github.com/kazegusuri/go-proto-anonymizer"
)

func main() {
	msg := &NestedMessage{
		Nested: &SimpleMessage{
			StringValule: "xxx",
			BoolValue: true,
		},
		MaskedNested: &SimpleMessage{
			StringValule: "yyy",
			BoolValue: true,
		},
	}
	anonymizer.Anonymize(msg)
	fmt.Printf("%v\n", msg.Nested.StringValue) // => ""
	fmt.Printf("%v\n", msg.Nested.BoolValue) // => false
	fmt.Printf("%v\n", msg.MaskedNested) // => nil
}
```

## Install and usage

First you need `protoc-gen-goanonymizer` of protoc plugin. You can get it by doing:

```bash
$ go get github.com/kazegusuri/go-proto-anonymizer/protoc-gen-goanonymizer
```

To generate anonymizer methods from proto, use `--goanonymizer_out` with `protoc` command. That runs `protoc-gen-zap-marshaler` internally and then generates `*.anon.go` files.

```
$ protoc --goanonymizer_out=. path/to/example.proto
```
