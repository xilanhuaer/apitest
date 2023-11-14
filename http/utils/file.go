package utils

import (
	"github.com/xilanhuaer/http-client/common/entity"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func ReadExcel(path string, data *[]entity.Data) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		panic(err)
	}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		content := entity.Data{}
		content.Project.Name = row[0]
		content.Project.TestingMethod = row[1]
		content.Project.TestingBasis = row[2]
		content.Project.TargetGene = row[3]
		content.Project.TechPlatform = row[4]
		content.Project.FlowID = parseString(row[5])
		content.Project.PricingMethod = row[6]
		content.Project.Price = parseString(row[7]) * 100
		content.Project.DetermineCondition = row[8]
		parts := strings.Split(row[9], ",")
		for _, v := range parts {
			content.SampleCategoryIDs = append(content.SampleCategoryIDs, parseString(v))
		}
		for i := 10; i < len(row); i++ {
			if row[i] != "" {
				operationStep := strings.Split(row[i], ",")
				if len(operationStep) != 3 {
					continue
				}
				content.OperationSteps = append(content.OperationSteps,
					entity.OperationSteps{
						StepName:   operationStep[0],
						FlowStepID: parseString(operationStep[1]),
						WidgetType: parseString(operationStep[2]),
					},
				)
			}
		}
		*data = append(*data, content)
	}
}

func parseString(str string) int {
	data, _ := strconv.Atoi(str)
	return data
}
