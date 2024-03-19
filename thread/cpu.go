package cxcputhread

import (
	"runtime"
	"github.com/zeebo/xxh3" // Updated import for xxh3
	"golang.org/x/sys/unix"
	"github.com/spf13/pflag"
)

var (
	NumCPU uint16
	CPUThread uint16 // Support up to 65534 CPU cores, "0" denotes master or default
)

// This function should be called from the main package or higher-level logic to define flags.

func init() {
	NumCPU = uint16(runtime.NumCPU())
	  pflag.Uint16VarP(&CPUThread, "CPUThread", "t", 0, "prefork child id")

	  //Parse it in main
//	  pflag.Parse() //parse at cxcputhread as it's very important, preparsing with another "flag" module before calling cxcputhread
}

//cpu core 0 == 1
//cpu core 1 == 2
func SetCPUAffinity(cpu uint16) error {
	var newMask unix.CPUSet
	newMask.Zero()
	newMask.Set(int(cpu) - 1)
	return unix.SchedSetaffinity(0, &newMask)
}

// Updated function to use xxh3 for hashing with AVX2 optimization when available.
func CPUHash(k []byte) uint16 {
	return uint16(xxh3.Hash(k) & uint64(runtime.NumCPU()-1))
}

func IsCurrentCPUID(id uint16) uint8 {
	if id == CPUThread-1 {
		return 1
	}
	return 0
}
