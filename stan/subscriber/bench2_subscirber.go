package main

import (
	"natsmicro/conf"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/nats-io/stan.go"
)

var donewg sync.WaitGroup

func main() {

	logger := glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":     "./",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})

	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, err := stan.Connect("test", "myI3D", stan.NatsURL(conf.NatsUrl))

	if err != nil {
		println(err.Error())
	}
	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming

	start := time.Now()

	NN := 100
	var subs []stan.Subscription
	for i := 0; i < NN; i++ {

		// Simple Async Subscriber
		sub1, _ := ns.QueueSubscribe("foo", "grouptest", func(m *stan.Msg) {
			//log.Printf("sub1: %s\n", string(m.Data))
			logger.Print(m.Data)

		}, stan.StartWithLastReceived(), stan.DurableName("testdurableName"))
		subs = append(subs, sub1)

	}
	N := 10
	donewg.Add(N)

	donewg.Wait()
	end := time.Now()

	span := end.Sub(start)
	sec := float64(span) / float64(time.Second)
	g.Dump(sec)
	g.Dump(float64(N) / sec)

	// Unsubscribe
	for i := 0; i < NN; i++ {
		subs[i].Unsubscribe()
	}
	// Close connection
	ns.Close()
}
