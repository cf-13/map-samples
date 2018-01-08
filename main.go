package main

import (
	"strings"
	"fmt"
	"github.com/nh038426/print-function-in-out/printio"
	"golang.org/x/tour/wc"
)

type Sample1 struct {
	Lat, Long float64
}

const (
	AnswerKey = "Answer"
)

func main() {
	printio.FuncIn(main)

	standardMap()
	standardMapLiteral1()
	standardMapLiteral2()
	mutatingMaps()
	exerciseMaps()

	printio.FuncOut(&printio.NotReturnFunction{main})
}

func standardMap() {
	printio.FuncIn(standardMap)

	m := make(map[string]Sample1)
	m["Bell Labs"] = Sample1 {
		40.68433, -74.39967,
	}
	fmt.Printf("m[\"Bell Labs\"]: %v\n", m["Bell Labs"])

	printio.FuncOut(&printio.NotReturnFunction{standardMap})
}

func standardMapLiteral1() {
	printio.FuncIn(standardMapLiteral1)

	m := map[string]Sample1 {
		"Bell Labs": Sample1{
			40.68433, -74.39967,
		},
		"Google": Sample1{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	printio.FuncOut(&printio.NotReturnFunction{standardMapLiteral1})
}

func standardMapLiteral2() {
	printio.FuncIn(standardMapLiteral2)

	m := map[string]Sample1 {
		"Bell Labs": { 40.68433, -74.39967 },
		"Google": { 37.42202, -122.08408 },
	}
	fmt.Println(m)

	printio.FuncOut(&printio.NotReturnFunction{standardMapLiteral2})
}

func mutatingMaps() {
	printio.FuncIn(mutatingMaps)

	m := make(map[string]int)

	m[AnswerKey] = 42
	fmt.Println("The Value:", m[AnswerKey])

	m[AnswerKey] = 48
	fmt.Println("The Value:", m[AnswerKey])

	delete(m, AnswerKey)
	fmt.Println("The Value:", m[AnswerKey])

	v, ok := m[AnswerKey]
	fmt.Println("The Value:", v, "Present?", ok)

	printio.FuncOut(&printio.NotReturnFunction{mutatingMaps})
}

func exerciseMaps() {
	printio.FuncIn(exerciseMaps)

	wc.Test(exerciseMapsTarget)

	printio.FuncOut(&printio.NotReturnFunction{exerciseMaps})
}

func exerciseMapsTarget(x string) map[string]int {
	printio.FuncIn(exerciseMapsTarget)

	s1 := strings.Fields(x)
	m := make(map[string]int)
	for _, v := range s1 {
		if m[v] == 0 {	// == 0: 未評価
			for _, tmp := range s1 {
				if strings.Compare(v, tmp) == 0 {
					m[v]++
				}
			}
			if m[v] == 0 {
				// 自身をカウント
				m[v] = 1
			}
		}
	}

	printio.FuncOut(&printio.ReturnFunction{exerciseMapsTarget, m})
	return m
}