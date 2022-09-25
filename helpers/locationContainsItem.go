package helpers

func LocationContainsItem(object string, items []string) bool {
	for _, item := range items {
		if object == item {
			return true
		}
	}
	return false
}
