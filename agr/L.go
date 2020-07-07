package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Log struct {
	Time   int
	Symbol string
	Nums   int
}

type Search struct {
	Start  int
	End    int
	Symbol string
}

func F(logs []Log, search []Search) {
	S := make([]int, 0)
	M := make(map[int]map[string]int, len(logs))
	sum := 0
	for idx, _ := range logs {
		S = append(S, logs[idx].Time)
		M[logs[idx].Time] = map[string]int{
			logs[idx].Symbol: logs[idx].Nums,
		}
	}
	//先根据二分查找确定范围,再在范围内遍历
	for idx, _ := range search {
		Start := BinarySearchLeft(S, search[idx].Start)
		for Start <= search[idx].End {
			sum += M[Start][search[idx].Symbol]
			Start++
		}
		fmt.Println(sum)
		sum = 0
	}
}

//二分查找
func BinarySearchLeft(s []int, Time int) int {
	Start, End := 0, len(s)-1
	for Start < End {
		mid := (Start + End) / 2
		if s[mid] < Time {
			Start = mid + 1
		} else {
			End = mid
		}
	}
	return s[Start]
}

func main() {
	//构造数据
	/*F([]Log{
		{
			Time:   1575562565,
			Symbol: "sqli",
			Nums:   3,
		}, {
			Time:   1575562567,
			Symbol: "xss",
			Nums:   10},
		{
			Time:   1575562572,
			Symbol: "csrf",
			Nums:   2},
		{
			Time:   1575562579,
			Symbol: "sqli",
			Nums:   5},
	},
		[]Search{
			{
				Start:  0,
				End:    1575562571,
				Symbol: "xss",
			}, {
				Start:  1575562565,
				End:    1575562579,
				Symbol: "sqli",
			}, {
				Start:  1575562565,
				End:    1575562579,
				Symbol: "none",
			},
		},
	)*/
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	lenLog, lenSearch := 0, 0
	str := strings.Split(s.Text(), " ")
	if len(str) != 2 {
		os.Exit(-1)
	}
	lenLog, err := strconv.Atoi(str[0])
	if err != nil {
		os.Exit(-1)
	}
	lenSearch, err = strconv.Atoi(str[1])
	if err != nil {
		os.Exit(-1)
	}
	logs, searchs := make([]Log, 0), make([]Search, 0)

	for i := 1; i <= lenLog; i++ {
		s.Scan()
		strLog := strings.Split(s.Text(), " ")
		if len(strLog) != 3 {
			os.Exit(-1)
		}
		timeStamp, err := strconv.Atoi(strLog[0])
		if err != nil {
			os.Exit(-1)
		}
		nums, err := strconv.Atoi(strLog[2])
		if err != nil {
			os.Exit(-1)
		}
		logs = append(logs, Log{
			Time:   timeStamp,
			Symbol: strLog[1],
			Nums:   nums,
		})
	}

	for i := 1; i <= lenSearch; i++ {
		s.Scan()
		strLog := strings.Split(s.Text(), " ")
		if len(strLog) != 3 {
			os.Exit(-1)
		}
		timeStampFrom, err := strconv.Atoi(strLog[0])
		if err != nil {
			os.Exit(-1)
		}
		timeStampTo, err := strconv.Atoi(strLog[1])
		if err != nil {
			os.Exit(-1)
		}
		searchs = append(searchs, Search{
			Start:  timeStampFrom,
			End:    timeStampTo,
			Symbol: strLog[2],
		})
	}
	F(logs[:], searchs[:])
}
