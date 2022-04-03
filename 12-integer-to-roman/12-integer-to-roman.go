type roman struct {
	s string
	v int
}

func intToRoman(num int) string {

	var m = []roman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	o := ""

	for r := num; r > 0; {
		for _, s := range m {

			if s.v > r {
				continue
			}

           
			r -= s.v
			o = fmt.Sprintf("%s%s", o, s.s)
            
            break
		}
	}

	return o
}