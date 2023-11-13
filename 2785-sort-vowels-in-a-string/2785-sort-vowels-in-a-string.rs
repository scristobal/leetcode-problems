use std::collections::VecDeque;

const CHARS: [char; 10] = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'];

impl Solution {

    pub fn sort_vowels(s: String) -> String {
        let mut vowels = s
            .chars()
            .filter(|a| CHARS.contains(a))
            .collect::<VecDeque<_>>();

        vowels.make_contiguous().sort();

        s.chars()
            .map(|c| {
                CHARS
                    .contains(&c)
                    .then(|| vowels.pop_front())
                    .unwrap_or(Some(c))
            })
            .filter_map(|f| f)
            .collect()
    }
}