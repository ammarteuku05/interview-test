package main

func JumpFrog(x1 int, v1 int, x2 int, v2 int) string {
	if x2 > x1 && x2 > v1 {
		return "NO"
	} else {
		if v2-v1 == 0 {
			return "NO"
		} else {
			res := (x1 - x2) % (v2 - v1)
			if res == 0 {
				return "YES"
			} else {
				return "NO"
			}
		}
	}
}
