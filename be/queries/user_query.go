package queries

const (
	InsertUser GoQuery = `
		INSERT INTO users (account_id, email, password, full_name, user_name, date_of_birth, status, gender, profile_picture_url)
		VALUES 
		    (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	SelectUserByEmail GoQuery = `
		SELECT id, account_id, email, password, full_name, user_name, date_of_birth, status, gender, profile_picture_url, created_at, updated_at, deleted_at
		FROM
		    users
		WHERE email = ?
	`
)
