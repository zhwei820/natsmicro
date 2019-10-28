package main

import (
	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	_ "github.com/nats-io/nats.go"
	_ "natsmicro/helloworld/rpc/proto/helloworld"
)

