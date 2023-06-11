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
	PersonalityID             uint64     `gorm:"column:personality_id" json:"personality_id"`
	FaceShapeID               uint64     `gorm:"column:face_shape_id" json:"face_shape_id"`
	IsVerified                bool       `gorm:"column:is_verified" json:"is_verified"`
	Avatar                    string     `gorm:"column:avatar" json:"avatar"`
	AddressID                 uint64     `gorm:"column:addresses_id" json:"addresses_id"`
	MerchantID                uint64     `gorm:"column:merchant_id" json:"merchant_id"`
	CreatedAt                 time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type UserWithPersonalityTag struct {
	ID                        uint64 `gorm:"column:id"`
	IsCompletePersonalityTest bool   `gorm:"column:is_complete_personality_test"`
	IsCompleteFaceTest        bool   `gorm:"column:is_complete_face_test"`
	PersonalityID             uint64 `gorm:"column:personality_id"`
	FaceShapeID               uint64 `gorm:"column:face_shape_id"`
	TagID                     string `gorm:"column:tag_ids"`
	TagIDs                    []uint64
}

func (User) TableName() string {
	return "users"
}

type UserSubscription struct {
	ID                uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID            uint64     `gorm:"column:user_id" json:"user_id"`
	Price             uint64     `gorm:"column:price" json:"price"`
	UniqueCode        uint64     `gorm:"column:unique_code" json:"unique_code"`
	SubscriptionStart *time.Time `gorm:"column:subscription_start" json:"subscription_start"`
	SubscriptionEnd   *time.Time `gorm:"column:subscription_end" json:"subscription_end"`
	Status            string     `gorm:"column:status" json:"status"`
	Message           string     `gorm:"column:message" json:"message"`
	Bank              string     `gorm:"column:bank" json:"bank"`
	PaidAt            *time.Time `gorm:"column:paid_at" json:"paid_at"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (UserSubscription) TableName() string {
	return "detail_user_subscriptions"
}

type UserSubscriptionProduct struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ProductID      uint64    `gorm:"column:product_id" json:"product_id"`
	SubscriptionID uint64    `gorm:"column:subscription_id" json:"subscription_id"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (UserSubscriptionProduct) TableName() string {
	return "detail_user_subscription_products"
}

type UserAuth struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	RoleID   uint64 `json:"role_id"`
}

type UserPersonality struct {
	ID        uint64    `gorm:"column:id;primary_key"`
	UserID    uint64    `gorm:"column:user_id"`
	IsActive  int       `gorm:"column:is_active"`
	TagIDs    string    `gorm:"column:tag_ids"`
	ExtResult float64   `gorm:"column:ext_result"`
	EstResult float64   `gorm:"column:est_result"`
	AgrResult float64   `gorm:"column:agr_result"`
	CsnResult float64   `gorm:"column:csn_result"`
	OpnResult float64   `gorm:"column:opn_result"`
	EXT1      int       `gorm:"column:EXT1"`
	EXT2      int       `gorm:"column:EXT2"`
	EXT3      int       `gorm:"column:EXT3"`
	EXT4      int       `gorm:"column:EXT4"`
	EXT5      int       `gorm:"column:EXT5"`
	EXT6      int       `gorm:"column:EXT6"`
	EXT7      int       `gorm:"column:EXT7"`
	EXT8      int       `gorm:"column:EXT8"`
	EXT9      int       `gorm:"column:EXT9"`
	EXT10     int       `gorm:"column:EXT10"`
	EST1      int       `gorm:"column:EST1"`
	EST2      int       `gorm:"column:EST2"`
	EST3      int       `gorm:"column:EST3"`
	EST4      int       `gorm:"column:EST4"`
	EST5      int       `gorm:"column:EST5"`
	EST6      int       `gorm:"column:EST6"`
	EST7      int       `gorm:"column:EST7"`
	EST8      int       `gorm:"column:EST8"`
	EST9      int       `gorm:"column:EST9"`
	EST10     int       `gorm:"column:EST10"`
	AGR1      int       `gorm:"column:AGR1"`
	AGR2      int       `gorm:"column:AGR2"`
	AGR3      int       `gorm:"column:AGR3"`
	AGR4      int       `gorm:"column:AGR4"`
	AGR5      int       `gorm:"column:AGR5"`
	AGR6      int       `gorm:"column:AGR6"`
	AGR7      int       `gorm:"column:AGR7"`
	AGR8      int       `gorm:"column:AGR8"`
	AGR9      int       `gorm:"column:AGR9"`
	AGR10     int       `gorm:"column:AGR10"`
	CSN1      int       `gorm:"column:CSN1"`
	CSN2      int       `gorm:"column:CSN2"`
	CSN3      int       `gorm:"column:CSN3"`
	CSN4      int       `gorm:"column:CSN4"`
	CSN5      int       `gorm:"column:CSN5"`
	CSN6      int       `gorm:"column:CSN6"`
	CSN7      int       `gorm:"column:CSN7"`
	CSN8      int       `gorm:"column:CSN8"`
	CSN9      int       `gorm:"column:CSN9"`
	CSN10     int       `gorm:"column:CSN10"`
	OPN1      int       `gorm:"column:OPN1"`
	OPN2      int       `gorm:"column:OPN2"`
	OPN3      int       `gorm:"column:OPN3"`
	OPN4      int       `gorm:"column:OPN4"`
	OPN5      int       `gorm:"column:OPN5"`
	OPN6      int       `gorm:"column:OPN6"`
	OPN7      int       `gorm:"column:OPN7"`
	OPN8      int       `gorm:"column:OPN8"`
	OPN9      int       `gorm:"column:OPN9"`
	OPN10     int       `gorm:"column:OPN10"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserPersonality) TableName() string {
	return "user_personalities"
}

type CronJobSubscription struct {
	ProductID       *uint64
	UserID          uint64
	SubscriptionEnd time.Time
}
