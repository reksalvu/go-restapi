package models

type Activity struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Schedule_Id int    `gorm:"not null" json:"schedule_id"`
	Activity    string `gorm:"type:varchar(255)" json:"activity"`
}
