package main

import (
	"fmt"
)

func main(){
	s := "pwwleweoritu"
	fmt.Println(lengthOfLongestSubstring(s))


}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256] int // extended ascii
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right<len(s) && freq[s[right] - 'a']==0{
			freq[s[right]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left)
	}
	return result
}
func max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}

