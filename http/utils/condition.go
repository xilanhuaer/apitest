package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Filter struct {
	Field   string
	Command string
	Value   string
}
type Condition struct {
	Page     int
	PageSize int
	Filters  []Filter
}

// ParseCondition return query, limit, offset, err
// paramsString: "page=1&page_size=10&filter=id|eq|1#created_at|between|2020-01-01,2020-01-02
func ParseCondition(query *gorm.DB, paramsString string) (*gorm.DB, error) {
	condition := Condition{
		Page:     1,
		PageSize: 10,
		Filters:  []Filter{},
	}

	if paramsString != "" {
		re := regexp.MustCompile(`(page|page_size|filter)=([^&]+)`)
		// matches format : [[page=1 page 1] [page_size=10 page_size 10] [filter=id|eq|1##created_at|between|2020-01-01,2020-01-02
		matches := re.FindAllStringSubmatch(paramsString, -1)
		for _, match := range matches {
			key := match[1]
			value := match[2]
			switch key {
			case "page":
				page, err := strconv.Atoi(value)
				if err != nil {
					return query, fmt.Errorf("page参数错误")
				}
				condition.Page = page

			case "page_size":
				pageSize, err := strconv.Atoi(value)
				if err != nil {
					return query, fmt.Errorf("page_size参数错误")
				}
				condition.PageSize = pageSize
			case "filter":
				condition.Filters = parseFilter(value)
			}
		}
	}
	for _, value := range condition.Filters {
		switch value.Command {
		case "eq":
			query = query.Where(fmt.Sprintf("%s = ?", value.Field), value.Value)
		case "like":
			query = query.Where(fmt.Sprintf("%s like ?", value.Field), fmt.Sprintf("%%%s%%", value.Value))
		case "between":
			// valueParts: ["2020-01-01" "2020-01-02"]
			valueParts := strings.Split(value.Value, ",")
			if len(valueParts) == 2 {
				query = query.Where(fmt.Sprintf("%s between ? and ?", value.Field), valueParts[0], valueParts[1])
			}
		}
	}
	return query.Limit(condition.PageSize).Offset((condition.Page - 1) * condition.PageSize), nil
}

func parseFilter(filterString string) []Filter {
	var filters []Filter
	// filterString: ["id|eq|1" "created_at|between|2020-01-01,2020-01-02"]
	filterParts := strings.Split(filterString, "#")
	// part: "id|eq|1" "created_at|between|2020-01-01,2020-01-02"
	for _, part := range filterParts {
		// conditionParts: ["id" "eq" "1"] ["created_at" "between" "2020-01-01,2020-01-02"]
		conditionParts := strings.Split(part, "|")
		if len(conditionParts) == 3 {
			filters = append(filters, Filter{
				Field:   conditionParts[0],
				Command: strings.ToLower(conditionParts[1]),
				Value:   conditionParts[2],
			})
		}
	}

	return filters
}
