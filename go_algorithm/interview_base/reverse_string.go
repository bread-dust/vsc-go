package main


func ReverString(s string) string {
	l := len(s)
	sa := make([]byte, 0)
	for _, v := range s {
		sa = append(sa, byte(v))
	}

	for i := 0; i < l/2; i++ {
		sa[i], sa[l-i-1] = sa[l-i-1], sa[i]
	}
	var res string
	for _, v := range sa {
		res += string(v)
	}
	return res
}

// 0 5
// 1 4
// 2 3

// 0 6
// 1 5
// 2 4
// 3 3