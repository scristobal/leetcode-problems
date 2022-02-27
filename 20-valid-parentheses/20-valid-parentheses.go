
func removeAtoms(s string) string {

	atoms := [3]string{"()", "{}", "[]"}
	t := s

	for _, atom := range atoms {

		t = strings.ReplaceAll(t, atom, "")
	}

	return t
}

func removeAllAtoms(s string) string {

	u := s
	t := removeAtoms(s)

	for u != t {
		t = u
		u = removeAtoms(t)
	}

	return t
}

func isValid(s string) bool {

	leftOvers := removeAllAtoms(s)
	return leftOvers == ""
}
