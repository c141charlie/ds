package main

import "ds_sort"
import "testing"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"ds_sort.TestGenericBubbleSort", ds_sort.TestGenericBubbleSort},
	{"ds_sort.TestBubbleSort", ds_sort.TestBubbleSort},
	{"ds_sort.TestSelectionSort", ds_sort.TestSelectionSort},
	{"ds_sort.TestInsertionSort", ds_sort.TestInsertionSort},
}
var benchmarks = []testing.InternalBenchmark{}

func main() {
	testing.Main(__regexp__.MatchString, tests)
	testing.RunBenchmarks(__regexp__.MatchString, benchmarks)
}
