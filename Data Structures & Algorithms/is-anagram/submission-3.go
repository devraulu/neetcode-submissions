func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    var lower func(rune) rune
    lower = func(b rune) rune{
        if b >= 'A' && b <= 'Z' {
            return b + ('a' - 'A')
        }
        return b
    }

    arr := [27]int{}
    for i, c := range s {
       arr[lower(c)-96]++
       c2 := rune(t[i])
       arr[lower(c2)-96]--
    }

    for _, c := range arr {
        if c != 0 {
            return false
        }
    }
    return true  
}
