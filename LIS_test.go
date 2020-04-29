package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestShouldReturnOneElementArrayWhenAllInputElementsAreInDescendingOrder(t *testing.T){
	input:=[]int{9,8,7,6,5,4,3,2,1}
	expectedResultSize:=1

	result:=FindLIS(input)

	assert.Equal(t, expectedResultSize, len(result))

}

func TestShouldReturnEmptyArrayWhenInputIsEmpty(t *testing.T){
	var input []int

	result:=FindLIS(input)

	assert.Equal(t, 0, len(result))

}

func TestShouldReturnOneElementArrayWhenAllInputElementsAreTheSame(t *testing.T){
	input:=[]int{5,5,5,5}
	expectedResult:=[]int{5}

	result:=FindLIS(input)

	assert.Equal(t, expectedResult, result)

}

func TestShouldReturnLISyWhenInputContainsNegativeNumbersElementsAreTheSame(t *testing.T){
	input:=[]int{-10,-5, 6,0,-4,3,0}
	expectedResult:=[]int{-10, -5,-4,0}

	result:=FindLIS(input)

	assert.Equal(t, expectedResult, result)

}

func TestShouldReturnLIS(t *testing.T){
	input:=[]int{1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16}
	expectedResult:=[]int{1, 3, 7, 10, 12, 16}

	result:=FindLIS(input)

	assert.Equal(t, expectedResult, result)

}
