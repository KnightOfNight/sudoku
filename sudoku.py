import argparse
import sys
import time

_SIZE = 9
_UNIT = 3
_UNASSIGNED = 0

def in_row(grid, row, num):
    return num in grid[row]

def in_col(grid, col, num):
    return num in [ x[col] for x in grid ]

def in_unit(grid, row, col, num):
    unit_row = row - (row % _UNIT)
    unit_col = col - (col % _UNIT)
    for r in range(unit_row, unit_row + _UNIT):
        if num in grid[r][unit_col:unit_col + _UNIT]:
            return True
    return False

def first_empty_square(grid):
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] == _UNASSIGNED:
                return (r, c)
    return (None, None)

def solve(grid):
    (r, c) = first_empty_square(grid)

    if r is None:
        return True

    for num in range(1, _SIZE + 1):
        if in_row(grid, r, num) or in_col(grid, c, num) or in_unit(grid, r, c, num):
            continue

        grid[r][c] = num

        if solve(grid):
            return True

        grid[r][c] = _UNASSIGNED

    return False

def nanotime():
    return time.monotonic_ns() / 1000000000

def timed_solve(grid, quiet):
    start = nanotime()

    if not solve(grid):
        print("no solution")
        return

    end = nanotime()

    diff = end - start

    if not quiet:
        print("time elapsed: %.6f" % diff)

def str_to_grid(grid, grid_str):
    r = 0
    c = 0
    for i in range(_SIZE * _SIZE):
        if grid_str[i] == ".":
            grid[r][c] = _UNASSIGNED
        else:
            grid[r][c] = int(grid_str[i])
        c += 1
        if c == _SIZE:
            c = 0
            r += 1

def main():
    sample_grid = [
        [5, 3, 0, 0, 7, 0, 0, 0, 0],
        [6, 0, 0, 1, 9, 5, 0, 0, 0],
        [0, 9, 8, 0, 0, 0, 0, 6, 0],
        [8, 0, 0, 0, 6, 0, 0, 0, 3],
        [4, 0, 0, 8, 0, 3, 0, 0, 1],
        [7, 0, 0, 0, 2, 0, 0, 0, 6],
        [0, 6, 0, 0, 0, 0, 2, 8, 0],
        [0, 0, 0, 4, 1, 9, 0, 0, 5],
        [0, 0, 0, 0, 8, 0, 0, 7, 9]
    ]

    grid = [
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
    ]

    parser = argparse.ArgumentParser(description="Solve sudoku")
    parser.add_argument("--sample", action="store_true", help="Solve a sample puzzle then exit")
    parser.add_argument("--quiet", action="store_true", help="Quiet output")
    args = parser.parse_args()

    if args.sample:
        print("puzzle: sample")
        timed_solve(sample_grid, False)
        sys.exit(0)

    with open("sudoku.txt", "r") as f:
        for line in f.readlines():
            str_to_grid(grid, line)
            if not args.quiet:
                print("puzzle: %s" % line.strip())
            timed_solve(grid, args.quiet)
            
    sys.exit(0)

main()
