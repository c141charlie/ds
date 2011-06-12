package strutil

import "os"

func SearchUsingBruteForce(pattern string, text string) (int, os.Error) {
    if text == "" || pattern == "" {
        return 0, os.NewError("text cannot be empty")
    }

    runes_pattern := []int(pattern)
    runes_text := []int(text)

    i := 0

    for i <= len(runes_text) - len(runes_pattern) {
        x := 0    
        for x < len(runes_pattern) {
            if runes_text[i + x] == runes_pattern[x]
        }

    }

    return 0, os.NewError("pattern not found")
    
}

