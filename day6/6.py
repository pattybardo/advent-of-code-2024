from itertools import cycle

path = "test2.txt"
# path = "test.txt"

with open(path) as f:
    lines = [l.strip() for l in f.readlines()]

ROWS = len(lines)
COLS = len(lines[0])

directions = {">": (0, 1), "v": (1, 0), "<": (0, -1), "^": (-1, 0)}

loop = cycle(directions)


def reset_loop(loop, char):
    while next(loop) != char:
        continue


initial_pos = ()
d = ()
for i in range(ROWS):
    for j in range(COLS):
        if lines[i][j] in directions:
            initial_pos = (i, j)
            d = directions[lines[i][j]]

            reset_loop(loop, lines[i][j])


def is_inside(lines, pos):
    return 0 <= pos[0] < ROWS and 0 <= pos[1] < COLS


def walk(lines, pos, d, obstacle, d_loop):
    i, j = pos
    x, y = d

    new_pos = (x + i, y + j)
    if is_inside(lines, new_pos) and (
        lines[new_pos[0]][new_pos[1]] == "#" or new_pos == obstacle
    ):
        return pos, directions[next(d_loop)]

    return new_pos, d


def evaluate_path(lines, initial_pos, initial_d, obstacle, d_loop):
    inside = True
    in_loop = False
    pos = initial_pos
    d = initial_d
    walked = set()
    loop_detec = set()
    while inside and not in_loop:
        walked.add(pos)
        loop_detec.add((pos, d))
        pos, d = walk(lines, pos, d, obstacle, d_loop)

        in_loop = (pos, d) in loop_detec
        inside = is_inside(lines, pos)

    return len(walked), in_loop


print("p1: ", evaluate_path(lines, initial_pos, d, None, loop)[0])

# P2 brute force?
p2 = 0
for row in range(ROWS):
    for col in range(COLS):
        reset_loop(loop, lines[initial_pos[0]][initial_pos[1]])

        size, in_loop = evaluate_path(lines, initial_pos, d, (row, col), loop)

        p2 += 1 if in_loop else 0
        print("{" + f"{col} {row}" + "}") if in_loop else None
        percentage = int((row * ROWS + col) / (ROWS * COLS) * 100)
