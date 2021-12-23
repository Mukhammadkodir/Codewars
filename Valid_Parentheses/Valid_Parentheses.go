/*
Write a function that takes a string of parentheses,
and determines if the order of the parentheses is valid.
The function should return true if the string is valid,
and false if it's invalid.

Examples

"()"              =>  true
")(()))"          =>  false
"("               =>  false
"(())((()())())"  =>  true

*/
package main

import (
	"fmt"
)

func main(){
fmt.Println(ValidParentheses("())("))

}
func ValidParentheses(parens string) bool {
 m := map[rune]int{
	'(':0,
	')':0,

 }

 for _,i := range parens{
		m[i]++
		if m[')']>m['(']{
			return false
		}
	}
	if m['('] != m[')']{
		return false
	}
	return true

}
