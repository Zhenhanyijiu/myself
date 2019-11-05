package main

/*
#include<stdio.h>
#include<stdlib.h>
typedef struct person{
	int age;
	char* name;
} person;
void print(char* pt){
	if(pt==NULL){
		printf("input is NULL,pt=%s\n",pt);
		return;
	}
	printf("%s\n",pt);
}
person pn ={23,""};
*/
import "C"
import "unsafe"

type Person struct {
	age  int
	name []byte
}

func main() {
	txt := []byte("hello cgo...")
	pt := (*C.char)(unsafe.Pointer(&txt[0]))
	C.print(pt)
}
