use std::io::prelude::*;
use std::fs::File;
use std::io::BufReader;
use std::collections::HashMap;

fn main() -> std::io::Result<()> {
    let mut amounts = HashMap::new();

    let file = File::open("./input.txt")?;
    let mut reader = BufReader::new(file);

    loop {
        let mut line = String::new();
        let len = reader.read_line(&mut line)?;
        if len == 0 {
            break;
        }
        let num = line.trim().parse::<i32>().unwrap();

        let need_num = 2020 - num;
        if amounts.contains_key(&need_num) {
            println!("The numbers are {} and {}, the product of which is {}", need_num, num, need_num*num);
            return Ok(());
        }

        amounts.insert(num, ());
    }

    Ok(())
}
