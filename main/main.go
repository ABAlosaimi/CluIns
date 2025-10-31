package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func getLocalIP() (string, error) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unable to get host's interface addresses: ", err
	}

	for _, addr := range addrs {

		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}

	}

	return "unable to get host's IPv4: ", fmt.Errorf("no non-loopback IPv4 found")
}

func main() {

	http.HandleFunc("/metric/resources", reportCpuAndMemData)

	ip, err := getLocalIP()
	if err != nil {
		log.Fatalf("Failed to get the IPv4 of the host: %v", err)
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

	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")

	memTime := time.NewTicker(time.Second)
	defer memTime.Stop()

	cpuTime := time.NewTicker(time.Second)
	defer cpuTime.Stop()

	clientGone := req.Context().Done()
	rc := http.NewResponseController(res)

	for {
		select{
		case <- clientGone:
			fmt.Println("the client is disconnected")
			return
		case <- memTime.C:
			m, err := mem.VirtualMemory()
			if err != nil {
				log.Fatalf("unable to get memory data: %v", err.Error())
				return
			}

			if _, err := fmt.Fprintf(res, "event:memory\n data: Total: %d, Used: %d, Pers: %.2f,\n\n", m.Total, m.Used, m.UsedPercent); err != nil {
					log.Fatalf("unable to write back to the client: %v", err.Error())
					return
			}

			rc.Flush()

		case <- cpuTime.C:
			cpu, err := cpu.Times(false) 
			if err != nil {
				log.Fatalf("unable to get memory data: %v", err.Error())
				return
			}

			if _, err := fmt.Fprintf(res, "event:cpu\n data: User: %.2f, System: %.2f, Idle : %.2f,\n\n", cpu[0].User, cpu[0].System, cpu[0].Idle); err != nil {
					log.Fatalf("unable to write back to the client: %v", err.Error())
					return
			}

			rc.Flush()
		}

	}

}