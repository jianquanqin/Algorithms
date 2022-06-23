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

//

func removeAnagrams(words []string) []string {

	var res []string
	res = append(res, words[0])

	for i := 1; i < len(words); i++ {
		if helper(words[i], words[i-1]) == false {
			res = append(res, words[i])
		}
	}
	return res
}

func helper(a string, b string) bool {

	HashMap1 := make(map[byte]int)
	HashMap2 := make(map[byte]int)

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			HashMap1[a[i]]++
		}
		for j := 0; j < len(b); j++ {
			HashMap2[b[j]]++
		}

		//fmt.Println("---", HashMap1, HashMap2)
		for i := 0; i < len(a); i++ {
			if HashMap1[a[i]] != HashMap2[a[i]] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

//剑指 Offer 50. 第一个只出现一次的字符

func firstUniqChar(s string) byte {

	//解法一

	hashMap := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if hashMap[s[i]] == true {
			continue
		}
		flag := true
		for j := 0; j < len(s); j++ {
			//fmt.Println(s[i], s[j])
			//怎么判断两个不是同一个，令他们index不同就行了
			if s[i] == s[j] && i != j {
				flag = false
				hashMap[s[i]] = true
				break
			}
		}
		if flag == true {
			return s[i]
		}
	}

	res := " "

	return res[0]
}

//offer 05
func ReplaceSpace(s string) string {

	var res []byte
	con := "%20"

	for i:=0;i<len(s);i++ {

		if s[i] == byte(' ') {
			res = append(res,con[:]...)
		}else {
			res = append(res,s[i])
		}
	}
	return string(res)
}

func ReverseLeftWords(s string, n int) string {

	var res []byte

	res = append(res,s[n:]...)
	res = append(res,s[:n]...)

	// fmt.Println(res)
	return string(res)
}