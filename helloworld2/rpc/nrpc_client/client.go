package nrpc_client

import (
	"natsmicro/common/natsConn"
	"natsmicro/conf"
	"natsmicro/helloworld2/rpc/proto/helloworld"
	"os"
)

var HelloWordNrpcCli *helloworld.Greeter2Client

func init() {
	Init()
}
func Init() {
	var natsURL = conf.NatsUrl
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	// Connect to the NATS server.
	nc := natsConn.GetNatConn(natsURL, "natsConn")
	//defer nc.Close()
	// This is our generated client.
	HelloWordNrpcCli = helloworld.NewGreeter2Client(nc)
}
