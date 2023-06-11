package repository

import (
	"net/http"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetListPaymentUser() (res []datastruct.UserSubscriptionResponse, statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db = Database()
	)

	if err := db.Table("detail_user_subscriptions dus").
		Select([]string{
			"users.full_name as user_name",
			"dus.*",
			"dusp.id as id_dusp",
		}).
		Joins("JOIN users ON users.id = dus.user_id").
		Joins("LEFT JOIN detail_user_subscription_products dusp ON dusp.subscription_id = dus.id").
		Where("dusp.id IS NULL").
		Find(&res).
		Error; err != nil {
		return res, http.StatusNotFound, err
	}

	return
}

func GetListPaymentMerchant() (res []datastruct.MerchantSubscriptionResponse, statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db = Database()
	)

	if err := db.Table("detail_user_subscriptions dus").
		Select("users.full_name as user_name,dus.*, group_concat(p.name) as products").
		Joins("JOIN users ON users.id = dus.user_id").
		Joins("LEFT JOIN detail_user_subscription_products dusp ON dusp.subscription_id = dus.id").
		Joins("left join products p on p.id = dusp.product_id").
		Where("dusp.id IS NOT NULL").
		Group("dusp.subscription_id").
		Find(&res).
		Error; err != nil {
		return res, http.StatusNotFound, err
	}

	return
}

func UserCreatePayment(userID uint64, data datastruct.UserCreatePaymentInput) (statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db          = Database()
		currentTime = time.Now()
	)

	payload := datastruct.UserSubscription{
		UserID:     userID,
		Price:      data.Price,
		UniqueCode: data.UniqueCode,
		Status:     constant.StatusWaitingPayment,
		Message:    data.Message,
		Bank:       data.Bank,
		PaidAt:     &currentTime,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	if err = db.Create(&payload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}

func PartnerCreatePayment(userID uint64, data datastruct.PartnerCreatePaymentInput) (statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db          = Database()
		currentTime = time.Now()
	)

	payload := datastruct.UserSubscription{
		UserID:     userID,
		Price:      data.Price,
		UniqueCode: data.UniqueCode,
		Status:     constant.StatusWaitingPayment,
		Message:    data.Message,
		Bank:       data.Bank,
		PaidAt:     &currentTime,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	if err = db.Create(&payload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	for _, v := range data.ProductIDs {
		payloadProduct := datastruct.UserSubscriptionProduct{
			ProductID:      v,
			SubscriptionID: payload.ID,
			CreatedAt:      currentTime,
			UpdatedAt:      currentTime,
		}

		if err = db.Create(&payloadProduct).Error; err != nil {
			return http.StatusInternalServerError, err
		}

		if err = db.Table("products").Where("id = ?", v).
			Update("status", constant.StatusWaitingPayment).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return
}

func VerifyPaymentUser(subsID uint64, data datastruct.VerifyPaymentUser) (statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db = Database()
	)

	status := constant.StatusRejected
	if data.Status {
		status = constant.StatusApproved
		if err = db.Table("detail_user_subscriptions").Where("id = ?", subsID).Updates(map[string]interface{}{
			"status":             status,
			"subscription_start": time.Now(),
			"subscription_end":   time.Now().AddDate(0, 1, 0),
		}).Error; err != nil {
			return http.StatusInternalServerError, err
		}

		var userID uint64
		if err := db.Table("detail_user_subscriptions dus").
			Select([]string{
				"dus.user_id",
			}).
			Where("dus.id = ?", subsID).
			Find(&userID).
			Error; err != nil {
			return http.StatusInternalServerError, err
		}

		if err = db.Table("users").Where("id = ?", userID).Updates(map[string]interface{}{
			"is_subscription_active": 1,
		}).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		if err = db.Table("detail_user_subscriptions").Where("id = ?", subsID).Update("status", status).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return
}

func VerifyPaymentMerchant(subsID uint64, data datastruct.VerifyPaymentMerchant) (statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db = Database()
	)

	status := constant.StatusRejected
	isActive := 0
	if data.Status {
		status = constant.StatusApproved
		isActive = 1
		if err = db.Table("detail_user_subscriptions").Where("id = ?", subsID).Updates(map[string]interface{}{
			"status":             status,
			"subscription_start": time.Now(),
			"subscription_end":   time.Now().AddDate(0, 1, 0),
		}).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		if err = db.Table("detail_user_subscriptions").Where("id = ?", subsID).Update("status", status).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	var productIDs []uint64
	if err = db.Table("detail_user_subscription_products").Select("product_id").
		Where("subscription_id = ?", subsID).Find(&productIDs).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	for _, v := range productIDs {
		if status == constant.StatusApproved {
			status = constant.StatusSubscribed
		}
		if err = db.Table("products").Where("id = ?", v).
			Updates(map[string]interface{}{
				"status":                 status,
				"rejected_note":          data.RejectedNote,
				"is_subscription_active": isActive,
			}).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return
}

func SubscriptionCronJob() (statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK
	var subscriptions []datastruct.CronJobSubscription

	if err := db.
		Table("detail_user_subscriptions dus").
		Select("dus.user_id, dusp.product_id, dus.subscription_end").
		Joins("left join detail_user_subscription_products dusp on dus.id = dusp.subscription_id").
		Find(&subscriptions).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	for _, v := range subscriptions {
		if time.Now().After(v.SubscriptionEnd) {
			if v.ProductID != nil {
				if err = db.Table("products").Where("id = ?", v.ProductID).
					Updates(map[string]interface{}{
						"status":                 constant.StatusApproved,
						"is_subscription_active": 0,
					}).Error; err != nil {
					return http.StatusInternalServerError, err
				}
			} else {
				if err = db.Table("users").Where("id = ?", v.UserID).Updates(map[string]interface{}{
					"is_subscription_active": 0,
				}).Error; err != nil {
					return http.StatusInternalServerError, err
				}
			}
		}
	}

	return
}
