package main

import (
	"fmt"
	"net/http"
	"strings"
)

func generateName(base string, suffix string) string {
	parts := []string{base, suffix}
	return strings.Join(parts, "-")
}

func processDevice(getName func(string, string) string, ip string) {
	base := "device"
	name := getName(base, ip)
	fmt.Println(name)
}

func makeCall(url string) (*http.Response, error) {
	resp, err := http.Get("example.com")
	if err != nil {
		return nil, fmt.Errorf("error in makeCall: %w", err)
	}

	return resp, nil
}

func main() {
	s := generateName("device", "01")

	// prints "device-01"
	fmt.Println(s)

	// prints "device-192.0.2.1"
	processDevice(generateName, "192.0.2.1")

}