# Graviton - A File Compressor

A simple and efficient file compression program written in Go. The program supports multi-threaded compression using goroutines, channels, and wait groups. By default, it compresses files using Gzip and saves the compressed files with a `.gz` extension.

## Features

- **Multi-threaded Compression**: Compress multiple files concurrently using a worker pool.
- **Customizable Directories**: Specify input and output directories via command-line arguments.
- **Gzip Compression**: Compress files into `.gz` format.
- **Scalable**: Configure the number of workers for optimal performance.

---

## Usage

### 1. **Installation**

Ensure you have Go installed on your system. Then clone the repository and initialize dependencies:

```bash
git clone https://github.com/yourusername/file-compressor.git
cd file-compressor
go mod tidy
```

### 2. **Prepare Input and Output Directories**

Create directories for input files (to be compressed) and output files (compressed results):

```bash
mkdir input output
echo "This is a test file." > input/test1.txt
echo "Another file to compress." > input/test2.txt
```

### 3. **Run the Program**

You can run the program with the default settings:

```bash
go run cmd/compressor/main.go
```

#### **Custom Paths**

Provide custom input and output directories using flags:

```bash
go run cmd/compressor/main.go -input=/path/to/input -output=/path/to/output -workers=8
```

| Flag        | Default Value   | Description                          |
|-------------|-----------------|--------------------------------------|
| `-input`    | `./input`       | Path to the directory of files to compress. |
| `-output`   | `./output`      | Path to save compressed files.       |
| `-workers`  | `4`             | Number of concurrent workers.        |

---

## Example Output

After running the program, you’ll find the compressed files in the output directory with a `.gz` extension:

```bash
ls output/
test1.txt.gz  test2.txt.gz
```

---

## Supported Compression Format

Currently, the program supports **Gzip** compression. Future versions may include support for additional formats like Brotli, Zstd, or LZ4.

---

## How It Works

1. **File Reading**: The program reads all files from the specified input directory.
2. **Worker Pool**: A configurable number of workers process the files concurrently.
3. **Compression**: Each file is compressed using Gzip and saved with a `.gz` extension in the output directory.

---

## Future Enhancements

- Support for additional compression formats (e.g., Brotli, Zstd).
- Detailed logs and metrics for compression performance.
- CLI improvements for better error handling and usage.

---

## Contributing

Feel free to fork this repository and submit pull requests for improvements or additional features. Feedback is always welcome!

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Author

**Your Name**  
- GitHub: [@yourusername](https://github.com/yourusername)  
- Email: your.email@example.com  

---

### Let me know if you'd like me to modify anything or generate additional sections!