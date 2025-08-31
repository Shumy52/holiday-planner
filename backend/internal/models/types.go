package models

import "time"

type Vacation struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"userId"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	EndDate   time.Time `db:"end_date" json:"endDate"`
	TotalDays int       `db:"total_days" json:"totalDays"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
