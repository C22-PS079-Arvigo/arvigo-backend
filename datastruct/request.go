package datastruct

type (
	LoginUserInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	UserRegisterInput struct {
		Email                string `json:"email" validate:"required,email"`
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
		FullName             string `json:"full_name" validate:"required"`
	}

	SearchCity struct {
		ProvinceID string `query:"province_id" validate:"required,numeric"`
	}

	SearchDistrict struct {
		CityID string `query:"city_id" validate:"required,numeric"`
	}

	SearchSubDistrict struct {
		DistrictID string `query:"district_id" validate:"required,numeric"`
	}

	SearchPostalCode struct {
		SubDistrictID string `query:"subdistrict_id" validate:"required,numeric"`
	}
)
