package main

import "strconv"

func isAdditiveNumber(num string) bool {
	lens := len(num)
	if lens < 3 {
		return false
	}

	f, s := "", ""
	for i := 0; i*2 < lens; i++ {
		f = num[0 : i+1]
		if num[0] == '0' && i > 0 {
			continue
		}
		for j := i + 1; (j-i)*2 <= lens; j++ {
			s = num[i+1 : j+1]
			if num[i+1] == '0' && j > i+1 {
				continue
			}
			tm := j + 1
			f_n, _ := strconv.Atoi(f)
			s_n, _ := strconv.Atoi(s)
			if tm == lens {
				return false
			}
			for tm < lens {
				sum_n := f_n + s_n
				sum := strconv.Itoa(sum_n)
				if tm+len(sum) > lens {
					break
				}
				if num[tm:tm+len(sum)] != sum {
					break
				}
				f_n, s_n = s_n, sum_n
				tm += len(sum)
			}
			if tm == lens {
				return true
			}
		}
	}

	return false
}
func main() {
	//println(isAdditiveNumber("112358"))
	//println(isAdditiveNumber("199100199"))
	//println(isAdditiveNumber("199111992"))
	println(isAdditiveNumber("111"))
}
