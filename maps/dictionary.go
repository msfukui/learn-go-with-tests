package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound  = errors.New("指定された単語を見つけることができませんでした")
	ErrExistWord = errors.New("既に存在しているため単語を追加できませんでした")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrExistWord
	default:
		return err
	}
	return nil
}
