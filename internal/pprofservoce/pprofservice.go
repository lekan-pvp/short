package pprofservoce

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func PprofService() {
	fmem, err := os.Create("profiles/result.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer fmem.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(fmem); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
