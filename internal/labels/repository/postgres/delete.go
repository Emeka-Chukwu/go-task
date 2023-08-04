package label

import (
	"context"

	"github.com/google/uuid"
)

// Delete implements Label.
func (lab *label) Delete(id uuid.UUID) error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmp := `delete from labels where id=$1`
	_, err := lab.DB.Exec(stmp, id)
	return err
}
