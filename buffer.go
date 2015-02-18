package logging

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

func (this *FreeableBuffer) Free() {
	this.Truncate(0)
	buffers.Put(this)
}
