package job

import "fmt"

func IntsToString(xs []int) string {
	id := ""
	for _, x := range xs {
		id += fmt.Sprintf("%v", x)
	}
	return id
}
