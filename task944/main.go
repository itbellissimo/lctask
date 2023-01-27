package main

func main() {

}

func minDeletionSize(strs []string) int {
	n := len(strs)    // number of words
	m := len(strs[0]) // length of word
	count := 0

	for j := 0; j < m; j++ {
		for i := 0; i < n-1; i++ {
			if strs[i][j] > strs[i+1][j] {
				count++
				break
			}
		}
	}

	return count
}
