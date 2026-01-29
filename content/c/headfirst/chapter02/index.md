# chapter 2: memory and pointers: what are you pointing at?

## C code includes pointers

- pointer::the address of a piece of data in memory

- pointers
    - pass a pointer instead of data
    - two pieces of code work on the same data, instead of a copy


    
The `*` and `&` operators are opposites. 
- The `&` operator takes a piece of data and tells you it's address. 
- The `*` operator takes an address and tells you what’s stored there.

- variables are allocated storage in memory
- pointers are variables that store memory addresses
- the `&` operator finds the address of a variable
- the `*` operator can read/write the contents of a memory address