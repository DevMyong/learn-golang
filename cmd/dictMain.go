package main

import (
	"../internal/mydict"
	"fmt"
)

// printResult error or normal
func printResult (word string, err error){
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(word)
	}
}

func main(){
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	word, err := dictionary.Search(baseWord)
	printResult(word, err)
	err = dictionary.Add(baseWord, "First")
	word, _ = dictionary.Search(baseWord)
	printResult(word, err)

	err = dictionary.Update(baseWord, "Second")
	word, _ = dictionary.Search(baseWord)
	printResult(word, err)
	err = dictionary.Update("qwerasdf", "Third")
	word, _ = dictionary.Search("qwerasdf")
	printResult(word, err)

	dictionary.Delete(baseWord)
	word, err = dictionary.Search(baseWord)
	printResult(word, err)
}

