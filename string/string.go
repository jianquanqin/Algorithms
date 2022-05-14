package string

//替换空格
func replaceSpace(s string) string {

	u := " %20"
	var res []byte
	for i := 0; i < len(s); i++ {

		if s[i] == u[0] {
			res = append(res, u[1:]...)
		} else {
			res = append(res, s[i])
		}
	}
	s = string(res)
	return s
}

//旋转数组

func reverseLeftWords(s string, n int) string {

	var res []byte
	res = append(res, s[n:]...)
	res = append(res, s[0:n]...)

	return string(res)
}
