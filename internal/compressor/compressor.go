package compressor

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressFile(inputPath, outputDir string) error {

	// open source file
	sourceFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", inputPath, err)
	}
	defer sourceFile.Close()

	// Create the destination file
	outputPath := filepath.Join(outputDir, filepath.Base(inputPath)+".gz")
	destFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer destFile.Close()

	// create a gzip writer
	gzipWriter := gzip.NewWriter(destFile)
	defer gzipWriter.Close()

	// Copy the contents of the source file to the zip writer
	_, err = io.Copy(gzipWriter, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to compress file %s: %w", inputPath, err)
	}

	// Successfully compressed, return no error
	return nil
}

// ReadFiles reads all files from the input directory and sends them to the files channel.
func ReadFiles(inputDir string, files chan<- string) {
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore directories, only allow files
		if !info.IsDir() {
			files <- path
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading files: %v\n", err)
	}

}
