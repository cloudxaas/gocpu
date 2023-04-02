package cxcputhread

var (
	CPUThread *int
)

func init() {
	CPUThread = flag.Int("t", -1, "prefork child id")
	flag.Parse()
}
