
all : csudoku gsudoku sudoku.class

csudoku : sudoku.c
	@cc -o csudoku -Wall sudoku.c

gsudoku : sudoku.go
	@go build -o gsudoku sudoku.go

sudoku.class: sudoku.java
	@javac sudoku.java

clean : 
	@echo "removing binaries"
	@rm -vf csudoku gsudoku sudoku.class 2>/dev/null

testall : all
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

fullrunfast : all
	@echo
	@echo "testing go"
	@time ./gsudoku --quiet
	@echo
	@echo "testing c"
	@time ./csudoku --quiet
	@echo
	@echo "testing java"
	@time ./jsudoku --quiet
	@echo
