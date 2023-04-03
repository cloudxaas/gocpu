package cxcputhread

import(
	"github.com/spf13/pflag"
	"golang.org/x/sys/unix"
)
var (
	//up to 65535 cpu cores supported
	CPUThread *uint16
)

func init() {
	CPUThread = flag.Uint16("t", -1, "prefork child id"))
	flag.Parse()
}

func SetCPUAffinity(cpu uint16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu))
	return unix.SchedSetaffinity(0, &newMask)
}
