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

func (r *GroupRepository) GetGroupsByUser(userID int) ([]models.Group, error) {
	query := `
		SELECT group.group_id, group.name
		FROM groups group
		JOIN group_members members ON group.group_id = members.group_id
		WHERE members.user_id = $1
	`
	rows, err := r.db.Query(context.Background(), query, userID)
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
