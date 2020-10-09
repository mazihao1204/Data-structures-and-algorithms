package main

type IFile interface {
	Read(buf []byte)(n int,err error)
	Write(buf []byte)(n int,err error)
	Seek(off int64,whence int)(pos int64,err error)
}

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString)findVowels() []rune{
	var vowels []rune
	for _,rune := range ms{
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u'{
			vowels = append(vowels,rune)
		}
	}
	return vowels
}

type tester interface {
	test()
	string() string
}

type node struct {
	data interface{
		string() string
	}
}

type data struct {

}

func (this *data) test(){

}

func (data) string() string{
	return ""
}

func main(){
	var n interface{
		string() string
	} = data{}

	var d data
	var t tester = &d
	t.test()
}

