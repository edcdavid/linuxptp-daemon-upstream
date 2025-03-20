package leakypipe

import (
	"github.com/golang/glog"
	"io"
	"sync"
)

type LeakyPipe struct {
	mu        sync.Mutex
	buffer    chan []byte
	closed    bool
	chunkSize int
}

func New(maxChunks, aChunkSize int) *LeakyPipe {
	return &LeakyPipe{
		chunkSize: aChunkSize,
		// Use a channel to simulate a queue of size maxLines
		buffer: make(chan []byte, maxChunks),
	}
}

func (p *LeakyPipe) Write(data []byte) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return 0, io.ErrClosedPipe
	}

	chunks := splitIntoChunks(data, p.chunkSize)

	for _, chunk := range chunks {
		select {
		// Make a copy of the data and send it to the channel
		// If the channel is full, instead of blocking it goes to default
		// and drops the data
		case p.buffer <- append([]byte{}, chunk...):
		default:
			// Dropping line
			glog.Warningf("Pipe buffer full, dropping %d bytes chunk!", p.chunkSize)
		}
	}

	return len(data), nil
}

// Pipe Read
func (p *LeakyPipe) Read(data []byte) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	select {
	case read := <-p.buffer:
		data = data[:0]
		data = append(data, read...)
		return p.chunkSize, nil
	default:
		return 0, nil
	}
}

// Pipe Close
func (p *LeakyPipe) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.closed {
		close(p.buffer)
		p.closed = true
	}
}

func splitIntoChunks(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, append([]byte{}, data[i:end]...))
	}
	return chunks
}
