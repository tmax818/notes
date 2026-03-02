# this is the 'makefile' for chapter 1. Check it out!!!


first: rocks.c
	@cat info.txt

rocks.c:
	@echo "#include <stdio.h>\n\nint main(){puts(\"C Rocks!\");return 0;}" > rocks.c

cards:
	gcc cards.c -o cards && ./cards

clean:
	rm -rf cards a.out rocks.c
