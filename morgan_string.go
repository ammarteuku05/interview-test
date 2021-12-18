package main

func MorganAndString(a string, b string) string {
	res := ""

	sizeA := len(a)
	sizeB := len(b)

	var indexA, indexB int
	for indexA < sizeA && indexB < sizeB {
		currA := a[indexA]
		currB := b[indexB]

		if currA <= currB {
			res += string(a[indexA])
			indexA++
		} else if currB < currA {
			res += string(b[indexB])
			indexB++
		}
	}

	if indexA == sizeA && indexB < sizeB {
		for indexB < sizeB {
			res += string(b[indexB])
			indexB++
		}
	} else if indexB == sizeB && indexA < sizeA {
		for indexA < sizeA {
			res += string(a[indexA])
			indexA++
		}
	}

	return res

}

// func yielda() chan string {
// 	yield := make(chan string)
// 	count := ""
// 	i := 0
// 	go func() {
// 		for {
// 			yield <- count
// 			i++
// 		}
// 	}()
// 	return yield
// }
