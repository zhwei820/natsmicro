package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/rs/zerolog/log"
	"github.com/zhwei820/natsmicro/pbservice"
	_ "github.com/zhwei820/natsmicro/routers"
	"github.com/zhwei820/natsmicro/utils/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// https://stackoverflow.com/questions/41432193/how-to-delete-a-file-using-golang-on-program-exit/41455960#41455960
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL) // 程序退出清理资源
		<-sigs
		destroy()
		os.Exit(0)
	}()

	debug, _ := beego.AppConfig.Bool("debug")
	beego.BConfig.WebConfig.StaticDir["/asset"] = "./views/templates"
	if debug {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
	}
	orm.DefaultTimeLoc = time.UTC
	beego.BConfig.ServerName = "snail server 1.0"
	fmt.Println("server listen at: ", beego.AppConfig.String("httpport"))
	beego.Run()

}

func destroy() {
	log.Info().Msg("清除资源")
	pbservice.NatsConn.Close()
	util.Destroy()
}
