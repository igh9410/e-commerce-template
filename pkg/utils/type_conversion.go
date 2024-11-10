package utils

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// UUIDToPgtype converts uuid.UUID to pgtype.UUID
func UUIDToPgtype(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

// PgtypeToUUID converts pgtype.UUID to uuid.UUID
func PgtypeToUUID(id pgtype.UUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}
	return id.Bytes
}

// StringToPgtype converts string to pgtype.Text
func StringToPgtype(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

// PgtypeToString converts pgtype.Text to string
func PgtypeToString(t pgtype.Text) string {
	if !t.Valid {
		return ""
	}
	return t.String
}

// TimeToPgtype converts time.Time to pgtype.Timestamp
func TimeToPgtype(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: t, Valid: !t.IsZero()}
}

// PgtypeToTime converts pgtype.Timestamp to time.Time
func PgtypeToTime(t pgtype.Timestamp) time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return t.Time
}

// Int32ToPgtype converts int32 to pgtype.Int4
func Int32ToPgtype(i int32) pgtype.Int4 {
	return pgtype.Int4{Int32: i, Valid: true}
}

// PgtypeToInt32 converts pgtype.Int4 to int32
func PgtypeToInt32(i pgtype.Int4) int32 {
	if !i.Valid {
		return 0
	}
	return i.Int32
}

// Int64ToPgtype converts int64 to pgtype.Int8
func Int64ToPgtype(i int64) pgtype.Int8 {
	return pgtype.Int8{Int64: i, Valid: true}
}

// PgtypeToInt64 converts pgtype.Int8 to int64
func PgtypeToInt64(i pgtype.Int8) int64 {
	if !i.Valid {
		return 0
	}
	return i.Int64
}

// Float64ToPgtype converts float64 to pgtype.Float8
func Float64ToPgtype(f float64) pgtype.Float8 {
	return pgtype.Float8{Float64: f, Valid: true}
}

// PgtypeToFloat64 converts pgtype.Float8 to float64
func PgtypeToFloat64(f pgtype.Float8) float64 {
	if !f.Valid {
		return 0
	}
	return f.Float64
}

// BoolToPgtype converts bool to pgtype.Bool
func BoolToPgtype(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

// PgtypeToBool converts pgtype.Bool to bool
func PgtypeToBool(b pgtype.Bool) bool {
	if !b.Valid {
		return false
	}
	return b.Bool
}

// StringToPgtypeOrNull converts a *string to pgtype.Text, handling null values.
func StringToPgtypeOrNull(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

// PgtypeToStringOrNull converts pgtype.Text to *string, handling null values.
func PgtypeToStringOrNull(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}

// TimeToPgtypeOrNull converts a *time.Time to pgtype.Timestamp, handling null values.
func TimeToPgtypeOrNull(t *time.Time) pgtype.Timestamp {
	if t == nil {
		return pgtype.Timestamp{Valid: false}
	}
	return pgtype.Timestamp{Time: *t, Valid: true}
}

// PgtypeToTimeOrNull converts pgtype.Timestamp to *time.Time, handling null values.
func PgtypeToTimeOrNull(t pgtype.Timestamp) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// TimeToPgtypeTimestamptz converts time.Time to pgtype.Timestamptz
func TimeToPgtypeTimestamptz(t time.Time) pgtype.Timestamptz {
	if t.IsZero() {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: t, Valid: true}
}

// TimeToPgtypeTimestamptzOrNull converts *time.Time to pgtype.Timestamptz, handling null values
func TimeToPgtypeTimestamptzOrNull(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: *t, Valid: true}
}

// PgtypeTimestamptzToTime converts pgtype.Timestamptz to time.Time
func PgtypeTimestamptzToTime(t pgtype.Timestamptz) time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return t.Time
}

// PgtypeTimestamptzToTimeOrNull converts pgtype.Timestamptz to *time.Time, handling null values
func PgtypeTimestamptzToTimeOrNull(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// StringToInt64 converts a string to int64, returning 0 on error
func StringToInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0 // Return default value if conversion fails
	}
	return value
}

// StringToInt32 converts a string to int32, returning 0 on error
func StringToInt32(s string) int32 {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0 // Return default value if conversion fails
	}
	return int32(value)
}

// StringToFloat64 converts a string to float64, returning 0.0 on error
func StringToFloat64(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0 // Return default value if conversion fails
	}
	return value
}

// StringToFloat32 converts a string to float32, returning 0.0 on error
func StringToFloat32(s string) float32 {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0 // Return default value if conversion fails
	}
	return float32(value)
}

// StringToBool converts a string to a bool, returning false on error
func StringToBool(s string) bool {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return false // Return default value if conversion fails
	}
	return value
}

// StringToTime converts a string to time.Time, returning time.Time{} on error
func StringToTime(s, layout string) time.Time {
	// Default to RFC3339 if no layout is provided
	if layout == "" {
		layout = time.RFC3339
	}
	value, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{} // Return default zero time value if conversion fails
	}
	return value
}

// StringToUint64 converts a string to uint64, returning 0 on error
func StringToUint64(s string) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0 // Return default value if conversion fails
	}
	return value
}

// StringToUint32 converts a string to uint32, returning 0 on error
func StringToUint32(s string) uint32 {
	value, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0 // Return default value if conversion fails
	}
	return uint32(value)
}
