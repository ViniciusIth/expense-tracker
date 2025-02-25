package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository struct {
	db *pgxpool.Pool
}

func NewGroupRepository(db *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(group *models.Group) error {
	query := `
		INSERT INTO groups (name)
		VALUES ($1)
		RETURNING group_id
	`
	err := r.db.QueryRow(context.Background(), query, group.Name).
		Scan(&group.GroupID)
	if err != nil {
		return fmt.Errorf("failed to create group: %w", err)
	}
	return nil
}

func (r *GroupRepository) GetGroupByID(groupID int) (*models.Group, error) {
	group := &models.Group{}
	query := `
		SELECT group_id, name
		FROM groups
		WHERE group_id = $1
	`
	err := r.db.QueryRow(context.Background(), query, groupID).
		Scan(&group.GroupID, &group.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	return group, nil
}

// Probably won't be useful in the future, maybe change to get all user repositories
func (r *GroupRepository) GetAllGroups() ([]models.Group, error) {
	query := `
		SELECT group_id, name
		FROM groups
	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.GroupID, &group.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}
