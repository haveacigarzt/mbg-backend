package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"regexp"
	"strings"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      password  `json:"-"`
	Activated     bool      `json:"activated"`
	RoleID        int64     `json:"role_id"`
	Version       int       `json:"-"`
	TerakhirAktif string    `json:"terakhir_aktif"`
	Jenis         string    `json:"jenis,omitempty"`
}

var AnonymousUser = &User{}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

var (
	hasLetter = regexp.MustCompile(`[A-Za-z]`)
	hasNumber = regexp.MustCompile(`[0-9]`)
)

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(len(password) >= 8, "password", "Password minimal 8 karakter")
	v.Check(len(password) <= 15, "password", "Password maksimal 15 karakter")
	v.Check(hasLetter.MatchString(password), "password", "Password harus mengandung huruf")
	v.Check(hasNumber.MatchString(password), "password", "Password harus mengandung angka")
}

type AkunInput struct {
	Name     string `json:"name"`
	RoleID   int64  `json:"role_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateAkunInput(v *validator.Validator, input AkunInput) {
	v.Check(strings.TrimSpace(input.Name) != "", "name", "Nama harus diisi")
	v.Check(len(strings.TrimSpace(input.Name)) >= 3, "name", "Nama minimal 3 karakter")

	v.Check(input.RoleID == 2 || input.RoleID == 3,
		"role_id",
		"Role harus Stakeholder atau SPPG",
	)

	v.Check(input.Email != "", "email", "Email harus diisi")
	v.Check(validator.Matches(input.Email, validator.EmailRX),
		"email",
		"Format email tidak valid",
	)

	v.Check(input.Password != "", "password", "Password harus diisi")
	v.Check(len(input.Password) >= 8,
		"password",
		"Password minimal 8 karakter",
	)
	v.Check(len(input.Password) <= 15,
		"password",
		"Password maksimal 15 karakter",
	)
	ValidatePassword(v, input.Password)
}

type AkunUpdate struct {
	ID        int64 `json:"id"`
	Activated *bool `json:"activated"`
}

func ValidateAkunUpdate(v *validator.Validator, input AkunUpdate) {
	v.Check(input.ID > 0, "id", "must be provided")
	v.Check(input.Activated != nil, "activated", "must be provided")
}

type AkunDelete struct {
	ID int64 `json:"id"`
}

func ValidateAkunDelete(v *validator.Validator, input AkunDelete) {
	v.Check(input.ID > 0, "id", "must be provided")
}

func (m UserModel) Insert(user *User) error {
	query := `
INSERT INTO users (name, email, password_hash, activated, role_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, version`
	args := []any{user.Name, user.Email, user.Password.hash, user.Activated, user.RoleID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrDuplicateEmail
			}
		}

		return err
	}
	return nil
}

func (m UserModel) InsertTx(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (name, email, password_hash, activated, role_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version
	`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
		user.RoleID,
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Version,
	)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) InsertAndGetID(user *User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password_hash,
			role_id,
			activated
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version
	`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.RoleID,
		user.Activated,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Version,
	)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrDuplicateEmail
			}
		}

		return err
	}
	return nil
}

func (m UserModel) InsertAndGetIDTx(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password_hash,
			role_id,
			activated
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version
	`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.RoleID,
		user.Activated,
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Version,
	)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrDuplicateEmail
			}
		}

		return err
	}
	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
SELECT id, created_at, name, email, password_hash, activated, version, role_id
FROM users
WHERE email = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
		&user.RoleID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (m UserModel) Get(id int64) (*User, error) {
	query := `
SELECT id, created_at, name, email, password_hash, activated, version, role_id
FROM users
WHERE id = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
		&user.RoleID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (m UserModel) Update(user *User) error {
	query := `
UPDATE users
SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
WHERE id = $5 AND version = $6
RETURNING version`
	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
		user.ID,
		user.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m UserModel) Delete(user *User) error {
	query := `
		UPDATE users
		SET
			deleted_at = NOW(),
			version = version + 1
		WHERE id = $1
		AND version = $2
		AND deleted_at IS NULL
		RETURNING version`

	args := []any{
		user.ID,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) DeleteTx(ctx context.Context, tx *sql.Tx, id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
DELETE FROM users
WHERE id = $1`

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
	// Calculate the SHA-256 hash of the plaintext token provided by the client.
	// Remember that this returns a byte *array* with length 32, not a slice.
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))
	// Set up the SQL query.
	query := `
SELECT users.id, users.created_at, users.name, users.email, users.password_hash, users.activated, users.version, users.role_id
FROM users
INNER JOIN tokens
ON users.id = tokens.user_id
WHERE tokens.hash = $1
AND tokens.scope = $2
AND tokens.expiry > $3`
	// Create a slice containing the query arguments. Notice how we use the [:] operator
	// to get a slice containing the token hash, rather than passing in the array (which
	// is not supported by the pq driver), and that we pass the current time as the
	// value to check against the token expiry.
	args := []any{tokenHash[:], tokenScope, time.Now()}
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Execute the query, scanning the return values into a User struct. If no matching
	// record is found we return an ErrRecordNotFound error.
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
		&user.RoleID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Return the matching user.
	return &user, nil
}

