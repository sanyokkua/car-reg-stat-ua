package models

type Department struct {
	Code string
	Name string
}

type Operation struct {
	Code int
	Name string
}

type Brand struct {
	BrandId int
	Name    string
}

type Model struct {
	ModelId int
	Name    string
}

type BodyType struct {
	BodyTypeId   int
	BodyTypeName string
}

type FuelType struct {
	FuelTypeId   int
	FuelTypeName string
}

type Color struct {
	ColorId int
	Name    string
}

type Kind struct {
	KindId   int
	KindName string
}

type Vehicle struct {
	VehicleId      int
	Kind           Kind
	Brand          Brand
	Model          Model
	BodyType       BodyType
	FuelType       FuelType
	Color          Color
	MakeYear       int
	EngineCapacity int
	OwnWeight      int
	TotalWeight    int
}

type Registration struct {
	PersonType            string
	RegistrationCityCode  string
	Department            Department
	Operation             Operation
	DateRegistration      string
	NumberRegistrationNew string
	Vin                   string
	Vehicle               Vehicle
	Purpose               string
}
