

all:
	pandoc main.txt -s -o notes.md
	pandoc main.txt -s -o notes.org
	pandoc main.txt -s -o notes.tex


reveal:
	pandoc -t revealjs -s -o index.html habits.txt \
  -V revealjs-url=./reveal.js \
  -V theme=sky \
  -V transition=cube \
  --slide-level 3

beamer:
	pandoc -t beamer habits.txt -o habits.pdf


clean:
	rm -rf notes.*