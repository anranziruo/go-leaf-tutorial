package main

import (
	"github.com/name5566/leaf"
	"leaf-tutorial/game"
	"leaf-tutorial/gate"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		f, _ := os.Create("cpu.pprof")
		defer f.Close()
		defer pprof.StopCPUProfile()
		time.Sleep(30 * time.Second)
		_ = pprof.StartCPUProfile(f)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f, _ := os.Create("heap.pprof")
		defer f.Close()
		time.Sleep(30 * time.Second)
		_ = pprof.WriteHeapProfile(f)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f, _ := os.Create("goroutine.pprof")
		defer f.Close()
		time.Sleep(30 * time.Second)
		_ = pprof.Lookup("goroutine").WriteTo(f, 0)
	}()

	leaf.Run(
		game.Module,
		gate.Module,
	)

	wg.Wait()
}
