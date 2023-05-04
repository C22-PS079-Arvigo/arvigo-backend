package datastruct

type (
	User struct {
		ID        uint64 `json:"id" gorm:"column:id"`
		FisrtName string `json:"first_name" gorm:"column:first_name"`
		LastName  string `json:"last_name" gorm:"column:last_name"`
		Email     string `json:"email" gorm:"column:email"`
		Password  string `json:"password" gorm:"column:password"`
		Role      string `json:"role" gorm:"column:role"`
		CreatedAt string `json:"created_at" gorm:"column:created_at"`
		UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
	}
)
