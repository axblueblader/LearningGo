package main

import "fmt"

type Context struct {
	strategy Strategy
}

func (context Context) Sort() {
	context.strategy.DoSort()
}

type Strategy interface {
	DoSort()
}

type MergeSort struct{}

func (msort MergeSort) DoSort() {
	fmt.Println("Merge sort algo is being used")
}

type BubbleSort struct{}

func (bsort BubbleSort) DoSort() {
	fmt.Println("Bubble sort is being used")
}

type QuickSort struct{}

func (qsort QuickSort) DoSort() {
	fmt.Println("Quick sort is being used")
}

func main() {
	fmt.Println("Strategy Pattern Demo")
	context := Context{strategy: QuickSort{}}
	context.Sort()
	context.strategy = MergeSort{}
	context.Sort()
	context.strategy = BubbleSort{}
	context.Sort()
}
