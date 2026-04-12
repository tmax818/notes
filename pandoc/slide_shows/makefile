all:
	pandoc -t revealjs -s -o index.html index.md -V revealjs-url=./reveal.js --slide-level 3


test:
	pandoc -t revealjs -s -o index.html index.md \
  -V revealjs-url=./reveal.js \
  -V theme=sky \
  -V transition=cube \
  --slide-level 3