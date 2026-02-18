# File I/O in Go

This lesson covers reading and writing files to disk in Go.

## Topics Covered

### 1. Simple File Operations
- `os.WriteFile()` - Write data to a file in one operation
- `os.ReadFile()` - Read entire file contents into memory
- Quick and convenient for small files

### 2. File Handles
- `os.Create()` - Create a new file (truncates if exists)
- `os.Open()` - Open file for reading
- `os.OpenFile()` - Open with specific flags and permissions
- Always use `defer file.Close()` to ensure proper cleanup

### 3. Buffered I/O
- `bufio.NewWriter()` - Buffered writing for better performance
- `bufio.NewReader()` - Buffered reading
- `bufio.Scanner` - Easy line-by-line reading
- Remember to call `Flush()` when using buffered writers

### 4. File Flags
- `os.O_RDONLY` - Read only
- `os.O_WRONLY` - Write only
- `os.O_APPEND` - Append to file
- `os.O_CREATE` - Create if doesn't exist
- `os.O_TRUNC` - Truncate file when opening

### 5. File Permissions
- Unix-style permissions (e.g., `0644`, `0755`)
- `0644` - Owner can read/write, others can read
- `0755` - Owner can read/write/execute, others can read/execute

### 6. Common Patterns
- Check if file exists using `os.Stat()`
- Copy files using `io.Copy()`
- Read line by line using `bufio.Scanner`
- Append data using `os.O_APPEND` flag

## Running the Code

```bash
go run main.go
```

This will create several test files in the current directory:
- `output.txt` - Simple write/read example
- `buffered.txt` - Buffered I/O example
- `output_copy.txt` - File copy example

## Key Takeaways

1. **Simple operations** (`os.ReadFile`/`os.WriteFile`) are great for small files
2. **Buffered I/O** improves performance for multiple reads/writes
3. **Always close files** using `defer` to prevent resource leaks
4. **Handle errors** properly - file operations can fail
5. **Use Scanner** for line-by-line reading - it's idiomatic and efficient

## Error Handling

File operations can fail for many reasons:
- File doesn't exist
- Insufficient permissions
- Disk full
- File is locked by another process

Always check and handle errors appropriately in production code.
