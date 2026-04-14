package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 1. Basic file operations with os package
func basicFileOperations() {
	fmt.Println("=== Basic File Operations ===")

	// Create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write to file
	content := "Hello, Go file I/O!\nThis is a test file.\n"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Println("File created and written successfully")
}

// 2. Reading entire file at once
func readEntireFile() {
	fmt.Println("\n=== Reading Entire File ===")

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s", string(data))
}

// 3. Reading file line by line with bufio.Scanner
func readFileLineByLine() {
	fmt.Println("\n=== Reading File Line by Line ===")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

// 4. Writing to file with bufio.Writer
func writeWithBufio() {
	fmt.Println("\n=== Writing with bufio.Writer ===")

	file, err := os.Create("buffered_write.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	lines := []string{
		"Line 1: Buffered writing is efficient",
		"Line 2: It reduces system calls",
		"Line 3: Great for large amounts of data",
	}

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing line: %v\n", err)
			return
		}
	}

	// Important: Flush the buffer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing buffer: %v\n", err)
		return
	}

	fmt.Println("Buffered writing completed")
}

// 5. Copying files with io.Copy
func copyFile() {
	fmt.Println("\n=== Copying Files with io.Copy ===")

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example_copy.txt")
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}

	fmt.Printf("Copied %d bytes successfully\n", bytesCopied)
}

// 6. Working with directories
func directoryOperations() {
	fmt.Println("\n=== Directory Operations ===")

	// Create directory
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	}

	// Create a file in the directory
	file, err := os.Create("test_dir/nested_file.txt")
	if err != nil {
		fmt.Printf("Error creating nested file: %v\n", err)
		return
	}
	file.WriteString("This is in a subdirectory")
	file.Close()

	// List directory contents
	files, err := os.ReadDir("test_dir")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, file := range files {
		fmt.Printf("  %s (dir: %v)\n", file.Name(), file.IsDir())
	}

	// Get file info
	info, err := os.Stat("test_dir/nested_file.txt")
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("Is directory: %v\n", info.IsDir())
	fmt.Printf("Permissions: %v\n", info.Mode())
}

// 7. File paths and path operations
func pathOperations() {
	fmt.Println("\n=== Path Operations ===")

	// Join paths (cross-platform)
	fullPath := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("Joined path: %s\n", fullPath)

	// Get absolute path
	absPath, err := filepath.Abs("example.txt")
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
	} else {
		fmt.Printf("Absolute path: %s\n", absPath)
	}

	// Get directory and filename
	dir := filepath.Dir(absPath)
	filename := filepath.Base(absPath)
	fmt.Printf("Directory: %s\n", dir)
	fmt.Printf("Filename: %s\n", filename)

	// Get file extension
	ext := filepath.Ext(filename)
	fmt.Printf("Extension: %s\n", ext)
}

// 8. Reading from stdin with bufio
func readFromStdin() {
	fmt.Println("\n=== Reading from Standard Input ===")
	fmt.Println("Enter some text (press Ctrl+D to finish):")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	fmt.Printf("You entered %d lines:\n", len(lines))
	for i, line := range lines {
		fmt.Printf("%d: %s\n", i+1, line)
	}
}

// 9. Writing to stdout and stderr
func writeToStdoutStderr() {
	fmt.Println("\n=== Writing to stdout and stderr ===")

	// Write to stdout
	fmt.Fprintln(os.Stdout, "This goes to stdout")

	// Write to stderr
	fmt.Fprintln(os.Stderr, "This goes to stderr")

	// Using io.Writer interface
	writer := os.Stdout
	writer.Write([]byte("Direct write to stdout\n"))
}

// 10. Error handling and cleanup with defer
func safeFileOperations() {
	fmt.Println("\n=== Safe File Operations with Defer ===")

	// Function to safely write to file
	writeToFile := func(filename, content string) error {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close() // Always close the file

		_, err = file.WriteString(content)
		if err != nil {
			return err
		}

		return nil
	}

	err := writeToFile("safe_write.txt", "This was written safely with defer cleanup")
	if err != nil {
		fmt.Printf("Error in safe write: %v\n", err)
	} else {
		fmt.Println("Safe write completed")
	}
}

// 11. Working with temporary files
func temporaryFiles() {
	fmt.Println("\n=== Temporary Files ===")

	// Create temporary file
	tempFile, err := os.CreateTemp("", "example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up temp file
	defer tempFile.Close()

	// Write to temp file
	tempFile.WriteString("This is a temporary file\n")
	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	// Read it back
	tempFile.Seek(0, 0) // Reset to beginning
	content, err := io.ReadAll(tempFile)
	if err != nil {
		fmt.Printf("Error reading temp file: %v\n", err)
		return
	}

	fmt.Printf("Temp file content: %s", string(content))
}

func main() {
	fmt.Println("=== File I/O in Go (os, io, bufio) ===\n")

	// Run all examples
	basicFileOperations()
	readEntireFile()
	readFileLineByLine()
	writeWithBufio()
	copyFile()
	directoryOperations()
	pathOperations()
	writeToStdoutStderr()
	safeFileOperations()
	temporaryFiles()

	// Note: readFromStdin() is commented out to avoid blocking
	// Uncomment to test interactive input
	readFromStdin()

	cleanup()
}

// cleanup removes all test files created during examples
func cleanup() {
	fmt.Println("\n=== Cleanup ===")
	// Clean up created files
	filesToClean := []string{
		"example.txt",
		"buffered_write.txt",
		"example_copy.txt",
		"safe_write.txt",
	}

	for _, file := range filesToClean {
		if err := os.Remove(file); err != nil {
			fmt.Printf("Warning: Could not remove %s: %v\n", file, err)
		}
	}

	// Remove directory
	if err := os.RemoveAll("test_dir"); err != nil {
		fmt.Printf("Warning: Could not remove test_dir: %v\n", err)
	}

	fmt.Println("File I/O examples completed!")
	fmt.Println("Note: All test files have been cleaned up. To see them, comment out the cleanup() call.")
}