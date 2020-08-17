package main

func isValidSerialization(preorder string) bool {
	jinhao, douhao := '#', ','
	c := 1
	for k, v := range preorder {
		if v == douhao {
			if c <= 0 {
				return false
			}
			c--
			if preorder[k-1] != uint8(jinhao) {
				c += 2
			}
		}
	}
	if preorder[len(preorder)-1] == uint8(jinhao) {
		c--
	} else {
		c++
	}
	return c == 0
}
func main() {
	println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")) //true
	println(isValidSerialization("1,#"))                       //false
	println(isValidSerialization("9,#,#,1"))                   //false
}
