package collect

import "strings"

func Join(args []interface{}, sep string) interface{} {
	l := make([]string, len(args))
	for index, v := range args {
		l[index] = v.(string)
	}
	return strings.Join(l, sep)
}
