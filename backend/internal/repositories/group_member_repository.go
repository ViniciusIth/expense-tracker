package repositories

import (
	"context"
	"fmt"

	"github.com/ViniciusIth/expanse_tracker/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupMemberRepository struct {
	db *pgxpool.Pool
}

func NewGroupMemberRepository(db *pgxpool.Pool) *GroupMemberRepository {
	return &GroupMemberRepository{db: db}
}

func (r *GroupMemberRepository) AddUserToGroup(groupID, userID int) error {
	query := `
		INSERT INTO expense_tracker.group_members (group_id, user_id)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(context.Background(), query, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %w", err)
	}
	return nil
}

func (r *GroupMemberRepository) RemoveUserFromGroup(groupID, userID int) error {
	query := `
		DELETE FROM expense_tracker.group_members
		WHERE group_id = $1 AND user_id = $2
	`
	_, err := r.db.Exec(context.Background(), query, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %w", err)
	}
	return nil
}

func (r *GroupMemberRepository) GetGroupMembers(groupID int) ([]models.GroupMember, error) {
	query := `
		SELECT group_id, user_id, joined_at
		FROM expense_tracker.group_members
		WHERE group_id = $1
	`
	rows, err := r.db.Query(context.Background(), query, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}
	defer rows.Close()

	var members []models.GroupMember
	for rows.Next() {
		var member models.GroupMember
		err := rows.Scan(&member.GroupID, &member.UserID, &member.JoinedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group member: %w", err)
		}
		members = append(members, member)
	}
	return members, nil
}
