package main

func MaxKey(m map[int]string) int {
	var c_max int
	for k, _ := range m {
		if k >= *&c_max {
			*&c_max = k
		}
	}

	return c_max
}
