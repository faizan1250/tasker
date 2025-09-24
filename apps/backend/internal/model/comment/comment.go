package comment

import (
	"github.com/faizan1250/tasker/internal/model"
	"github.com/google/uuid"
)

type Comment struct {
	model.Base
	TodoID  uuid.UUID `json:"todoId" db:"todo_id"`
	UserID  string    `json:"userId" db:"user_id"`
	Comment string    `json:"comment" db:"comment"`
}
