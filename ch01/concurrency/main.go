package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

type Result struct {
	Device string `json:"device"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

const (
	maxRetries       = 3
	timeoutPerDevice = 2 * time.Second
	maxConcurrent    = 5
	logFilePath      = "connect.log"
	jsonOutputPath   = "results.json"
)

func init() {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// getConfig simulates connecting to a device with random delay and possible failure
func getConfig(ctx context.Context, device string) (string, error) {
	select {
	case <-time.After(time.Millisecond * time.Duration(rand.Intn(1000))):
		if rand.Intn(10) < 2 { // 20% 실패 확률
			return "", fmt.Errorf("failed to connect to device %q", device)
		}
		return fmt.Sprintf("Connected to device %q", device), nil
	case <-ctx.Done():
		return "", fmt.Errorf("timeout while connecting to device %q", device)
	}
}

// connectDevices handles concurrent device connection with limited concurrency
func connectDevices(devices []string) []Result {
	var (
		wg      sync.WaitGroup
		results = make([]Result, 0, len(devices))
		mu      sync.Mutex
		sem     = make(chan struct{}, maxConcurrent)
	)

	for _, d := range devices {
		d := d // avoid closure capture issue
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			var (
				output string
				err    error
			)

			for attempt := 1; attempt <= maxRetries; attempt++ {
				ctx, cancel := context.WithTimeout(context.Background(), timeoutPerDevice)
				defer cancel()

				log.Printf("Connecting to %s (attempt %d)", d, attempt)
				output, err = getConfig(ctx, d)
				if err == nil {
					break
				}
				log.Printf("Error on %s (attempt %d): %v", d, attempt, err)
			}

			result := Result{Device: d}
			if err != nil {
				result.Error = err.Error()
			} else {
				result.Output = output
			}

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}

// saveResults writes the results as a JSON file
func saveResults(results []Result) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonOutputPath, data, 0644)
}

// loadDevices returns a list of devices (hardcoded or from file)
func loadDevices() []string {
	// 필요한 경우 파일에서 읽을 수 있도록 확장 가능
	return []string{"leaf01", "leaf02", "spine01", "leaf03", "spine02"}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	devices := loadDevices()

	fmt.Println("Starting device connection...")
	results := connectDevices(devices)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Device < results[j].Device
	})

	fmt.Println("Connection results:")
	for _, res := range results {
		if res.Error != "" {
			fmt.Printf("[ERROR] %s: %s\n", res.Device, res.Error)
		} else {
			fmt.Printf("[OK] %s: %s\n", res.Device, res.Output)
		}
	}

	if err := saveResults(results); err != nil {
		log.Printf("Failed to save results: %v", err)
	} else {
		log.Println("Results saved to", jsonOutputPath)
	}
}
