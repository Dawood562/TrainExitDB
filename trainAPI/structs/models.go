package structs

// Struct represents a railway operator
type Operator struct {
	Code	string	`json:"operatorCode"`
	Name	string	`json:"operatorName"`
}

// Struct represents a railway route
type Route struct {
	Identifier	string	`json:"identifier"`
	Name		string	`json:"routeName"`
	Beginning	string	`json:"beginning"`
	End			string	`json:"end"`
}

// Represents the stops on a given route
type RouteStops struct {
	Identifier	string	`json:"identifier`
	Station		string	`json:"Station"`
	Position	int	`json:"position"`
	Platform	int	`json:"platform"`
}

// Struct represents a train station.
type Station struct {
    Code        string  `json:"code"`
    Name        string  `json:"name"`
    Operator    string  `json:"operator"`
	Platforms   int     `json:"platforms"`
}

// Struct represents an exit at a station
type Exit struct {
	Platform	int		`json:"platform"`
	Carriage	int		`json:"carriage"`
	Door		string	`json:"door"`
	ExitType	string	`json:"exitType"`
}