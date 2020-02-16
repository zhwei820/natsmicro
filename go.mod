module natsmicro

require (
	github.com/gogf/gf v1.11.4
	github.com/golang/protobuf v1.3.2
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	github.com/nats-io/nats-streaming-server v0.16.2 // indirect
	github.com/nats-io/nats.go v1.9.1
	github.com/nats-io/stan.go v0.6.0
	github.com/nats-rpc/nrpc v0.0.0-20190920042445-3ae2c6c6aad4
)

replace github.com/nats-rpc/nrpc v0.0.0-20190920042445-3ae2c6c6aad4 => github.com/zhwei820/nrpc v0.0.0-20191028151421-a74606762646

go 1.13
