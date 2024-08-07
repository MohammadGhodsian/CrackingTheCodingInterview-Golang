package binary_to_string

/*
5.2 Binary to String
Given a real number between 0 and 1 (e.g., 0.72) that is passed in as a double, print the binary representation.
If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".
*/

func FloatToBinary(num float32) string {
	if num < 0 || num > 1 {
		return "ERROR"
	}
	if num == 0 {
		return "0.0"
	}
	num64 := float64(num)
	resultStr := "0."
	for num64 > 0 {
		if len(resultStr) > 32 {
			return "ERROR"
		}
		num64 *= 2
		if num64 >= 1 {
			resultStr += "1"
			num64--
		} else {
			resultStr += "0"
		}
	}
	return resultStr
}
