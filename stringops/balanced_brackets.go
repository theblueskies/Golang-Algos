package stringops

// CheckValidBalance checks if the order of brackets is valid and balanced
func CheckValidBalance(s string) string {
	brOpenMap := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	brCloseMap := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	var stack []string
	var d string

	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, string(v))
			continue
		} else if _, ok := brOpenMap[string(v)]; ok {
			stack = append(stack, string(v))
		} else if v1, ok := brCloseMap[string(v)]; ok {
			// This means that the current char is one of the brackets.
			// We need to check whether the corresponding bracket is in the stack
			d = stack[len(stack)-1]
			if d != v1 {
				return "Invalid"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return "Invalid"
	}
	return "Valid"
}
