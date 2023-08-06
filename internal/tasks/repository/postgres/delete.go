package tasks

import (
	"context"

	"github.com/google/uuid"
)

// Delete implements Task.
func (t *task) DeleteTask(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `delete from tasks where id=$1`
	_, err := t.db.ExecContext(ctx, stmp, id)
	return err
}
