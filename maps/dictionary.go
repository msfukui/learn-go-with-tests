package maps

type Dictionary map[string]string

const (
	ErrNotFound  = DictionaryErr("指定された単語を見つけることができませんでした")
	ErrExistWord = DictionaryErr("既に存在しているため単語を追加できませんでした")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
