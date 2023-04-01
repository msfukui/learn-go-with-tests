package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("指定された単語を見つけることができませんでした")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
