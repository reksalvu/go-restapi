package models

type Schedule struct {
	Id         int        `gorm:"primaryKey" json:"id"`
	Title      string     `gorm:"type:varchar(255)" json:"title"`
	Start_At   string     `gorm:"type:timestamp with time zone" json:"start_at"`
	End_At     string     `gorm:"type:timestamp with time zone" json:"end_at"`
	Activities []Activity `gorm:"foreignKey:Schedule_Id;constraint:OnDelete:CASCADE" json:"activities"`
}
