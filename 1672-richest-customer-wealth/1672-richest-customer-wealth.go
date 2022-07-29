func maximumWealth(accounts [][]int) int {
    max := 0
    for _, acc := range accounts {
        wlth := 0
        for _, clnt := range acc{
            wlth += clnt
        }
        if wlth > max {
            max = wlth
        }
    }
    return max
}