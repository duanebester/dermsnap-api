package api

import (
	"context"
	"dermsnap/api/http"
	"dermsnap/api/public"
	"dermsnap/models"
)

func (a API) Me(ctx context.Context, request http.MeRequestObject) (http.MeResponseObject, error) {
	user, _ := models.NewUser(models.User{
		Email:    "test@test.com",
		Password: "test",
		Role:     models.Client,
	})

	return http.Me200JSONResponse(*user), nil
}

func (a API) Register(ctx context.Context, request public.RegisterRequestObject) (public.RegisterResponseObject, error) {
	_, err := a.services.UserService.RegisterUser(request.Body.Email, request.Body.Password)
	if err != nil {
		return public.Register500JSONResponse{}, err
	}

	token, err := a.services.AuthService.LoginUser(request.Body.Email, request.Body.Password)
	if err != nil {
		return public.Register500JSONResponse{}, err
	}

	return public.Register200JSONResponse{
		Token: token,
	}, nil
}
