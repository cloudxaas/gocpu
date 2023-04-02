package cxcputhread

import(
	"flag"
	"syscall"
)
var (
	CPUThread *int
)

func init() {
	CPUThread = flag.Int("t", -1, "prefork child id")
	flag.Parse()
}

func SetCPUAffinity(cpu int) error {
	var newMask unix.CPUSet
	newMask.Set(cpu)
	return unix.SchedSetaffinity(0, &newMask)
}
