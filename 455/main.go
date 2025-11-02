func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    res := 0
    var i int
    var j int
    for {
        if i >= len(g) || j >= len(s) {
            break
        }
        if g[i] <= s[j] {
            res += 1
            i++
            j++
            continue
        } else {
            j++
            continue
        }
    }
    return res
}
