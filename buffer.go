package tidy

import (
	"bytes"
	"sync"
)

var (
	buffers *sync.Pool
)

func init() {
	buffers = &sync.Pool{
		New: func() interface{} {
			return &FreeableBuffer{}
		},
	}
}

func NewBuffer() *FreeableBuffer {
	return buffers.Get().(*FreeableBuffer)
}

type FreeableBuffer struct {
	bytes.Buffer
}

var digits = []byte("0123456789")

func (this *FreeableBuffer) WriteTwoDigits(value int) {
	this.Write([]byte{
		digits[value/10],
		digits[value%10],
	})
}

func (this *FreeableBuffer) Free() {
	this.Truncate(0)
	buffers.Put(this)
}
