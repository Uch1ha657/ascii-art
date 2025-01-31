package function

func Rida(input string, asciiarr map[rune][]string) []string {
	ret := make([]string, 8)
	for _, v := range input {
		if ris, booln := asciiarr[v]; booln {
			for i := 0; i < 8; i++ {
				ret[i] = ret[i] + ris[i]
			}
		}
	}
	return ret
}
