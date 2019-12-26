import argparse
import sys
import time

_SIZE = 9
_UNIT = int(_SIZE / 3)
_TOTAL_SIZE = _SIZE * _SIZE
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
    for i in range(_TOTAL_SIZE):
        square = grid_str[i]
        if square == ".":
            grid[r][c] = _UNASSIGNED
        else:
            grid[r][c] = int(square)
        c += 1
        if c == _SIZE:
            c = 0
            r += 1

def print_solution(grid):
    sys.stdout.write("solution: ")
    for r in range(_SIZE):
        for c in range(_SIZE):
            sys.stdout.write(str(grid[r][c]))
    print()

def read_grid_strs(filename):
    lines = []

    with open(filename, "r") as f:
        lines = [ l.strip() for l in f.readlines() ]

    for line in lines:
        if len(line) != _TOTAL_SIZE:
            print("ERROR: incorrect length %d, invalid line '%s'" % (len(line), line))
            sys.exit(1)

        for i in range(_TOTAL_SIZE):
            square = line[i]
            if square == ".":
                continue
            try:
                num = int(square)
                continue
            except ValueError:
                print("ERROR: incorrect character '%s' at index %d, invalid line '%s'" % (square, i, line))
                sys.exit(1)

    return lines

def main():
    sample_grid = [ [5, 3, 0, 0, 7, 0, 0, 0, 0], [6, 0, 0, 1, 9, 5, 0, 0, 0], [0, 9, 8, 0, 0, 0, 0, 6, 0], [8, 0, 0, 0, 6, 0, 0, 0, 3], [4, 0, 0, 8, 0, 3, 0, 0, 1], [7, 0, 0, 0, 2, 0, 0, 0, 6], [0, 6, 0, 0, 0, 0, 2, 8, 0], [0, 0, 0, 4, 1, 9, 0, 0, 5], [0, 0, 0, 0, 8, 0, 0, 7, 9] ]
    empty_grid = [ [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], [0, 0, 0, 0, 0, 0, 0, 0, 0], ]

    parser = argparse.ArgumentParser(description="Solve sudoku")
    parser.add_argument("--sample", action="store_true", help="Solve a sample puzzle then exit")
    parser.add_argument("--quiet", action="store_true", help="Quiet output")
    args = parser.parse_args()

    if args.sample:
        grid = sample_grid
        print("puzzle: sample")
        timed_solve(grid, False)
        print_solution(grid)
        sys.exit(0)

    grid_strs = read_grid_strs("sudoku.txt")

    count = len(grid_strs)

    start = nanotime()

    count = 0
    for grid_str in grid_strs:
        grid = empty_grid
        str_to_grid(grid, grid_str)
        if not args.quiet:
            print("puzzle: %s" % grid_str)
        timed_solve(grid, args.quiet)
        if not args.quiet:
            print_solution(grid)

        end = nanotime()

        count += 1

        if (end - start) > 10:
            break

    end = nanotime()

    diff = end - start

    print("puzzles: %d" % count)
    print("total time elapsed: %.6f" % diff)
            
    sys.exit(0)

main()
