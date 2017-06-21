OUTFILE=some-go-pointers.pdf
INFILE=some-go-pointers.md

all:
	pandoc -t beamer -o ${OUTFILE} ${INFILE}

.PHONY: all
