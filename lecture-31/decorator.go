package main

import (
	"io"
	"log"
	"strings"
)

type CountWriter struct {
	Count int
}

func (c *CountWriter) Write(data []byte) (int, error) {
	c.Count += len(data)
	return len(data), nil
}

func (c *CountWriter) GetReadCount() int {
	return c.Count
}

type CountDecorator interface {
	io.Reader

	GetReadCount() int
}

type CountDecoratorImpl struct {
	io.Reader

	*CountWriter
}

func NewDecorator(r io.Reader) CountDecorator {
	wr := &CountWriter{}
	result := &CountDecoratorImpl{
		Reader:      io.TeeReader(r, wr),
		CountWriter: wr,
	}
	return result
}

func main() {
	reader := strings.NewReader("Hello world!")

	decorator := NewDecorator(reader)

	buf := make([]byte, 5)

	decorator.Read(buf)

	log.Printf("Buf: %s", buf)

	log.Printf("Read countt: %d", decorator.GetReadCount())
}
