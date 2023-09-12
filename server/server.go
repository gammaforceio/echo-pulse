package server

import (
	"fmt"
	"net"
	"strings"

	"github.com/SiirRandall/echo-pulse/logger"
)

type UDPEchoServer struct {
	LogDir    string
	UniqueIPs map[string]struct{}
	Blacklist []string
}

func NewUDPEchoServer(logDir string, blacklist []string) *UDPEchoServer {
	return &UDPEchoServer{
		LogDir:    logDir,
		UniqueIPs: make(map[string]struct{}),
		Blacklist: blacklist,
	}
}

func (s *UDPEchoServer) Start(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Listening on %s for UDP packets...\n", addr)

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		clientData := strings.ReplaceAll(string(buf[:n]), "\n", "")

		// Check if data contains any blacklisted keywords
		blacklisted := false
		for _, keyword := range s.Blacklist {
			if strings.Contains(clientData, keyword) {
				blacklisted = true
				break
			}
		}
		if blacklisted {
			continue // Do not process this packet
		}

		logger.LogToFile(
			s.LogDir,
			"udp_all_traffic.log",
			fmt.Sprintf("Received from %s: %s\n", clientAddr, clientData),
		)
		// Check for unique IP and log if needed
		if _, exists := s.UniqueIPs[clientAddr.IP.String()]; !exists {
			s.UniqueIPs[clientAddr.IP.String()] = struct{}{}
			logger.LogToFile(s.LogDir, "unique_ips.log", clientAddr.IP.String()+"\n")
		}

		// Echo the data back to the client
		conn.WriteToUDP(buf[:n], clientAddr)
	}
}
