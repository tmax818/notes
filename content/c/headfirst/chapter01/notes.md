

# C is a language for small, fast programs

The C language is designed to create small, fast programs. It’s
lower-level than most other languages; that means it creates code
that’s a lot closer to what machines really understand.

## the way C works

Computers really only understand one language: machine code, a
binary stream of 1s and 0s. You convert your C code into machine
code with the aid of a compiler.

[[image]]

C is used where speed, space, and
portability are important. Most
operating systems are written in C.
Most other computer languages are
also written in C. And most game
software is written in C.

There are three C standards that you may
stumble across. ANSI C is from the late 1980s
and is used for the oldest code. A lot of things
were fixed up in the C99 standard from 1999. And
some cool new language features were added in
the current standard, C11, released in 2011. The
differences between the different versions aren’t
huge, and we’ll point them out along the way.

- what does the following code do?

```c

void problem_one(){
    int card_count = 11;
    if (card_count > 10) 
    puts("The deck is hot. Increase bet.");
}

void problem_two() {
    int c = 10;
    while (c > 0) {
        puts("I must not write code in class");
        c = c - 1;
    }
}

void problem_three() {
    /* Assume name shorter than 20 chars. */
    char ex[20];
    puts("Enter boyfriend's name: ");
    scanf("%19s", ex);
    printf("Dear %s.\n\n\tYou're history.\n", ex);
}

void problem_four() {
    char suit = 'H';
    switch(suit) {
    case 'C':
        puts("Clubs");
        break;
    case 'D':
        puts("Diamonds");
        break;
    case 'H':
        puts("Hearts");
        break;
    default:
        puts("Spades");
    }
}

```

# But what does a complete C program look like?

To create a full program, you need to enter your code into a C source file. C source files can be created by any text editor, and their filenames usually end with `.c`.

## C programs normally begin with a comment.
The comment describes the purpose of the code in the file, and might
include some license or copyright information. There’s no absolute need
to include a comment here—or anywhere else in the file—but it’s good
practice and what most C programmers will expect to find.


## Next comes the include section.
C is a very, very small
language and it can do
almost nothing without
the use of external
libraries. You will need
to tell the compiler what
external code to use by
including header files
for the relevant libraries.
The header you will see
more than any other
is stdio.h. The stdio
library contains code
that allows you to read
and write data from and
to the terminal.

## The last thing you find in a source file are the functions.
All C code runs inside functions. The most important function you will find in any C program is called the **main() function**. The `main()` function is the starting point for all of the code in your program.

## the `main()` function up close

The computer will start running your program from the `main()`
function. The name is important: if you don’t have a function called `main()`, your program won’t be able to start.

The `main()` function has a return type of `int`. So what does this mean?
Well, when the computer runs your program, it will need to have some way of deciding if the program ran successfully or not. It does this by checking the return value of the main() function. If you tell your main() function to return 0, this means that the program was successful. If you tell it to return any other value, this means that there was a problem.

# But how do you run the program?

C is a compiled language. That means the computer will not interpret the code directly. Instead, you will need to convert—or compile—the human-readable source code into machine-readable machine code.

To compile the code, you need a program called a compiler. One of the most popular C compilers is the *GNU Compiler Collection* or **gcc**. **gcc** is available on a lot of operating systems, and it can compile lots of languages other than C. Best of all, it’s completely free.

Here’s how you can compile and run the program using **gcc**:


## The C language doesn’t support strings out of the box.

C is more low-level than most other languages, so instead
of strings, it normally uses something similar: an array of
single characters. If you’ve programmed in other languages,
you’ve probably met an array before. An array is just a list of
things given a single name. So card_name is just a variable
name you use to refer to the list of characters entered at
the command prompt. You defined card_name to be a
two-character array, so you can refer to the first and second
character as char_name[0] and char_name[1]. To see
how this works, let’s take a deeper dive into the computer’s
memory and see how C handles text…




## strings way up close

Strings are just character arrays. When C sees a string like this:

```c
s = "Shatner"
```
This is how you define an array in C.
it reads it like it was just an array of separate characters:

```c
s = {'S', 'h', 'a', 't', 'n', 'e', 'r'}
```

Each of the characters in the string is just an element in an array, which is why you can refer to the individual characters in the string by using an index,
like s[0] and s[1].

But what happens when C wants to read the contents of the string? Say
it wants to print it out. Now, in a lot of languages, the computer keeps
pretty close track of the size of an array, but C is more low-level than most
languages and can’t always work out exactly how long an array is. If C is going
to display a string on the screen, it needs to know when it gets to the end of
the character array. And it does this by adding a sentinel character.
'\0'
C knows
to stop
when it
sees \0.
The sentinel character is an additional character at the end of the string that
has the value \0. Whenever the computer needs to read the contents of the
string, it goes through the elements of the character array one at a time, until
it reaches \0. 
\0 is the ASCII character
with value 0.
That’s why in our code we had to define the card_name variable like this:

```c
char card_name[3];
```

The card_name string is only ever going to record one or two characters, but because
strings end in a sentinel character we have to allow for an extra character in the array.

# Two types of command

## do something

## do something only if something is true

# Here’s the code so far


```c
/*
* Program to evaluate face values.
* Released under the Vegas Public License.
* (c)2014 The College Blackjack Team.
*/
#include <stdio.h>
#include <stdlib.h>

int main(){
    char card_name[3];
    puts("Enter the card_name: ");
    scanf("%2s", card_name);
    int val = 0;

    if (card_name[0] == 'K') {
        val = 10;
    } else if (card_name[0] == 'Q') {
        val = 10;
    } else if (card_name[0] == 'J') {
        val = 10;
    } else if (card_name[0] == 'A') {
        val = 11;
    } else {
        val = atoi(card_name);
    }

printf("The card value is: %i\n", val);
return 0;

}


```

# Card counting? In C?

# There’s more to booleans than equals…

## `&&` checks if two conditions are true

## `||` checks if one of two conditions is true

## `!` flips the value of a condition


# What’s the code like now?

```c

int main() {
char card_name[3];
puts("Enter the card_name: ");
scanf("%2s", card_name);
int val = 0;
if (card_name[0] == 'K') {
val = 10;
} else if (card_name[0] == 'Q') {
val = 10;
} else if (card_name[0] == 'J') {
val = 10;
} else if (card_name[0] == 'A') {
val = 11;
} else {
val = atoi(card_name);
}
/* Check if the value is 3 to 6 */
if ((val > 2) && (val < 7))
puts("Count has gone up");
/* Otherwise check if the card was 10, J, Q, or K */
else if (val == 10)
puts("Count has gone down");
return 0;
}

```

# Pulling the ol’ switcheroo

# Sometimes once is not enough...

## using while loops in C

# Loops often follow the same structure...

## ...and the `for` loop makes this easy

# You use break to break out...

## ...and `continue` to continue

# Your C Toolbox