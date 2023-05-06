package datastruct

type (
	Province struct {
		ProvinceID   uint64 `json:"province_id" gorm:"column:prov_id"`
		ProvinceName string `json:"province_name" gorm:"column:prov_name"`
	}

	City struct {
		CityID       uint64 `json:"city_id" gorm:"column:city_id"`
		ProvinceName string `json:"city_name" gorm:"column:city_name"`
		// ProvinceID   uint64 `json:"province_id" gorm:"column:prov_id"`
	}

	District struct {
		DistrictID   uint64 `json:"district_id" gorm:"column:dis_id"`
		ProvinceName string `json:"district_name" gorm:"column:dis_name"`
		// CityID       uint64 `json:"city_id" gorm:"column:city_id"`
	}

	SubDistrict struct {
		SubDistrictID uint64 `json:"subdistrict_id" gorm:"column:subdis_id"`
		ProvinceName  string `json:"subdistrict_name" gorm:"column:subdis_name"`
		// DistrictID    uint64 `json:"district_id" gorm:"column:dis_id"`
	}

	PostalCode struct {
		PostalCodeNumber uint64 `json:"postal_code" gorm:"column:postal_code"`
		// SubDistrictID uint64 `json:"subdistrict_id" gorm:"column:subdis_id"`
	}
)

func (Province) TableName() string {
	return "provinces"
}

func (City) TableName() string {
	return "cities"
}

func (District) TableName() string {
	return "districts"
}

func (SubDistrict) TableName() string {
	return "subdistricts"
}

func (PostalCode) TableName() string {
	return "postal_codes"
}
