package int_to_roman

func intToRoman(num int) string {
	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := ""
	for num > 0 {
		for _, val := range nums {
			if num >= val {
				s += m[val]
				num -= val
				break
			}

		}
	}
	return s
}
