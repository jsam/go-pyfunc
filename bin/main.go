package main

import (
	"fmt"

	gopyfunc "github.com/jsam/go-pyfunc/pkg"
)

func main() {
	gopyfunc.PyInit()
	s, e := gopyfunc.Py_GetPath()
	fmt.Println(s, e)
	gopyfunc.Py_SetPath(fmt.Sprintf("%s:%s", "/Users/sam/justsam.io/go-pyfunc/", s))
	s, e = gopyfunc.Py_GetPath()
	fmt.Println(s, e)

	ptr, err := gopyfunc.FnPtr("examples", "hey")
	fmt.Println(ptr, err)
	status, response, err := gopyfunc.Request(ptr, "/path", map[string]string{"key": "value"})
	fmt.Println(status, response, err)
}
