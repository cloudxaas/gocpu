package cxcputhread

import(
	"flag"
	"golang.org/x/sys/unix"
)
var (
	CPUThread *int16
)

func init() {
	CPUThread = int16(flag.Int("t", -1, "prefork child id"))
	flag.Parse()
}

func SetCPUAffinity(cpu int16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu))
	return unix.SchedSetaffinity(0, &newMask)
}
