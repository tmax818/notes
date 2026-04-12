

# 1 getting started with C Diving in

Want to get inside the computer’s head? 

Need to write high-performance code for a new game? Program an Arduino? Or use that advanced third-party library in your iPhone app? If so, then C’s here to help. C works at a much lower level than most other languages, so understanding C gives you a much better idea of what’s really going on. C can even help you better understand other languages as well. So dive in and grab your compiler, and you’ll soon get started in no time.

## C is a language for small, fast programs 2
## But what does a complete C program look like? 5
## But how do you run the program? 9
## Two types of command 14
## Here’s the code so far 15
## Card counting? In C? 17
## There’s more to booleans than equals… 18
## What’s the code like now? 25
## Pulling the ol’ switcheroo 26
## Sometimes once is not enough… 29
## Loops often follow the same structure… 30
## You use break to break out… 31
## Your C Toolbox 40

# 2 memory and pointers What are you pointing at?

If you really want to kick butt with C, you need to understand how C handles memory.

The C language gives you a lot more control over how your program uses the computer’s memory. In this chapter  you’ll strip back the covers and see exactly what happens when you read and write variables. You’ll learn how arrays work, how to avoid some nasty memory SNAFUs, and most of all, you’ll see how mastering pointers and memory addressing is key to becoming a kick-ass C programmer.

## C code includes pointers 42
## Digging into memory 43
## Set sail with pointers 44
## Try passing a pointer to the variable 47
## Using memory pointers 48
## How do you pass a string to a function? 53
## Array variables are like pointers… 54
## What the computer thinks when it runs your code 55
## But array variables aren’t quite pointers 59
## Why arrays really start at 0 61
## Why pointers have types 62
## Using pointers for data entry 65
## Be careful with `scanf()` 66
## `fgets()` is an alternative to `scanf()` 67
## String literals can never be updated 72
## If you’re going to change a string, make a copy 74
## Memory memorizer 80
## Your C Toolbox 81

# 2.5 strings String theory

There’s more to strings than reading them.

You’ve seen how strings in C are actually char arrays but what does C allow you to do with them? That’s where string.h comes in. `string.h` is p art of the C Standard Library that’s dedicated to string manipulation. If you want to concatenate strings together, copy one string to another, or compare two strings, the functions in string.h are there to help. In this chapter, you’ll see how to create an array of strings, and then take a close look at how to search within strings using the `strstr()` function.

## Desperately seeking Frank 84
## Create an array of arrays 85
## Find strings containing the search text 86
## Using the `strstr()` function 89
## It’s time for a code review 94
## Array of arrays vs. array of pointers 98
## Your C Toolbox 101


# 3 creating small tools Do one thing and do it well

Every operating system includes small tools.

Small tools written in C perform specialized small tasks, such as reading and writing files, or filtering data. If you want to perform more complex tasks, you can even link several tools together. But how are these small tools built? In this chapter, you’ll look at the building blocks of creating small tools. You’ll learn how to control command-line options, how to manage streams of information, and redirection, getting tooled up in no time.

## Small tools can solve big problems 104
## Here’s how the program should work 108
## But you’re not using files… 109
## You can use redirection 110
## Introducing the Standard Error 120
## By default, the Standard Error is sent to the display 121
## `fprintf()` prints to a data stream 122
## Let’s update the code to use `fprintf()` 123
## Small tools are flexible 128
## Don’t change the geo2json tool 129
## A different task needs a different tool 130
## Connect your input and output with a pipe 131
## The bermuda too l132
## But what if you want to output to more than one file? 137
## Roll your own data streams 138
## There’s more to main() 141
## Let the library do the work for you 149
## Your C Toolbox 156


# 4 using multiple source files Break it down, build it up

If you create a big program, you don’t want a big source file.

Can you imagine how difficult and time-consuming a single source file for an enterprise- level program would be to maintain? In this chapter, you’ll learn how C allows you to break your source code into small, manageable chunks and then rebuild them into one huge program. Along the way, you’ll learn a bit more about data type subtleties and get to meet your new best friend: make.

