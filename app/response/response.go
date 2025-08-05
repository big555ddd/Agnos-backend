package response

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// conventionalMarshallerFromPascal wraps a value for custom JSON marshalling
type conventionalMarshallerFromPascal struct {
	Value interface{}
}

// Regex patterns for naming convention conversion
var (
	keyMatchRegex    = regexp.MustCompile(`"(\w+)":`)
	wordBarrierRegex = regexp.MustCompile(`([a-z\d])([A-Z])`)
)

// convertToCamelCase converts PascalCase to camelCase in JSON
func convertToCamelCase(data []byte) []byte {
	return keyMatchRegex.ReplaceAllFunc(data, func(match []byte) []byte {
		key := string(match[1 : len(match)-2]) // Remove quotes and colon
		if len(key) > 0 {
			// First convert any underscores to camelCase (snake_case to camelCase)
			if bytes.Contains([]byte(key), []byte("_")) {
				parts := bytes.Split([]byte(key), []byte("_"))
				result := string(parts[0])
				for i := 1; i < len(parts); i++ {
					if len(parts[i]) > 0 {
						runes := []rune(string(parts[i]))
						runes[0] = unicode.ToUpper(runes[0])
						result += string(runes)
					}
				}
				key = result
			}

			// Then ensure first letter is lowercase (PascalCase to camelCase)
			runes := []rune(key)
			runes[0] = unicode.ToLower(runes[0])
			return []byte(`"` + string(runes) + `":`)
		}
		return match
	})
}

// MarshalJSON implements custom JSON marshalling with naming convention support
func (c conventionalMarshallerFromPascal) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	if err != nil {
		return nil, err
	}
	naming := viper.GetString("HTTP_JSON_NAMING")

	val := reflect.TypeOf(c.Value)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {

		field, ok := val.FieldByName("json")
		if ok {

			if field.Tag.Get("naming") != "" {
				naming = field.Tag.Get("naming")
			}
		}
	}

	var converted []byte
	switch naming {
	case "snake_case":

		converted = keyMatchRegex.ReplaceAllFunc(
			marshalled,
			func(match []byte) []byte {
				return bytes.ToLower(wordBarrierRegex.ReplaceAll(
					match,
					[]byte(`${1}_${2}`),
				))
			},
		)
	case "camel_case":
		converted = convertToCamelCase(marshalled)
	case "pascal_case":
		return marshalled, nil
	default:
		return nil, err
	}

	return converted, nil
}

// NewConventionalMarshaller creates a new conventionalMarshallerFromPascal
func NewConventionalMarshaller(value interface{}) conventionalMarshallerFromPascal {
	return conventionalMarshallerFromPascal{Value: value}
}

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponsePaginate struct {
	Code       int64      `json:"code"`
	Message    string     `json:"message"`
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ResponsePaginate0 struct {
	Code       int64  `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Pagination any    `json:"pagination"`
}

// Success ส่งผลลัพธ์เมื่อสำเร็จ
func Success(ctx *gin.Context, data any) {
	response := Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusOK, marshalled)
}



// InternalError ส่งผลลัพธ์เมื่อมีข้อผิดพลาดภายใน
func InternalError(ctx *gin.Context, message any, data any) {
	response := Response{
		Code:    500,
		Message: message.(string), // Set the message directly here
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusInternalServerError, marshalled)
}

func NotFound(ctx *gin.Context, message any, data any) {
	response := Response{
		Code:    404,
		Message: message.(string), // Set the message directly here
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusNotFound, marshalled)
}

// BadRequest ส่งผลลัพธ์เมื่อมีข้อผิดพลาดจากการขอข้อมูลที่ไม่ถูกต้อง
func BadRequest(ctx *gin.Context, message any, data any) {
	response := Response{
		Code:    400,
		Message: message.(string), // Set the message directly here
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusBadRequest, marshalled)
}

func Unauthorized(ctx *gin.Context, message any, data any) {
	response := Response{
		Code:    401,
		Message: message.(string), // Set the message directly here
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusUnauthorized, marshalled)
}

type Pagination struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func SuccessWithPaginate(ctx *gin.Context, data any, page, size, count int) {

	pagination := Pagination{
		Page:  page,
		Size:  size,
		Total: count,
	}

	if pagination.Total == 0 {
		response := ResponsePaginate0{
			Code:       200,
			Message:    "Success",
			Data:       []any{},
			Pagination: gin.H{},
		}

		marshalled := NewConventionalMarshaller(response)
		ctx.JSON(http.StatusOK, marshalled)
		return
	} else {
		response := ResponsePaginate{
			Code:       200,
			Message:    "Success",
			Data:       data,
			Pagination: pagination,
		}

		marshalled := NewConventionalMarshaller(response)
		ctx.JSON(http.StatusOK, marshalled)
	}
}

func Forbidden(ctx *gin.Context, message any, data any) {
	response := Response{
		Code:    403,
		Message: message.(string), // Set the message directly here
		Data:    data,
	}

	marshalled := NewConventionalMarshaller(response)
	ctx.JSON(http.StatusForbidden, marshalled)
}
