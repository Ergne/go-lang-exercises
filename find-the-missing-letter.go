package kata

func FindMissingLetter(chars []rune) rune {
	var i rune

	for i = chars[0]; i <= chars[len(chars)-1]; i++ {
		if i != chars[i-chars[0]] {
			break
		}
	}

	return i
}
