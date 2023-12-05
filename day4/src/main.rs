use std::fs;
use std::collections::HashSet;
use std::collections::HashMap;
use std::iter::FromIterator;

fn parse(cards: &str) -> Vec<Vec<Vec<i32>>> {
    cards.lines().map(|card: &str| {
        card.split("|").map(
            |sequence: &str| {
                sequence.trim().replace("  ", " ").split(' ')
                .map(|num: &str| num.parse::<i32>().unwrap()).collect()
            }
        ).collect()
    }).collect()
}

fn score_cards(cards: Vec<Vec<Vec<i32>>>) -> usize {
    let mut score : usize = 0;
    for card in cards {
        let winning_nums: HashSet<i32> = HashSet::from_iter(card[0].iter().cloned());
        let my_nums: HashSet<i32> = HashSet::from_iter(card[1].iter().cloned());
        let matching_count : usize = winning_nums.intersection(&my_nums).count();
        if matching_count > 0 {
            score += ((matching_count > 0) as usize) * (1 << (matching_count-1));
        }
    }
    score
}

fn count_copies(cards: Vec<Vec<Vec<i32>>>) -> usize {
    let mut copies: HashMap<i32, usize> = HashMap::from([(0, 1)]);
    let mut copy_count : usize = 0;
    for (i, card) in cards.iter().enumerate() {
        let winning_nums: HashSet<i32> = HashSet::from_iter(card[0].iter().cloned());
        let my_nums: HashSet<i32> = HashSet::from_iter(card[1].iter().cloned());
        let matching_count : usize = winning_nums.intersection(&my_nums).count();
        for j in 1..(matching_count+1) {
            let card_to_copy : i32 = (i + j) as i32;
            let mut card_copies = 1;
            if copies.contains_key(&(i as i32)) {
                card_copies = copies[&(i as i32)];
            }
            let card_copy_count = copies.entry(card_to_copy).or_insert(1);
            *card_copy_count += card_copies;
        }
        let card_copy_count = copies.entry(i as i32).or_insert(1);
        copy_count += *card_copy_count;
    }
    copy_count
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let cards = parse(&input);
    println!("Task 1: {}", score_cards(cards.clone()));
    println!("Task 2: {}", count_copies(cards));
}
