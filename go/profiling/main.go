package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	var cpuprofile, memprofile string
	flag.StringVar(&cpuprofile, "cpuprofile", "", "cpu_profile")
	flag.StringVar(&memprofile, "memprofile", "", "mem_profile")
	flag.Parse()

	if memprofile != "" {
		runtime.MemProfileRate = 1
	}

	if memprofile != "" {
		f, err := os.Create("./mem.prof")
		defer f.Close()
		if err != nil {
			panic(err)
		}
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(err)
		}
	}
}
