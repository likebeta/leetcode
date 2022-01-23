package main

import (
	"container/heap"
)

// 股票价格波动
type StockPrice struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}

func (sp *StockPrice) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, pair{-price, timestamp})
	heap.Push(&sp.minPrice, pair{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}

func (sp *StockPrice) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	stockPrice := Constructor()
	stockPrice.Update(1, 10) // 时间戳为 [1] ，对应的股票价格为 [10] 。
	stockPrice.Update(2, 5)  // 时间戳为 [1,2] ，对应的股票价格为 [10,5] 。
	stockPrice.Current()     // 返回 5 ，最新时间戳为 2 ，对应价格为 5 。
	stockPrice.Maximum()     // 返回 10 ，最高价格的时间戳为 1 ，价格为 10 。
	stockPrice.Update(1, 3)  // 之前时间戳为 1 的价格错误，价格更新为 3 。
	// 时间戳为 [1,2] ，对应股票价格为 [3,5] 。
	stockPrice.Maximum()     // 返回 5 ，更正后最高价格为 5 。
	stockPrice.Update(4, 2)  // 时间戳为 [1,2,4] ，对应价格为 [3,5,2] 。
	stockPrice.Minimum()     // 返回 2 ，最低价格时间戳为 4 ，价格为 2 。
}
