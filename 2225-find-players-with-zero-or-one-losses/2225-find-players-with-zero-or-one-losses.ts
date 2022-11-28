function findWinners(matches: number[][]): number[][] {
    
    const lossScore = new Map<number, number>();
    
    for (let [winner, loser] of matches){
        let losses = lossScore.get(loser)!==undefined ? lossScore.get(loser) : 0;
        lossScore.set(loser, losses + 1);
        
        losses = lossScore.get(winner)!==undefined ? lossScore.get(winner) : 0;
        lossScore.set(winner, losses)
    }
      
    const noLosses = [];
    const oneLoss = [];
    
    for (let [player, losses] of lossScore){
        if (losses == 0) {
            noLosses.push(player)
            continue
        }
        if (losses == 1){
            oneLoss.push(player)
            continue
        }
    }
    
    noLosses.sort((a, b) => a -b);
    oneLoss.sort((a,b) => a - b);
    
    return [noLosses, oneLoss]
    
};