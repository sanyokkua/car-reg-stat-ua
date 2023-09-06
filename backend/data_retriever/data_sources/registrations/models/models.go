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

type DataPackageJson struct {
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

type CsvRecord struct {
    DepartmentName            string // "DEP"   -- string Can contain departmentCode "XXXX code"
    DepartmentCode            string // "DEP_CODE"  -- string Should be string, there are numbers that starts from 0XXX
    OperationName             string // "OPER_NAME" -- string Can contain opCode "code - XXXXX"
    OperationCode             int    // "OPER_CODE" -- int Number
    Brand                     string // "BRAND" -- string Can contain mode "Brand MODEL"
    Model                     string // "MODEL" -- String
    MakeYear                  int    // "MAKE_YEAR" -- Number
    Color                     string // "COLOR" -- String
    Body                      string // "BODY"  -- String
    Fuel                      string // "FUEL"  -- String, Nullable
    Capacity                  int    // "CAPACITY"  -- Number, Nullable
    OwnWeight                 int    // "OWN_WEIGHT"    -- Float, Nullable
    TotalWeight               int    // "TOTAL_WEIGHT"  -- Number, Nullable
    Person                    string // "PERSON"    -- String
    RegistrationAddressKoatuu string // "REG_ADDR_KOATUU"   -- String, Nullable
    DateRegistration          string // "D_REG" -- String, Date
    Kind                      string // "KIND"  -- String
    Purpose                   string // "PURPOSE"   -- String
    NumberRegistrationNew     string // "N_REG_NEW" -- String, Nullable
    Vin                       string // "VIN"   -- String, Nullable
}
