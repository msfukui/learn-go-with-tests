package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("任意のサイズのスライスを指定する場合", func(t *testing.T) {
		members := []int{1, 2, 3, 4, 5}
		got := Sum(members)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, members)
		}
	})
}
