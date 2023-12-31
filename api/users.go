package api

import (
	"context"
	"dermsnap/api/http"
	"dermsnap/middleware"
	"dermsnap/models"
)

func (a API) Me(ctx context.Context, request http.MeRequestObject) (http.MeResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	return http.Me200JSONResponse(*user), nil
}

func (a API) GetUserInfo(ctx context.Context, request http.GetUserInfoRequestObject) (http.GetUserInfoResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	if user.ID != request.UserId {
		return http.GetUserInfo401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}
	userInfo, err := a.services.UserService.GetUserInfo(request.UserId)
	if err != nil {
		return http.GetUserInfo500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.GetUserInfo200JSONResponse(*userInfo), nil
}

func (a API) CreateUserInfo(ctx context.Context, request http.CreateUserInfoRequestObject) (http.CreateUserInfoResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	if user.ID != request.UserId {
		return http.CreateUserInfo401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}

	userInfo, err := a.services.UserService.CreateUserInfo(request.UserId, *request.Body)
	if err != nil {
		return http.CreateUserInfo500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.CreateUserInfo200JSONResponse(*userInfo), nil
}

func (a API) GetDoctorInfo(ctx context.Context, request http.GetDoctorInfoRequestObject) (http.GetDoctorInfoResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	if user.Role != models.Admin && user.ID != request.UserId {
		return http.GetDoctorInfo401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}
	doctorInfo, err := a.services.UserService.GetDoctorInfo(request.UserId)
	if err != nil {
		return http.GetDoctorInfo500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.GetDoctorInfo200JSONResponse(*doctorInfo), nil
}

func (a API) CreateDoctorInfo(ctx context.Context, request http.CreateDoctorInfoRequestObject) (http.CreateDoctorInfoResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	if user.Role != models.Admin && user.ID != request.UserId {
		return http.CreateDoctorInfo401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}

	userInfo, err := a.services.UserService.CreateDoctorInfo(request.UserId, *request.Body)
	if err != nil {
		return http.CreateDoctorInfo500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.CreateDoctorInfo200JSONResponse(*userInfo), nil
}
