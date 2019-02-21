package say

import (
	. "github.com/zhwei820/natsmicro/pb/natsmicro/say"
	"time"
)

func SayHello(in []byte) ([]byte, error) {
	args := &SayInput{}
	args.Unmarshal(in)
	output := &SayOutput{Url: "yyyyyyyyyyy say OKxxxxxxxxxxxxxxxxxxx", Title: args.Query,}
	time.Sleep(300 * time.Millisecond)
	println(args.PageNumber)
	return output.Marshal()
}
