package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"regexp"
)

func getAsnRange(asn string) []string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://ipinfo.io/"+asn, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0")
	r, _ := client.Do(req)
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	subnets := findSubNets(string(b))
	return subnets

}

func findSubNets(input string) []string {
	regexPattern := "[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\/[0-9][0-9]"

	regEx := regexp.MustCompile(regexPattern)
	matches := regEx.FindAllString(input, -1)

	uniqueMap := map[string]bool{}
	subnets := []string{}

	for _, s := range matches {
		if _, value := uniqueMap[s]; !value {
			uniqueMap[s] = true
			subnets = append(subnets, s)
		}
	}
	return subnets
}
func subnetsFromFile(file string) []string {
	var subnets []string
	subfile, _ := os.Open(file)
	defer subfile.Close()
	scanner := bufio.NewScanner(subfile)
	for scanner.Scan() {
		subnets = append(subnets, scanner.Text())
	}

	return subnets
}
