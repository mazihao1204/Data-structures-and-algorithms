package string


//有效的字母异位词
//异位词：字母相同，排列不同的字符串，所输入的字母都为小写
func isAnagram(s string,t string) bool{
	alphabet := make([]int,26)
	if len(s) != len(t){
		return false
	}
	for i := 0;i < len(s);i++{
		alphabet[s[i]-'a']++
	}
	for i := 0;i < len(t);i++{
		alphabet[t[i]-'a']--
	}
	for i := 0;i < 26;i++ {
		if alphabet[i] != 0{
			return false
		}
	}
	return true
}
