package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
)

func getCerts(ip string, port int, rf *os.File) {
	conn, err := tls.Dial("tcp", ip+":"+strconv.Itoa(port), conf)
	if err != nil {
		fmt.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		if port == 3389 {
			fmt.Printf("Issuer Name: %s\n", cert.Issuer)
			rf.WriteString("Source: " + ip + ":" + strconv.Itoa(port) + " " + "Issuer: " + cert.Issuer.CommonName)
			continue
		}
		if cert.DNSNames != nil {
			fmt.Printf("Issuer Name: %s\n", cert.DNSNames)
			rf.WriteString("Source: " + ip + ":" + strconv.Itoa(port) + " " + "DNSname: " + fmt.Sprintln(cert.DNSNames))
		}
	}
}
