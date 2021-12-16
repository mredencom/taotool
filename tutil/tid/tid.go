package tid

import (
	"github.com/google/uuid"
)

// TId id生成器
type TId struct {
}

func (TId) RandomUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
