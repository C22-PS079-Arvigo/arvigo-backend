package datastruct

import "time"

type (
	FaceShape struct {
		ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name      string    `gorm:"column:name" json:"name"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailFaceShapeTag struct {
		ID          uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name        string    `gorm:"column:name" json:"name"`
		FaceShapeID uint64    `gorm:"column:face_shape_id" json:"face_shape_id"`
		TagID       uint64    `gorm:"column:tag_id" json:"tag_id"`
		CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	FaceShapeMachineLearningPayload struct {
		Image string `json:"image"`
	}

	IsHumanRes struct {
		Result bool `json:"result"`
	}

	FaceTestRes struct {
		IsHuman bool   `json:"is_human"`
		Shape   string `json:"shape"`
	}
)

func (FaceShape) TableName() string {
	return "face_shape"
}

func (DetailFaceShapeTag) TableName() string {
	return "detail_face_shape_tags"
}
