package helper

import "strings"

//isContain search element exists in sets
//@param element element
//@param sets elements sets
//@param i ignore case
//@return bool
func IsContain(element string, sets []string, i bool) bool {
	maps := slice2MapWithContain(sets, i)
	if i {
		element = strings.ToLower(element)
	}
	return maps[element]
}
