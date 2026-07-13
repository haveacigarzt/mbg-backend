package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type SPPGInvitations struct {
	ID        int64        `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	Token     string       `json:"token"`
	ExpiresAt sql.NullTime `json:"expires_at"`
	UsedAt    sql.NullTime `json:"used_at"`
	NamaSPPG  string       `json:"nama_sppg"`
}

type SPPGInvitationsModel struct {
	DB *sql.DB
}

func ValidateSPPGInvitations(v *validator.Validator, invitation *SPPGInvitations) {
	v.Check(invitation.Token != "", "token", "harus diisi")
	v.Check(len(invitation.Token) <= 255, "token", "maksimal 255 karakter")

	v.Check(invitation.NamaSPPG != "", "nama_sppg", "harus diisi")
	v.Check(len(invitation.NamaSPPG) <= 200, "nama_sppg", "maksimal 200 karakter")

	if invitation.ExpiresAt.Valid {
		v.Check(
			invitation.ExpiresAt.Time.After(time.Now()),
			"expires_at",
			"harus berada di masa depan",
		)
	}
}

func (m SPPGInvitationsModel) Insert(invitation *SPPGInvitations) error {
	query := `
		INSERT INTO sppg_invitations (
			token,
			nama_sppg,
			expires_at
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	args := []any{
		invitation.Token,
		invitation.NamaSPPG,
		invitation.ExpiresAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&invitation.ID,
		&invitation.CreatedAt,
	)
}

func (m SPPGInvitationsModel) GetByToken(token string) (*SPPGInvitations, error) {
	query := `
		SELECT
			id,
			created_at,
			token,
			expires_at,
			used_at,
			nama_sppg
		FROM sppg_invitations
		WHERE token = $1
		AND used_at IS NULL
		AND (
			expires_at IS NULL
			OR expires_at > NOW()
		)
	`

	var invitation SPPGInvitations

	err := m.DB.QueryRow(query, token).Scan(
		&invitation.ID,
		&invitation.CreatedAt,
		&invitation.Token,
		&invitation.ExpiresAt,
		&invitation.UsedAt,
		&invitation.NamaSPPG,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &invitation, nil
}

func (m SPPGInvitationsModel) DeleteByToken(token string) error {
	query := `
		DELETE FROM
		sppg_invitations
		WHERE token = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, token)
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

func (m SPPGInvitationsModel) GetByTokenTx(ctx context.Context, tx *sql.Tx, token string) (*SPPGInvitations, error) {
	query := `
		SELECT
			id,
			created_at,
			token,
			expires_at,
			used_at,
			nama_sppg
		FROM sppg_invitations
		WHERE token = $1
		AND used_at IS NULL
		AND (
			expires_at IS NULL
			OR expires_at > NOW()
		)
	`

	var invitation SPPGInvitations

	err := tx.QueryRowContext(
		ctx,
		query, token,
	).Scan(
		&invitation.ID,
		&invitation.CreatedAt,
		&invitation.Token,
		&invitation.ExpiresAt,
		&invitation.UsedAt,
		&invitation.NamaSPPG,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &invitation, nil
}

func (m SPPGInvitationsModel) GetAll(filters Filters) ([]*SPPGInvitations, Metadata, error) {
	query := fmt.Sprintf(`
	SELECT
		count(*) OVER(),
		id,
		created_at,
		token,
		expires_at,
		used_at,
		nama_sppg
	FROM sppg_invitations
	ORDER BY %s %s
	LIMIT $1 OFFSET $2`,
		filters.sortColumn(),
		filters.sortDirection(),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	invitations := []*SPPGInvitations{}

	for rows.Next() {

		var invitation SPPGInvitations

		err := rows.Scan(
			&totalRecords,
			&invitation.ID,
			&invitation.CreatedAt,
			&invitation.Token,
			&invitation.ExpiresAt,
			&invitation.UsedAt,
			&invitation.NamaSPPG,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		invitations = append(invitations, &invitation)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return invitations, metadata, nil

}

func (m SPPGInvitationsModel) MarkAsUsed(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		UPDATE sppg_invitations
		SET used_at = NOW()
		WHERE id = $1
	`
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

// GET /register/:token
// ↓
// GetByToken(token)
// ↓
// Token valid
// ↓
// Tampilkan form registrasi
// ↓
// POST registrasi
// ↓
// Insert pending_sppg
// ↓
// MarkAsUsed(invitation.ID)
// ↓
// Menunggu approval admin
