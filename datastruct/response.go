package datastruct

type (
	LoginRegisterResponse struct {
		UserID uint64 `json:"user_id"`
		Token  string `json:"token"`
	}
)
