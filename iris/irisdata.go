package main

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var labelStrs = []string{
	"Iris-setosa",
	"Iris-versicolor",
	"Iris-virginica",
}

func split(data [][]float32, labels []int, rate float32) (trainData [][]float32, trainLabels []int, testData [][]float32, testLabels []int) {
	rand.Seed(time.Now().UnixNano())
	numTrain := int(float32(len(data)) * rate)
	numTest := len(data) - numTrain

	trainData = make([][]float32, numTrain)
	trainLabels = make([]int, numTrain)
	for i := 0; i < numTrain; i++ {
		j := rand.Intn(len(data))
		trainData[i] = data[j]
		trainLabels[i] = labels[j]
	}

	testLabels = make([]int, numTest)
	testData = make([][]float32, numTest)
	for i := 0; i < numTest; i++ {
		j := rand.Intn(len(data))
		testData[i] = data[j]
		testLabels[i] = labels[j]
	}

	return
}

func load(filePath string) (data [][]float32, labels []int, err error) {

	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	//reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	//reader := csv.NewReader(transform.NewReader(f, japanese.EUCJP.NewDecoder()))
	reader := csv.NewReader(f) //utf8
	reader.LazyQuotes = true   // ダブルクオートを厳密にチェックしない

	data = [][]float32{}
	labels = []int{}
	var cols []string
	for {
		cols, err = reader.Read() // 1行読み出す
		if err != nil {
			break
		}
		data = append(data, strsToFloat32s(cols[0:4]))
		labels = append(labels, strToLabelID(cols[4]))
	}

	if err == io.EOF {
		err = nil
	}
	return
}

func strsToFloat32s(xs []string) (r []float32) {
	if len(xs) == 0 {
		return
	}
	r = make([]float32, len(xs))
	for i, x := range xs {
		tmp, err := strconv.ParseFloat(x, 32)
		if err != nil {
			break
		}
		r[i] = float32(tmp)
	}
	return
}

func strToLabelID(str string) (labelID int) {
	//fmt.Println("# str =", str)
	for i, labelStr := range labelStrs {
		if str == labelStr {
			labelID = i
			return
		}
	}

	labelID = numLabels // error
	return
}
