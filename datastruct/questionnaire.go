package datastruct

type (
	Questionnaire struct {
		ID       uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Type     string `gorm:"column:type" json:"type"`
		Question string `gorm:"column:question" json:"question"`
	}

	QuestionnaireRequest struct {
		EXT1  int `json:"EXT1"`
		EXT2  int `json:"EXT2"`
		EXT3  int `json:"EXT3"`
		EXT4  int `json:"EXT4"`
		EXT5  int `json:"EXT5"`
		EXT6  int `json:"EXT6"`
		EXT7  int `json:"EXT7"`
		EXT8  int `json:"EXT8"`
		EXT9  int `json:"EXT9"`
		EXT10 int `json:"EXT10"`

		EST1  int `json:"EST1"`
		EST2  int `json:"EST2"`
		EST3  int `json:"EST3"`
		EST4  int `json:"EST4"`
		EST5  int `json:"EST5"`
		EST6  int `json:"EST6"`
		EST7  int `json:"EST7"`
		EST8  int `json:"EST8"`
		EST9  int `json:"EST9"`
		EST10 int `json:"EST10"`

		AGR1  int `json:"AGR1"`
		AGR2  int `json:"AGR2"`
		AGR3  int `json:"AGR3"`
		AGR4  int `json:"AGR4"`
		AGR5  int `json:"AGR5"`
		AGR6  int `json:"AGR6"`
		AGR7  int `json:"AGR7"`
		AGR8  int `json:"AGR8"`
		AGR9  int `json:"AGR9"`
		AGR10 int `json:"AGR10"`

		CSN1  int `json:"CSN1"`
		CSN2  int `json:"CSN2"`
		CSN3  int `json:"CSN3"`
		CSN4  int `json:"CSN4"`
		CSN5  int `json:"CSN5"`
		CSN6  int `json:"CSN6"`
		CSN7  int `json:"CSN7"`
		CSN8  int `json:"CSN8"`
		CSN9  int `json:"CSN9"`
		CSN10 int `json:"CSN10"`

		OPN1  int `json:"OPN1"`
		OPN2  int `json:"OPN2"`
		OPN3  int `json:"OPN3"`
		OPN4  int `json:"OPN4"`
		OPN5  int `json:"OPN5"`
		OPN6  int `json:"OPN6"`
		OPN7  int `json:"OPN7"`
		OPN8  int `json:"OPN8"`
		OPN9  int `json:"OPN9"`
		OPN10 int `json:"OPN10"`
	}

	PersonalityPercentages struct {
		Agreeable     float64 `json:"percentage_of_agreeable"`
		Conscientious float64 `json:"percentage_of_conscientious"`
		Extraversion  float64 `json:"percentage_of_extraversion"`
		Neurotic      float64 `json:"percentage_of_neurotic"`
		Openness      float64 `json:"percentage_of_openess"`
	}
)

func (Questionnaire) TableName() string {
	return "personality_questionnaires"
}
