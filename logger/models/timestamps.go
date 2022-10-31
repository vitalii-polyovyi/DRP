package models

import "time"

type Timestamps struct {
	CreatedAt int64  `bson:"created_at,omitempty" json:"created_at" validate:"omitempty,gte=0"`
	UpdatedAt *int64 `bson:"updated_at,omitempty" json:"update_at" validate:"omitempty,gte=0"`
}

func (m *Timestamps) SetOnCreate() {
	t := time.Now().Unix()
	m.CreatedAt = t
	m.UpdatedAt = &t
}

func (m *Timestamps) SetOnUpdate() {
	updatedAt := time.Now().Unix()
	m.UpdatedAt = &updatedAt
}
