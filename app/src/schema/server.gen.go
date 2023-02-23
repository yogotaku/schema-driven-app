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

	"H4sIAAAAAAAC/8RWa28bRRf+K9Z5++llHa+d0Db7qU0LKAKlUUv5QDDSdHdsT9lbZ8dtjWWJ9UIppKhV",
	"BAkpFTdBVAiUIEq5BfpjJnbbf4HOzPq6zkVKVT55vXvmXJ95ntMEO/DCwKe+iMBqQmTXqEfU4wucBxwf",
	"Qh6ElAtG1Ws7cCj+OjSyOQsFC3ywtHFOfTNANEIKFjBf0Crl0DLAo1FEqrue633uH40EZ34VWi0DOL1U",
	"Z5w6YC1B6r9nXm4ZsECvLFKRTdMnngrnkauvUL8qamAdNcBjfu9f0YCQCEE5pvHm0sn86+VmsbVE8m+X",
	"/38km4kBglTHHBaf38sjenruyL41qUS197Se8xGd0HeHCHqmMsc4hmoCvUq80EW/JbM4mzeP50szYEAl",
	"4B4RYCnzSVVQjzAXHfQt9ZsJphXGI7HQ6+NInRlblxzQdKz6QYwhH70sywYIJlSRvbb0/QUXLlJbYOh0",
	"/MR1z1TAWmrCEU4rYMH/CgNoF1JcF1K4tIzx9jInC06Z3JZJIpPrMr4nkw2ZLMtkW7Z/mj8NavDMq3tg",
	"mVnAt9D/cJnMgXILx9ub7YGzVQey6dqcEkFP45AnpL2h8nwgk+1u8l7ny59lfG/n7zvd67e6a99219vY",
	"4MPA5zXKWYVRZ+/QMr4hk6/wIdl8/P1Hj+9uy/ZK596NnT+vyXhNxhsyfle2l2W8LOPvZHwNH95pv+Fn",
	"TG7IeEvwOlVfB0ldCAKXEh+zYvumMja/KTjY0MYLNoYbjxNFPDO/Emhm9AWxFRjTWwa1oEpPpL2esgMP",
	"jJSYQJC36g2SawTVACsYzb1klqYL5nTBLMl488nqcqf9g4zXZXu5s/xBZ/u3ne3bjz6+i0W1f5XJLzJZ",
	"k8nmycV5GNwXjaC8w9ll6udJGIIBlymPtP/ilIlBg5D6JGRgwfSUOWWC4q+awlchpFoQqvp27XIxUni1",
	"V7rf3Hl0/2udpBoSgpWg+bwDFrzIfGcRPWIITjwqKI8U/Ecd99ysyfgLnP3wBWw/lO0tHRCw6WDBpTrl",
	"jeGeVjGCrl1dakE9VUWWzPULwjlpgJr8pER2/nrQ/WQLb8/vHz5Zv7VLWJd5TIzEzUCrjKwXhYEf6etb",
	"Ms0eYqiv6SsMXWarjhUuRpjDEMcvaX5CvUljLgZ2jaWyYYETVBVHKKOjfaNzNcaDyUbH+kavEo/0bWwi",
	"AJOd0MK9eEpT6nhTM7A+87KGeoXUXXGA+gdZ7BVc7ykTwtV9ejWktqBOjqY2BkR1zyO8ARacpaLO/ShH",
	"XDcXamwqAFlLEFIBZbQu1CMF1CaEQST2IZn2Snd16/HGTU21+i5kbsIpxR+pliHb0EjMBU7jqbWjrxkt",
	"rbUjqCs+tTC9GONNx/c5XaSD0545ENIPOek54uTO6l6ODVlnklugV3K9/SGdMU52eMiFJv7MO61dWY85",
	"Mt6UyeeKkuKUpLJa217p3Fzt/LO2CwCQCjGVuYZaIw5DC6NrwKiSj6yLUJydPZYvmvnpYn+5soC4zKZT",
	"kcdE7UQV36UKNabxKL0jyyCcxIOgZbc4Uxre/uAcusPBHw5GmiqeFXgmY+clKhRocvN+JchdaKR/TmdB",
	"tI+qyeR9BZKHMllXD3/IeLPz46edO3c1SMZ2FRVCKQ3q8UBoNEBheIPWwxl0oL+VFo2JKhQSYdf2XZYG",
	"sv7Z/e7q1i5IPh86KZX1sfyM6WwmW8lCkDuVxv3v2Ud3aAAiXK7xTJaF8Bjllyfjxw1s4tZQgAyoc7VX",
	"ChFahUL/g3XcnDWVcqeOm8OYUbKf/kdpa5Vb/wYAAP//d0vFvfgPAAA=",
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
