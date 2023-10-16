package main

import(
"net/http/http.util"

)

type SimpleServer struct {

	addr string
	proxy *http.util.ReverseProxy


}

func NewSimpleServer(addr string) *SimpleServer {
	ServerUrl, _ := url.Parse(addr)
	
}