package main

import (
	"math"
	"os"
)

type bowData struct {
	numCat int // 0<=catID<numCat

	cats      []catData      // カテゴリのデータ len(cats) == numCat
	words     []string       // 全文書に出現した単語のリスト 単語ID→単語文字列
	wordIndex map[string]int // 全文書に出現した単語のリスト 単語文字列→単語ID

	docs []docData // 文書のリスト
}

type catData struct {
	catID    int         // カテゴリID
	numDoc   int         // カテゴリに属する文書の個数
	wordFreq map[int]int // カテゴリ内の単語の出現回数
	numWord  int         // カテゴリ内の単語の合計
}

type docData struct {
	catID    int         // 文書の属するカテゴリ
	seq      []int       // 文書内の単語IDのリスト
	wordFreq map[int]int // 文書内の単語の出現回数
}

// P(Ci) カテゴリCi出現率
func (b *bowData) pCat(catID int) (r float64) {
	r = math.Log(float64(b.cats[catID].numDoc) / float64(len(b.docs)))
	return
}

// P(Wj|Ci) カテゴリCi内の単語Wjの出現率
func (b *bowData) pWordCat(catID, wordID int) (r float64) {
	cat := b.cats[catID]
	//r = float64(cat.wordFreq[wordID]) / float64(cat.numWord)
	r = math.Log(float64(cat.wordFreq[wordID]+1) / float64(cat.numWord+len(b.words)))
	return
}

// P(Ci|D) 文書DがカテゴリCiに属する確率
func (b *bowData) pCatDoc(catID int, wordIDs []int) (r float64) {
	tmp := b.pCat(catID)
	for _, wordID := range wordIDs {
		tmp += b.pWordCat(catID, wordID)
	}
	return
}

func main() {
	os.Exit(run())
}

func run() int {

	return 0
}
