package ds_sort

import "fmt"

type Sortable interface {
    len() int
    less(i, j int) bool
    swap(i, j int)
}


func Print(array []int) {
    for _, val := range(array) {
        fmt.Printf("%d", val)
    }
    fmt.Printf("\n")
}

func GenericBubbleSort(data Sortable) {
    for i := 0; i < data.len(); i++ {
        for j := 0; j < data.len()-i-1; j++ {
            if data.less(j+1, j) {
                data.swap(j+1, j) 
            }
        }
    }
}

func BubbleSort(array []int) {
    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array)-i-1; j++ {
            if array[j] > array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
}

func SelectionSort(array []int) {
    for i := 0; i < len(array); i++ {
        small := i
        for j := i; j<len(array); j++ {
            if array[j] < array[small] {
                small = j
            }
        }
        if small != i {
            array[i], array[small] = array[small], array[i]
        }
    }
}

func InsertionSort(array []int) {
    for i := 0; i < len(array); i++ {
        elem := array[i]
        j := i-1
        for j >= 0 && array[j] > elem {
            array[j+1] = array[j]
            j--
        }
        array[j+1] = elem
    }
}
