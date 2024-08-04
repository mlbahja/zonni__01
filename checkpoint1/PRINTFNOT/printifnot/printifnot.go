package printifnot

func PrintIfNot(str string) string {
	if len(str) < 3 || len(str) == 0 {
		return "G\n"
	}
	return "Invalid Input\n"
}
