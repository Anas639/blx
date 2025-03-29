package task

import (
	"reflect"
	"testing"
)

func TestParseStatusSlice(t *testing.T) {
	cases := []struct {
		has  []string
		want []TaskStatus
	}{
		{
			has: []string{"new", "ongoing", "paused", "ended"},
			want: []TaskStatus{
				TASK_NEW, TASK_ONGOING, TASK_PAUSED, TASK_ENDED,
			},
		},
		{
			has: []string{"new", "paused", "ended"},
			want: []TaskStatus{
				TASK_NEW, TASK_PAUSED, TASK_ENDED,
			},
		},
		{
			has: []string{"new", "new", "new"},
			want: []TaskStatus{
				TASK_NEW,
			},
		},
		{
			has:  []string{"foo", "bar"},
			want: []TaskStatus{},
		},
	}

	for _, c := range cases {
		got := StatusesFromSlice(c.has)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("Want %s but got %s", c.want, c.has)
		}
	}
}
