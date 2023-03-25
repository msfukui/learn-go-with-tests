package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("整数が要素の任意のサイズのスライスを指定するとスライス内の数を合計して返す", func(t *testing.T) {
		members := []int{1, 2, 3, 4, 5}
		got := Sum(members)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, members)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("整数が要素の任意のサイズのスライスを複数指定するとスライス毎に数を合計したスライスを返す", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(y *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("整数が要素の任意のサイズのスライスを複数指定するとスライス毎の最初を除く合計の要素のスライスを返す", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("安全性のため空のスライスには0を返す", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
