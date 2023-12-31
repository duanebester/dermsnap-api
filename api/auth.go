package api

import (
	"context"
	"dermsnap/api/public"
	"errors"
)

func (a API) Login(context.Context, public.LoginRequestObject) (public.LoginResponseObject, error) {
	return public.Login500JSONResponse{}, errors.New("not implemented")
}
