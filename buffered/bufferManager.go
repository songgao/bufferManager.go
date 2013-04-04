package buffered

import (
	"sync"
)

type Token struct {
	Data  Data
	owner chan<- *Token
}

func (t *Token) Return() {
	t.owner <- t
}

func (t *Token) WaitAndReturn(wg *sync.WaitGroup) {
	wg.Wait()
	t.owner <- t
}

type BufferManager struct {
	buffer chan *Token
}

func NewBufferManager(size int) *BufferManager {
	buffer := make(chan *Token, size)
	ret := &BufferManager{buffer}
	for i := 0; i < size; i++ {
		buffer <- &Token{Data: Data{}, owner: buffer}
	}
	return ret
}

func (b *BufferManager) GetToken() *Token {
	return <-b.buffer
}
