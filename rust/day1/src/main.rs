use shared::parse_input;
use std::collections::HashMap;

fn main() {
    let mut iter = 0;
    let mut net_val = 0;
    let mut repeated_value: Option<isize> = None;
    let mut freq: HashMap<isize, usize> = HashMap::new();

    let numbers = parse_input(|line| line.parse::<isize>().unwrap());

    while repeated_value == None {
        for value in &numbers {
            net_val += value;

            *freq.entry(net_val).or_insert(0) += 1;

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
