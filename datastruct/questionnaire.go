package datastruct

type Questionnaire struct {
	ID       uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Type     string `gorm:"column:type" json:"type"`
	Question string `gorm:"column:question" json:"question"`
	// CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Questionnaire) TableName() string {
	return "personality_questionnaires"
}
