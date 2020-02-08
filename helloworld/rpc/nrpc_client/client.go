package nrpc_client

import (
	"natsmicro/common/natsConn"
	"natsmicro/conf"
	"natsmicro/helloworld/rpc/proto/helloworld"
	"os"
)

var HelloWordNrpcCli *helloworld.GreeterClient

func init() {
	Init()
}
func Init() {
	var natsURL = conf.NatsUrl
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	nc := natsConn.GetNatConn(natsURL, "natsConn")
	//defer nc.Close()

	// This is our generated client.
	HelloWordNrpcCli = helloworld.NewGreeterClient(nc)
}
