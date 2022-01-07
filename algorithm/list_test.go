package main

import (
	"fmt"
	"testing"
)

func TestCreateSingleList(t *testing.T) {
	sl := createSingleList(1, 2, 3, 4, 5)
	fmt.Println(sl.checkCircle())
	rl := createRoundSingleList(1, 2, 3, 4, 5)
	fmt.Println(rl.checkCircle())
	//sl.Println()
	//sl.Reverse()
	//sl.Println()
	/*dl := createDoubleList(1,2,3,4,5)
	dl.Println()
	dl.Reverse()
	dl.Println()*/
}
