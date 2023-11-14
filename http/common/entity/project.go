package entity

type Data struct {
	Project           Project
	OperationSteps    []OperationSteps
	SampleCategoryIDs []int
}
type Project struct {
	Name               string
	TestingMethod      string
	TestingBasis       string
	TargetGene         string
	TechPlatform       string
	FlowID             int
	PricingMethod      string
	Price              int
	DetermineCondition string
}
type OperationSteps struct {
	StepName   string
	FlowStepID int
	WidgetType int
}
