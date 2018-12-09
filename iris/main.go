package main

import (
	"fmt"
	"os"

	"github.com/bunji2/knn"
	"github.com/bunji2/metrics"
)

const (
	k          = 5
	numDataElm = 4
	numLabels  = 3
)

func main() {
	os.Exit(run())
}

func run() int {
	// iris データの読み出し
	data, labels, err := load("iris.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	//fmt.Println(data, labels)

	// 学習データと評価データを用意
	trainData, trainLabels, testData, testLabels := split(data, labels, 0.8)
	//fmt.Println(trainData, trainLabels, testData, testLabels)

	// k-NN 分類器に学習データをセット
	cls := knn.New(k, numDataElm, numLabels)
	err = cls.Add(trainData, trainLabels)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	// メトリクスオブジェクトの用意
	md := metrics.New(numLabels)

	// 評価作業
	for i := 0; i < len(testData); i++ {
		predLabel := cls.Predict(testData[i])
		err = md.AddClassID(predLabel, testLabels[i])
		//fmt.Println(testData[i], testLabels[i], predLabel)
	}

	// メトリクスの表示
	printMetrics(md)
	return 0
}

func printMetrics(md *metrics.Data) {

	microPrecision, microRecall, microFMeasure, overallAccuracy := md.MicroMetrics()
	macroPrecision, macroRecall, macroFMeasure, averageAccuracy := md.MacroMetrics()
	fmt.Printf("Micro Precision:  %f\n", microPrecision)
	fmt.Printf("Micro Recall:     %f\n", microRecall)
	fmt.Printf("Micro F-Measure:  %f\n", microFMeasure)
	fmt.Printf("Overall Accuracy: %f\n", overallAccuracy)
	fmt.Printf("Macro Precision:  %f\n", macroPrecision)
	fmt.Printf("Macro Recall:     %f\n", macroRecall)
	fmt.Printf("Macro F-Measure:  %f\n", macroFMeasure)
	fmt.Printf("Average Accuracy: %f\n", averageAccuracy)
}
