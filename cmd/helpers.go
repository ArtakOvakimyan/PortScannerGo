package cmd

func countGroups(matches []string) int {
	capturedCount := 0
	for _, match := range matches[1:] {
		if match != "" {
			capturedCount++
		}
	}
	return capturedCount
}
