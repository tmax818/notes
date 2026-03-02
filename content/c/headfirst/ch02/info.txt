#include <stdio.h>

int y = 1;

int main() {
    int x = 4;
    int *addr_of_x = &x;
    printf("x is stored at %p\n", &x);
    printf("y is stored at %p\n", &y);
    printf("x is equal to %d\n", *addr_of_x);
    return 0;
}
