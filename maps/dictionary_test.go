package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"テスト": "これがテストだ!"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("テスト")
		want := "これがテストだ!"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "テスト"
		definition := "これがテストだ!"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "テスト"
		definition := "これがテストだ!"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrExistWord)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "テスト"
		definition := "これがテストだ!"
		dictionary := Dictionary{word: definition}
		newDefinition := "新しい定義"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "テスト"
		definition := "これがテストだ!"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
