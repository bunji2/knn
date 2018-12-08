package main

import (
	"fmt"
	"os"

	"github.com/bunji2/knn"
)

func main() {
	os.Exit(run())
}

/*

  x,  y, label
 11  11  0
  1   1  0
-11  11  1
 -1   1  1
-11 -11  2
 -1  -1  2
 11 -11  3
  1  -1  3
*/

func run() int {

	data := [][]float32{
		[]float32{11, 11},
		[]float32{1, 1},
		[]float32{-11, 11},
		[]float32{-1, 1},
		[]float32{-11, -11},
		[]float32{-1, -1},
		[]float32{11, -11},
		[]float32{1, -1},
	}
	labels := []int{
		0,
		0,
		1,
		1,
		2,
		2,
		3,
		3,
	}

	k := 2
	numDataElm := 2
	numLabels := 4
	cls := knn.New(k, numDataElm, numLabels)
	err := cls.Add(data, labels)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	xx := []float32{10, 10}
	fmt.Println(xx, "==>", cls.Predict(xx))
	xx = []float32{-10, 10}
	fmt.Println(xx, "==>", cls.Predict(xx))
	xx = []float32{-10, -10}
	fmt.Println(xx, "==>", cls.Predict(xx))
	xx = []float32{10, -10}
	fmt.Println(xx, "==>", cls.Predict(xx))
	xx = []float32{5, -5}
	fmt.Println(xx, "==>", cls.Predict(xx))

	return 0
}
