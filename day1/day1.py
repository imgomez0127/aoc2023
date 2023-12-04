import string
import re

digits = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']
digit_map = {k:str(i) for i, k in enumerate(digits)}
digit_map.update({str(i):str(i) for i in range(10)})

def parse(path):
    with open(path, 'r') as f:
        return [[s for s in line if s in string.digits] for line in f]

def parse2(path):
    num_regex = re.compile(r'([1-9]|one|two|three|four|five|six|seven|eight|nine)')
    with open(path, 'r') as f:
        return [
            [
                digit_map[num_regex.match(line[i:]).group(0)] 
                for i in range(len(line)) if num_regex.match(line[i:])
            ]
            for line in f
        ]

def get_val(encoding):
    return sum([
        int(line[0] + line[-1])
        for line in encoding
    ])

print(get_val(parse('input.txt')))
print(get_val(parse2('input.txt')))