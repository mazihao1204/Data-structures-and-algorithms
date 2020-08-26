package main

//leetcode17 电话号码的字母组合
//解法：直接深度优先

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}  
var res []string

func letterCombinations(digits string) []string {
	if digits == ""{
		return []string{}
	}
	res = []string{}
	backtrack(digits,0,"")
	return res
}

func backtrack(digits string,index int,s string){
	if index == len(digits){
		res = append(res,s)
	}else {
		num := string(digits[index])
		letter := phoneMap[num]
		for i := 0;i < len(letter);i++{
			backtrack(digits,index+1,s+string(letter[i]))
		}
	}
}