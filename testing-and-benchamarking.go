// go test command runs test
// Testing code usually lies in the same package as the code it tests


package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int{
	// Typically the code we're testing would be in a source file named something like
	// intutils.go and the test file for it would then be na,ed intutils_test.go

	if a < b{
		return a
	}
	return b
}

// A test   is created by writing the function name beginning with Test
func TestIntMinBasic(t *testing.T){
	ans := IntMin(2, -2)
	if ans != -2{
		// t.Error* will report test failures but continue executing the test.
		// t.Fatal* will report test failures and stop the test immediately
		t.Errorf("IntMin(2,-2) = %d; want -2", ans)
	}


}


// Writing tests can be repetitive, so it is idomatic to use a table-driven style, where 
// test input and outputs are listed in a table and single loop walks over them
// and perform the test logic.
func TestIntMinTableDriven(t *testing.T){
	var tests =[]struct{
		a,b int
		want int
	}

	{
	{0,1,0}
	{1,0,0}
	{2,-2,-2}
	{0,-1,-1}
	{-1,0,-1}
	}


	for _, tt := range tests{
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T){
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want{
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}


// Benchmark tests typically go in _test.go files and are named beginning with benchmark
func BenchmarkIntMin(b *testing.B){
	// Any code thats required for the benchmark to run but should not be measures goes before

	for b.Loop(){
		IntMin(1,2)
	}
}

// The benchmark runner will automatically execute this loop body many times to
// determine a reasonable estimate of the run-time of a single iteration