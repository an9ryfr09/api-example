package helper

//Contain check a value inside sets
func Contain(v interface{}, sets map[interface{}]bool) bool {
	return sets[v]
}
