package ds_sort


import (
    "testing"
)

type IntSlice []int 

func (i IntSlice) len() int {
    return len(i)
}

func (i IntSlice) less(a, b int) bool {
    if i[a] < i[b] {
        return true
    }
    return false
}

func (i IntSlice) swap(a, b int) {
    i[a], i[b] = i[b], i[a]
}

func TestGenericBubbleSort(t *testing.T) {
    var array IntSlice = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}

    
    GenericBubbleSort(array)
    
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            break
        }
    }

}


func TestBubbleSort(t *testing.T) {
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}
    
    BubbleSort(array)
    
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            break
        }
    }
}

func TestSelectionSort(t *testing.T) {
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}
    SelectionSort(array)
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            break
        }
    }

}

func TestInsertionSort(t *testing.T) {
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}
    InsertionSort(array)
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            break
        }
    }
}

func TestShellSort(t *testing.T) {
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}

    ShellSort(array)
    
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            return
        }
    }

}

func TestQuickSort(t *testing.T) {
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}
    QuickSort(array)
    for pos, val := range(array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            return
        }
    }
}


func TestMergeSort(t *testing.T) {
    //var array []int = []int{1,1,1}
    //var sorted_array []int = []int{1,1,1}
    
    var array []int = []int{4, 3, 5, 1, 4, 0, 7, 0, 4, 8}
    var sorted_array []int = []int{0, 0, 1, 3, 4, 4, 4, 5, 7, 8}
    Print(array) 
    merged_array := MergeSort(array)
    Print(merged_array) 
    for pos, val := range(merged_array) {
        if val != sorted_array[pos] {
            t.Errorf("Arrays do not equal each other.")
            return
        }
    }
}
