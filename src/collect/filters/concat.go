package collect

import "strings"

func Concat(args ...string) interface{} {
	return strings.Join(args, "")
}
