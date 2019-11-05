package cgo

/*
#include<stdio.h>
#include<stdlib.h>
void print(char* pt){
	if(pt==NULL){
		printf("input is NULL,pt=%s\n",pt);
		return;
	}
	printf("%s\n",pt);
}
*/
import "C"
