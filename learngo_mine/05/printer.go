package printer

import (
	"fmt"
	"runtime"
)

//Version prints runtime golang version
func Version() {
	fmt.Println(runtime.Version())
}
