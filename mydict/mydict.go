package mydict

import "errors"

//type의 Alias 를 만듦

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

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

// Update
func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {

	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {

	delete(d, word)
}
