package datastruct

import "time"

type (
	LoginRegisterResponse struct {
		UserID uint64 `json:"user_id"`
		Token  string `json:"token"`
	}

	UserDetailResponse struct {
		UserDetail
		Address UserAddress `json:"address"`
	}

	FaceShapeResponse struct {
		ImageUrl string `json:"image_url"`
		Result   string `json:"result"`
	}

	UserDetail struct {
		ID                        uint64     `json:"id"`
		Email                     string     `json:"email"`
		RoleID                    uint64     `json:"role_id"`
		RoleName                  string     `json:"role_name"`
		FullName                  string     `json:"full_name"`
		Gender                    string     `json:"gender"`
		DateOfBirth               *time.Time `json:"date_of_birth"`
		PlaceOfBirth              string     `json:"place_of_birth"`
		IsCompletePersonalityTest bool       `json:"is_complete_personality_test"`
		IsCompleteFaceTest        bool       `json:"is_complete_face_test"`
		PersonalityID             bool       `json:"personality_id"`
		FaceShapeTagID            bool       `json:"face_shape_tag_id"`
		IsVerified                bool       `json:"is_verified"`
		Avatar                    string     `json:"avatar"`
		AddressID                 uint64     `json:"addresses_id"`
		MerchantID                uint64     `json:"merchant_id"`
	}

	UserAddress struct {
		Street      string `json:"street"`
		Province    string `json:"province"`
		City        string `json:"city"`
		District    string `json:"district"`
		SubDistrict string `json:"sub_district"`
		PostalCode  uint64 `json:"postal_code"`
	}
)
