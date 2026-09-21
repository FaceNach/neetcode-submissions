func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	n := make(map[byte]int,0)
	m := make(map[byte]int,0)

	for i:=0; i < len(s); i++{
		n[s[i]]++
		m[t[i]]++
	}

	for k, v := range n {
		if m[k] != v {
			return false
		}
	}

return true
}
