package server

import (
	"fmt"
	"net"
	"strings"

	"github.com/gammaforceio/echo-pulse/logger"
)

type EchoServer struct {
	LogDir    string              // directory to save log files
	UniqueIPs map[string]struct{} // holds unique IPs connected to the server
	Blacklist []string            // list of blacklisted keywords
}

func NewEchoServer(logDir string, blacklist []string) *EchoServer {
	return &EchoServer{
		LogDir:    logDir,
		UniqueIPs: make(map[string]struct{}),
		Blacklist: blacklist,
	}
}

func (s *EchoServer) StartUDP(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to start UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Listening on %s for UDP packets...\n", addr)

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		clientData := strings.ReplaceAll(string(buf[:n]), "\n", "")

		if s.isBlacklisted(clientData) {
			continue
		}

		logger.LogToFile(
			s.LogDir,
			"udp_all_traffic.log",
			fmt.Sprintf("Received from %s: %s\n", clientAddr, clientData),
		)

		if _, exists := s.UniqueIPs[clientAddr.IP.String()]; !exists {
			s.UniqueIPs[clientAddr.IP.String()] = struct{}{}
			logger.LogToFile(s.LogDir, "unique_ips.log", clientAddr.IP.String()+"\n")
		}

		conn.WriteToUDP(buf[:n], clientAddr)
	}
}

func (s *EchoServer) StartTCP(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to start TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on %s for TCP connections...\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.handleTCPConnection(conn)
	}
}

func (s *EchoServer) handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from TCP:", err)
			break
		}

		clientData := strings.ReplaceAll(string(buf[:n]), "\n", "")

		if s.isBlacklisted(clientData) {
			continue
		}

		logger.LogToFile(
			s.LogDir,
			"tcp_all_traffic.log",
			fmt.Sprintf("Received from %s: %s\n", conn.RemoteAddr(), clientData),
		)

		ip := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		if _, exists := s.UniqueIPs[ip]; !exists {
			s.UniqueIPs[ip] = struct{}{}
			logger.LogToFile(s.LogDir, "unique_ips.log", ip+"\n")
		}

		conn.Write(buf[:n])
	}
}

func (s *EchoServer) isBlacklisted(data string) bool {
	for _, keyword := range s.Blacklist {
		if strings.Contains(data, keyword) {
			return true
		}
	}
	return false
}
