func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToTMapping := make(map[byte]byte)
	tToSMapping := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]

		if mappedT, exists := sToTMapping[charS]; exists {
			if mappedT != charT {
				return false
			}
		} else {
			sToTMapping[charS] = charT
		}

		if mappedS, exists := tToSMapping[charT]; exists {
			if mappedS != charS {
				return false
			}
		} else {
			tToSMapping[charT] = charS
		}
	}

	return true
}

func main() {
	s1, t1 := "egg", "add"
	fmt.Println(isIsomorphic(s1, t1))

	s2, t2 := "foo", "bar"
	fmt.Println(isIsomorphic(s2, t2))
}
