package strutil

import "os"
import "fmt"

func SearchUsingBruteForce(pattern string, text string) (os.Error, int) {
    if text == "" || pattern == "" {
        return os.NewError("text cannot be empty"), 0
    }

    runes_pattern = []int(pattern)
    runes_text = []int(text)

    i := 0

    for i <= len(runes_text) - len(runes_pattern) {
         
    }
    
}
