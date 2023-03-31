package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"テスト": "これがテストだ!"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("テスト")
		want := "これがテストだ!"

		AssertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "指定された単語を見つけることができませんでした"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		AssertStrings(t, err.Error(), want)
	})
}

func AssertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
