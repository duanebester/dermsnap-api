// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package http

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/url"
	"path"
	"strings"

	models "dermsnap/models"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// BodyLocation defines model for BodyLocation.
type BodyLocation = models.BodyLocation

// CreateDermsnap defines model for CreateDermsnap.
type CreateDermsnap = models.CreateDermsnap

// CreateDoctorInfo defines model for CreateDoctorInfo.
type CreateDoctorInfo = models.CreateDoctorInfo

// CreateUserInfo defines model for CreateUserInfo.
type CreateUserInfo = models.CreateUserInfo

// Dermsnap defines model for Dermsnap.
type Dermsnap = models.Dermsnap

// DermsnapImage defines model for DermsnapImage.
type DermsnapImage = models.DermsnapImage

// DoctorInfo defines model for DoctorInfo.
type DoctorInfo = models.DoctorInfo

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// IdentifierType defines model for IdentifierType.
type IdentifierType = models.IdentifierType

// Role defines model for Role.
type Role = models.Role

// UpdateDermsnap defines model for UpdateDermsnap.
type UpdateDermsnap = models.UpdateDermsnap

// User defines model for User.
type User = models.User

// UserInfo defines model for UserInfo.
type UserInfo = models.UserInfo

// DermsnapId defines model for dermsnap_id.
type DermsnapId = openapi_types.UUID

// UserId defines model for user_id.
type UserId = openapi_types.UUID

// UploadDermsnapImageMultipartBody defines parameters for UploadDermsnapImage.
type UploadDermsnapImageMultipartBody struct {
	DermsnapId *openapi_types.UUID `json:"dermsnap_id,omitempty"`
	File       *openapi_types.File `json:"file,omitempty"`
}

// CreateDermsnapJSONRequestBody defines body for CreateDermsnap for application/json ContentType.
type CreateDermsnapJSONRequestBody = CreateDermsnap

// GetDermsnapByIdJSONRequestBody defines body for GetDermsnapById for application/json ContentType.
type GetDermsnapByIdJSONRequestBody = UpdateDermsnap

// UpdateDermsnapByIdJSONRequestBody defines body for UpdateDermsnapById for application/json ContentType.
type UpdateDermsnapByIdJSONRequestBody = UpdateDermsnap

// UploadDermsnapImageMultipartRequestBody defines body for UploadDermsnapImage for multipart/form-data ContentType.
type UploadDermsnapImageMultipartRequestBody UploadDermsnapImageMultipartBody

// CreateDoctorInfoJSONRequestBody defines body for CreateDoctorInfo for application/json ContentType.
type CreateDoctorInfoJSONRequestBody = CreateDoctorInfo

// CreateUserInfoJSONRequestBody defines body for CreateUserInfo for application/json ContentType.
type CreateUserInfoJSONRequestBody = CreateUserInfo

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get dermsnaps
	// (GET /dermsnaps)
	GetDermsnaps(c *fiber.Ctx) error
	// Create dermsnap
	// (POST /dermsnaps)
	CreateDermsnap(c *fiber.Ctx) error
	// Delete dermsnap
	// (DELETE /dermsnaps/{dermsnap_id})
	DeleteDermsnapById(c *fiber.Ctx, dermsnapId DermsnapId) error
	// Get dermsnap
	// (GET /dermsnaps/{dermsnap_id})
	GetDermsnapById(c *fiber.Ctx, dermsnapId DermsnapId) error
	// Update dermsnap
	// (PUT /dermsnaps/{dermsnap_id})
	UpdateDermsnapById(c *fiber.Ctx, dermsnapId DermsnapId) error
	// Upload dermsnap image
	// (POST /dermsnaps/{dermsnap_id}/images)
	UploadDermsnapImage(c *fiber.Ctx, dermsnapId DermsnapId) error
	// Get current user
	// (GET /me)
	Me(c *fiber.Ctx) error
	// Get doctor info
	// (GET /users/{user_id}/doctor-info)
	GetDoctorInfo(c *fiber.Ctx, userId UserId) error
	// Create doctor info
	// (POST /users/{user_id}/doctor-info)
	CreateDoctorInfo(c *fiber.Ctx, userId UserId) error
	// Get user info
	// (GET /users/{user_id}/user-info)
	GetUserInfo(c *fiber.Ctx, userId UserId) error
	// Create user info
	// (POST /users/{user_id}/user-info)
	CreateUserInfo(c *fiber.Ctx, userId UserId) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// GetDermsnaps operation middleware
