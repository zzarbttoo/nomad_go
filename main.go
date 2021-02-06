package main

import (
	"fmt"

	"github.com/zzarbttoo/learngo/mydict"
)

func main() {

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	definition := "First"

	//새로운 단어 추가
	dictionary.Add(baseWord, definition)
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println(word)

	/*
		updateErr := dictionary.Update("gi", "Second")

		if updateErr != nil {
			fmt.Println(updateErr)
		}

		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)

	*/

	//같은 단어 또 추가
	/*
		hello, _ := dictionary.Search(word)
		fmt.Println("found", word, "definition:", hello)
		err2 := dictionary.Add(word, definition)

		if err2 != nil {
			fmt.Println(err2)
		}

	*/

}
