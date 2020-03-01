package nrpc_client

import (
	"github.com/nats-io/nats.go"
	"natsmicro/common/natsConn"

	"natsmicro/apps/helloworld2/rpc/proto/helloworld"
	"os"
)

var HelloWordNrpcCli *helloworld.Greeter2Client
var closeChan = make(chan interface{}, 0)

func init() {
	Init()
}
func Init() {
	var natsURL = nats.DefaultURL
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	// Connect to the NATS server.
	nc := natsConn.GetNatConn(natsURL, "natsConn")
	defer nc.Close()
	// This is our generated client.
	HelloWordNrpcCli = helloworld.NewGreeter2Client(nc)
	<-closeChan
}

func Close() {
	closeChan <- nil
}
