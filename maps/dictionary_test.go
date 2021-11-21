package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"ball": "circular thing"}

	// Key values  must be comparable
	// https://golang.org/ref/spec#Comparison_operators
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("ball")
		want := "circular thing"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("apple")

		if err == nil {
			t.Fatal("expected an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := "no"
		word := "yes"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "yes"
		definition := "no"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new no")

		assertAddError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "something"
		definition := "is true"
		newDefinition := "is false"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "def"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)

	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "something"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

//////////////////////
// Assertion helpers
//////////////////////

func assertAddError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected an error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find the word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

}
