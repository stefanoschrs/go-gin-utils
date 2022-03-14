package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractParamString(c *gin.Context, name string) (value string, err error) {
	paramStr, ok := c.Params.Get(name)
	if !ok {
		err = errors.New("cannot get param: " + name)
		return
	}

	value = paramStr
	return
}

func ExtractParamUint(c *gin.Context, name string) (value uint, err error) {
	paramStr, err := ExtractParamString(c, name)
	if err != nil {
		err = fmt.Errorf("ExtractParamString: %w", err)
		return
	}

	param, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("strconv.ParseUint: %w", err)
		return
	}

	value = uint(param)
	return
}

func ExtractParamInt(c *gin.Context, name string) (value int, err error) {
	paramStr, err := ExtractParamString(c, name)
	if err != nil {
		err = fmt.Errorf("ExtractParamString: %w", err)
		return
	}

	param, err := strconv.ParseInt(paramStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("strconv.ParseInt: %w", err)
		return
	}

	value = int(param)
	return
}

// TODO: Move to github.com/stefanoschrs/go-utils
func String(str string) *string {
	return &str
}

// TODO: Move to github.com/stefanoschrs/go-utils
func SliceUintToString(arr []uint) (newArr []string) {
	for _, item := range arr {
		newArr = append(newArr, fmt.Sprint(item))
	}

	return
}
