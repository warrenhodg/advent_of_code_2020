use std::io::prelude::*;
use std::fs::File;
use std::io::BufReader;
use regex::Regex;

fn main() -> std::io::Result<()> {
    let file = File::open("./input.txt")?;
    let mut reader = BufReader::new(file);

    let re = Regex::new(r"^([0-9]+)-([0-9]+) (.): (.*)$").unwrap();

    let mut count = 0;
    loop {
        let mut line = String::new();
        let len = reader.read_line(&mut line)?;
        if len == 0 {
            break;
        }

        let line = line.trim();

        let captures = re.captures(line).unwrap();
        let first = captures.get(1).unwrap().as_str();
        let second = captures.get(2).unwrap().as_str(); 
        let ch = captures.get(3).unwrap().as_str().chars().next().unwrap();
        let password = captures.get(4).unwrap().as_str();

        let first = first.parse::<usize>().unwrap();
        let second = second.parse::<usize>().unwrap();

        let password = password.chars();

        let mut i = 0;
        let mut ch_count = 0;
        password.for_each(|c| {
            i += 1;
            if i == first || i == second {
                if c == ch {
                    ch_count += 1;
                }
            }
        });
        if ch_count == 1 {
            count += 1;
        }
    }

    println!("{}", count);

    Ok(())
}
