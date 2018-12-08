package knn

import (
	"errors"
	"math"
	"sort"
)

// Classifier : kNN を使った分類器
type Classifier struct {
	k          int         `json:"k"`            // 調べる近傍点の個数
	numLabels  int         `json:"num_labels"`   // ラベルの個数（0<=ラベル番号<numLabels）
	data       [][]float32 `json:"data"`         // データ
	labels     []int       `json:"labels"`       // データごとのラベル
	numDataElm int         `json:"num_data_elm"` // データの要素の個数 (!=データの個数)
	distances  []distance
}

type distance struct {
	idx      int
	distance float32
}

func (cls *Classifier) Len() int {
	return len(cls.distances)
}

func (cls *Classifier) Less(i, j int) bool {
	return cls.distances[i].distance < cls.distances[j].distance
}

func (cls *Classifier) Swap(i, j int) {

	cls.distances[i], cls.distances[j] = cls.distances[j], cls.distances[i]
}

// New : 分類器の作成。
// Params:
//       k --- 調べる近傍点の個数
//       numDataElm --- データの要素の個数 (!=データの個数)
//       numLabels --- ラベルの個数（0<=ラベル番号<numLabels）
func New(k, numDataElm, numLabels int) (r *Classifier) {
	return &Classifier{
		k:          k,
		numLabels:  numLabels,
		data:       [][]float32{},
		labels:     []int{},
		numDataElm: numDataElm,
	}
}

// Add : データとラベルの追加
func (cls *Classifier) Add(data [][]float32, labels []int) (err error) {
	// num : 処理する個数。data と labels のそれぞれの個数が異なる時はエラー。
	num := len(data)
	if num != len(labels) {
		err = errors.New("knn.Classifier.Add: size of data and labels are different")
		return
	}
	for i := 0; i < num; i++ {
		if len(data[i]) != cls.numDataElm {
			// データの要素の個数と合わないものはエラー
			err = errors.New("knn.Classifier.Add: len(data[i]) != cls.numDataElem")
			break
		}
		if labels[i] >= cls.numLabels {
			// ラベル番号の上限を超えるものはエラー
			err = errors.New("knn.Classifier.Add: labels[i] >= cls.numLabels")
			break
		}
		cls.data = append(cls.data, data[i])
		cls.labels = append(cls.labels, labels[i])
	}

	// fmt.Println(cls.data)
	// fmt.Println(cls.labels)
	return
}

// Predict : ラベルの予測
func (cls *Classifier) Predict(data []float32) (label int) {
	// 全データとの距離を計算
	cls.calcDistances(data)
	// 距離を昇順でソート
	sort.Sort(cls)
	// k 個の近傍点のラベルの個数を集計
	labelCounts := make([]int, cls.numLabels)
	for i := 0; i < cls.k; i++ {
		idx := cls.distances[i].idx
		label := cls.labels[idx]
		labelCounts[label] = labelCounts[label] + 1
	}
	// 最も多かったラベルを決定
	maxCount := 0
	for i, count := range labelCounts {
		if count > maxCount {
			label = i
		}
	}
	return
}

func (cls *Classifier) calcDistances(xx []float32) (err error) {
	if len(xx) != cls.numDataElm {
		err = errors.New("knn.Classifier.calcDistances: len(xx) != cls.numDataElm")
		return
	}
	cls.distances = make([]distance, len(cls.data))
	for i, xx2 := range cls.data {
		cls.distances[i] = distance{idx: i, distance: calcDistance(xx, xx2)}
	}
	return
}

// calcDistance : calculate L2
func calcDistance(xx, xx2 []float32) (r float32) {
	s := float64(0)
	for i := 0; i < len(xx); i++ {
		delta := float64(xx[i] - xx2[i])
		s += delta * delta
	}
	r = float32(math.Sqrt(s))
	return
}

