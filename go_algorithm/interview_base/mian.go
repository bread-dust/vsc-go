package main

import "fmt"

type People interface{
	Show()
}

type Student struct{
	Name string
}

func (s *Student) Show(){
	
	fmt.Println("success")
}

func main() {
	// Goroutine_alternately()

	// result1 := IsUniqueString1("abca")
	// fmt.Println(result1)

	// result2 := IsUniqueString2("abc")
	// fmt.Println(result2)

	// result3 := IsUniqueString3("abca")
	// fmt.Println(result3)

	// result4 := ReverString("abcdef")
	// fmt.Println(result4)

	// result5 := IsRegroup("abcd Ff","caFfbd")
	// fmt.Println(result5)

	// result6 := ReplaceBlank("ab2cd Ff")
	// fmt.Println(result6)

	// var s = new(Student)
	// s.Show()

	// GetRandom()

	// IpVisit()
	// g_ticker()
	// g_timer()

	// WriteClosedChan()
	// ReadClosedChan()
	fmt.Println(f1())
	
}

func f1()(i int){
	i=1
	defer func(){
		i=i+1
	}()
	return i
}