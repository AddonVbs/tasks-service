package task

type Task struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task string `json:"task"`
	//ALTER
	UserID int `json:"user_id"`
}
