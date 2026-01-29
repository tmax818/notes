# sample makefile

hello:
	echo "Hello, World"
	@echo "The echo didn't print for this one because I used the @ symbol!"
	@echo $^

target: prereq
	echo "this is the:" $@
	echo $^



prereq:
	echo "this is the" $@ "it will run first"



