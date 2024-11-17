package compressor

import (
	"log"
	"sync"
)

// Worker is a gor that reads from files chan, compresses files, sends the results to result chan
func Worker(files <-chan string, results chan<- error, outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range files {
		log.Printf("Worker started compressing: %s", filePath)

		// Compress the file
		err := CompressFile(filePath, outputDir)
		if err != nil {
			log.Printf("Error compressing file %s: %v", filePath, err)
		}

		// Send the result (nil for success, error for failure) (error is nil if compression succeeds)
		results <- err
	}
}
