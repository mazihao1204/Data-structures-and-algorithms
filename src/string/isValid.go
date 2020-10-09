package string

//判断有效括号
func isValid(s string) bool{
	if len(s) == 0 || len(s)%2 != 0{
		return false
	}
	arr := []byte{}
	for i := 0;i < len(s);i++{
		if s[i] == '['{
			arr = append(arr,']')
		}else if s[i] == '{'{
			arr = append(arr,'}')
		}else if s[i] == '('{
			arr = append(arr,')')
		}else if arr[len(arr)-1] == s[i]{
			arr = arr[:len(arr)-1]
		}
	}
	return arr == nil
}

