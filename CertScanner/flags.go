package main

import (
	"flag"
)

type Flag struct {
	maxThreads int
	resultFile string
	subnetFile string
	ASN        string
}

//Bind and Prase Flags
// Bind the flag

func init_flags() Flag {

	v := Flag{}
	flag.StringVar(&v.subnetFile, "in", "subnets.txt", "A file contining subnets")
	flag.StringVar(&v.resultFile, "out", "results.txt", "Output file")
	flag.StringVar(&v.ASN, "asn", "", "ASN Number to scan")
	flag.IntVar(&v.maxThreads, "m", 25, "Max Threads")

	// Parse the flag
	flag.Parse()
	return v

}
