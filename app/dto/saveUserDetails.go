package dto

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type SaveUserDetailRequest struct {
	UserID   int64  `json:"userid"`
	UserName string `json:"username" validate:"required"`
	Mail     string `json:"mail" validate:"required"`
	Address  string `json:"address" validate:"required"`
	City     string `json:"city" validate:"required"`
	Pincode  int64  `json:"pincode" validate:"required"`
	Phone    int64  `json:"phonenumber" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SaveUserResponse struct {
	UserId int64 `json:"userid"`
}

func (args *SaveUserDetailRequest) Parse(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func (args *SaveUserDetailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(args)
	if err != nil {
		return err
	}
	return nil
}
