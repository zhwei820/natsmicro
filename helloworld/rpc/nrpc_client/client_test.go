package nrpc_client

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	"natsmicro/conf"
	"natsmicro/helloworld/rpc/proto/helloworld"

	nats "github.com/nats-io/nats.go"
)

var wg sync.WaitGroup

func work(cli *helloworld.GreeterClient) {
	for ii := 0; ii < 20; ii += 1 {
		t1 := time.Now()
		// Contact the server and print out its response.
		resp, err := cli.SayHello(helloworld.HelloRequest{Name: "world"})
		if err != nil {
			//log.Fatal(err)
		}
		println(time.Now().Sub(t1).Seconds())

		fmt.Printf("Greeting: %s %v \n", resp.GetMessage(), ii)
	}
	wg.Done()
}

func TestInit(t *testing.T) {

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
	// This is our generated client.
	cli := helloworld.NewGreeterClient(nc)
	t1 := time.Now()
	for ii := 0; ii < 2000; ii += 1 {
		go work(cli)
		wg.Add(1)
	}
	wg.Wait()
	println(time.Now().Sub(t1).Seconds())
	// print
}
