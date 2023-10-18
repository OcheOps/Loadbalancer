package main

import(
	"fmt"
	"net/http/http.util"
	"net/url"
	"os"
	"net/http"

)


type Server interface {

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
 type Loadbalancer struct {
	port string
	roundRobinCount int
	servers []Server
 }

 


func NewLoadbalancer(port string, servers[]Server ) *Loadbalancer{
	return &Loadbalancer{
		port: port,
		roundRobinCount: 0,
		servers: make([]Server,0),
	}
}

//}
func handleErr(err error){
	if err !=nil{
		fmt.Print("error :%v/n",err)
		os.Exit(1)
	
	}
}

 func (lb *Loadbalancer)  	getNextAvailableServer() Server{
 	nextAvailableServer := lb.servers[lb.roundRobinCount]
 	lb.roundRobinCount = (lb.roundRobinCount + 1) % len(lb.servers)
 	return nextAvailableServer
 }

func (lb *Loadbalancer) serveProxy(rw http.ResponseWriter, req *http.Request){
	nextAvailableServer := lb.getNextAvailableServer()
	nextAvailableServer.ServeHTTP(rw, req)
}


func Address(s *SimpleServer) string{
	return s.addr
}

func IsAlive(s *SimpleServer) bool{
	_, err := http.Get(s.addr)
	if err != nil{
		return false
	}
}


func Serve(s *SimpleServer, rw http.ResponseWriter, req *http.Request){
	targetServer := lb.getNextAvailableServer()
	targetServer.ServeHTTP(rw, req)
	s.proxy.ServeHTTP(rw, req)
}



func main(){
	servers:=[]Server{
		NewSimpleServer("http://localhost:3000"),
		NewSimpleServer("http://localhost:3001"),
		NewSimpleServer("http://localhost:3002"),

	}
	lb := NewLoadbalancer(":8080", servers)
	http.HandleFunc("/", lb.serveProxy)
	http.ListenAndServe(":8080", nil)

	fmt.Println("serving at localhost")
}