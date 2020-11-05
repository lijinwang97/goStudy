package http

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/service"
	"git.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/inkelogic/daenerys"
	httpserver "git.inke.cn/inkelogic/daenerys/http/server"
	httpplugin "git.inke.cn/inkelogic/daenerys/plugins/http"
)

var (
	svc *service.Service

	httpServer httpserver.Server
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = daenerys.HTTPServer()

	// add namespace plugin
	httpServer.Use(httpplugin.Namespace)

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()

}

func Shutdown() {
	if httpServer != nil {
		httpServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}

