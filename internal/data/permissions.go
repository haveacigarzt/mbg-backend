package data

import (
	"context"
	"database/sql"
	"slices"
	"time"

	"github.com/lib/pq"
)

// Define a Permissions slice, which we will use to hold the permission codes (like
// "sekolah:read" and "sekolah:write") for a single user.
type Permissions []string

// Add a helper method to check whether the Permissions slice contains a specific
// permission code.
func (p Permissions) Include(code string) bool {
	return slices.Contains(p, code)
}

// Define the PermissionModel type.
type PermissionModel struct {
	DB *sql.DB
}

func (m PermissionModel) GetAllForUser(roleID int64) (Permissions, error) {
	query := `
		SELECT permissions
		FROM roles
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var raw []string

	err := m.DB.QueryRowContext(ctx, query, roleID).Scan(
		pq.Array(&raw),
	)
	if err != nil {
		return nil, err
	}

	return Permissions(raw), nil
}

// func (m PermissionModel) AddForUser(userID int64, codes ...string) error {
// 	query := `
// INSERT INTO users_permissions
// SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)`
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
// 	_, err := m.DB.ExecContext(ctx, query, userID, pq.Array(codes))
// 	return err
// }
