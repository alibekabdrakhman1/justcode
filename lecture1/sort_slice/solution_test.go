package sort_slice

import (
	"reflect"
	"testing"
)

func TestSortSlice_Sort(t *testing.T) {
	type fields struct {
		Slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "first test",
			fields: fields{Slice: []int{51, 52, 12, 38}},
			want:   []int{12, 38, 51, 52},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SortSlice{
				Slice: tt.fields.Slice,
			}
			if got := s.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
