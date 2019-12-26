package main

import (
    "fmt"
    "time"
    "os"
    "strings"
    "io/ioutil"
    "strconv"
)

const SIZE = 9
const UNIT = SIZE / 3
const TOTAL_SIZE = SIZE * SIZE
const UNASSIGNED = 0

func in_row(grid [][]int, r int, num int) bool {
    for c,_ := range grid[r] {
        if grid[r][c] == num {
            return true
        }
    }
    return false
}

func in_col(grid [][]int, c int, num int) bool {
    for r,_ := range grid {
        if grid[r][c] == num {
            return true
        }
    }
    return false
}

func in_unit(grid [][]int, r int, c int, num int) bool {
    unit_row := r - (r % UNIT)
    unit_col := c - (c % UNIT)
    for ur := unit_row; ur < (unit_row + UNIT); ur++ {
        for uc := unit_col; uc < (unit_col + UNIT); uc++ {
            if grid[ur][uc] == num {
                return true
            }
        }
    }
    return false
}

func first_empty_square(grid [][]int) (int, int) {
    for r,row := range grid {
        for c,_ := range row {
            if grid[r][c] == UNASSIGNED {
                return r, c
            }
        }
    }

    return -1, -1
}

func solve(grid [][]int) bool {
    r, c := first_empty_square(grid)

    if r == -1 {
        return true
    }

    for num := 1; num < 10; num++ {
        if in_row(grid, r, num) || in_col(grid, c, num) || in_unit(grid, r, c, num) {
            continue
        }

        grid[r][c] = num

        if solve(grid) {
            return true
        }

        grid[r][c] = UNASSIGNED
    }

    return false
}

func nanotime() float64 {
    return float64(time.Now().UnixNano()) / 1000000000.0
}

func timed_solve(grid [][]int, quiet bool) {
    start := nanotime()

    if ! solve(grid) {
        fmt.Println("no solution")
        return
    }

    end := nanotime()

    diff := end - start

    if ! quiet {
        fmt.Printf("time elapsed: %.6f\n", diff)
    }
}

func str_to_grid(grid [][]int, grid_str string) {
    r := 0
    c := 0
    for i := 0; i < TOTAL_SIZE; i++ {
        square := string(grid_str[i])
        if square == "." {
            grid[r][c] = UNASSIGNED
        } else {
            num,_ := strconv.Atoi(square)
            grid[r][c] = num
        }
        c += 1
        if c == SIZE {
            c = 0
            r += 1
        }
    }
}

func print_solution(grid [][]int) {
    fmt.Printf("solution: ")
    for r := 0; r < SIZE; r++ {
        for c := 0; c < SIZE; c++ {
            fmt.Printf("%d", grid[r][c])
        }
    }
    fmt.Printf("\n")
}

func read_grid_strs(filename string) []string {
    bytes,_ := ioutil.ReadFile(filename)

    line := strings.TrimSpace(string(bytes))

    lines := strings.Split(line, "\n")

    for l := 0; l < len(lines); l++ {
        line := lines[l]

        if len(line) != TOTAL_SIZE {
            fmt.Printf("ERROR: incorrect length %d, invalid line '%s'\n", len(line), line)
            os.Exit(1)
        }

        for i := 0; i < TOTAL_SIZE; i++ {
            square := string(line[i])

            if square == "." {
                continue
            }

            num,_ := strconv.Atoi(square)
            if num == 0 {
                fmt.Printf("ERROR: incorrect character '%s' at index %d, invalid line '%s'\n", square, i, line)
                os.Exit(1)
            }
        }
    }

    return lines
}

func main() {
    sample_grid := [][]int {
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }

    grid := [][]int {
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
    }

    args := os.Args[1:]
    sample := false;
    quiet := false;
    if len(args) > 0 {
        if args[0] == "--sample" {
            sample = true;
        } else if args[0] == "--quiet" {
            quiet = true;
        }
    }

    if sample {
        fmt.Printf("puzzle: sample\n")
        timed_solve(sample_grid, false)
        print_solution(sample_grid)
        os.Exit(0)
    }

    grid_strs := read_grid_strs("sudoku.txt")

    count := len(grid_strs)

    start := nanotime()

    for _,line := range grid_strs {
        str_to_grid(grid, line)
        if ! quiet {
            fmt.Printf("puzzle: %s\n", line)
        }
        timed_solve(grid, quiet)
        if ! quiet {
            print_solution(grid)
        }
    }

    end := nanotime()

    diff := end - start;

    fmt.Printf("puzzles: %d\n", count);
    fmt.Printf("total time elapsed: %.6f\n", diff);

    os.Exit(0)
}