func (m UserModel) UpdateLastActive(id int64) error {
	query := `
		UPDATE users
		SET last_active_at = NOW()
		WHERE id = $1
		AND (
			last_active_at IS NULL
			OR last_active_at < NOW() - INTERVAL '5 minutes'
		)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func (m UserModel) GetAll(name string, status string, roleID string, filters Filters) ([]*User, Metadata, error) {

	query := `
	SELECT
    count(*) OVER(),
    u.id,
    COALESCE(s.nama, u.name) AS name,
    u.email,
    u.activated,
    u.created_at,
    u.role_id,
    COALESCE(
        TO_CHAR(u.last_active_at, 'YYYY-MM-DD HH24:MI:SS'),
        ''
    ) AS terakhir_aktif,
    r.name AS role_name
FROM users u
JOIN roles r ON r.id = u.role_id
LEFT JOIN sppg s
    ON s.user_id = u.id
    AND u.role_id = 3
`

	var args []any
	var whereClauses []string

	whereClauses = append(whereClauses, "u.deleted_at IS NULL")
	whereClauses = append(whereClauses, "u.role_id IN (2, 3)")

	if status == "aktif" {
		whereClauses = append(whereClauses, "u.activated = true")
	}

	if status == "nonaktif" {
		whereClauses = append(whereClauses, "u.activated = false")
	}

	if roleID != "" {
		args = append(args, roleID)
		whereClauses = append(whereClauses, fmt.Sprintf("u.role_id = $%d", len(args)))
	}

	if name != "" {
		args = append(args, "%"+strings.ToLower(name)+"%")
		whereClauses = append(
			whereClauses,
			fmt.Sprintf("LOWER(u.name) LIKE $%d", len(args)),
		)
	}

	query += " WHERE " + strings.Join(whereClauses, " AND ")

	// ======================
	// SORTING (aman dari safelist)
	// ======================

	sortColumn := "u.id"
	sortDirection := "DESC"

	switch filters.Sort {
	case "id":
		sortColumn = "u.id"
	case "-id":
		sortColumn = "u.id"
		sortDirection = "DESC"
	case "name":
		sortColumn = "u.name"
	case "-name":
		sortColumn = "u.name"
		sortDirection = "DESC"
	case "email":
		sortColumn = "u.email"
	case "-email":
		sortColumn = "u.email"
		sortDirection = "DESC"
	case "role_id":
		sortColumn = "u.role_id"
	case "-role_id":
		sortColumn = "u.role_id"
		sortDirection = "DESC"
	case "terakhir_aktif":
		sortColumn = "u.last_active_at"
	case "-terakhir_aktif":
		sortColumn = "u.last_active_at"
		sortDirection = "DESC"
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortColumn, sortDirection)

	// ======================
	// PAGINATION
	// ======================

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d",
		len(args)+1,
		len(args)+2,
	)

	args = append(args, filters.limit(), filters.offset())

	// ======================
	// EXEC QUERY
	// ======================

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	// ======================
	// SCAN RESULT
	// ======================

	totalRecords := 0
	users := []*User{}

	for rows.Next() {
		var u User
		var terakhirAktif string

		err := rows.Scan(
			&totalRecords,
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Activated,
			&u.CreatedAt,
			&u.RoleID,
			&terakhirAktif,
			&u.Jenis,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		u.TerakhirAktif = terakhirAktif
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return users, metadata, nil
}

type UsersSummary struct {
	Total       int `json:"total"`
	SPPG        int `json:"sppg"`
	Stakeholder int `json:"stakeholder"`
	Nonaktif    int `json:"nonaktif"`
}

func (m UserModel) GetSummary() (UsersSummary, error) {
	query := `
	SELECT
		COUNT(*) FILTER (
			WHERE role_id IN (2, 3)
		) AS total,
		COUNT(*) FILTER (
			WHERE role_id = 3
			AND activated = true
		) AS sppg,
		COUNT(*) FILTER (
			WHERE role_id = 2
			AND activated = true
		) AS stakeholder,
		COUNT(*) FILTER (
			WHERE activated = false
		) AS nonaktif
	FROM users
	WHERE deleted_at IS NULL
`

	var summary UsersSummary

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query).Scan(
		&summary.Total,
		&summary.SPPG,
		&summary.Stakeholder,
		&summary.Nonaktif,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return UsersSummary{}, ErrRecordNotFound
		default:
			return UsersSummary{}, err
		}
	}

	return summary, nil
}
