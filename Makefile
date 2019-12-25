
all: csudoku gsudoku sudoku.class sudoku.pyc

reports: all c-sol.txt g-sol.txt j-sol.txt

forcereports: cleanreports reports

clean: cleanbin cleanreports

cleanbin:
	@echo "removing binaries"
	@rm -vf csudoku gsudoku sudoku.class sudoku.pyc 2>/dev/null

cleanreports: 
	@echo "removing reports"
	@rm -f *-sol.txt 2>/dev/null

csudoku: sudoku.c
	@echo "compiling csudoku"
	@cc -o csudoku -Wall sudoku.c
	@rm -f c-sol.txt 2>/dev/null

gsudoku: sudoku.go
	@echo "compiling gsudoku"
	@go build -o gsudoku sudoku.go
	@rm -f g-sol.txt 2>/dev/null

sudoku.class: sudoku.java
	@echo "compiling jsudoku"
	@javac sudoku.java
	@rm -f j-sol.txt 2>/dev/null

sudoku.pyc: sudoku.py
	@echo "compiling psudoku"
	@python3 -c "import py_compile; py_compile.compile('sudoku.py', cfile='sudoku.pyc', doraise=True)"
	@rm -f p-sol.txt 2>/dev/null

c-sol.txt:
	@echo "getting csudoku report"
	@./csudoku > c-sol.txt

g-sol.txt:
	@echo "getting gsudoku report"
	@./gsudoku > g-sol.txt

j-sol.txt:
	@echo "getting jsudoku report"
	@./jsudoku > j-sol.txt

sample: all
	@echo "sample puzzle test: all languages"
	@echo "language: c"
	@./csudoku --sample
	@echo "language: go"
	@./gsudoku --sample
	@echo "language: java"
	@./jsudoku --sample
	@echo "language: python"
	@./psudoku --sample

summary: reports
	@echo "report summary:  fast languages"
	@echo "language: c"
	@tail -2 c-sol.txt
	@echo "language: go"
	@tail -2 g-sol.txt
	@echo "language: java"
	@tail -2 j-sol.txt

