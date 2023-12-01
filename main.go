package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/lambada/aoc-golang/Year2023/Day01/Part1"
	"github.com/lambada/aoc-golang/Year2023/Day01/Part2"
)

func main() {

	year := flag.String("year", "2023", "AoC Year to run")
	day := flag.String("day", "01", "Day [Day01 - 25]")
	part := flag.String("part", "1", "Which part [Part1 - Part2]")
	rawInput := flag.String("input", "", "The input to give")

	flag.Parse()

	input := strings.Split(*rawInput, "\\n")

	packages := map[string]interface{}{
		"y2023d01p1": y2023d01p01.Calculate,
		"y2023d01p2": y2023d01p02.Calculate,
	}

	packageName := "y" + *year + "d" + *day + "p" + *part

	fn := reflect.ValueOf(packages[packageName])

	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(input)

	var res []reflect.Value
	res = fn.Call(args)
	result := res[0].Interface()

	fmt.Println(result)
}
