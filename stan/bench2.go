package main

import (
	"natsmicro/conf"
	"strconv"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/nats-io/stan.go"
)

var donewg sync.WaitGroup

func main() {
	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, err := stan.Connect("test", "myID3", stan.NatsURL(conf.NatsUrl))
	println(conf.NatsUrl)
	//ns, err := stan.Connect("test-cluster", "myID3", stan.NatsURL(nats.DefaultURL))

	if err != nil {
		println(err.Error())
	}
	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming

	start := time.Now()

	//// Simple Async Subscriber
	//sub, _ := ns.QueueSubscribe("foo", "grouptest",func(m *stan.Msg) {
	//	log.Printf("sub: %s\n", string(m.Data))
	//}, stan.StartWithLastReceived())
	//sub1, _ := ns.QueueSubscribe("foo", "grouptest", func(m *stan.Msg) {
	//	log.Printf("sub1: %s\n", string(m.Data))
	//}, stan.StartWithLastReceived())

	N := 1000000
	//N := 100000
	donewg.Add(N)

	// Run Subscribers first
	for i := 0; i < N; i++ {
		//time.Sleep(time.Second)
		var i = i
		go func() {
			for ii := 0; ii < 1; ii++ {

				ns.Publish("foo", []byte("Hello World "+strconv.Itoa(i)))
			}
			donewg.Done()
		}()
	}
	donewg.Wait()
	end := time.Now()

	span := end.Sub(start)
	sec := float64(span) / float64(time.Second)
	g.Dump(sec)
	g.Dump(float64(N) / sec)

	// Unsubscribe
	//sub.Unsubscribe()
	//sub1.Unsubscribe()

	// Close connection
	ns.Close()
}
