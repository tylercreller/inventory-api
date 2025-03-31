package pubsub

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Notifier struct {
	db *bun.DB
}

func NewNotifier(db *bun.DB) *Notifier {
	return &Notifier{db: db}
}

func Notify(ctx context.Context, db *bun.DB, channel, payload string) error {
	err := pgdriver.Notify(ctx, db, channel, payload)
	return err
}
