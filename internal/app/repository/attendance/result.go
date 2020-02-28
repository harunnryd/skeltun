package attendance

// SkeltunGetAttendance ...
type SkeltunGetAttendance struct {
	AttendanceID   *string `db:"attendance_id" json:"attendance-id"`
	Latitude       *string `db:"latitude" json:"latitude"`
	Longitude      *string `db:"longitude" json:"longitude"`
	UserID         string  `db:"user_id" json:"user-id"`
	Fullname       string  `db:"fullname" json:"fullname"`
	Phone          string  `db:"phone" json:"phone"`
	Email          string  `db:"email" json:"e-mail"`
	AttendanceStat *string `db:"attendance_stat" json:"attendance-stat"`
	FileURI        *string `db:"file_uri" json:"file-uri"`
	CheckinAt      *string `db:"checkin_at" json:"checkin-at"`
	CheckoutAt     *string `db:"checkout_at" json:"checkout-at"`
	Text           *string `db:"text" json:"text"`
	Type           *int    `db:"type" json:"type"`
}
