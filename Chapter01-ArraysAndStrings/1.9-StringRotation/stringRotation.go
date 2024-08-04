package stringRotation

import "strings"

/*
String Rotation:
Assume you have amethod isSubstring which checks if one word is a substring of another.
Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring
(e.g.,"waterbottle"is a rotation of"erbottlewat").
*/

func CheckStringRotation(s1, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
