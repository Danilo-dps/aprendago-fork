package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {

	// A função embutida `new` aloca uma nova variável inicializada e retorna um ponteiro para ela.
	p := new(SyncedBuffer) // type *SyncedBuffer
	r := SyncedBuffer{}    // type  SyncedBuffer
	var v SyncedBuffer     // type  SyncedBuffer

	fmt.Println(p)
	fmt.Println(r)
	fmt.Println(v)
}
