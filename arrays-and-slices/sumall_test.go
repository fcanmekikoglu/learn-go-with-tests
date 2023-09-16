package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("test sumall", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// its like comparing objects in jest. they look same but not equal
		// it doesnt have a type check, you can expect string for example

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
