package service

import (
	"fmt"
	"net/http"
	"pjt1/app/dto"
	"pjt1/app/repo"
	"pjt1/pkg/e"
	"pjt1/pkg/jwt"

	"github.com/rs/zerolog/log"
)

type UserService interface {
	SaveUserDetails(r *http.Request) (*dto.SaveUserResponse, error)
	LoginUser(r *http.Request) (*dto.LoginResponse, error)
	//ExampleHandler(r *http.Request) error
}

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) SaveUserDetails(r *http.Request) (*dto.SaveUserResponse, error) {
	args := &dto.SaveUserDetailRequest{}

	// parsing the req.body
	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	//validation
	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	userID, err := s.userRepo.SaveUserDetails(args)
	if err != nil {
		return nil, e.NewError(e.ErrSaveUserDetails, "error while creating user", err)
	}
	log.Info().Msgf("Successfully created user with id %d", userID)

	return &dto.SaveUserResponse{
		UserId: userID,
	}, nil
}

func (s *userServiceImpl) LoginUser(r *http.Request) (*dto.LoginResponse, error) {
	args := &dto.LoginRequest{}

	// parsing the req.body
	err := args.Parse(r)
	if err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "error while parsing", err)
	}

	//validation
	err = args.Validate()
	if err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "error while validating", err)
	}
	log.Info().Msg("Successfully completed parsing and validation of request body")

	// Fetching user from database
	user, err := s.userRepo.GetUserByUsername(args.Username)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "user not found", err)
	}

	// Check if user is nil
	if user == nil {
		return nil, e.NewError(e.ErrResourceNotFound, "user not found", err)
	}
	log.Info().Msgf("the user is %s", user.Username)

	// Validate password
	if user.Password != args.Password {
		err := fmt.Errorf("invalid password for user %s", user.Username)
		return nil, e.NewError(e.ErrInvaliPassword, "invalid password", err)
	}

	// Generating JWT Token
	token, err := jwt.GenerateToken(user.ID, user.Username) //userid and username in the token
	if err != nil {
		return nil, e.NewError(e.ErrTokenNotGenerated, "failed to generate token", err)
	}

	fmt.Printf("the token is %s : \n ", token)

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

// func (s *userServiceImpl) ExampleHandler(r *http.Request) error {

// 	userID, err := helper.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		return err
// 	}

// 	// You can also get the username if needed
// 	username, err := helper.GetUsernameFromContext(r.Context())
// 	if err != nil {
// 		return err
// 	}

// 	// userID or username
// 	fmt.Printf("UserID: %d, Username: %s\n", userID, username)

// 	return nil
// }
