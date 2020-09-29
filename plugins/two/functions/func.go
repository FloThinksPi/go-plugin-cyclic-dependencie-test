package functions

import (
	"fmt"
	one "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/one/functions"
	"time"
)

func TestFunc()  {
	fmt.Print("one: testFunc called")
	time.Sleep(1)
	one.TestFunc()
}