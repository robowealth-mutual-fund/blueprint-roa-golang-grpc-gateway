package entity

type Users struct {
	ID          int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	FullName    string `gorm:"column:full_name"`
	Address     string `gorm:"column:address"`
	PhoneNumber string `gorm:"column:phone_number"`
	Gender      string `gorm:"column:gender"`
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
	CreatedBy   string `gorm:"column:created_by"`
	UpdatedBy   string `gorm:"column:updated_by"`
}

func (Users) TableName() string {
	return "users"
}
