package common

import "strings"

//CombinePath path 组合
func CombinePath(isRoot bool, paths ...string) string {
	str := strings.Builder{}
	if isRoot {
		str.WriteString("/")
	}
	for i, v := range paths {
		str.WriteString(v)
		if i != len(paths)-1 {
			str.WriteString("/")
		}
	}
	return str.String()
}
