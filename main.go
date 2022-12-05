package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getStatus(url string) *http.Response {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Failed to initialize new HTTP request: %v", err)
	}

	if resp.StatusCode > 299 {
		resp.Body.Close()
		log.Fatalf("Response failed with status code: %d\n", resp.StatusCode)
	}

	return resp
}

func parseStatus(services []string, resp *http.Response) {
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		status := scanner.Text()

		for _, serviceName := range services {
			if !strings.Contains(status, serviceName) {
				continue
			}
			fmt.Println(status)
		}

		if len(services) == 0 {
			fmt.Println(status)
		}
	}

	if err := scanner.Err(); err != nil {
		resp.Body.Close()
		log.Fatalf("Scanner encountered an error: %v", err)
	}
}

func main() {
	flag.Parse()
	services := flag.Args()

	resp := getStatus("https://status.plaintext.sh/t")
	defer resp.Body.Close()

	parseStatus(services, resp)
}