## Your quick guide to data types 162
## Don’t put something big into something small 163
## Use casting to put floats into whole numbers 164
## Oh no…it’s the out-of-work actors… 168
## Let’s see what’s happened to the code 169
## Compilers don’t like surprises 171
## Split the declaration from the definition 173
## Creating your first header file 174
## If you have common features… 182
## You can split the code into separate files 183
## Compilation behind the scenes 184
## The shared code needs its own header file 186
## It’s not rocket science…or is it? 189
## Don’t recompile every file 190
## First, compile the source into object files 191
## It’s hard to keep track of the files 196
## Automate your builds with the make tool 198
## How make works 199
## Tell make about your code with a makefile 200
## Liftoff ! 205
## Your C Toolbox 206


# C Lab 1 Arduino

Ever wished your plants could tell you when they need watering? Well, with an Arduino, they can! In this lab, you’ll build an Arduino-powered plant monitor, all coded in C.


# 5 structs, unions, and bitfields Rolling your own structures

Most things in life are more complex than a simple number.

So far, you’ve looked at the basic data types of the C language, but what if you want to go beyond numbers and pieces of text, and model things in the real world? structs allow you to model real-world complexities by writing your own structures. In this chapter, you’ll learn how to combine the basic data types into structs, and even handle life’s uncertainties with unions. And if you’re after a simple yes or no, bitfields may be just what you need.


## Sometimes you need to hand around a lot of data 218
## Cubicle conversation 219
## Create your own structured data types with a struct 220
## Just give them the fish 221
## Read a struct’s fields with the “.” operator 222
## Can you put one struct inside another? 227
## How do you update a struct? 236
## The code is cloning the turtle 238
## You need a pointer to the struct 239
## (*t).age vs. *t.age 240
## Sometimes the same type of thing needs different types of data 246
## A union lets you reuse memory space 247
## How do you use a union? 248
## An enum variable stores a symbol 255
## Sometimes you want control at the bit level 261
## Bitfields store a custom number of bits 262
## Your C Toolbox 266

# 6 data structures and dynamic memory Building bridges

Sometimes, a single struct is simply not enough.

To model complex data requirements, you often need to link structs together. In
this chapter, you’ll see how to use struct pointers to connect custom data types into
large, complex data structures. You’ll explore key principles by creating linked lists.
You’ll also see how to make your data structures cope with flexible amounts of data by
dynamically allocating memory on the heap, and freeing it up when you’re done. And
if good housekeeping becomes tricky, you’ll also learn how valgrind can help.

## Do you need flexible storage? 268
## Linked lists are like chains of data 269
## Linked lists allow inserts 270
## Create a recursive structure 271
## Create islands in C… 272
## Inserting values into the list 273
## Use the heap for dynamic storage 278
## Give the memory back when you’re done 279
## Ask for memory with `malloc()`… 280
## Let’s fix the code using the `strdup()` function 286
## Free the memory when you’re done 290
## An overview of the SPIES system 300
## Software forensics: using valgrind 302
## Use valgrind repeatedly to gather more evidence 303
## Look at the evidence 304
## The fix on trial 307
## Your C Toolbox 309



# 7 advanced functions Turn your functions up to 11

Basic functions are great, but sometimes you need more.

So far, you’ve focused on the basics, but what if you need even more power and flexibility to achieve what you want? In this chapter, you’ll see how to up your code’s IQ by passing functions as parameters. You’ll find out how to get things sorted with comparator functions. And finally, you’ll discover how to make your code super stretchy with variadic functions.

## Looking for Mr. Right… 312
## Pass code to a function 316
## You need to tell `find()` the name of a function 317
## Every function name is a pointer to the function… 318
## …but there’s no function data type 319
## How to create function pointers 320
## Get it sorted with the C Standard Library 325
## Use function pointers to set the order 326
## Automating the Dear John letters 334
## Create an array of function pointers 338
## Make your functions streeeeeetchy 343
## Your C Toolbox 350



# 8 static and dynamic libraries Hot-swappable code

You’ve already seen the power of standard libraries.

Now it’s time to use that power for your own code. In this chapter, you’ll see how to
create your own libraries and reuse the same code across several programs.
What’s more, you’ll learn how to share code at runtime with dynamic libraries. You’ll
learn the secrets of the coding gurus. And by the end of the chapter, you’ll be able to
write code that you can scale and manage simply and efficiently.


