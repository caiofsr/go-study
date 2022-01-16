package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is just a test"

		compareString(t, result, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("desconhecida")

		compareError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		expected := "this is just a test"

		compareError(t, err, nil)
		compareDefinition(t, dictionary, "test", expected)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		compareError(t, err, ErrExistingWord)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: defintion}

		err := dictionary.Update(word, newDefinition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, defintion)

		compareError(t, err, ErrInexistentWord)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	compareError(t, err, ErrNotFound)
}

func compareString(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Expected '%s' but got '%s', given '%s'", expected, result, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if result != definition {
		t.Errorf("Expected '%s' but got '%s'", definition, result)
	}
}
