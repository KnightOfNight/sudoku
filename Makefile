
all: csudoku gsudoku sudoku.class sudoku.pyc

reports: all c-sol.txt g-sol.txt j-sol.txt

forcereports: cleanreports reports

clean: cleanbin cleanreports

cleanbin:
	@echo "removing binaries"
	@rm -vf csudoku gsudoku sudoku.class sudoku.pyc 2>/dev/null

cleanreports: 
	@echo "removing reports"
	@rm -vf *-sol.txt 2>/dev/null

csudoku: sudoku.c
	@echo "compiling csudoku"
	@cc -o csudoku -Wall sudoku.c

gsudoku: sudoku.go
	@echo "compiling gsudoku"
	@go build -o gsudoku sudoku.go

sudoku.class: sudoku.java
	@echo "compiling jsudoku"
	@javac sudoku.java

sudoku.pyc: sudoku.py
	@echo "compiling psudoku"
	@python3 -c "import py_compile; py_compile.compile('sudoku.py', cfile='sudoku.pyc', doraise=True)"

c-sol.txt:
	@echo "running csudoku"
	@./csudoku > c-sol.txt

g-sol.txt:
	@echo "running gsudoku"
	@./gsudoku > g-sol.txt

j-sol.txt:
	@echo "running jsudoku"
	@./jsudoku > j-sol.txt

sample: all
	@echo "sample puzzle tested on all languages"
	@echo "language: c"
	@./csudoku --sample
	@echo "language: go"
	@./gsudoku --sample
	@echo "language: java"
	@./jsudoku --sample
	@echo "language: python"
	@./psudoku --sample

summary: all reports
	@echo "summary on (fast) languages"
	@echo "language: c"
	@tail -2 c-sol.txt
	@echo "language: go"
	@tail -2 g-sol.txt
	@echo "language: java"
	@tail -2 j-sol.txt

