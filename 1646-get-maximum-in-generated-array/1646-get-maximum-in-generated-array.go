func getMaximumGenerated(n int) int {
    ga := make([]int, n + 1, n + 1);
    
    if n == 0 { return 0 }
    
    ga[1] = 1;
    
    m := 1
    
    for i := 2; i < len(ga); i++ {
        
        ga[i] = ga[i/2]
        
        if i % 2 != 0 { 
            ga[i] += ga[i/2 + 1]
        }
                
         if m < ga[i] {
            m = ga[i]
        }
    }
    
    return m; 
}