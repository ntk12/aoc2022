import string

d = {k: v for k, v in zip(string.ascii_letters, range(1, 53))}


def solve(line: str):
    w = int(len(line) / 2)
    first = line[:w]
    second = line[w:]

    uniq = set(first).intersection(second)
    c = uniq.pop()
    return d[c]


def solve2(lines):
    s = map(set, lines)
    uniq = set.intersection(*s)
    c = uniq.pop()
    return d[c]


with open("./input2.txt") as f:
    s = 0
    group = []
    for l in f.readlines():
        group.append(l.strip())
        if len(group) == 3:
            s += solve2(group)
            group = list()
    print(s)
