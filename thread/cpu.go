package cxcputhread

import(
	"github.com/spf13/pflag"
	"golang.org/x/sys/unix"
)
var (
	//Support up to 32000 cpu cores 
	//because of the way other functions wraps it, 
        //negative value denotes no change in other func
	CPUThread *int16
)

func init() {
	CPUThread = flag.Int16("t", -1, "prefork child id")
	flag.Parse()
}

func SetCPUAffinity(cpu Int16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu))
	return unix.SchedSetaffinity(0, &newMask)
}
