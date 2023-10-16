package main

import(
	"fmt"
	"net/http/http.util"
	"net/url"
	"os"

)


type Server struct {

	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type SimpleServer struct {

	addr string
	proxy *http.util.ReverseProxy


}

func NewSimpleServer(addr string) *SimpleServer {
	ServerUrl, err := url.Parse(addr)
	//if err != nil {
	//	panic(err)
	//}

	handleErr(err)

	return &SimpleServer{
		addr: addr,
		proxy: http.util.NewSingleHostReverseProxy(ServerUrl),

	}

}

//func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	s.proxy.ServeHTTP(w, r)
//}
 type LoadBalance struct {
	port string
	roundRobinCount int
	servers []Server
 }


func handleErr(err error){
	if err !=nil{
		fmt.Print("error :%v/n",err)
		os.Exit(1)
	
	}
}