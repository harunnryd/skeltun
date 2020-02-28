package attendance

const (
	skeltunGetAttendance = `
	SELECT
		attendance.id AS attendance_id,
		attendance.latitude,
		attendance.longitude,
		usr.id AS user_id,
		usr.fullname,
		usr.phone,
		usr.email,
		LOWER(attendance.attendanceable_type) AS attendance_stat,
	(SELECT file_uri FROM attachments WHERE attachmentable_id = attendance.id AND attachmentable_type = 'Attendance') AS file_uri,
	(SELECT checkin_at FROM presents WHERE id = attendance.attendanceable_id AND attendance.attendanceable_type = 'Present') AS checkin_at,
	(SELECT checkout_at FROM presents WHERE id = attendance.attendanceable_id AND attendance.attendanceable_type = 'Present') AS checkout_at,
	(SELECT text FROM permits WHERE id = attendance.attendanceable_id AND attendance.attendanceable_type = 'Permit') AS text,
	(SELECT type FROM permits WHERE id = attendance.attendanceable_id AND attendance.attendanceable_type = 'Permit') AS type
	FROM (SELECT * FROM attendances WHERE CAST(created_at AS DATE) = :date) AS attendance
	RIGHT JOIN users usr ON attendance.user_id = usr.id
	WHERE CAST(usr.created_at AS DATE) <= :date -- condition 'user must regist before current date'
	ORDER BY usr.fullname
	`
)
