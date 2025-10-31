func distributeCandies(candyType []int) int {
    m := make(map[int]bool)
    for _, ct := range candyType {
        m[ct] = true
    }
    numType := len(m)
    if numType < len(candyType)/2 {
        return numType
    } else {
        return len(candyType)/2
    }
}
