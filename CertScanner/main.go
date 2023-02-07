package main

import (
	//"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var tlsPorts = []int{443, 465, 995, 993, 3389, 8443}

var activeThreads = 0
var subnets []string
var conf = &tls.Config{
	InsecureSkipVerify: true,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	v := init_flags()
	var rf, _ = os.Create(v.resultFile)
	defer rf.Close()

	doneChannel := make(chan bool)

	if v.subnetFile != " " {
		subnets = subnetsFromFile(v.subnetFile)
	}

	if v.ASN != "" {
		subnets = getAsnRange(v.ASN)
	}

	for _, s := range subnets {
		ips, _ := Hosts(s)
		for _, ip := range ips {
			for _, port := range tlsPorts {
				if activeThreads >= v.maxThreads {
					<-doneChannel // Wait
					activeThreads--
				}
				go testTcpConnection(ip, port, doneChannel, rf)
				activeThreads++
			}
		}
		// Wait for all threads to finish
		for activeThreads > 0 {
			<-doneChannel
			activeThreads--
		}
	}
}

func testTcpConnection(ip string, port int, doneChannel chan bool, rf *os.File) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*2)
	if err == nil {
		fmt.Printf("Host: %s Port %d: Open\n", ip, port)
		getCerts(ip, port, rf)
	}
	doneChannel <- true
}
