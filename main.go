package main

import (
	"./buffered"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

const N = 500000
const M = 16

func main() {
	list_T := make([]*buffered.Token, M)
	list_D := make([]*buffered.Data, M)
	t := time.Now()
	f, _ := os.Create("with_bufferManager.prof")
	pprof.StartCPUProfile(f)
	m := buffered.NewBufferManager(M * 2)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			t := m.GetToken()
			t.Data.Str = "haha"
			list_T[j] = t
		}
		for j := 0; j < M; j++ {
			list_T[j].Return()
		}
	}
	fmt.Printf("With bufferManager:                    %v\n", time.Since(t))
	pprof.StopCPUProfile()

	t = time.Now()
	f, _ = os.Create("without_bufferManager.prof")
	pprof.StartCPUProfile(f)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			d := new(buffered.Data)
			d.Str = "haha"
			list_D[j] = d
		}
	}
	fmt.Printf("Without buffermanager (Relying on GC): %v\n", time.Since(t))
	pprof.StopCPUProfile()

}
