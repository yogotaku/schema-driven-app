// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams)
	// Create New User
	// (POST /users)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Get User Info by User ID
	// (GET /users/{userId})
	FindUserByID(w http.ResponseWriter, r *http.Request, userId int)
	// Update User Information
	// (PATCH /users/{userId})
	UpdateUserByID(w http.ResponseWriter, r *http.Request, userId int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// FindPets operation middleware
func (siw *ServerInterfaceWrapper) FindPets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams

	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tags", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPets(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindUserByID operation middleware
func (siw *ServerInterfaceWrapper) FindUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, chi.URLParam(r, "userId"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindUserByID(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUserByID operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, chi.URLParam(r, "userId"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUserByID(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pets", wrapper.FindPets)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{userId}", wrapper.FindUserByID)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/users/{userId}", wrapper.UpdateUserByID)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RXW28bxRf/Ktb5/x838SVBbfapTQsoQkqjovJS8jDdHdtTdne2s+O2xrLEeqEUUtQq",
	"goSUipsgKgRKEKXcAv0wE7vtt0BnZte3tZPeVJ6ynjlzbr/f+c2kBQ73Qx7QQEZgtyBy6tQn+vNlIbjA",
	"j1DwkArJqF52uEvxr0sjR7BQMh6AbYwLes8C2Qwp2MACSWtUQNsCn0YRqU09l233j0ZSsKAG7bYFgl5o",
	"MEFdsM9C6j8zX21bsEwvrVCZTzMgvg435tACSWoT1scC6dOp+zMRndAGl0h6qrrIhKzjT3qZ+KGHHiul",
	"8sJM6ehMZR4sqHLhEwm2Ns+XZwH1CfPQQd/SrEwwrTIRyeVpZXlk6uZYbQM/Q6eyTFYtkEzqQrLS+6nw",
	"c+epIzFY2nHieaeqYJ9twf8FrYIN/ysO2FRMqVRMEWpb4y1kbp4PKrmpkkQlV1V8RyXbKllTyZ7q/LR0",
	"cgKv2uhzuDTmwmp71crQiHRuGKb8kpUyAla4U2eQ8gBcXoM2Ap2h/Ng16QP5ohxBiaQnEe4JxW3rau6p",
	"ZK+XvNf98mcV39n/+1bv6o3e5re9rQ5Yz0akN6hgVUbdg0Or+JpKvsKPZOfh9x89vL2nOuvdO9f2/7yi",
	"4k0Vb6v4XdVZU/Gair9T8RX8eKfzZpAzuabiXSkaVO8OkjrHuUdJgFmxQ1MZQ3n2MWEeLzgP+zAQo70c",
	"GV0oLywcmSmXZubK/SGwgXjMobORz2T9WA3XZh3u54LaWPzIYMJxPAim8PJ8ZXgu4XV0h2zDiWRBlRs5",
	"DSRx9Dhlweu8Ro+ltaRxU+5K8lajSQpNXuPY3dG+VkqVuWJprliqqHjn0cZat/ODirdUZ6279kF377f9",
	"vZsPPr6NDe/8qpJfVLKpkp3jK0swmHge0oCEOBwXqYiM1/JsCUNlWzbMzeKSBSGRdc34YkjN3VEzqjBl",
	"oFPCd9Z739x6cPdrk5qmDY4PQfMlF2x4hQXuCnrEEIL4VFJhEB11nLnZVPEXyMZh4ejcV51dExCxQPML",
	"DSqaw52sYQQzz1qMJPWjyReGWSBCkCZoLk5KZP+ve71PdnGef//w0daNKWE95jM5EjdH9lVU6yjkQWQE",
	"pVIqZTyhgZHdMPSYoztWPB9hDq1JhRykX0aQx0vLUerUa4ZmVdLw5BNlcVBw87CYEK4R0MshdSR1CzS1",
	"sSBq+D4RTbDhNJUNEUQF4nmF0DBEw2ifhZBKWEXrYiPSdGlByCN5iPh01nsbuw+3rxsJNozM8fGE1pH0",
	"JkQVopFc5G7zubWjf5e0zU09gn35uYXJYow3HdcLpkgX0Z5/Qr49FdKLxC2cNr0cA9lkUlimlwrZ6yPF",
	"GJEdBrnYwj9Lbnuq9jBXxTsq+VwLQ5xKRf4O7qx3r290/9mcQgAUJExlsakfIc80nE8DkBnCFwXLZFRe",
	"pVLDUVgKqrxwrpn+OJmH5xDVVsn7uv33VbKlP/5Q8U73x0+7t26b9o+9DnQIraR43wyE1EAPwy9bcxkf",
	"rKwhkU790CfJ4Kr67G5vY3cKL86EbioMfWa8YHGYz1eyzAsn0rj//SybDg2Ig09YPJOfaTxGxcXJnPG4",
	"Q7w6yrkFDaFfSFKGdrHY37CPlhZKgACnjlvDPNGv9PQ3XhRod3kmkjz0WK0us39FwL+0QMQRcvTt8xc9",
	"B9rtfwMAAP//iTx4RRwPAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