func (siw *ServerInterfaceWrapper) GetDermsnaps(c *fiber.Ctx) error {

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.GetDermsnaps(c)
}

// CreateDermsnap operation middleware
func (siw *ServerInterfaceWrapper) CreateDermsnap(c *fiber.Ctx) error {

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.CreateDermsnap(c)
}

// DeleteDermsnapById operation middleware
func (siw *ServerInterfaceWrapper) DeleteDermsnapById(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "dermsnap_id" -------------
	var dermsnapId DermsnapId

	err = runtime.BindStyledParameter("simple", false, "dermsnap_id", c.Params("dermsnap_id"), &dermsnapId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter dermsnap_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.DeleteDermsnapById(c, dermsnapId)
}

// GetDermsnapById operation middleware
func (siw *ServerInterfaceWrapper) GetDermsnapById(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "dermsnap_id" -------------
	var dermsnapId DermsnapId

	err = runtime.BindStyledParameter("simple", false, "dermsnap_id", c.Params("dermsnap_id"), &dermsnapId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter dermsnap_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.GetDermsnapById(c, dermsnapId)
}

// UpdateDermsnapById operation middleware
func (siw *ServerInterfaceWrapper) UpdateDermsnapById(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "dermsnap_id" -------------
	var dermsnapId DermsnapId

	err = runtime.BindStyledParameter("simple", false, "dermsnap_id", c.Params("dermsnap_id"), &dermsnapId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter dermsnap_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.UpdateDermsnapById(c, dermsnapId)
}

// UploadDermsnapImage operation middleware
func (siw *ServerInterfaceWrapper) UploadDermsnapImage(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "dermsnap_id" -------------
	var dermsnapId DermsnapId

	err = runtime.BindStyledParameter("simple", false, "dermsnap_id", c.Params("dermsnap_id"), &dermsnapId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter dermsnap_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.UploadDermsnapImage(c, dermsnapId)
}

// Me operation middleware
func (siw *ServerInterfaceWrapper) Me(c *fiber.Ctx) error {

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.Me(c)
}

// GetDoctorInfo operation middleware
func (siw *ServerInterfaceWrapper) GetDoctorInfo(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Params("user_id"), &userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.GetDoctorInfo(c, userId)
}

// CreateDoctorInfo operation middleware
func (siw *ServerInterfaceWrapper) CreateDoctorInfo(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Params("user_id"), &userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.CreateDoctorInfo(c, userId)
}

// GetUserInfo operation middleware
func (siw *ServerInterfaceWrapper) GetUserInfo(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Params("user_id"), &userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.GetUserInfo(c, userId)
}

// CreateUserInfo operation middleware
func (siw *ServerInterfaceWrapper) CreateUserInfo(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Params("user_id"), &userId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user_id: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	return siw.Handler.CreateUserInfo(c, userId)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Get(options.BaseURL+"/dermsnaps", wrapper.GetDermsnaps)

	router.Post(options.BaseURL+"/dermsnaps", wrapper.CreateDermsnap)

	router.Delete(options.BaseURL+"/dermsnaps/:dermsnap_id", wrapper.DeleteDermsnapById)

	router.Get(options.BaseURL+"/dermsnaps/:dermsnap_id", wrapper.GetDermsnapById)

	router.Put(options.BaseURL+"/dermsnaps/:dermsnap_id", wrapper.UpdateDermsnapById)

	router.Post(options.BaseURL+"/dermsnaps/:dermsnap_id/images", wrapper.UploadDermsnapImage)

	router.Get(options.BaseURL+"/me", wrapper.Me)

	router.Get(options.BaseURL+"/users/:user_id/doctor-info", wrapper.GetDoctorInfo)

	router.Post(options.BaseURL+"/users/:user_id/doctor-info", wrapper.CreateDoctorInfo)

	router.Get(options.BaseURL+"/users/:user_id/user-info", wrapper.GetUserInfo)

	router.Post(options.BaseURL+"/users/:user_id/user-info", wrapper.CreateUserInfo)

}

type GetDermsnapsRequestObject struct {
}

type GetDermsnapsResponseObject interface {
	VisitGetDermsnapsResponse(ctx *fiber.Ctx) error
}

type GetDermsnaps200JSONResponse []Dermsnap

func (response GetDermsnaps200JSONResponse) VisitGetDermsnapsResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetDermsnaps401JSONResponse Error

func (response GetDermsnaps401JSONResponse) VisitGetDermsnapsResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type GetDermsnaps500JSONResponse Error

func (response GetDermsnaps500JSONResponse) VisitGetDermsnapsResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type CreateDermsnapRequestObject struct {
	Body *CreateDermsnapJSONRequestBody
}

type CreateDermsnapResponseObject interface {
	VisitCreateDermsnapResponse(ctx *fiber.Ctx) error
}

type CreateDermsnap200JSONResponse Dermsnap

func (response CreateDermsnap200JSONResponse) VisitCreateDermsnapResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type CreateDermsnap401JSONResponse Error

func (response CreateDermsnap401JSONResponse) VisitCreateDermsnapResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type CreateDermsnap500JSONResponse Error

func (response CreateDermsnap500JSONResponse) VisitCreateDermsnapResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type DeleteDermsnapByIdRequestObject struct {
	DermsnapId DermsnapId `json:"dermsnap_id"`
}

type DeleteDermsnapByIdResponseObject interface {
	VisitDeleteDermsnapByIdResponse(ctx *fiber.Ctx) error
}

type DeleteDermsnapById200JSONResponse Dermsnap

func (response DeleteDermsnapById200JSONResponse) VisitDeleteDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type DeleteDermsnapById401JSONResponse Error

func (response DeleteDermsnapById401JSONResponse) VisitDeleteDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type DeleteDermsnapById500JSONResponse Error

func (response DeleteDermsnapById500JSONResponse) VisitDeleteDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type GetDermsnapByIdRequestObject struct {
	DermsnapId DermsnapId `json:"dermsnap_id"`
	Body       *GetDermsnapByIdJSONRequestBody
}

type GetDermsnapByIdResponseObject interface {
	VisitGetDermsnapByIdResponse(ctx *fiber.Ctx) error
}

type GetDermsnapById200JSONResponse Dermsnap

func (response GetDermsnapById200JSONResponse) VisitGetDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetDermsnapById401JSONResponse Error

func (response GetDermsnapById401JSONResponse) VisitGetDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type GetDermsnapById500JSONResponse Error

func (response GetDermsnapById500JSONResponse) VisitGetDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type UpdateDermsnapByIdRequestObject struct {
	DermsnapId DermsnapId `json:"dermsnap_id"`
	Body       *UpdateDermsnapByIdJSONRequestBody
}

type UpdateDermsnapByIdResponseObject interface {
	VisitUpdateDermsnapByIdResponse(ctx *fiber.Ctx) error
}

type UpdateDermsnapById200JSONResponse Dermsnap

func (response UpdateDermsnapById200JSONResponse) VisitUpdateDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type UpdateDermsnapById401JSONResponse Error

func (response UpdateDermsnapById401JSONResponse) VisitUpdateDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type UpdateDermsnapById500JSONResponse Error

func (response UpdateDermsnapById500JSONResponse) VisitUpdateDermsnapByIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type UploadDermsnapImageRequestObject struct {
	DermsnapId DermsnapId `json:"dermsnap_id"`
	Body       *multipart.Reader
}

type UploadDermsnapImageResponseObject interface {
	VisitUploadDermsnapImageResponse(ctx *fiber.Ctx) error
}

type UploadDermsnapImage200JSONResponse DermsnapImage

func (response UploadDermsnapImage200JSONResponse) VisitUploadDermsnapImageResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type UploadDermsnapImage401JSONResponse Error

func (response UploadDermsnapImage401JSONResponse) VisitUploadDermsnapImageResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type UploadDermsnapImage500JSONResponse Error

func (response UploadDermsnapImage500JSONResponse) VisitUploadDermsnapImageResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type MeRequestObject struct {
}

type MeResponseObject interface {
	VisitMeResponse(ctx *fiber.Ctx) error
}

type Me200JSONResponse User

func (response Me200JSONResponse) VisitMeResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type Me401JSONResponse Error

func (response Me401JSONResponse) VisitMeResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type Me500JSONResponse Error

func (response Me500JSONResponse) VisitMeResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type GetDoctorInfoRequestObject struct {
	UserId UserId `json:"user_id"`
}

type GetDoctorInfoResponseObject interface {
	VisitGetDoctorInfoResponse(ctx *fiber.Ctx) error
}

type GetDoctorInfo200JSONResponse DoctorInfo

func (response GetDoctorInfo200JSONResponse) VisitGetDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetDoctorInfo401JSONResponse Error

func (response GetDoctorInfo401JSONResponse) VisitGetDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type GetDoctorInfo500JSONResponse Error

func (response GetDoctorInfo500JSONResponse) VisitGetDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type CreateDoctorInfoRequestObject struct {
	UserId UserId `json:"user_id"`
	Body   *CreateDoctorInfoJSONRequestBody
}

type CreateDoctorInfoResponseObject interface {
	VisitCreateDoctorInfoResponse(ctx *fiber.Ctx) error
}

type CreateDoctorInfo200JSONResponse DoctorInfo

func (response CreateDoctorInfo200JSONResponse) VisitCreateDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type CreateDoctorInfo401JSONResponse Error

func (response CreateDoctorInfo401JSONResponse) VisitCreateDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type CreateDoctorInfo500JSONResponse Error

func (response CreateDoctorInfo500JSONResponse) VisitCreateDoctorInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type GetUserInfoRequestObject struct {
	UserId UserId `json:"user_id"`
}

type GetUserInfoResponseObject interface {
	VisitGetUserInfoResponse(ctx *fiber.Ctx) error
}

type GetUserInfo200JSONResponse UserInfo

func (response GetUserInfo200JSONResponse) VisitGetUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetUserInfo401JSONResponse Error

func (response GetUserInfo401JSONResponse) VisitGetUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type GetUserInfo500JSONResponse Error

func (response GetUserInfo500JSONResponse) VisitGetUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type CreateUserInfoRequestObject struct {
	UserId UserId `json:"user_id"`
	Body   *CreateUserInfoJSONRequestBody
}

type CreateUserInfoResponseObject interface {
	VisitCreateUserInfoResponse(ctx *fiber.Ctx) error
}

type CreateUserInfo200JSONResponse UserInfo

func (response CreateUserInfo200JSONResponse) VisitCreateUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type CreateUserInfo401JSONResponse Error

func (response CreateUserInfo401JSONResponse) VisitCreateUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type CreateUserInfo500JSONResponse Error

func (response CreateUserInfo500JSONResponse) VisitCreateUserInfoResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get dermsnaps
	// (GET /dermsnaps)
	GetDermsnaps(ctx context.Context, request GetDermsnapsRequestObject) (GetDermsnapsResponseObject, error)
	// Create dermsnap
	// (POST /dermsnaps)
	CreateDermsnap(ctx context.Context, request CreateDermsnapRequestObject) (CreateDermsnapResponseObject, error)
	// Delete dermsnap
	// (DELETE /dermsnaps/{dermsnap_id})
	DeleteDermsnapById(ctx context.Context, request DeleteDermsnapByIdRequestObject) (DeleteDermsnapByIdResponseObject, error)
	// Get dermsnap
	// (GET /dermsnaps/{dermsnap_id})
	GetDermsnapById(ctx context.Context, request GetDermsnapByIdRequestObject) (GetDermsnapByIdResponseObject, error)
	// Update dermsnap
	// (PUT /dermsnaps/{dermsnap_id})
	UpdateDermsnapById(ctx context.Context, request UpdateDermsnapByIdRequestObject) (UpdateDermsnapByIdResponseObject, error)
	// Upload dermsnap image
	// (POST /dermsnaps/{dermsnap_id}/images)
	UploadDermsnapImage(ctx context.Context, request UploadDermsnapImageRequestObject) (UploadDermsnapImageResponseObject, error)
	// Get current user
	// (GET /me)
	Me(ctx context.Context, request MeRequestObject) (MeResponseObject, error)
	// Get doctor info
	// (GET /users/{user_id}/doctor-info)
	GetDoctorInfo(ctx context.Context, request GetDoctorInfoRequestObject) (GetDoctorInfoResponseObject, error)
	// Create doctor info
	// (POST /users/{user_id}/doctor-info)
	CreateDoctorInfo(ctx context.Context, request CreateDoctorInfoRequestObject) (CreateDoctorInfoResponseObject, error)
	// Get user info
	// (GET /users/{user_id}/user-info)
	GetUserInfo(ctx context.Context, request GetUserInfoRequestObject) (GetUserInfoResponseObject, error)
	// Create user info
	// (POST /users/{user_id}/user-info)
	CreateUserInfo(ctx context.Context, request CreateUserInfoRequestObject) (CreateUserInfoResponseObject, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetDermsnaps operation middleware
func (sh *strictHandler) GetDermsnaps(ctx *fiber.Ctx) error {
	var request GetDermsnapsRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetDermsnaps(ctx.UserContext(), request.(GetDermsnapsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDermsnaps")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetDermsnapsResponseObject); ok {
		if err := validResponse.VisitGetDermsnapsResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateDermsnap operation middleware
func (sh *strictHandler) CreateDermsnap(ctx *fiber.Ctx) error {
	var request CreateDermsnapRequestObject

	var body CreateDermsnapJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.CreateDermsnap(ctx.UserContext(), request.(CreateDermsnapRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDermsnap")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(CreateDermsnapResponseObject); ok {
		if err := validResponse.VisitCreateDermsnapResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteDermsnapById operation middleware
func (sh *strictHandler) DeleteDermsnapById(ctx *fiber.Ctx, dermsnapId DermsnapId) error {
	var request DeleteDermsnapByIdRequestObject

	request.DermsnapId = dermsnapId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteDermsnapById(ctx.UserContext(), request.(DeleteDermsnapByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteDermsnapById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(DeleteDermsnapByIdResponseObject); ok {
		if err := validResponse.VisitDeleteDermsnapByIdResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetDermsnapById operation middleware
func (sh *strictHandler) GetDermsnapById(ctx *fiber.Ctx, dermsnapId DermsnapId) error {
	var request GetDermsnapByIdRequestObject

	request.DermsnapId = dermsnapId

	var body GetDermsnapByIdJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetDermsnapById(ctx.UserContext(), request.(GetDermsnapByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDermsnapById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetDermsnapByIdResponseObject); ok {
		if err := validResponse.VisitGetDermsnapByIdResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateDermsnapById operation middleware
func (sh *strictHandler) UpdateDermsnapById(ctx *fiber.Ctx, dermsnapId DermsnapId) error {
	var request UpdateDermsnapByIdRequestObject

	request.DermsnapId = dermsnapId

	var body UpdateDermsnapByIdJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateDermsnapById(ctx.UserContext(), request.(UpdateDermsnapByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateDermsnapById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(UpdateDermsnapByIdResponseObject); ok {
		if err := validResponse.VisitUpdateDermsnapByIdResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UploadDermsnapImage operation middleware
func (sh *strictHandler) UploadDermsnapImage(ctx *fiber.Ctx, dermsnapId DermsnapId) error {
	var request UploadDermsnapImageRequestObject

	request.DermsnapId = dermsnapId

	request.Body = multipart.NewReader(bytes.NewReader(ctx.Request().Body()), string(ctx.Request().Header.MultipartFormBoundary()))

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.UploadDermsnapImage(ctx.UserContext(), request.(UploadDermsnapImageRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UploadDermsnapImage")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(UploadDermsnapImageResponseObject); ok {
		if err := validResponse.VisitUploadDermsnapImageResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Me operation middleware
func (sh *strictHandler) Me(ctx *fiber.Ctx) error {
	var request MeRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.Me(ctx.UserContext(), request.(MeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Me")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(MeResponseObject); ok {
		if err := validResponse.VisitMeResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetDoctorInfo operation middleware
func (sh *strictHandler) GetDoctorInfo(ctx *fiber.Ctx, userId UserId) error {
	var request GetDoctorInfoRequestObject

	request.UserId = userId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetDoctorInfo(ctx.UserContext(), request.(GetDoctorInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDoctorInfo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetDoctorInfoResponseObject); ok {
		if err := validResponse.VisitGetDoctorInfoResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateDoctorInfo operation middleware
func (sh *strictHandler) CreateDoctorInfo(ctx *fiber.Ctx, userId UserId) error {
	var request CreateDoctorInfoRequestObject

	request.UserId = userId

	var body CreateDoctorInfoJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.CreateDoctorInfo(ctx.UserContext(), request.(CreateDoctorInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDoctorInfo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(CreateDoctorInfoResponseObject); ok {
		if err := validResponse.VisitCreateDoctorInfoResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUserInfo operation middleware
func (sh *strictHandler) GetUserInfo(ctx *fiber.Ctx, userId UserId) error {
	var request GetUserInfoRequestObject

	request.UserId = userId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserInfo(ctx.UserContext(), request.(GetUserInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserInfo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetUserInfoResponseObject); ok {
		if err := validResponse.VisitGetUserInfoResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateUserInfo operation middleware
func (sh *strictHandler) CreateUserInfo(ctx *fiber.Ctx, userId UserId) error {
	var request CreateUserInfoRequestObject

	request.UserId = userId

	var body CreateUserInfoJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUserInfo(ctx.UserContext(), request.(CreateUserInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUserInfo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(CreateUserInfoResponseObject); ok {
		if err := validResponse.VisitCreateUserInfoResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
