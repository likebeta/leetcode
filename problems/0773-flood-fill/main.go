package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 图像渲染
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		dx  = []int{1, 0, 0, -1}
		dy  = []int{0, 1, -1, 0}
		dfs func(int, int)
	)

	currColor := image[sr][sc]
	if currColor != newColor {
		row, col := len(image), len(image[0])
		dfs = func(x, y int) {
			if image[x][y] == currColor {
				image[x][y] = newColor
				for i := 0; i < 4; i++ {
					mx, my := x+dx[i], y+dy[i]
					if mx >= 0 && mx < row && my >= 0 && my < col {
						dfs(mx, my)
					}
				}
			}
		}
		dfs(sr, sc)
	}
	return image
}

func testOne(in string, sr, sc, newColor int) {
	var inArr [][]int
	if err := json.Unmarshal([]byte(in), &inArr); err != nil {
		log.Fatalln("parse failed", err)
	}

	dumpImage := func(inArr [][]int) []string {
		var inStr []string
		for i := range inArr {
			buf := strings.Builder{}
			for j := range inArr[i] {
				buf.WriteString(fmt.Sprintf("%7d", inArr[i][j]))
			}
			inStr = append(inStr, buf.String())
		}
		return inStr
	}

	inStr := dumpImage(inArr)
	floodFill(inArr, sr, sc, newColor)
	outStr := dumpImage(inArr)
	for i := range inStr {
		fmt.Printf("%s ---->  %s\n", inStr[i], outStr[i])
	}
}

func main() {
	testOne("[[1,1,1],[1,1,0],[1,0,1]]", 1, 1, 2)  // [[2,2,2],[2,2,0],[2,0,1]]
}
