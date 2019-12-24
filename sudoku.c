#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/time.h>

#define SIZE 9
#define UNIT 3
#define UNASSIGNED 0

int in_row(int grid[SIZE][SIZE], int row, int num) {
    for (int col = 0; col < SIZE; col++) {
        if (grid[row][col] == num) {
            return(1);
        }
    }
    return(0);
}

int in_col(int grid[SIZE][SIZE], int col, int num) {
    for (int row = 0; row < SIZE; row++) {
        if (grid[row][col] == num) {
            return(1);
        }
    }
    return(0);
}

int in_unit(int grid[SIZE][SIZE], int row, int col, int num) {
    int start_row = row - (row % UNIT);
    int end_row = start_row + UNIT;
    int start_col = col - (col % UNIT);
    int end_col = start_col + UNIT;

    for (int r = start_row; r < end_row; r++) {
        for (int c = start_col; c < end_col; c++) {
            if (grid[r][c] == num) {
                return(1);
            } 
        }
    }
    return(0);
}

int first_empty_square(int grid[SIZE][SIZE], int *row, int *col) {
    for (*row = 0; *row < SIZE; (*row)++) {
        for (*col = 0; *col < SIZE; (*col)++) {
            if (grid[*row][*col] == 0) {
                return(1);
            }
        }
    }
    return(0);
}

int solve(int grid[SIZE][SIZE]) {
    int row = 0;
    int col = 0;
    
    if (! first_empty_square(grid, &row, &col)) {
        return(1);
    }
    
    for (int num = 1; num <= SIZE; num++ ) {
        if (in_row(grid, row, num) || in_col(grid, col, num) || in_unit(grid, row, col, num)) {
            continue;
        }

        grid[row][col] = num;
            
        if (solve(grid)) {
            return(1);
        }
            
        grid[row][col] = UNASSIGNED;
    }
    
    return(0);
}

double nanotime() {
    struct timespec tsp;
    double secs;

    clock_gettime(CLOCK_MONOTONIC, &tsp);

    secs = (double)tsp.tv_sec + ((double)tsp.tv_nsec / 1000000000.0);

    return(secs);
}

void timed_solve(int grid[SIZE][SIZE], int quiet) {
    double start, end, diff;

    start = nanotime();

    if (! solve(grid)) {
        printf("no solution\n");
        return;
    }

    end = nanotime();

    diff = end - start;

    if (! quiet) {
        printf("time elapsed: %.6f\n", diff);
    }
}

void str_to_grid(int grid[SIZE][SIZE], char grid_str[256]) {
    int r = 0;
    int c = 0;

    for (int i = 0; i < (SIZE * SIZE); i++) {
        if (grid_str[i] == '.') {
            grid[r][c] = UNASSIGNED;
        } else {
            grid[r][c] = grid_str[i] - '0';
        }
        c++;
        if (c == SIZE) {
            c = 0;
            r++;
        }
    }
}

int main(int argc, char *argv[]) {
    int sample_grid[SIZE][SIZE] = {
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9}
    };
    int grid[SIZE][SIZE] = {
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0}
    };
    FILE* file;
    char line[256];
    int sample = 0;
    int quiet = 0;

    if (argc > 1) {
        if (! strcmp(argv[1], "--sample")) {
            sample = 1;
        } else if (! strcmp(argv[1], "--quiet")) {
            quiet = 1;
        }
    }

    if (sample) {
        printf("puzzle: sample\n");
        timed_solve(sample_grid, 0);
        return(0);
    }

    file = fopen("sudoku.txt", "r");
    if (file == NULL) {
        printf("ERROR: unable to open file\n");
        return(1);
    }
    while (fgets(line, sizeof(line), file)) {
        str_to_grid(grid, line);
        if (! quiet) {
            printf("puzzle: %s", line);
        }
        timed_solve(grid, quiet);
    }

    return(0);
}
