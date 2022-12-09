package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getStatus(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Failed to initialize new HTTP request: %v", err)
	}

	if resp.StatusCode > 299 {
		return resp, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
	}

	return resp, nil
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
		log.Printf("scanner encountered an error: %v\n", err)
	}
}

func main() {
	flag.Parse()
	services := flag.Args()

	resp, err := getStatus("https://status.plaintext.sh/t")
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return
	}

	parseStatus(services, resp)
}
