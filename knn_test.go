package knn_test

func Example() {
	// import "github.com/bunji2/knn"

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

	k := 3
	numDataElm := 2
	numLabels := 4
	cls := knn.New(k, numDataElm, numLabels)
	err := cls.Add(data, labels)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	xx := []float32{10, 10}
	fmt.Println(xx, "==>", cls.Predict(xx)) // 0

	xx = []float32{-10, 10}
	fmt.Println(xx, "==>", cls.Predict(xx)) // 1

	xx = []float32{-10, -10}
	fmt.Println(xx, "==>", cls.Predict(xx)) // 2

	xx = []float32{10, -10}
	fmt.Println(xx, "==>", cls.Predict(xx)) // 3

	xx = []float32{5, -5}
	fmt.Println(xx, "==>", cls.Predict(xx)) // 3

}
