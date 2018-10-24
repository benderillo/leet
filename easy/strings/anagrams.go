package strings

import "fmt"

func isAnagram_two_arrays(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	array1 := make([]byte, 26, 26)
	array2 := make([]byte, 26, 26)

	for i := 0; i < len(s); i++ {
		array1[s[i]-'a']++
		array2[t[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func isAnagram_one_array(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	diffs := make([]byte, 26, 26)

	for i := 0; i < len(s); i++ {
		diffs[s[i]-'a']++
		diffs[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if diffs[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagram_map(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	diffs := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		diffs[s[i]-'a']++
		diffs[t[i]-'a']--
	}

	for _, diff := range diffs {
		if diff != 0 {
			return false
		}
	}
	return true
}

func isAnagram_sort(s string, t string) bool {
	// First sort both input strings
	// Then go through them and compare symbols
	return false
}

func main() {
	fmt.Println(isAnagram_map("anagram", "nagaram"))
	fmt.Println(isAnagram_map("car", "cat"))
}
