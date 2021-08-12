#include <stdio.h>
#include "ccode.h"

void helloFromC(){
    printf("Hello from C code in independent file\n");
}
void passingArgumentsToC(char* go_message){
    printf("Receive a GO parameter, send me: %s\n", go_message);
}