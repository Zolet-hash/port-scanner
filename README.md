# **TCP & UDP Port Scanner**

A simple, concurrent port scanner written in **Go** that performs **TCP** and **UDP** port scanning. This project was built as a learning exercise to better understand how TCP and UDP work under the hood while exploring Go's concurrency model.

**Note:** This is an educational project designed to deepen my understanding of networking and concurrent programming in Go.

---

## **Features**

* TCP port scanning  
* UDP port scanning  
* Concurrent scanning using Go goroutines  
* Configurable target host  
* Configurable port range  
* Fast execution through worker-based concurrency  
* Simple command-line interface  
* Lightweight with no external dependencies

---

## **Why I Built This**

The main objective of this project was not simply to build a scanner, but to understand the networking concepts behind it.

While developing this project, I learned:

* How TCP connections are established using the three-way handshake.  
* Why a successful TCP connection can indicate an open port.  
* How UDP differs from TCP since it is connectionless.  
* Why UDP scanning is generally less reliable than TCP scanning.  
* How operating systems respond to closed versus open UDP ports.  
* How socket programming works in Go.  
* How network timeouts affect scan results.  
* How concurrency can significantly improve scanning performance.

---

## **Concepts Explored**

### **TCP Scanning**

The TCP scanner attempts to establish a TCP connection with each target port.

Typical results:

* Successful connection → Port is likely **open**  
* Connection refused → Port is **closed**  
* Timeout → Port may be **filtered** or unreachable

This mimics the behavior of a basic TCP Connect Scan.

---

### **UDP Scanning**

UDP does not establish a connection before sending data.

Instead, the scanner sends UDP packets and interprets the response:

* No response → Port may be open or filtered  
* ICMP Port Unreachable → Port is closed

Because UDP is connectionless, determining the exact state of a port is inherently more difficult than with TCP.

---

## **Concurrency**

One of the primary goals of this project was to practice Go's concurrency primitives.

The scanner uses:

* Goroutines  
* Channels  
* Worker pools  
* Synchronization primitives (where appropriate)

Instead of scanning ports sequentially, multiple workers scan ports concurrently, resulting in significantly faster execution.

Example workflow:
```shell
Ports  
   │  
   ▼  
Channel  
   │  
   ▼  
Worker 1 ── Scan Port  
Worker 2 ── Scan Port  
Worker 3 ── Scan Port  
...  
Worker N ── Scan Port
```
This approach allows efficient utilization of system resources while keeping the implementation simple and scalable.

---

## **Project Structure**
```shell
.
├── Dockerfile
├── go.mod
├── LICENSE
├── main.go
├── README.md
└── tcp
    ├── tcp.go
    └── udp.go
```
---

## **Installation**

Clone the repository:
```shell
git clone https://github.com/Zolet-hash/port-scanner.git
```
Move into the project directory:
```shell
cd port-scanner
```shell
Run the program:
```shell
go run .
```shell
Or build the executable:
```shell
go build
```
---

## **Usage**

Example:
```shell
go run . \[OPTIONS\] \[HOST\] \[PORT\]
```shell
UDP scan:
```shell
go run . -u 127.0.0.1
```shell
Example output:

Scanning scanme.nmap.org...
```shell
\[OPEN\] TCP 22  
\[OPEN\] TCP 80  
\[OPEN\] TCP 443
```shell
Scan completed in 1.83s

---

## **Technologies Used**

* Go  
* net package  
* Goroutines  
* Channels  
* WaitGroups  
* Socket programming

---

## **Limitations**

This project intentionally focuses on the fundamentals.

It does **not** include advanced scanning techniques such as:

* SYN scans  
* FIN scans  
* ACK scans  
* Service/version detection  
* OS fingerprinting  
* Banner grabbing  
* Stealth scanning  
* IPv6 support

Those features are outside the scope of this educational project.

---

## **What I Learned**

Building this project improved my understanding of:

* TCP/IP networking  
* TCP and UDP communication  
* Socket programming  
* Timeouts and network errors  
* Concurrent programming in Go  
* Goroutines and channels  
* Designing worker pools  
* Writing maintainable Go code

---

## **Future Improvements**

Possible enhancements include:

* CIDR/network scanning  
* IPv6 support  
* Banner grabbing  
* Service detection  
* Progress indicator  
* Scan result export (JSON/CSV)  
* Adjustable worker count  
* Better timeout tuning  
* Context cancellation support

---

## **Disclaimer**

This tool is intended for educational purposes and for scanning systems that you own or are explicitly authorized to test. Always obtain permission before scanning networks or hosts that you do not control.

