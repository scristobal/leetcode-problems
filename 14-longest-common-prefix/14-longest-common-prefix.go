

func longestCommonPrefix(strs []string) (chain string) {

   defer func() {
		recover()
	}()

	chain = ""

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; true; i++ {

		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == 0 {
				return ""
			}

			chain = strs[j][:i]

			if strs[j-1][i] != strs[j][i] {

				return chain
			}
		}

	}

	return chain

}
