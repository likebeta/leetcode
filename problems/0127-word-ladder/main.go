package main

import (
	"leetcode/helper"
)

// 单词接龙
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordLen := len(beginWord)
	comboMap := make(map[string][]string)
	var (
		combo string
		bCan  bool
		tmp   []byte
	)
	for _, word := range wordList {
		if word == endWord {
			bCan = true
		}
		if word != beginWord {
			tmp = []byte(word)
			for i := 0; i < wordLen; i++ {
				tmp[i] = '*'
				combo = string(tmp)
				comboMap[combo] = append(comboMap[combo], word)
				tmp[i] = word[i]
			}
		}
	}
	if bCan {
		visitedMap := make(map[string]bool)
		queue := []string{beginWord}
		for ans := 1; len(queue) > 0; ans++ {
			var newQueue []string
			for _, word := range queue {
				tmp = []byte(word)
				for i := 0; i < wordLen; i++ {
					tmp[i] = '*'
					combo = string(tmp)
					for _, wd := range comboMap[combo] {
						if wd == endWord {
							return ans + 1
						}
						if !visitedMap[wd] {
							visitedMap[wd] = true
							newQueue = append(newQueue, wd)
						}
					}
					tmp[i] = word[i]
				}
			}
			queue = newQueue
		}
	}
	return 0
}

func main() {
	helper.Log(ladderLength("a", "c", []string{"a", "b", "c"}), 2)
	helper.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}), 5)
	helper.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}), 0)
}
