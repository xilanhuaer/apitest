package utils

import (
	"strconv"
	"strings"

	"github.com/xilanhuaer/http-client/common/entity"
	"github.com/xuri/excelize/v2"
)

func ReadExcel(path string, data *[]entity.Data, maps map[string]interface{}) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		panic(err)
	}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	for index, row := range rows {

		if index == 0 {
			continue
		}

		content := entity.Data{}
		content.Project.Name = row[0]
		content.Project.TestingMethod = row[1]
		content.Project.TestingBasis = row[2]
		content.Project.TargetGene = row[3]
		content.Project.TechPlatform = row[4]
		content.Project.FlowID = parseString(isOk("flow", row[5], maps))
		content.Project.PricingMethod = isOk("price", row[6], maps)
		content.Project.Price = parseString(row[7]) * 100
		content.Project.DetermineCondition = row[8]
		parts := strings.Split(row[9], ",")
		for _, v := range parts {
			content.SampleCategoryIDs = append(content.SampleCategoryIDs, parseString(maps["sample"].(map[string]string)[v]))
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
						FlowStepID: parseString(maps["control"].(map[string]string)[operationStep[1]]),
						WidgetType: parseString(maps["step"].(map[string]string)[operationStep[2]]),
					},
				)
			}
		}
		*data = append(*data, content)
	}
}

func isOk(locate, key string, data map[string]interface{}) string {
	value, ok := data[locate].(map[string]string)[key]
	if ok {
		return value
	}
	return ""
}

func parseString(str string) int {
	data, _ := strconv.Atoi(str)
	return data
}
