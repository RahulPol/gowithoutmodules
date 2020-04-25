package dup

// GetDuplicates := returns duplicate string 
func GetDuplicates(data []string) []string{
	if len(data) == 0{
		return []string{}
	}

	var result = []string{}
	count := make(map[string]int)	

	for _, s := range data{
		count[s]++
		if n := count[s]; n > 1{
			result = append(result, s)
		}
	}

	return result
}