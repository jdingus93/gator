// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: unfollow.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE FROM feed_follows WHERE user_id = $1 and feed_id = $2
`

type UnfollowFeedParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.UserID, arg.FeedID)
	return err
}
