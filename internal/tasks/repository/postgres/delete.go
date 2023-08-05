package tasks

import (
	"context"

	"github.com/google/uuid"
)

// Delete implements Task.
func (t *task) DeleteTask(id uuid.UUID) error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `delete from tasks where id=$1`
	_, err := t.DB.Exec(stmp, id)
	return err
}
