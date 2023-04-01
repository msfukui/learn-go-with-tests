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
}
