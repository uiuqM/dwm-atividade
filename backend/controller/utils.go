package controller

func sliceString(str string) ([]string){
	var number1, number2, operator string
	number1 = ""
	number2 = ""
	operator = ""

	for _, char := range str {
		char := string(char)
		if ((char != "+" || char != "-" || char != "/" || char != "*") && operator == "") {
			number1 += char
		} else if (char == "+" || char == "-" || char == "/" || char == "*"){
			operator = char
		} else if ((char != "+" || char != "-" || char != "/" || char != "*") && operator != "")  {
			number2 += char
		}
	}
	slicedString := [] string {number1, number2, operator}

	return slicedString
}
