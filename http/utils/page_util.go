package utils

import (
	"fmt"
	"strconv"
)

func PageUtil(pageString, sizeString string) (limit, offset int, err error) {
	page, err := strconv.Atoi(pageString)
	if err != nil || page <= 0 {
		return 0, 0, fmt.Errorf("page is invalid, please check your request")
	}
	limit, err = strconv.Atoi(sizeString)
	if err != nil || limit <= 0 {
		return 0, 0, fmt.Errorf("page_size is invalid, please check your request")
	}
	offset = (page - 1) * limit
	return limit, offset, err
}
