package providers

// GetString returns the string value of a pointer to string or a default empty string.
func GetString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// GetInt64 returns the int value of a pointer to int64 or a default of 0.
func GetInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}
