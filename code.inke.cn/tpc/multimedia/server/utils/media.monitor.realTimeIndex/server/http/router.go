
// Generated by the daenerys tool.  DO NOT EDIT!
package http
import (
	httpserver "git.inke.cn/inkelogic/daenerys/http/server"
)


func initRoute(s httpserver.Server) {

	s.ANY("/ping", ping)



}
