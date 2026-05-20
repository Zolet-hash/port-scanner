## PortScanner — Multi-Technique Port Scanner in Golang

A lightweight yet powerful port scanner written in Go, implementing multiple scanning techniques inspired by concepts from Nmap while making use of Go's networking power and goroutines for high-speed concurrent scanning.

Features
TCP Connect Scanning
SYN Scanning
FIN Scanning
NULL Scanning
UDP Scanning
Service Banner Grabbing
Concurrent Scanning using Goroutines
Configurable Timeout & Thread Count
CIDR / Host Range Support
Fast and Lightweight
Cross-platform Support
Clean CLI Interface
Why This Project?

This project was built to deepen understanding of:

Network protocols
TCP/IP internals
Socket programming
Concurrent programming in Go
Low-level reconnaissance techniques
Cybersecurity tooling development

It demonstrates practical knowledge of:

Goroutines
Channels
Raw sockets
Packet crafting
CLI application development
Error handling and concurrency management in Go
Tech Stack
Golang
Go Net Package
Raw Socket Programming
Goroutines & Channels
Scanning Techniques Implemented
TCP Connect Scan

Performs a full TCP handshake to determine whether a port is open.

Best for:

Reliability
User-space scanning
Non-privileged environments
SYN Scan (Half-Open Scan)

Sends SYN packets and analyzes responses without completing the full TCP connection.

Advantages:

Faster scanning
Stealthier than full connect scans
Commonly used in professional reconnaissance
FIN Scan

Sends FIN packets to detect closed/open ports based on RFC behavior.

Useful for:

Firewall evasion
Advanced reconnaissance
NULL Scan

Sends packets with no TCP flags set.

Can help:

Identify firewall rules
Detect filtered ports
UDP Scan

Attempts communication over UDP ports to identify open services.

Useful for discovering:

DNS
SNMP
DHCP
NTP
Other UDP-based services
Architecture
            +------------------+
            |     CLI Input     |
            +------------------+
                      |
                      v
            +------------------+
            |  Target Parser    |
            +------------------+
                      |
                      v
            +------------------+
            | Scan Dispatcher   |
            +------------------+
               /    |     \
              /     |      \
             v      v       v
        TCP Scan  SYN Scan  UDP Scan
             \      |       /
              \     |      /
               v    v     v
            +------------------+
            | Result Aggregator |
            +------------------+
                      |
                      v
            +------------------+
            | Output Formatter  |
            +------------------+
