package binary

//import "fmt"

func Search(str []string, letter string) int {
    low := 0
    high := len(str)-1
        
    

    for low <= high {
        index := low + (high - low)/2
        if letter == str[index] {
            return index
        } else if letter < str[index] {
            high = index - 1
        } else {
            low = index + 1
        }
    }
    if low == 0 {
        return low -1
    }
    return high+1
}
