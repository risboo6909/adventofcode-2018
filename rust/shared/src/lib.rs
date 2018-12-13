use std::io::{self, BufRead};


pub fn parse_input<F, T>(func: F) -> Vec<T> 
    where F: Fn(&str) -> T {

    let mut output = Vec::new();
    for line in io::stdin().lock().lines() {
        let val = line.unwrap();
        output.push(func(&val));
    }

    output
}
