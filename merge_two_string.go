package main

func MergeTwoString(s1 string, s2 string) string {
	var res string

	res = ""
	var i int
	for i = 0; (i < len(s1)) || (i < len(s2)); i++ {
		if i < len(s1) {
			res += string(s1[i])
		}

		if i < len(s2) {
			res += string(s2[i])
		}
	}

	return res
}
