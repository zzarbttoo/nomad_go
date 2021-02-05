package mydict

import "errors"

//type의 Alias 를 만듦

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {

	//string, boolean
	value, exist := d[word]

	if exist {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to dictionary
func (d Dictionary) Add(word, def string) error {

	_, err := d.Search(word)

	/*
		if err == errNotFound {

			d[word] = def
		} else if err == nil {

			return errWordExists
		}

		return nil

	*/

	switch err {

	case errNotFound:
		d[word] = def

	case nil:
		return errWordExists
	}
	return nil

}
