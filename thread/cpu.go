package cxcputhread

import(
	"golang.org/x/sys/unix"
	"github.com/cloudxaas/gohash/fnv1a32"
	flag "github.com/spf13/pflag"
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

func SetCPUAffinity(cpu int16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu))
	return unix.SchedSetaffinity(0, &newMask)
}

//hash based on fnv1a32, getting hashed cpu id
func CPUHash(k []byte) int16 {
	return int16(fnv1a32.Hash(k) % uint32(runtime.NumCPU()))
}
