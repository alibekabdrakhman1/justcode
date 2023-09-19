package longest_common_prefix

func Prefix(strs []string) string {
	for i := range strs[0] {
		for _, str := range strs {
			if i >= len(str) {
				return strs[0][:i]
			}
			if strs[0][i] != str[i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
