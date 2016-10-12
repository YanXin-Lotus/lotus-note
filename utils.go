package utils

//RemoveDuplicate remove duplicate string in string array
func RemoveDuplicate(slis *[]string) {
    found := make(map[string]bool)
    j := 0
    for i, val := range *slis {
        if _, ok := found[val]; !ok {
            found[val] = true
            (*slis)[j] = (*slis)[i]
            j++
        }
    }
    *slis = (*slis)[:j]
}
