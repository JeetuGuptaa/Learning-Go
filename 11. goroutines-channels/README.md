# Goroutines and Channels

This example demonstrates concurrent programming in Go using goroutines and channels.

## Concepts Covered

### Goroutines
- Lightweight threads managed by Go runtime
- Started with the `go` keyword
- Run concurrently with other goroutines

### Channels
- Typed conduits for communication between goroutines
- Created with `make(chan Type)`
- Send data: `ch <- value`
- Receive data: `value := <-ch`

### Buffered Channels
- Channels with a capacity: `make(chan Type, capacity)`
- Sends block only when buffer is full
- Receives block only when buffer is empty

### Channel Operations
- **Close**: `close(ch)` - signals no more values will be sent
- **Range**: `for v := range ch` - receives values until channel is closed
- **Select**: multiplexes multiple channel operations

## Examples in main.go

1. **Basic Goroutines**: Two functions running concurrently
2. **Channel Communication**: Sending and receiving data between goroutines
3. **Buffered Channels**: Using channels with capacity
4. **Worker Pool Pattern**: Multiple workers processing jobs concurrently
5. **Select Statement**: Handling multiple channel operations

## Running the Code

```bash
go run main.go
```

## Key Points

- Goroutines are cheap to create (thousands can run concurrently)
- Channels provide safe communication between goroutines
- Always close channels when done sending (sender's responsibility)
- Receiving from a closed channel returns the zero value
- Sending to a closed channel causes a panic
- Use `sync.WaitGroup` for more complex synchronization (not shown here)
- Select allows waiting on multiple channel operations

## Common Patterns

1. **Fan-out**: Multiple goroutines reading from same channel
2. **Fan-in**: Multiple goroutines writing to same channel
3. **Pipeline**: Chain of stages connected by channels
4. **Worker Pool**: Fixed number of workers processing jobs from a queue

## Additional Resources

- Go Blog: [Share Memory By Communicating](https://go.dev/blog/codelab-share)
- Effective Go: [Concurrency](https://go.dev/doc/effective_go#concurrency)
- Go Tour: [Goroutines](https://go.dev/tour/concurrency/1)
