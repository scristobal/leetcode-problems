use std::collections::HashMap;
impl Solution {
    pub fn max_points(points: Vec<Vec<i32>>) -> i32 {

        if points.len()==1{ return 1}

        let mut visited: Vec<Vec<i32>> = vec![];
        let mut unvisited = points;

        let mut lines_count: HashMap<[Vec<i32>; 2], i32> = HashMap::new();

        while let Some(point) = unvisited.pop() {
            let lines = lines_count.keys().collect::<Vec<_>>();

            let mut found = vec![];

            for line in lines {
                if aligned(&point, &line[0], &line[1]) {
                    let line = line.clone();

                    found.push(line)
                }
            }

            for line in &found {
                lines_count.entry(line.clone()).and_modify(|v| *v += 1);
            }

            if found.is_empty() {
                for visited_point in &visited {
                    lines_count.insert([point.clone(), visited_point.clone()], 2);
                }
            }

            visited.push(point);
        }

        *lines_count.values().max().unwrap()
    }
}

fn aligned(point_a: &Vec<i32>, point_b: &Vec<i32>, point_c: &Vec<i32>) -> bool {
    point_a[0] * (point_b[1] - point_c[1])
        + point_b[0] * (point_c[1] - point_a[1])
        + point_c[0] * (point_a[1] - point_b[1])
        == 0
}