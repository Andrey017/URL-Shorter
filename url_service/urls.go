package url_service

import "errors"

type Urls struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Real_url    string `json:"real_url" binding:"required"`
	Shorter_url string `json:"shorter_url"`
	User_id     int    `json:"-" db:"user_id"`
}

type UpdateURL struct {
	Name        *string `json:"name"`
	Real_url    *string `json:"real_url"`
	Shorter_url *string `json:"shorter_url"`
}

func (input UpdateURL) Validate() error {
	if input.Name == nil && input.Real_url == nil {
		return errors.New("Update stucture has no validate")
	}

	return nil
}
