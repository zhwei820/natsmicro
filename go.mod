module natsmicro

require (
	github.com/VictoriaMetrics/fastcache v1.5.7
	github.com/go-redis/cache/v8 v8.0.0-beta.4
	github.com/go-redis/redis/v7 v7.2.0
	github.com/gogf/gf v1.11.4
	github.com/golang/protobuf v1.4.2
	github.com/nacos-group/nacos-sdk-go v0.0.0-20191128082542-fe1b325b125c
	github.com/nats-io/nats-streaming-server v0.24.3 // indirect
	github.com/nats-io/nats.go v1.13.1-0.20220308171302-2f2f6968e98d
	github.com/nats-io/stan.go v0.10.2
	github.com/nats-rpc/nrpc v0.0.0-20190920042445-3ae2c6c6aad4
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/ratelimit v0.1.0
)

replace (
	github.com/nacos-group/nacos-sdk-go v0.0.0-20191128082542-fe1b325b125c => github.com/jjz/nacos-sdk-go v0.0.0-20191122033133-945e8c5a4346
	github.com/nats-rpc/nrpc v0.0.0-20190920042445-3ae2c6c6aad4 => github.com/zhwei820/nrpc v0.0.0-20191028151421-a74606762646
)

go 1.13
