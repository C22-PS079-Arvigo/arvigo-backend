package datastruct

type Questionnaire struct {
	ID       uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Type     string `gorm:"column:type" json:"type"`
	Question string `gorm:"column:question" json:"question"`
}

func (Questionnaire) TableName() string {
	return "personality_questionnaires"
}
