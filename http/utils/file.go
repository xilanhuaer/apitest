package utils

import (
	"github.com/gogf/gf/v2/util/gconv"
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
		content.Project.FlowID = gconv.Int(maps["flow"].(map[string]string)[row[5]])
		content.Project.PricingMethod = gconv.Int(maps["price"].(map[string]string)[row[6]])
		content.Project.Price = gconv.Int(row[7]) * 100
		content.Project.DetermineCondition = row[8]
		parts := strings.Split(row[9], ",")
		for _, v := range parts {
			content.SampleCategoryIDs = append(content.SampleCategoryIDs, gconv.Int(maps["sample"].(map[string]string)[v]))
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
						FlowStepID: gconv.Int(maps["control"].(map[string]string)[operationStep[1]]),
						WidgetType: gconv.Int(maps["widget"].(map[string]string)[operationStep[2]]),
					},
				)
			}
		}
		*data = append(*data, content)
	}
}
