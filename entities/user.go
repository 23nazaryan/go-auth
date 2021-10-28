package entities

type User struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(50)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}