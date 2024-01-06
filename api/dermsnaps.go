package api

import (
	"context"
	"dermsnap/api/http"
	"dermsnap/middleware"
	"dermsnap/models"
)

func (a API) GetDermsnaps(ctx context.Context, request http.GetDermsnapsRequestObject) (http.GetDermsnapsResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	dermsnaps, err := a.services.DermsnapService.GetUserDermsnaps(user.ID)
	if err != nil {
		return http.GetDermsnaps500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.GetDermsnaps200JSONResponse(dermsnaps), nil
}

func (a API) CreateDermsnap(ctx context.Context, request http.CreateDermsnapRequestObject) (http.CreateDermsnapResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	dermsnap, err := a.services.DermsnapService.CreateDermsnap(user.ID, *request.Body)
	if err != nil {
		return http.CreateDermsnap500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	return http.CreateDermsnap200JSONResponse(*dermsnap), nil
}

func (a API) GetDermsnapById(ctx context.Context, request http.GetDermsnapByIdRequestObject) (http.GetDermsnapByIdResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	dermsnap, err := a.services.DermsnapService.GetDermsnapById(request.DermsnapId)
	if err != nil {
		return http.GetDermsnapById500JSONResponse{
			Message: err.Error(),
		}, nil
	}

	if dermsnap.UserID != user.ID {
		return http.GetDermsnapById401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}

	return http.GetDermsnapById200JSONResponse(*dermsnap), nil
}

func (a API) UpdateDermsnapById(ctx context.Context, request http.UpdateDermsnapByIdRequestObject) (http.UpdateDermsnapByIdResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	dermsnap, err := a.services.DermsnapService.GetDermsnapById(request.DermsnapId)
	if err != nil {
		return http.UpdateDermsnapById500JSONResponse{
			Message: err.Error(),
		}, nil
	}

	if dermsnap.UserID != user.ID {
		return http.UpdateDermsnapById401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}

	_, err = a.services.DermsnapService.UpdateDermsnap(dermsnap.ID, *request.Body)
	if err != nil {
		return http.UpdateDermsnapById500JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return http.UpdateDermsnapById200JSONResponse(*dermsnap), nil
}

func (a API) DeleteDermsnapById(ctx context.Context, request http.DeleteDermsnapByIdRequestObject) (http.DeleteDermsnapByIdResponseObject, error) {
	user := ctx.Value(middleware.UserKey).(*models.User)
	dermsnap, err := a.services.DermsnapService.GetDermsnapById(request.DermsnapId)
	if err != nil {
		return http.DeleteDermsnapById500JSONResponse{
			Message: err.Error(),
		}, nil
	}
	if dermsnap.UserID != user.ID {
		return http.DeleteDermsnapById401JSONResponse{
			Message: "Unauthorized",
		}, nil
	}

	_, err = a.services.DermsnapService.DeleteDermsnap(dermsnap)
	if err != nil {
		return http.DeleteDermsnapById500JSONResponse{
			Message: err.Error(),
		}, nil
	}

	return http.DeleteDermsnapById200JSONResponse(*dermsnap), nil
}

func (a API) UploadDermsnapImage(ctx context.Context, request http.UploadDermsnapImageRequestObject) (http.UploadDermsnapImageResponseObject, error) {
	return http.UploadDermsnapImage500JSONResponse{
		Message: "unimplemented",
	}, nil
}
