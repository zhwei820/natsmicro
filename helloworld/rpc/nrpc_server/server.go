package nrpc_server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"log"
	"os"
	"time"

	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	"natsmicro/conf"
	"natsmicro/helloworld/rpc/proto/helloworld"

	"github.com/nats-io/nats.go"
)

// server implements the helloworld.GreeterServer interface.
type server struct{}

// SayHello is an implementation of the SayHello method from the definition of
// the Greeter service.
func (s *server) SayHello(ctx context.Context, req helloworld.HelloRequest) (resp helloworld.HelloReply, err error) {
	resp.Message = "Hello " + req.Name
	glog.Info(resp.Message)
	time.Sleep(50 * time.Millisecond)
	return
}

func StartServer() {
	var natsURL = conf.NatsUrl
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	// Connect to the NATS server.
	nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Our server implementation.
	s := &server{}

	// The NATS handler from the helloworld.nrpc.proto file.
	h := helloworld.NewGreeterHandler(context.TODO(), nc, s)

	// Start a NATS subscription using the handler. You can also use the
	// QueueSubscribe() method for a load-balanced set of servers.
	//sub, err := nc.Subscribe(h.Subject(), h.Handler)
	for ii := 0; ii <= 1000; ii += 1 {
		sub, err := nc.QueueSubscribe(h.Subject(), "groupname", h.Handler)

		if err != nil {
			log.Fatal(err)
		}
		defer sub.Unsubscribe()

	}
	// Keep running until ^C.
	fmt.Println("server is running, ^C quits.")
	select {}
}
