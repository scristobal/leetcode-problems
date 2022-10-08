impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut result = Vec::new();
        generate_parenthesis_rec(n, n, String::new(), &mut result);
        result 
    }
}

fn generate_parenthesis_rec(open: i32, close: i32, s: String, result: &mut Vec<String>) {
        if open == 0 && close == 0 {
            result.push(s);
            return;
        }

        if open > 0 {
            generate_parenthesis_rec(open - 1, close, s.clone() + "(", result);
        }

        if close > open {
            generate_parenthesis_rec(open, close - 1, s.clone() + ")", result);
        }
    }