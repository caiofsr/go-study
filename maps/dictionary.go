package dictionary

import "errors"

type Dictionary map[string]string

type ErrDictionary string

var (
	ErrExistingWord   = errors.New("cannot add existing word")
	ErrInexistentWord = errors.New("cannot update inexistent word")
	ErrNotFound       = errors.New("could not find the word you were looking for")
)

func (e ErrDictionary) Error() string {
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
		return ErrExistingWord
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrInexistentWord
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
