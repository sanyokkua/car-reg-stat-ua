package models

type Contributor struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	Title string `json:"title"`
}

type License struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Resource struct {
	Mimetype string `json:"mimetype"`
	Profile  string `json:"profile"`
	Name     string `json:"name"`
	Format   string `json:"format"`
	Encoding string `json:"encoding"`
	Path     string `json:"path"`
}

type DataPackage struct {
	Profile      string        `json:"profile"`
	Name         string        `json:"name"`
	Contributors []Contributor `json:"contributors"`
	Created      string        `json:"created"`
	Title        string        `json:"title"`
	Keywords     []string      `json:"keywords"`
	Version      string        `json:"version"`
	Licenses     []License     `json:"licenses"`
	Homepage     string        `json:"homepage"`
	Id           string        `json:"id"`
	Resources    []Resource    `json:"resources"`
	Description  string        `json:"description"`
}

type Department struct {
	Name string //Department                string // "DEP"   -- Can contain departmentCode "XXXX code"
	Code string //DepartmentCode            string // "DEP_CODE"  -- Should be string, there are numbers that starts from 0XXX
}

type Operation struct {
	Name string //OperationName             string // "OPER_NAME" -- Can contain opCode "code - XXXXX"
	Code int    //OperationCode             int    // "OPER_CODE" -- Number
}

type Vehicle struct {
	Brand          string //Brand                     string // "BRAND" -- Can contain mode "Brand MODEL"
	Model          string //Model                     string // "MODEL" -- String
	Characteristic Characteristic
}

type Characteristic struct {
	MakeYear    int    // "MAKE_YEAR" -- Number
	Color       string // "COLOR" -- String
	Body        string // "BODY"  -- String
	Fuel        string // "FUEL"  -- String, Nullable
	Capacity    int    // "CAPACITY"  -- Number, Nullable
	OwnWeight   int    // "OWN_WEIGHT"    -- Float, Nullable
	TotalWeight int    // "TOTAL_WEIGHT"  -- Number, Nullable
}
type CsvRecord struct {
	Person                    string // "PERSON"    -- String
	RegistrationAddressKoatuu string // "REG_ADDR_KOATUU"   -- String, Nullable
	Operation                 Operation
	DateRegistration          string // "D_REG" -- String, Date TODO:
	Department                Department
	Vehicle                   Vehicle
	Kind                      string // "KIND"  -- String
	Purpose                   string // "PURPOSE"   -- String
	NumberRegistrationNew     string // "N_REG_NEW" -- String, Nullable
	Vin                       string // "VIN"   -- String, Nullable
}
