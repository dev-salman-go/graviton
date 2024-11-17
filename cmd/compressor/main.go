package main

import (
	"graviton/internal/compressor"
	"log"
	"os"
	"sync"
)

func main() {
	//parsed command line args
	inputDir := "../../input"
	outputDir := "../../output"
	const maxWorkers uint8 = 4

	// validate directories
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		log.Fatalf("Input directory does not exist: %v", inputDir)
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		log.Fatalf("Output directory does not exist: %v", outputDir)
	}

	// initialize channels
	files := make(chan string, maxWorkers)
	results := make(chan error, maxWorkers)

	// create wait group
	var wg sync.WaitGroup

	// start workers

	for i := uint8(0); i < maxWorkers; i++ {
		wg.Add(1)
		go compressor.Worker(files, results, outputDir, &wg)
	}

	// Read files and send to channel
	go func() {
		compressor.ReadFiles(inputDir, files)
		close(files)
	}()

	// wait for all workers to finish
	wg.Wait()
	close(results)

	// Handle results
	for result := range results {
		if result != nil {
			log.Printf("Error: %v", result)
		}
	}
	log.Println("File compression completed.")

}
