use std::io::{self, BufRead};
use std::collections::HashMap;


fn parse_input() -> Vec<isize> {
    let mut numbers = Vec::new();
    let stdin = io::stdin();

    for line in stdin.lock().lines() {
        let val = line.unwrap();
        numbers.push(val.parse::<isize>().unwrap());
    }

    numbers
}

fn main() {

    let mut iter = 0;
    let mut net_val = 0;
    let mut repeated_value: Option<isize> = None;
    let mut freq: HashMap<isize, usize> = HashMap::new();

    let numbers = parse_input();

    while repeated_value == None {        
        for value in &numbers {
            
            net_val += value;
            
            *freq
            .entry(net_val)
            .or_insert(0) += 1;

            if freq[&net_val] > 1 {                
                repeated_value = Some(net_val);
                break;
            }
        }

        if iter == 0 {
            println!("Result value: {}", net_val);
        }

        iter += 1;
    }

    println!("Repeated value: {}", repeated_value.unwrap());

}
