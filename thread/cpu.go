package cxcputhread

import(
	"flag"
	"golang.org/x/sys/unix"
)
var (
	CPUThread *uint16
)

func init() {
	CPUThread = uint16(flag.Int("t", -1, "prefork child id"))
	flag.Parse()
}

func SetCPUAffinity(cpu uint16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu))
	return unix.SchedSetaffinity(0, &newMask)
}
