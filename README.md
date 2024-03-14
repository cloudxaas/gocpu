# gocpu
Golang CPU

```
package main

import (
	"github.com/cloudxaas/gocpu"
	"github.com/spf13/pflag"
)

func main() {
	// Define flags from all packages
	cxcputhread.DefineFlags()

	// Parse flags once after all definitions
	pflag.Parse()

	// Your application logic here
}
```
