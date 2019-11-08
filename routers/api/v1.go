package v1

type Router struct {
	path    string
	handler func()
}

var routerGroup []Router{
	path: "/test1"
	handler: test1()
}{

}