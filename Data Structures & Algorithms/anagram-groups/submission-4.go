import "slices"
func groupAnagrams(strs []string) [][]string {
    if len(strs) < 1 {
        return [][]string{}
    }
    m := map[string][]string{}
    for _, unsorted := range strs {
        runes := []rune(unsorted)
        slices.Sort(runes)
        str := string(runes)
        m[str] = append(m[str], unsorted)
    }

    arr := [][]string{}
   for _, v := range m {
    arr = append(arr, v)
   }
   return arr
}
