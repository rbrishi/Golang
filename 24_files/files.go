// Package main is the entry point for this Go program that demonstrates various file operations
package main

// Import required packages:
// - bufio: Provides buffered I/O operations
// - fmt: Implements formatted I/O with functions analogous to C's printf and scanf
// - os: Provides platform-independent interface to operating system functionality
import (
	"bufio"
	"fmt"
	"os"
)

// main function serves as the entry point for the program
// It demonstrates various file operations including:
// - Reading file information
// - Reading file contents using different approaches
// - Creating and writing to files
// - Copying files
// - Directory operations
// - File deletion
func main() {
	// file, err := os.Open("ex.txt")
	// if err != nil {
	// 	// Handle the error
	// 	panic(err)
	// }

	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	// Handle the error
	// 	panic(err)
	// }
	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("Size in bytes:", fileInfo.Size())
	// fmt.Println("Last Modified:", fileInfo.ModTime())
	// fmt.Println("Is Directory:", fileInfo.IsDir())
	// fmt.Println("Mode:", fileInfo.Mode())
	// file.Close()




	// APPROACH 1: Basic file reading using a buffer
	// Open the file for reading
	// os.Open returns a file descriptor and an error (if any)
	file, err := os.Open("ex.txt")
	if err != nil {
		// If there's an error opening the file, panic and stop execution
		panic(err)
	}
	// defer ensures the file is closed after we're done with it
	// This is a good practice to prevent resource leaks
	defer file.Close()

	// Create a buffer to store the file contents
	// make([]byte, 100) creates a byte slice with capacity of 100 bytes
	buffer := make([]byte, 100)
	// Read up to 100 bytes from the file into the buffer
	bytesRead, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	// Print statistics and content
	fmt.Println("Bytes read:", bytesRead) // Print the number of bytes read
	fmt.Println("Content:", string(buffer[:bytesRead])) // Convert bytes to string and print

	// APPROACH 2: Reading and displaying file contents byte by byte
	// This approach demonstrates how to process each byte individually
	for i := 0; i < bytesRead; i++ {
		// Using Printf with %c format specifier to print each byte as a character
		// %d prints the byte position, %c prints the character representation
		fmt.Printf("Byte %d: %c\n", i, buffer[i])
		// Note: Using buffer[i] directly would print ASCII values instead of characters
	}

	// APPROACH 3: Using os.ReadFile for simple file reading
	// This method is convenient but has limitations:
	// - Reads entire file into memory at once
	// - Not suitable for large files (memory constraints)
	// - No control over reading in chunks
	// - Best used for small files where simplicity is preferred over memory efficiency
	data, err := os.ReadFile("ex.txt")	
	if err != nil {
		panic(err)
	}
	// Convert the entire byte slice to a string and print
	fmt.Println("Content read using ReadFile:", string(data))


	//Approach: 4
	// for big files


	// DIRECTORY OPERATIONS
	// Open the current directory (".")
	// Note: "../" would reference the parent directory
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	// Ensure directory handle is closed after we're done
	defer dir.Close()
	
	// Read directory contents
	// ReadDir(5) limits to reading 5 entries. Use 0 to read all entries
	// Returns a slice of DirEntry interfaces containing file/directory information
	info, err := dir.ReadDir(5)
	// Iterate through directory entries and print names
	for _, fi := range info {
		fmt.Println(fi.Name()) // Print each file/directory name
	}



	// FILE CREATION AND WRITING
	// Create a new file using os.Create
	// If file exists, it will be truncated. If it doesn't exist, it will be created
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	// Ensure file is closed after we're done with it
	defer newFile.Close()

	// Write strings directly to the file
	// WriteString is a convenient way to write string content
	newFile.WriteString("Hello, this is a new file created by Go!\n")
	newFile.WriteString("This file is created using os.Create method.\n")

	// WRITING BYTES TO FILE
	// Convert string to byte slice and write to file
	// This demonstrates how to write raw bytes instead of strings
	byte := []byte("This is some additional text added to the file.\n")
	newFile.Write(byte) // Write accepts a byte slice as argument




	// FILE COPYING USING BUFFERED I/O
	// This demonstrates how to copy files in a memory-efficient streaming fashion
	
	// Open source file for reading
	sourceFile, err := os.Open("ex.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	// Create destination file
	destFile, err := os.Create("copy_ex.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// Create buffered reader and writer
	// bufio provides buffering for better performance when reading/writing
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	// Copy content byte by byte
	// This loop demonstrates streaming file copy:
	// - Reads one byte at a time
	// - Checks for EOF (End Of File)
	// - Writes each byte to destination
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {	// If error is not EOF, it's an actual error
				panic(err)
			}
			break	// Break loop when EOF is reached
		}
		// Write the byte to destination file
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	// Flush ensures any buffered data is written to the underlying writer
	writer.Flush()
	fmt.Println("File copied successfully from ex.txt to copy_ex.txt")



	// FILE DELETION
	// os.Remove deletes a file or empty directory
	// Be careful with this operation as it's irreversible
	err = os.Remove("newfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
	// Note: For removing directories that aren't empty,
	// use os.RemoveAll() instead of os.Remove()
}