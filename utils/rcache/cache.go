package rcache

import ("fmt"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/go-redis/redis/v7"

	"github.com/go-redis/cache/v8")

func Example_advancedUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6381",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: fastcache.New(100 << 20), // 100 MB
	})

	obj := new(Object)
	err := mycache.Once(&cache.Item{
		Key:   "mykey",
		Value: obj, // destination
		Do: func(*cache.Item) (interface{}, error) {
			return &Object{
				Str: "mystring",
				Num: 42,
			}, nil
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	// Output: &{mystring 42}
}