## Code you can take to the bank 352
## Angle brackets are for standard headers 354
## But what if you want to share code? 355
## Sharing `.h` header files 356
## Share `.o` object files by using the full pathname 357
## An archive contains `.o` files 358
## Create an archive with the ar command… 359
## Finally, compile your other programs 360
## The Head First Gym is going global 365
## Calculating calories 366
## But things are a bit more complex… 369
## Programs are made out of lots of pieces… 370
## Dynamic linking happens at runtime 372
## Can you link `.a` at runtime? 373
## First, create an object file 374
## What you call your dynamic library depends on your platform 375
## Your C Toolbox 387


# C Lab 2 OpenCV

Imagine if your computer could keep an eye on your house while you’re out, and tell you who’s been prowling around. In this lab, you’ll build a C-powered intruder detector using the cleverness of OpenCV.



# 9 processes and system calls Breaking boundaries

It’s time to think outside the box.

You’ve already seen that you can build complex applications by connecting small tools together on the command line. But what if you want to use other programs from inside your own code? In this chapter, you’ll learn how to use system services to create and control processes. That will give your programs access to email, the Web, and any other tool you’ve got installed. By the end of the chapter, you’ll have the power to go beyond C.


## System calls are your hotline to the OS 398
## Then someone busted into the system… 402
## Security’s not the only problem 403
## The `exec()` functions give you more control 404
## There are many `exec()` functions 405
## The array functions: `execv()`, `execvp()`, `execve()` 406
## Passing environment variables 407
## Most system calls go wrong in the same way 408
## Read the news with RSS 416
## `exec()` is the end of the line for your program 420
## Running a child process with `fork()` + `exec()` 421
## Your C Toolbox 427


# 10 interprocess communication It’s good to talk

Creating processes is just half the story.

What if you want to control the process once it’s running? What if you want to send it data? Or read its output? Interprocess communication lets processes work together to get the job done. We’ll show you how to multiply the power of your code by letting it
talk to other programs on your system.

## Redirecting input and output 430
## A look inside a typical process 431
## Redirection just replaces data streams 432
## `fileno()` tells you the descriptor 433
## Sometimes you need to wait… 438
## Stay in touch with your child 442
## Connect your processes with pipes 443
## Case study: opening stories in a browser 444
## In the child 445
## In the parent 445
## Opening a web page in a browser 446
## The death of a process 451
## Catching signals and running your own code 452
## sigactions are registered with `sigaction()` 453
## Rewriting the code to use a signal handler 454
## Use kill to send signals 457
## Sending your code a wake-up call
## Your C Toolbox


# 11 sockets and networking There’s no place like 127.0.0.1

Programs on different machines need to talk to each other.

You’ve learned how to use I/O to communicate with files and how processes on the
same machine can communicate with each other. Now you’re going to reach out
to the rest of the world, and learn how to write C programs that can talk to other
programs across the network and across the world. By the end of this chapter,
you’ll be able to create programs that behave as servers and programs that
behave as clients.

## The Internet knock-knock server 468
## Knock-knock server overview 469
## BLAB: how servers talk to the Internet 470
## A socket’s not your typical data stream 472
## Sometimes the server doesn’t start properly 476
## Why your mom always told you to check for errors 477
## Reading from the client 478
## The server can only talk to one person at a time 485
## You can `fork()` a process for each client 486
## Writing a web client 490
## Clients are in charge 491
## Create a socket for an IP address 492
## `getaddrinfo()` gets addresses for domains 493
## Your C Toolbox 500

# 12 threads It’s a parallel world

Programs often need to do several things at the same time.

POSIX threads can make your code more responsive by spinning off several pieces of code to run in parallel. But be careful! Threads are powerful tools, but you don’t want them crashing into each other. In this chapter, you’ll learn how to put up traffic signs and lane markers that will prevent a code pileup. By the end, you will know how to create POSIX threads and how to use synchronization mechanisms to protect the integrity of sensitive data.

## Tasks are sequential…or not… 502
## …and processes are not always the answer 503
## Simple processes do one thing at a time 504
## Employ extra staff: use threads 505
## How do you create threads? 506
## Create threads with `pthread_create` 507
## The code is not thread-safe 512
## You need to add traffic signals 513
## Use a mutex as a traffic signal 514
## Your C Toolbox
