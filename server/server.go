// This is the server package
package server

// Import necessary packages and libraries
import (
	"fmt"     // For printing to console
	"net"     // To handle networking aspects
	"strings" // To manipulate strings

	"github.com/gammaforceio/echo-pulse/logger" // To log errors in a log file
)

// UDPEchoServer encapsulates the properties needed for the echo server
type UDPEchoServer struct {
	LogDir    string              // directory to save log files
	UniqueIPs map[string]struct{} // holds unique IPs connected to the server
	Blacklist []string            // list of blacklisted keywords
}

// NewUDPEchoServer returns a new instance of a UDP echo server
func NewUDPEchoServer(logDir string, blacklist []string) *UDPEchoServer {
	return &UDPEchoServer{
		LogDir:    logDir,
		UniqueIPs: make(map[string]struct{}), // Initialize UniqueIPs as an empty map
		Blacklist: blacklist,
	}
}

// Start begins to start listening on given IP and port
func (s *UDPEchoServer) Start(ip string, port int) {
	// Format the address using provided IP and port
	addr := fmt.Sprintf("%s:%d", ip, port)

	// Resolve UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Failed to resolve address:", err) // print error if unable to resolve address
		return
	}

	// Begin listening on the resolved UDP address
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to start server:", err) // print error if unable to start the server
		return
	}
	defer conn.Close() // Close the connection when function exits

	fmt.Printf("Listening on %s for UDP packets...\n", addr)

	buf := make([]byte, 1024) // Create a buffer to hold incoming data
	for {
		// Read from the UDP connection
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err) // print error if unable to read data
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

		// If data contains blacklisted keywords, do not process this packet.
		if blacklisted {
			continue
		}

		// Log the received data from the client
		logger.LogToFile(
			s.LogDir,
			"udp_all_traffic.log",
			fmt.Sprintf("Received from %s: %s\n", clientAddr, clientData),
		)

		// Check if the client's IP is unique and log it if it is
		if _, exists := s.UniqueIPs[clientAddr.IP.String()]; !exists {
			s.UniqueIPs[clientAddr.IP.String()] = struct{}{}

			// Logs the unique IP
			logger.LogToFile(s.LogDir, "unique_ips.log", clientAddr.IP.String()+"\n")
		}

		// Echo the data back to the client
		conn.WriteToUDP(buf[:n], clientAddr)
	}
}
