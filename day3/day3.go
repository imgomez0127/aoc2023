package main

import (
    "fmt"
	"log"
	"bufio"
	"os"
	"unicode"
	"strconv"
)

type Pair struct {
	x int
	y int
}


func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(array[] int) int {
    result := 0
    for _,
    v := range array {
        result += v
    }
    return result
}

func has_adjacent_symbol(data []string, i int, j int, max_len int) bool {
	var has_symbol = false
	for a := i - 1; a <= i + 1; a++ {
		for b := j - 1; b <= j + 1; b++ {
			var c = max(min(a, max_len-1), 0)
			var d = max(min(b, max_len-1), 0)
			has_symbol = has_symbol || (!unicode.IsDigit(rune(data[c][d])) && data[c][d] != '.')
		}
	}
	return has_symbol
}

func has_adjacent_symbol2(data []string, i int, j int, max_len int, mapping map[string]map[int]bool, num_loc int) bool {
	var has_symbol = false
	for a := i - 1; a <= i + 1; a++ {
		for b := j - 1; b <= j + 1; b++ {
			var c = max(min(a, max_len-1), 0)
			var d = max(min(b, max_len-1), 0)
			if data[c][d] == '*' {
				var symbol_loc = strconv.Itoa(c) + " " + strconv.Itoa(d)
				if mapping[symbol_loc] == nil {
					mapping[symbol_loc] = make(map[int]bool, 100)
				}
				mapping[symbol_loc][num_loc] = true
				has_symbol = true
			}
		}

	}

	return has_symbol
}

func compute(data []string, max_len int) int {
	var valid_nums = make([]int, 1000)
	for i, line := range data {
		var num = ""
		var is_valid = false
		for j, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)
				is_valid = is_valid || has_adjacent_symbol(data, i, j, max_len)
			}  else if !unicode.IsDigit(char) && is_valid {
				num_to_add, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				valid_nums = append(valid_nums, num_to_add)
				is_valid = false
				num = ""
			} else {
				is_valid = false
				num = ""
			}
		}
		if num != "" && is_valid {
			num_to_add, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			valid_nums = append(valid_nums, num_to_add)
		}
	}
	return sum(valid_nums)
}

func compute2(data []string, max_len int) int {
	var valid_nums = make([]int, 1000)
	var num_count = 0
	var mapping = make(map[string]map[int]bool, 100)
	for i, line := range data {
		var num = ""
		var is_valid = false
		for j, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)
				is_valid = is_valid || has_adjacent_symbol2(data, i, j, max_len, mapping, num_count)
			}  else if !unicode.IsDigit(char) && is_valid {
				num_to_add, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				valid_nums[num_count] = num_to_add
				num_count++
				is_valid = false
				num = ""
			} else {
				is_valid = false
				num = ""
			}
		}
		if num != "" && is_valid {
			num_to_add, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			valid_nums[num_count] = num_to_add
			num_count++
		}
	}
	var total = 0
	for _, gears := range mapping {
		if len(gears) == 2 {
			var prod = 1;
			for gear, _ := range gears {
				prod *= valid_nums[gear] 
			}
			total += prod
		}
	}
	return total
}

func main() {
    // Get a greeting message and print it.
	var data = make([]string, 140);
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		data[i] = scanner.Text()
		i++
	}
	var output = compute(data, i)
	var output2 = compute2(data, i)
    fmt.Println("score: ", output)
	fmt.Println("score2: ", output2)
}