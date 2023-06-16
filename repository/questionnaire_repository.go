package repository

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/utils"
)

func GetQuestionnaire() (res []datastruct.Questionnaire, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GenerateQuestionnaireResult(data datastruct.QuestionnaireRequest, userID uint64) (res datastruct.PersonalityPercentages, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	response, err := utils.FetchMachineLearningAPI("POST", "/detect_personality", data)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	err = json.Unmarshal(response, &res)
	if err != nil {
		fmt.Printf("Error unmarshaling response body: %v", err)
		return
	}

	top2 := GetTop2FieldNames(res)
	var resTagIDs []string
	for _, v := range top2 {
		tag := constant.GetPersonalityTag[v]
		tagSplit := strings.Split(tag, ",")
		resTagIDs = append(resTagIDs, tagSplit...)
	}

	resTagIDsStr := strings.Join(resTagIDs, ",")
	payload := datastruct.UserPersonality{
		UserID:    userID,
		IsActive:  1,
		TagIDs:    resTagIDsStr,
		ExtResult: res.Extraversion,
		EstResult: res.Neurotic,
		AgrResult: res.Agreeable,
		CsnResult: res.Conscientious,
		OpnResult: res.Openness,
		EXT1:      data.EXT1,
		EXT2:      data.EXT2,
		EXT3:      data.EXT3,
		EXT4:      data.EXT4,
		EXT5:      data.EXT5,
		EXT6:      data.EXT6,
		EXT7:      data.EXT7,
		EXT8:      data.EXT8,
		EXT9:      data.EXT9,
		EXT10:     data.EXT10,
		EST1:      data.EST1,
		EST2:      data.EST2,
		EST3:      data.EST3,
		EST4:      data.EST4,
		EST5:      data.EST5,
		EST6:      data.EST6,
		EST7:      data.EST7,
		EST8:      data.EST8,
		EST9:      data.EST9,
		EST10:     data.EST10,
		AGR1:      data.AGR1,
		AGR2:      data.AGR2,
		AGR3:      data.AGR3,
		AGR4:      data.AGR4,
		AGR5:      data.AGR5,
		AGR6:      data.AGR6,
		AGR7:      data.AGR7,
		AGR8:      data.AGR8,
		AGR9:      data.AGR9,
		AGR10:     data.AGR10,
		CSN1:      data.CSN1,
		CSN2:      data.CSN2,
		CSN3:      data.CSN3,
		CSN4:      data.CSN4,
		CSN5:      data.CSN5,
		CSN6:      data.CSN6,
		CSN7:      data.CSN7,
		CSN8:      data.CSN8,
		CSN9:      data.CSN9,
		CSN10:     data.CSN10,
		OPN1:      data.OPN1,
		OPN2:      data.OPN2,
		OPN3:      data.OPN3,
		OPN4:      data.OPN4,
		OPN5:      data.OPN5,
		OPN6:      data.OPN6,
		OPN7:      data.OPN7,
		OPN8:      data.OPN8,
		OPN9:      data.OPN9,
		OPN10:     data.OPN10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = db.Table("user_personalities").
		Create(&payload).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Model(&datastruct.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"is_complete_personality_test": 1,
			"updated_at":                   time.Now(),
		}).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetTop2FieldNames(p datastruct.PersonalityPercentages) []string {
	// Initialize the top 2 highest values to the lowest possible value
	top1 := math.Inf(-1)
	top2 := math.Inf(-1)

	// Initialize the top 2 field names
	topFields := make([]string, 2)

	// Iterate through each field and update the top 2 highest values and field names accordingly
	val := reflect.ValueOf(p)
	typeOfP := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typeOfP.Field(i).Name
		fieldValue := field.Float()

		if fieldValue > top1 {
			top2 = top1
			top1 = fieldValue
			topFields[1] = topFields[0]
			topFields[0] = fieldName
		} else if fieldValue > top2 {
			top2 = fieldValue
			topFields[1] = fieldName
		}
	}

	return topFields
}
