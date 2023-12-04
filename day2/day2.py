from math import prod
m = {
    'blue': 0,
    'green': 1,
    'red': 2
}
def parse(path):
    with open(path, 'r') as f:
        return [(line.strip().split(':')[0].split(' '), [[z.strip().split(' ') for z in y.strip().split(',')] for y in line.strip().split(':')[1].strip().split(';')]) for line in f ]

def compute(rep):
    caps = [14, 13, 12]
    game = []
    for line in rep: 
        total_counts = [0, 0, 0]
        for pull in line[1]:
            pull_counts = [0, 0, 0]
            for item in pull:
                color = m[item[1].strip()]
                count = int(item[0].strip())
                pull_counts[color] += count
            for i, x in enumerate(zip(total_counts, pull_counts)):
                total_counts[i] = max(x)
        is_valid = True
        for cap, count in zip(caps, total_counts):
            is_valid &= cap >= count
        game.append(is_valid)
    return sum([(i+1) * validity for i, validity in enumerate(game)])
    
def compute2(rep):
    caps = [14, 13, 12]
    game = []
    for line in rep: 
        total_counts = [0, 0, 0]
        for pull in line[1]:
            pull_counts = [0, 0, 0]
            for item in pull:
                color = m[item[1].strip()]
                count = int(item[0].strip())
                pull_counts[color] += count
            for i, x in enumerate(zip(total_counts, pull_counts)):
                total_counts[i] = max(x)
        is_valid = True
        game.append(prod(total_counts))
    return sum(game)

r = parse('input.txt')
print(compute(r))
print(compute2(r))