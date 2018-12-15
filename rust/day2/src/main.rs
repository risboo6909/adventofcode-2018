use shared::parse_input;


#[derive(Debug)]
struct Container {
    max_value: Option<usize>,
    two: usize,
    three: usize,
}

impl Container {
    pub(crate) fn new(max_value: Option<usize>) -> Self {
        Container {
            max_value: max_value,
            two: 0,
            three: 0,
        }
    }

    fn inc_two(&mut self, delta: usize) {
        self.two = match self.max_value {
            Some(x) => std::cmp::min(x, self.two + delta),
            None => self.two + delta,
        }
    }

    fn inc_three(&mut self, delta: usize) {
        self.three = match self.max_value {
            Some(x) => std::cmp::min(x, self.three + delta),
            None => self.three + delta,
        }
    }

    pub(crate) fn inc(&mut self, val: usize) {
        match val {
            2 => self.inc_two(1),
            3 => self.inc_three(1),
            _ => {}
        }
    }

    pub(crate) fn update(&mut self, other: Container) {
        self.inc_two(other.two);
        self.inc_three(other.three);
    }
}

#[inline]
fn put_to_bucket(letter: char, buckets: &mut Vec<usize>) {
    buckets[letter as usize - 97] += 1;
}

#[inline]
fn scan_buckets(buckets: &Vec<usize>) -> Container {
    let mut res = Container::new(Some(1));
    for val in buckets {
        res.inc(*val);
    }
    res
}

#[inline]
fn str_diff(word1: &str, word2: &str) -> (usize, String) {
    let mut diffs = 0;
    let mut common_substr = Vec::new();
    for (a, b) in word1.chars().zip(word2.chars()) {
        if a != b {
            diffs += 1;
        } else {
            common_substr.push(a);
        }
    }

    (diffs, common_substr.into_iter().collect::<String>())
}


fn main() {

    let input_lines: Vec<String> = parse_input(|line| line.to_owned());

    // First part

    let alphabet = "abcdefghijklmnopqrstuvwxyz";

    let mut res = Container::new(None);

    for line in &input_lines {
        let mut buckets = vec![0; alphabet.len()];

        for c in line.chars() {
            put_to_bucket(c, &mut buckets);
        }

        res.update(scan_buckets(&buckets));
    }

    println!("Checksum: {}", res.two * res.three);

    // Second part

    for (idx, word1) in input_lines.iter().enumerate() {

        for word2 in input_lines.iter().skip(idx) {
            let (diff_count, common_substr) = str_diff(word1, word2);
            if diff_count == 1 {
                println!("Common sub-string: {}", common_substr);
                break;
            }
        }

    }

}
