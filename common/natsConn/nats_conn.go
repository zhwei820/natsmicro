package natsConn

import (
	"github.com/gogf/gf/os/glog"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func GetNatConn(natsURL, natCliName string) *nats.Conn {
	option := make([]nats.Option, 0)
	option = append(option, nats.Name(natCliName),
		nats.PingInterval(30*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.ReconnectWait(3*time.Second),
		nats.MaxReconnects(200),
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			glog.Error("nat连接断开,等待重新连接时间:3s")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			glog.Error("nat重新连接，URL:%v", conn.ConnectedUrl())
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			glog.Error("nat退出,err:%v", conn.LastError())
		}),
	)

	nc, err := nats.Connect(natsURL, option...)
	glog.Info("连接成功" + nc.ConnectedUrl())

	// Connect to the NATS server.
	//nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	return nc
}
