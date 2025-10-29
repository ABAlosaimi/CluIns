package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func getLocalIP() (string, error) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {

		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}

	}

	return "", fmt.Errorf("no non-loopback IP found")
}

func main() {

	http.HandleFunc("/metric/resources", reportCpuAndMemData)

	ip, err := getLocalIP()
	if err != nil {
		log.Fatalf("Failed to get local IP of the host: %v", err)
	}

	bind := fmt.Sprintf("%s:0", ip)
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	fmt.Printf("Server listening at http://%s\n", addr.String())

	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("the service unable to start: %v", err)
	}
	
}

func reportCpuAndMemData(res http.ResponseWriter, req *http.Request) {


}