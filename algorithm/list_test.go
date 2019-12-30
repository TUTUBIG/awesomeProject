package algorithm

import (
	"fmt"
	"testing"
)

func TestCreateSingleList(t *testing.T) {
	sl := createSingleList(1, 2, 3, 4, 5)
	fmt.Println(sl.CheckCircle())
	sl.GenerateCircle()
	fmt.Println(sl.CheckCircle())
	//sl.Println()
	//sl.Reverse()
	//sl.Println()
	/*dl := createDoubleList(1,2,3,4,5)
	dl.Println()
	dl.Reverse()
	dl.Println()*/
}
