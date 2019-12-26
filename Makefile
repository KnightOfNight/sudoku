
# binaries
binaries: csudoku gsudoku jsudoku psudoku

# reports
reports: c-sol.txt g-sol.txt j-sol.txt p-sol.txt

# everything 
all: binaries reports

# binary
csudoku: sudoku.c
	@echo "compiling csudoku"
	@cc -o csudoku -Wall sudoku.c
	@rm -f c-sol.txt 2>/dev/null
# binary
gsudoku: sudoku.go
	@echo "compiling gsudoku"
	@go build -o gsudoku sudoku.go
	@rm -f g-sol.txt 2>/dev/null
# alias
jsudoku: sudoku.class
# binary
sudoku.class: sudoku.java
	@echo "compiling jsudoku"
	@javac sudoku.java
	@rm -f j-sol.txt 2>/dev/null
# alias
psudoku: sudoku.pyc
# binary
sudoku.pyc: sudoku.py
	@echo "compiling psudoku"
	@python3 -c "import py_compile; py_compile.compile('sudoku.py', cfile='sudoku.pyc', doraise=True)"
	@rm -f p-sol.txt 2>/dev/null

# report file
c-sol.txt: csudoku
	@echo "getting csudoku report"
	@./csudoku > c-sol.txt
# report file
g-sol.txt: gsudoku
	@echo "getting gsudoku report"
	@./gsudoku > g-sol.txt
# report file
j-sol.txt: jsudoku
	@echo "getting jsudoku report"
	@./jsudoku > j-sol.txt
# report file
p-sol.txt: psudoku
	@echo "getting psudoku report"
	@./psudoku > p-sol.txt

# clean everything
clean: cleanbinaries cleanreports
# clean binaries
cleanbinaries:
	@echo "removing binaries"
	@rm -vf csudoku gsudoku sudoku.class sudoku.pyc 2>/dev/null
# clean reports
cleanreports: 
	@echo "removing reports"
	@rm -vf *-sol.txt 2>/dev/null

# alias
quicktest: sample
# sample run
sample: binaries
	@echo "sample puzzle test"
	@echo "language: c"
	@./csudoku --sample
	@echo "language: go"
	@./gsudoku --sample
	@echo "language: java"
	@./jsudoku --sample
	@echo "language: python"
	@./psudoku --sample
# summary report
report: all reports
	@echo "summary report"
	@echo "language: c"
	@tail -2 c-sol.txt
	@echo "language: go"
	@tail -2 g-sol.txt
	@echo "language: java"
	@tail -2 j-sol.txt
	@echo "language: python"
	@tail -2 p-sol.txt
# code report
code:
	@echo "code report"
	@echo "language: c"
	@egrep -e "^int" -e "^void" -e "^double" -e "^char" sudoku.c
	@egrep -e "^int" -e "^void" -e "^double" -e "^char" sudoku.c | wc -l
	@echo "language: go"
	@grep func sudoku.go
	@grep func sudoku.go | wc -l
	@echo "language: java"
	@grep "public static" sudoku.java
	@grep "public static" sudoku.java | wc -l
	@echo "language: python"
	@grep def sudoku.py
	@grep def sudoku.py | wc -l

# list all targets
list:
	@egrep -B 1 -e "^[a-zA-Z]" Makefile | cut -f 1 -d :
