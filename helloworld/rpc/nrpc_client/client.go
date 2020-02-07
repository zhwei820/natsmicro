package nrpc_client

import (
	"log"
	"natsmicro/conf"
	"natsmicro/helloworld/rpc/proto/helloworld"
	"os"
	"time"

	"github.com/nats-io/nats.go"
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
	// Connect to the NATS server.
	nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()

	// This is our generated client.
	HelloWordNrpcCli = helloworld.NewGreeterClient(nc)
}
