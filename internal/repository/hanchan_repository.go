package repository

import (
	"context"
	"errors"
	"fmt"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HanchanRepository interface {
	CreateWithPlayers(ctx context.Context, hanchan *domain.Hanchan, seating []domain.PlayerSeating) error
	GetByID(ctx context.Context, id int) (*domain.Hanchan, error)
	ListByGroup(ctx context.Context, groupID int) ([]*domain.Hanchan, error)
	AssignPlayer(ctx context.Context, hp *domain.HanchanPlayer) error
	ListPlayers(ctx context.Context, hanchanID int) ([]*domain.HanchanPlayer, error)
	Close(ctx context.Context, hanchanID int, results []domain.HanchanPlayer) error
}

type hanchanRepo struct {
	db *pgxpool.Pool
}

func NewHanchanRepo(db *pgxpool.Pool) HanchanRepository {
	return &hanchanRepo{db: db}
}

func (r *hanchanRepo) CreateWithPlayers(ctx context.Context, hanchan *domain.Hanchan, seating []domain.PlayerSeating) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx,
		`INSERT INTO hanchans (group_id, name, date, base_score, uma) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`,
		hanchan.GroupID, hanchan.Name, hanchan.Date, hanchan.BaseScore, hanchan.Uma,
	).Scan(&hanchan.ID, &hanchan.CreatedAt)

	if err != nil {
		return fmt.Errorf("create hanchan: %w", err)
	}

	for _, seat := range seating {
		query := `INSERT INTO hanchan_players (hanchan_id, player_id, initial_seat) VALUES ($1, $2, $3)`
		_, err := tx.Exec(ctx, query, hanchan.ID, seat.PlayerID, seat.InitialSeat)
		if err != nil {
			return fmt.Errorf("unable to insert hanchan_players %d: %w", seat.PlayerID, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

func (r *hanchanRepo) GetByID(ctx context.Context, id int) (*domain.Hanchan, error) {
	h := &domain.Hanchan{}

	err := r.db.QueryRow(ctx,
		`SELECT id, group_id, name, date, status, base_score, uma, created_at FROM hanchans WHERE id = $1`,
		id,
	).Scan(&h.ID, &h.GroupID, &h.Name, &h.Date, &h.Status, &h.BaseScore, &h.Uma, &h.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get hanchan: %w", err)
	}

	return h, nil
}

func (r *hanchanRepo) ListByGroup(ctx context.Context, groupID int) ([]*domain.Hanchan, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, group_id, name, date, status, base_score, uma, created_at
		FROM hanchans
		WHERE group_id = $1`,
		groupID,
	)
	if err != nil {
		return nil, fmt.Errorf("list hanchans: %w", err)
	}
	defer rows.Close()

	var hanchans []*domain.Hanchan
	for rows.Next() {
		h := &domain.Hanchan{}
		if err := rows.Scan(&h.ID, &h.GroupID, &h.Name, &h.Date, &h.Status, &h.BaseScore, &h.Uma, &h.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan hanchan: %w", err)
		}
		hanchans = append(hanchans, h)
	}
	return hanchans, rows.Err()
}

func (r *hanchanRepo) AssignPlayer(ctx context.Context, hp *domain.HanchanPlayer) error {
	query := `INSERT INTO hanchan_players (hanchan_id, player_id, initial_seat) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, hp.HanchanID, hp.PlayerSeat.PlayerID, hp.PlayerSeat.InitialSeat)
	return err
}

func (r *hanchanRepo) ListPlayers(ctx context.Context, hanchanID int) ([]*domain.HanchanPlayer, error) {
	rows, err := r.db.Query(ctx,
		`SELECT hanchan_id, player_id, initial_seat, final_score, placement
		FROM hanchan_players
		WHERE hanchan_id = $1`,
		hanchanID,
	)
	if err != nil {
		return nil, fmt.Errorf("list hanchan players: %w", err)
	}
	defer rows.Close()

	var players []*domain.HanchanPlayer
	for rows.Next() {
		hp := &domain.HanchanPlayer{}
		if err := rows.Scan(&hp.HanchanID, &hp.PlayerSeat.PlayerID, &hp.PlayerSeat.InitialSeat, &hp.FinalScore, &hp.Placement); err != nil {
			return nil, fmt.Errorf("scan hanchan player: %w", err)
		}
		players = append(players, hp)
	}
	return players, rows.Err()
}

func (r *hanchanRepo) Close(ctx context.Context, hanchanID int, results []domain.HanchanPlayer) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	// Write final_score and placement for each player.
	for _, r := range results {
		tag, err := tx.Exec(ctx, `
				UPDATE hanchan_players
				SET final_score = $1, placement = $2
				WHERE hanchan_id = $3 AND player_id = $4`,
			r.FinalScore, r.Placement, hanchanID, r.PlayerSeat.PlayerID,
		)
		if err != nil {
			return fmt.Errorf("update hanchan_player %d: %w", r.PlayerSeat.PlayerID, err)
		}
		if tag.RowsAffected() == 0 {
			return fmt.Errorf("player %d not found in hanchan: %w", r.PlayerSeat.PlayerID, ErrNotFound)
		}
	}

	// Mark the hanchan as closed.
	_, err = tx.Exec(ctx, `UPDATE hanchans SET status = 'CLOSED' WHERE id = $1`, hanchanID)
	if err != nil {
		return fmt.Errorf("close hanchan: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
