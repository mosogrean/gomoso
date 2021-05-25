package gomoso

import "fmt"

func main() {
	fmt.Println(`1111`)
}

type Helper struct {

}

func (Helper) Println() {
	fmt.Println(11111)
}
