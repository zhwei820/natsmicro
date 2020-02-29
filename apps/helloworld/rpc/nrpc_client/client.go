package nrpc_client

import (
	"github.com/nats-io/nats.go"
	"natsmicro/common/natsConn"

	"natsmicro/apps/helloworld/rpc/proto/helloworld"
	"os"
)

var HelloWordNrpcCli *helloworld.GreeterClient

func init() {
	Init()
}
func Init() {
	var natsURL = nats.DefaultURL
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	nc := natsConn.GetNatConn(natsURL, "natsConn")
	//defer nc.Close()

	// This is our generated client.
	HelloWordNrpcCli = helloworld.NewGreeterClient(nc)
}
