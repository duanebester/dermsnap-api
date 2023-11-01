// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message,omitempty"`
}

// Login defines model for Login.
type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = Login

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody = Login

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Login
	// (POST /login)
	Login(c *fiber.Ctx) error
	// Register
	// (POST /register)
	Register(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(c *fiber.Ctx) error {

	return siw.Handler.Login(c)
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(c *fiber.Ctx) error {

	return siw.Handler.Register(c)
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

	router.Post(options.BaseURL+"/login", wrapper.Login)

	router.Post(options.BaseURL+"/register", wrapper.Register)

}

type LoginRequestObject struct {
	Body *LoginJSONRequestBody
}

type LoginResponseObject interface {
	VisitLoginResponse(ctx *fiber.Ctx) error
}

type Login200JSONResponse LoginResponse

func (response Login200JSONResponse) VisitLoginResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type Login400JSONResponse Error

func (response Login400JSONResponse) VisitLoginResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(400)

	return ctx.JSON(&response)
}

type Login401JSONResponse Error

func (response Login401JSONResponse) VisitLoginResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(401)

	return ctx.JSON(&response)
}

type Login500JSONResponse Error

func (response Login500JSONResponse) VisitLoginResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type RegisterRequestObject struct {
	Body *RegisterJSONRequestBody
}

type RegisterResponseObject interface {
	VisitRegisterResponse(ctx *fiber.Ctx) error
}

type Register200JSONResponse LoginResponse

func (response Register200JSONResponse) VisitRegisterResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type Register400JSONResponse Error

func (response Register400JSONResponse) VisitRegisterResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(400)

	return ctx.JSON(&response)
}

type Register500JSONResponse Error

func (response Register500JSONResponse) VisitRegisterResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Login
	// (POST /login)
	Login(ctx context.Context, request LoginRequestObject) (LoginResponseObject, error)
	// Register
	// (POST /register)
	Register(ctx context.Context, request RegisterRequestObject) (RegisterResponseObject, error)
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

// Login operation middleware
func (sh *strictHandler) Login(ctx *fiber.Ctx) error {
	var request LoginRequestObject

	var body LoginJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.Login(ctx.UserContext(), request.(LoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Login")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(LoginResponseObject); ok {
		if err := validResponse.VisitLoginResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Register operation middleware
func (sh *strictHandler) Register(ctx *fiber.Ctx) error {
	var request RegisterRequestObject

	var body RegisterJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.Register(ctx.UserContext(), request.(RegisterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Register")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(RegisterResponseObject); ok {
		if err := validResponse.VisitRegisterResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xUQW8TTQz9K5G/77jpbilc5tZKIAWQqFIQB8Rhsutmp90dD7a3EKL978izSXooXFCC",
	"OHBar8fjZ/s9zxZWXvDaawsOyjSsulBDATX1iSJGFXBbkLrF3mfzJTOxGYkpIWvA7O5RxK/RTN0kBAei",
	"HOIaCvg2X9PcnHO5D2lOSQNF380ThajI4JQHHMdif5FWd1grjAW8pXWIT7Gw96H7faQCkhf5StycoNol",
	"SqIo+LRqpXuMx0QcCxCsBw66uTF+JpgVeka+HIzP/d8r4t4rOHj98T0UE5uWaTqFQ+ZWNcFoiUO8pVxr",
	"0M5OLq8XUMADsgSK4OD8rDqrrGlKGH0K4OAiu2y42uZKyu5AH4na18bhrbtFA27HbgGMXwYUvaJmY0E1",
	"RcWY431KXajzjfJOKD5K0az/GW/BwX/lo1bLnVDLKfeYe+EdJbmoZ1V1XJAD4RmrQak5ZAbBwbs3NqLn",
	"R4Sctu8nUFe+mS2nQU6Y56fH/BD9oC1x+I6Ngb74E40ubCei72Y3yA/Is31gATL0vefNQVnmKxnXQfIW",
	"/UqGy33EPyWeQol/jSgOPI/Ty2mBAu7TFgbudo+fK8uOat+1JOouqqqC8fP4IwAA//8euozjIgcAAA==",
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
