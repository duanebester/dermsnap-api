package api

import (
	"context"
	"dermsnap/api/public"
)

func (a API) Login(ctx context.Context, request public.LoginRequestObject) (public.LoginResponseObject, error) {
	token, err := a.services.AuthService.LoginUser(request.Body.Email, request.Body.Password)
	if err != nil {
		return public.Login500JSONResponse{}, err
	}

	return public.Login200JSONResponse{Token: token}, nil
}
