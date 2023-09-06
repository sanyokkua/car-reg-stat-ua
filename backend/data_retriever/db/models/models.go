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

type Vehicle struct {
    VehicleId      int
    Kind           string
    Brand          Brand
    Model          Model
    BodyType       BodyType
    FuelType       string // "",".","НЕ ВИЗНАЧЕНО","NULL","ВІДСУТНЄ"
    Color          string
    MakeYear       int
    EngineCapacity int
    OwnWeight      int
    TotalWeight    int
}

type Registration struct {
    PersonType            string
    RegistrationCityCode  string // "","NULL"
    Department            Department
    Operation             Operation
    DateRegistration      string
    NumberRegistrationNew string
    Vin                   string
    Vehicle               Vehicle
    Purpose               string
}
