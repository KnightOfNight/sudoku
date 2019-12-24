package main

import (
    "fmt"
    "time"
    "os"
    "strings"
    "io/ioutil"
    "strconv"
)

const _SIZE = 9
const _UNIT = 3
const _UNASSIGNED = 0

func in_row(grid [][]int, r int, num int) (bool) {
    for c,_ := range grid[r] {
        if grid[r][c] == num {
            return true
        }
    }
    return false
}

func in_col(grid [][]int, c int, num int) (bool) {
    for r,_ := range grid {
        if grid[r][c] == num {
            return true
        }
    }
    return false
}

func in_unit(grid [][]int, r int, c int, num int) (bool) {
    unit_row := r - (r % _UNIT)
    unit_col := c - (c % _UNIT)
    for ur := unit_row; ur < (unit_row + _UNIT); ur++ {
        for uc := unit_col; uc < (unit_col + _UNIT); uc++ {
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
            if grid[r][c] == _UNASSIGNED {
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

        grid[r][c] = _UNASSIGNED
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
    for i := 0; i < (_SIZE * _SIZE); i++ {
        square := string(grid_str[i])
        if square == "." {
            grid[r][c] = _UNASSIGNED
        } else {
            num,_ := strconv.Atoi(square)
            grid[r][c] = num
        }
        c += 1
        if c == _SIZE {
            c = 0
            r += 1
        }
    }
}

func print_solution(grid [][]int) {
    fmt.Printf("solution: ")
    for r := 0; r < _SIZE; r++ {
        for c := 0; c < _SIZE; c++ {
            fmt.Printf("%d", grid[r][c])
        }
    }
    fmt.Printf("\n")
}

func readlines(filename string) []string {
    bytes,_ := ioutil.ReadFile(filename)
    lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
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

    lines := readlines("sudoku.txt")

    for _,line := range lines {
        str_to_grid(grid, line)
        if ! quiet {
            fmt.Printf("puzzle: %s\n", line)
        }
        timed_solve(grid, quiet)
        if ! quiet {
            print_solution(grid)
        }
    }

    os.Exit(0)
}
