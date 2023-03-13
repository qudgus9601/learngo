package dict

import (
	"errors"
)

type Dictionary map[string]string

type Money int

var errNotFound error = errors.New("not found")
var errWordExists error = errors.New("already exist")

/*
Dev
*/
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

/*
Dev	
*/
func (d Dictionary) Add(word , def string) error {
	_,err := d.Search(word);
	switch err {
	case errNotFound:
		d[word]=def
	case nil:
		return errWordExists
	}
	return nil;
}

/*
Dev
*/
func (d Dictionary) Update(word, newDef string) (string,error) {
	_,err := d.Search(word)
	if err == errNotFound {
		return "", errNotFound
	}
	d[word]=newDef
	return "Update Success" , nil
}

/*
Dev
*/
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word);
	if err == errNotFound {
		return errNotFound
	}
	delete(d,word)
	return nil;
}