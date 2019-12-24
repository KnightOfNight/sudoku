
all : csudoku gsudoku sudoku.class sudoku.pyc

csudoku : sudoku.c
	@cc -o csudoku -Wall sudoku.c

gsudoku : sudoku.go
	@go build -o gsudoku sudoku.go

sudoku.class: sudoku.java
	@javac sudoku.java

sudoku.pyc: sudoku.py
	@python3 -c "import py_compile; py_compile.compile('sudoku.py', cfile='sudoku.pyc', doraise=True)"

clean : 
	@echo "removing binaries"
	@rm -vf csudoku gsudoku sudoku.class sudoku.pyc 2>/dev/null

quicktest : all
	@echo
	@echo "quick test"
	@echo
	@echo "testing go"
	@./gsudoku --sample
	@echo
	@echo "testing c"
	@./csudoku --sample
	@echo
	@echo "testing java"
	@./jsudoku --sample
	@echo
	@echo "testing python"
	@./psudoku --sample
	@echo

longtest: all
	@echo
	@echo "long test"
	@echo
	@echo "testing go"
	@./gsudoku | grep -v "time elapsed" > g-sol.txt
	@echo
	@echo "testing c"
	@./csudoku | grep -v "time elapsed" > c-sol.txt
	@echo
	@echo "testing java"
	@./jsudoku | grep -v "time elapsed" > j-sol.txt
	@echo
	@echo "testing python"
	@./psudoku | grep -v "time elapsed" > p-sol.txt
	@echo
