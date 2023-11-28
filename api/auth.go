package api

import (
	"context"
	"dermsnap/api/public"
)

func (a API) Login(ctx context.Context, request public.LoginRequestObject) (public.LoginResponseObject, error) {

	return public.Login200JSONResponse{Token: "abc"}, nil
}
