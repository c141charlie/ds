package main

import "list"
import "testing"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"list.TestInsertEmptyList", list.TestInsertEmptyList},
	{"list.TestInsertMultipleElements", list.TestInsertMultipleElements},
	{"list.TestInsertBetweenElements", list.TestInsertBetweenElements},
	{"list.TestInsertBeforeFirstElement", list.TestInsertBeforeFirstElement},
	{"list.TestInsertAfterLastElement", list.TestInsertAfterLastElement},
	{"list.TestInsertOutOfBounds", list.TestInsertOutOfBounds},
	{"list.TestAdd", list.TestAdd},
	{"list.TestSet", list.TestSet},
	{"list.TestDeleteOnlyElement", list.TestDeleteOnlyElement},
	{"list.TestDeleteFirstElement", list.TestDeleteFirstElement},
	{"list.TestDeleteLastElement", list.TestDeleteLastElement},
	{"list.TestDeleteMiddleElement", list.TestDeleteMiddleElement},
	{"list.TestDeleteOutOfBounds", list.TestDeleteOutOfBounds},
	{"list.TestDeleteByValue", list.TestDeleteByValue},
}
var benchmarks = []testing.InternalBenchmark{}

func main() {
	testing.Main(__regexp__.MatchString, tests)
	testing.RunBenchmarks(__regexp__.MatchString, benchmarks)
}
