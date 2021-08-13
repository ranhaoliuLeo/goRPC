package gorpc

import (
	"log"
	"net"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		go server.ServeConn(conn)
	}
}
 
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

// rpc 序列化函数
func (server *Server) ServeConn(conn net.Conn) {

}
