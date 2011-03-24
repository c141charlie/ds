package binary

func Search(str []string, letter string) int {

    low := 0
    high := len(str)-1

    for low <= high {
        index := low + (high-low)/2
        if letter == str[index] {
            return index
        } else if str[index] > letter {
            high = index - 1
        } else {
            low = index + 1
        }
    }

    if low == 0 {
        return low - 1 
    }
    return high + 1
}

/*func Search(str []string, letter string) int {
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
*/
func Rsearch(str []string, letter string) int {
    return rec_search(str, letter, 0, len(str)-1)
}

func rec_search(str []string, letter string, low int, high int) int {
    if low > high {
        if low == 0 {
            return low -1
        }
        return high+1
    }
    
    index := low + (high - low)/2

    if letter < str[index] {
        return rec_search(str, letter, low, index-1)
    } else if letter > str[index] {
        return rec_search(str, letter, index+1, high)
    }
    return index
} 
