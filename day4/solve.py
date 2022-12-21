def to_range(l: str):
    parts = l.split("-")
    return int(parts[0]), int(parts[1])


def fully_contain(line):
    parts = line.split(",")
    first, second = parts[0], parts[1]

    a, b = to_range(first)
    c, d = to_range(second)

    return (a >= c and b <= d) or (a <= c and b >= d)


def overlap(line):
    parts = line.split(",")
    first, second = parts[0], parts[1]

    a, b = to_range(first)
    c, d = to_range(second)

    return a <= d and c <= b

    # 64-64,2-64


s = 0
# with open("./input1.txt") as f:
with open("./input2.txt") as f:

    for l in f.readlines():
        l = l.strip()
        # if fully_contain(l):
        if overlap(l):
            s += 1
            print(l, "*")
        else:
            print(l)

print(s)
