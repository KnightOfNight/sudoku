import java.nio.file.*;

class sudoku {
    final static int SIZE = 9;
    final static int UNIT = SIZE / 3;
    final static int TOTAL_SIZE = SIZE * SIZE;
    final static int UNASSIGNED = 0;

    public static boolean in_row(int grid[][], int row, int num) {
        for (int col = 0; col < SIZE; col++) {
            if (grid[row][col] == num) {
                return(true);
            }
        }
        return(false);
    }

    public static boolean in_col(int grid[][], int col, int num) {
       for (int row = 0; row < SIZE; row++) {
            if (grid[row][col] == num) {
                return(true);
            }
        }
        return(false);
    }

    public static boolean in_unit(int grid[][], int row, int col, int num) {
        int start_row = row - (row % UNIT);
        int end_row = start_row + UNIT;
        int start_col = col - (col % UNIT);
        int end_col = start_col + UNIT;

        for (int r = start_row; r < end_row; r++) {
            for (int c = start_col; c < end_col; c++) {
                if (grid[r][c] == num) {
                    return(true);
                }
            }
        }

        return(false);
    }

    public static int[] first_empty_square(int grid[][]) {
        for (int row = 0; row < SIZE; (row)++) {
            for (int col = 0; col < SIZE; (col)++) {
                if (grid[row][col] == 0) {
                    int rc[] = { row, col };
                    return rc;
                }
            }
        }

        int rc[] = { -1, -1 };
        return rc;
    }

    public static boolean solve(int grid[][]) {
        int rc[] = first_empty_square(grid);
        if (rc[0] == -1) {
            return(true);
        }
        int row = rc[0];
        int col = rc[1];

        for (int num = 1; num <= SIZE; num++ ) {
            if (in_row(grid, row, num) || in_col(grid, col, num) || in_unit(grid, row, col, num)) {
                continue;
            }

            grid[row][col] = num;

            if (solve(grid)) {
                return(true);
            }

            grid[row][col] = UNASSIGNED;
        }

        return(false);
    }

    public static double nanotime() {
        long time = System.nanoTime();

        double seconds = time / 1000000000.0;

        return(seconds);
    }

    public static void timed_solve(int grid[][], boolean quiet) {
        double start = nanotime();

        if (! solve(grid)) {
            System.out.printf("no solution\n");
            return;
        }

        double end = nanotime();

        double diff = end - start;

        if (! quiet) {
            System.out.printf("time elapsed: %.6f\n", diff);
        }
    }

    public static void str_to_grid(int grid[][], String grid_str) {
        int r = 0;
        int c = 0;

        for (int i = 0; i < TOTAL_SIZE; i++) {
            char square = grid_str.charAt(i);
            if (square == '.') {
                grid[r][c] = UNASSIGNED;
            } else {
                grid[r][c] = Character.getNumericValue(square);
            }
            c++;
            if (c == SIZE) {
                c = 0;
                r++;
            }
        }
    }

    public static void print_solution(int grid[][]) {
        System.out.printf("solution: ");
        for (int r = 0; r < SIZE; r++) {
            for (int c = 0; c < SIZE; c++) {
                System.out.printf("%d", grid[r][c]);
            }
        }
        System.out.printf("\n");
    }

    public static String[] read_grid_strs(String filename) throws Exception {
        String line = new String(Files.readAllBytes(Paths.get(filename)));

        line = line.trim();

        String lines[] = line.split("\n");

        for (int l = 0; l < lines.length; l++) {
            line = lines[l];

            if (line.length() != TOTAL_SIZE) {
                System.out.printf("ERROR: incorrect length %d, invalid line '%s'\n", line.length(), line);
                System.exit(1);
            }

            for (int i = 0; i < TOTAL_SIZE; i++) {
                char square = line.charAt(i);

                if (square == '.') {
                    continue;
                }

                int num = Character.getNumericValue(square);
                if ((num < 1) || (num > SIZE)) {
                    System.out.printf("ERROR: incorrect character '%c' at index %d, invalid line '%s'\n", square, i, line);
                    System.exit(1);
                }
            }
        }

        return(lines);
    }

    public static void main(String args[]) {
        boolean sample = false;
        boolean quiet = false;

        int sample_grid[][] = {
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

        int grid[][] = {
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

        if (args.length > 0) {
            if (args[0].equals("--sample")) {
                sample = true;
            } else if (args[0].equals("--quiet")) {
                quiet = true;
            }
        }

        if (sample) {
            System.out.printf("puzzle: sample\n");
            timed_solve(sample_grid, false);
            print_solution(sample_grid);
            return;
        }

        String grid_strs[];
        try {
            grid_strs = read_grid_strs("sudoku.txt");
        } catch(Exception e) {
            return;
        }

        int count = grid_strs.length;

        double start = nanotime();

        for (int i = 0; i < grid_strs.length; i++) {
            String line = grid_strs[i];
            str_to_grid(grid, line);
            if (! quiet) {
                System.out.printf("puzzle: %s\n", line);
            }
            timed_solve(grid, quiet);
            if (! quiet) {
                print_solution(grid);
            }
        }

        double end = nanotime();

        double diff = end - start;

        System.out.printf("puzzles: %d\n", count);
        System.out.printf("total time elapsed: %.6f\n", diff);

        return;
    }
}
