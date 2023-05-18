package repository

import (
	"net/http"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetQuestionnaire() (res []datastruct.Questionnaire, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}
