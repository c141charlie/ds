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

func ShellSort(array []int) {
    increments := []int{4,3,2,1}
    for i:=0; i < len(increments); i++ {
        increment := increments[i]
        hsort(array, increment)
    }
}

func hsort(array []int, increment int) {
    if (len(array) <  increment * 2) {
        return
    }
    for i:=0; i<increment; i++ {
        sortsublist(array, i, increment)
    }
}

func sortsublist(array []int, start_index int, increment int) {
    for i:=start_index+increment; i < len(array); i+= increment {
        value := array[i]
        var j int
        for j=i; j > start_index; j = j-increment {
            previous_value  := array[j-increment]
            if value > previous_value {
                break
            }
            array[j] = previous_value 
        }
        array[j] = value
    }
}

func QuickSort(array []int) {
    quicksort(array, 0, len(array)-1)   
}

func partition(array []int, left int, right int) int {
    i, j := left, right

    pivot := array[(left + right)/2]

    for i < j {
        for array[i] < pivot {
            i++
        }
        for array[j] > pivot {
            j --
        }
        if i <= j {
            array[i], array[j] = array[j], array[i]
            i++
            j--
        }
        
    }
    return i
}

func quicksort(array []int, left int, right int) {
    index := partition(array, left, right)
    
    if left < index -1 {
        quicksort(array, left, index-1)
    }
    if index < right {
        quicksort(array, index, right)
    }
}

func MergeSort(array []int) []int{
    l := len(array)
    if l <= 1 {
        return array
    }
    left := MergeSort(array[0:l/2])
    right := MergeSort(array[l/2:l])
    result := merge(left, right)
    
    return result
}

func merge(left []int, right []int) []int {
    lenl, lenr := len(left), len(right)
    result := make([]int, lenl+lenr)
    i := 0
    l, r := lenl, lenr
    
    for l > 0 || r > 0 {
        if l == 0 && r > 0 {
            result[i] = right[lenr-r]
            r --
        } else if l > 0 && r == 0 {
            result[i] = left[lenl-l]
            l--
        } else if left[lenl-l] < right[lenr-r] {
            result[i] = left[lenl-l]
            l--
        } else {
            result[i] = right[lenr-r]
            r--
        }
        i++
    }
    return result
}
