package main

import(
"net/http/http.util"
"net/url"
)

type SimpleServer struct {

	addr string
	proxy *http.util.ReverseProxy


}

func NewSimpleServer(addr string) *SimpleServer {
	ServerUrl, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	return &SimpleServer{
		addr: addr,
		proxy: http.util.NewSingleHostReverseProxy(ServerUrl),

	}

}

//func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	s.proxy.ServeHTTP(w, r)
//}