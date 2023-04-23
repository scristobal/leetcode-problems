function map(arr: number[], fn: (n: number, i: number) => number): number[] {
    const res = [];
    for (const [i,v ] of arr.entries()) {
        res.push(fn(v, i))
    }
    return res
};