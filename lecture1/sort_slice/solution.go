package sort_slice

type SortSlice struct {
	Slice []int
}

func (s *SortSlice) Sort() []int {
	for i := 0; i < len(s.Slice)-1; i++ {
		for j := 0; j < len(s.Slice)-i-1; j++ {
			if s.Slice[j] > s.Slice[j+1] {
				s.Slice[j], s.Slice[j+1] = s.Slice[j], s.Slice[j+1]
			}
		}
	}
	return s.Slice
}
