package datastruct

import "time"

type User struct {
	ID                        uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email                     string     `gorm:"column:email" json:"email"`
	Password                  string     `gorm:"column:password" json:"-"`
	RoleID                    uint64     `gorm:"column:role_id" json:"role_id"`
	FullName                  string     `gorm:"column:full_name" json:"full_name"`
	Gender                    string     `gorm:"column:gender" json:"gender"`
	DateOfBirth               *time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	PlaceOfBirth              string     `gorm:"column:place_of_birth" json:"place_of_birth"`
	IsCompletePersonalityTest bool       `gorm:"column:is_complete_personality_test" json:"is_complete_personality_test"`
	IsCompleteFaceTest        bool       `gorm:"column:is_complete_face_test" json:"is_complete_face_test"`
	PersonalityID             bool       `gorm:"column:personality_id" json:"personality_id"`
	FaceShapeID               bool       `gorm:"column:face_shape_id" json:"face_shape_id"`
	IsVerified                bool       `gorm:"column:is_verified" json:"is_verified"`
	Avatar                    string     `gorm:"column:avatar" json:"avatar"`
	AddressID                 uint64     `gorm:"column:addresses_id" json:"addresses_id"`
	MerchantID                uint64     `gorm:"column:merchant_id" json:"merchant_id"`
	CreatedAt                 time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserSubscription struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID            int       `gorm:"column:user_id" json:"user_id"`
	Price             int       `gorm:"column:price" json:"price"`
	UniqueCode        int       `gorm:"column:unique_code" json:"unique_code"`
	SubscriptionStart time.Time `gorm:"column:subscription_start" json:"subscription_start"`
	SubscriptionEnd   time.Time `gorm:"column:subscription_end" json:"subscription_end"`
	IsVerified        int       `gorm:"column:is_verified" json:"is_verified"`
	Message           string    `gorm:"column:message" json:"message"`
	PaidAt            time.Time `gorm:"column:paid_at" json:"paid_at"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (UserSubscription) TableName() string {
	return "detail_user_subscriptions"
}

type UserAuth struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	RoleID   uint64 `json:"role_id"`
}
