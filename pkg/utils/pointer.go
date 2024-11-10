package utils

import "strconv"

// ToStringPointer converts a string to *string
func ToStringPointer(s string) *string {
	return &s
}

// ToIntPointer converts an int to *int
func ToIntPointer(i int) *int {
	return &i
}

// ToInt64Pointer converts an int64 to *int64
func ToInt64Pointer(i int64) *int64 {
	return &i
}

// ToBoolPointer converts a bool to *bool
func ToBoolPointer(b bool) *bool {
	return &b
}

// ToFloat64Pointer converts a float64 to *float64
func ToFloat64Pointer(f float64) *float64 {
	return &f
}

// FromStringPointer converts *string to string, returns default value if nil
func FromStringPointer(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// FromIntPointer converts *int to int, returns default value if nil
func FromIntPointer(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// FromInt64Pointer converts *int64 to int64, returns default value if nil
func FromInt64Pointer(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// FromBoolPointer converts *bool to bool, returns default value if nil
func FromBoolPointer(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// FromFloat64Pointer converts *float64 to float64, returns default value if nil
func FromFloat64Pointer(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// Int64ToStringPointer converts an int64 to *string
func Int64ToStringPointer(i int64) *string {
	s := strconv.FormatInt(i, 10)
	return &s
}
