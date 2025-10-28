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

	http.HandleFunc("/", nil)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	
	addr := listener.Addr().(*net.TCPAddr)
	port := addr.Port

	
	ip, err := getLocalIP()
	if err != nil {
		log.Fatalf("Failed to get local IP: %v", err)
	}

	serverAddr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("Server starting at http://%s\n", serverAddr)

	// Close the listener and let http.ListenAndServe listen on the same port
	listener.Close()
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		fmt.Printf("the service unable to start: %v", err.Error())
	}
}

func reportCpuAndRamData(res http.ResponseWriter, req *http.Request) {

}