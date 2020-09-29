package functions

import (
	"fmt"
	two "github.com/FloThinksPi/go-plugin-cyclic-dependencie-test/plugins/two/functions"
	"time"
)

func TestFunc()  {
	fmt.Print("one: testFunc called")
	time.Sleep(1)
	two.TestFunc()
}
