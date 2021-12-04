use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader, Read};

/*
https://adventofcode.com/2021/day/2
Solution for part 1 and part 2
*/

/*
Solution is much like day 1. Read in the values, split apart the direction and value, match on the
direction, and follow the rules.
 */
fn part1() -> io::Result<()> {
    //open the file
    let f = File::open("./input").expect("Unable to open file");
    //use a buffered reader to read it line by line
    let mut br = BufReader::new(f);
    //initialize movement counters
    let mut horizontal = 0;
    let mut vertical = 0;
    //for each line in the input
    for line in br.lines() {
        let input_line = line?;
        //split the direction from the movement delta
        let split_pieces: Vec<&str> = input_line.split(" ").collect();
        //match on the direction and follow the associated rule
        match split_pieces[0] {
            //if we move "forward" we add the movement to horizontal movement counter
            "forward" => horizontal += split_pieces[1].parse::<u32>().unwrap(),
            //if we move "up" we subtract the movement from the aim
            "up" => vertical -= split_pieces[1].parse::<u32>().unwrap(),
            //if we move "down" we add the movement to the aim
            "down" => vertical += split_pieces[1].parse::<u32>().unwrap(),
            _ => panic!("unknown value")
        }
    }
    //show the result
    println!("{} : {} = {}",horizontal,vertical, horizontal*vertical);
    return Ok(())
}

/*
Solution is like part1() but with slightly different rules
 */
fn part2() -> io::Result<()> {
    //open the file
    let f = File::open("./input").expect("Unable to open file");
    //use a buffered reader to read it line by line
    let mut br = BufReader::new(f);
    //initialize movement counters
    let mut horizontal: u64 = 0;
    let mut vertical: u64 = 0;
    let mut aim: u64 = 0;
    //for each line in the input
    for line in br.lines() {
        let input_line = line?;
        //split the direction from the movement delta
        let split_pieces: Vec<&str> = input_line?.split(" ").collect();
        //convert movement delta to integer
        let movement_int: u64 = split_pieces[1].parse::<u64>().unwrap();
        //match on the direction and follow the associated rule
        match split_pieces[0] {
            /*
            if we move "forward" we add the movement to horizontal and add the "forward" movement
            multiplied by the "aim" to the vertical movement
             */
            "forward" => {horizontal += movement_int; vertical += aim * movement_int },
            //if we move "up" we subtract the movement from the aim
            "up" => aim -= movement_int,
            //if we move "down" we add the movement to the aim
            "down" => aim += movement_int,
            _ => panic!("unknown value")
        }
    }
    //show the result
    println!("{} : {} = {}",horizontal,vertical, horizontal*vertical);
    return Ok(())
}

fn main() -> io::Result<()>  {
    Ok(part2()?)
}
