package main

type Dictionary map[string]string

type DictionaryErr string

//https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrNotFound         = DictionaryErr("no item matches the key")
	ErrWordExists       = DictionaryErr("item already exists")
	ErrWordDoesNotExist = DictionaryErr("item to update does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Note that maps can be modified
// Without using a pointer because they are a reference type
// https://go.dev/blog/maps
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	// Built in delete function for maps
	delete(d, key)
}
