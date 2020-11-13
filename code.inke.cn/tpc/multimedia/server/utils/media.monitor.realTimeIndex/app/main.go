package main

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/server"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/server/http"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/service"
	"git.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/inkelogic/daenerys"
)

func init() {
	configS := flag.String("config", "app/config/ali-test/config.toml", "Configuration file")
	appS := flag.String("app", "", "App dir")
	flag.Parse()
	
	daenerys.Init(
		daenerys.ConfigPath(*configS),
	)
	
	if *appS != "" {
		daenerys.InitNamespace(*appS)
	}
}


func main() {

	defer daenerys.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	srv := service.New(cfg)

	// init and start http server
	http.Init(srv, cfg)

	//dao.InitDB()

	server.StartServer()

	defer http.Shutdown()


	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("media.monitor.realTimeIndex server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}

