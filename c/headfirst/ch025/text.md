There’s more to strings than reading them.

You’ve seen how strings in C are actually char arrays but what does C allow you to do with them? That’s where `string.h` comes in. string.h is p art of the C Standard Library that’s dedicated to string manipulation. If you want to concatenate strings together, copy one string to another, or compare two strings, the functions in string.h are there to help. In this chapter, you’ll see how to create an array of strings, and then take a close look at how to search within strings using the `strstr()` function.

# Desperately seeking Frank 84

There are so many tracks on the retro jukebox that people can’t find
the music they are looking for. To help the customers, the guys in the
Head First Lounge want you to write another program.

This is the track list:

Tracks from the new album “Little Known Sinatra.
Track list:
I left my heart in Harvard Med School
Newark, Newark - a wonderful town
Dancing with a Dork
From here to maternity
The girl from Iwo Jima

The guys say that there will be lots more
tracks in the future, but they’ll never be
more than 79 characters long.

The list is likely to get longer, so there’s just the first few tracks for
now. You’ll need to write a C program that will ask the user which
track she is looking for, and then get it to search through all of the
tracks and display any that match.

There’ll be lots of strings in this program. How do you think you can
record that information in C?

# Create an array of arrays 85

There are several track names that you need to record. You can record
several things at once in an array. But remember: each string is itself an
array. That means you need to create an array of arrays, like this:



That means that you’ll be able to find an individual track name
like this:


But you can also read the individual characters of each of the
strings if you want to:


So now that you know how to record the data in C, what do you
need to do with it?

# Find strings containing the search text 86

The guys have helpfully given you a spec.

Well, you know how to record the tracks. You also know how to
read the value of an individual track name, so it shouldn’t be too
difficult to loop through each of them. You even know how to ask
the user for a piece of text to search for. But how do you look to
see if the track name contains a given piece of text?

## Using string.h

The C Standard Library is a bunch of useful code that you
get for free when you install a C compiler. The library code
does useful stuff like opening files, or doing math, or managing
memory. Now, chances are, you won’t want to use the whole of the
Standard Library at once, so the library is broken up into several
sections, and each one has a header file. The header file lists all
of the functions that live in a particular section of the library.


So far, you have only really used the stdio.h header file. stdio.h lets
you use the standard input/output functions like printf and
`scanf`.

But the Standard Library also contains code to process strings.
String processing is required by a lot of the programs, and the
string code in the Standard Library is tested, stable, and fast.



You include the string code into your program using the string.h
header file. You add it at the top of your program, just like you
include stdio.h.

#include <stdio.h>
#include <string.h>



# Using the `strstr()` function 89

So how exactly does the strstr() function work? Let’s look at an
example. Let’s say you’re looking for the string “fun” inside a larger
string, “dysfunctional.” You’d call it like this:

The strstr() function will search for the second string in the first
string. If it finds the string, it will return the address of the located
string in memory. In the example here, the function would find that
the fun substring begins at memory location 4,000,003.
But what if the strstr() can’t find the substring? What then? In
that case, strstr() returns the value 0. Can you think why that
is? Well, if you remember, C treats zero as false. That means you
can use strstr() to check for the existence of one string inside
another, like this:

char s0[] = "dysfunctional";
char s1[] = "fun";
if (strstr(s0, s1))
puts("I found the fun in dysfunctional!");

Let’s see how we can use strstr() in the jukebox program.



# It’s time for a code review 94

Let’s bring the code together and review what you’ve got so far:

You still need stdio.h for the
printf() and scanf() functions.
You’ll set the tracks array
outside of the main() and
find_track() functions; that
way, the tracks will be usable
everywhere in the program.
This is your new find_track()
function. You’ll need to declare it
here before you call it from main().
This code will display all
the matching tracks.
And this is your main() function,
which is the starting point of
the program.

```c
#include <stdio.h>
#include <string.h>

char tracks[][80] = {
"I left my heart in Harvard Med School",
"Newark, Newark - a wonderful town",
"Dancing with a Dork",
"From here to maternity",
"The girl from Iwo Jima",
};

void find_track(char search_for[])
{
i++ means “increase
int i;
the value of i by 1.”
for (i = 0; i < 5; i++) {
if (strstr(tracks[i], search_for))
printf("Track %i: '%s'\n", i, tracks[i]);
}
}
int main()
{
e
You're asking for .th
char search_for[80];
search text here
printf("Search for: ");
scanf("%79s", search_for);
search_for[strlen(search_for) - 1] = '\0';
find_track(search_for); Now you call your new
find_track() function and
return 0;
display the matching tracks.
}
```

It’s important that you assemble the code in this order. The headers are
included at the top so that the compiler will have all the correct functions
before it compiles your code. Then you define the tracks before you write
the functions. This is called putting the tracks array in global scope.
A global variable is one that lives outside any particular function. Global
variables like tracks are available to all of the functions in the program.

Finally, you have the functions: find_track() first, followed by
main(). The find_track() function needs to come first, before you
call it from main().

# Array of arrays vs. array of pointers 98

# Your C Toolbox 101