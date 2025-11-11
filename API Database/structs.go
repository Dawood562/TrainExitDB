package main

// Struct represents a railway operator
type operator struct {
	Code	string	`json:"operatorCode"`
	Name	string	`json:"operatorName"`
}

// Struct represents a railway route
type route struct {
	Identifier	string	`json:"identifier"`
	Name		string	`json:"routeName"`
	Beginning	string	`json:"beginning"`
	End			string	`json:"end"`
}

// Represents the stops on a given route
type routeStops struct {
	Identifier	string	`json:"identifier`
	Station		string	`json:"Station"`
	Position	int	`json:"position"`
	Platform	int	`json:"platform"`
}

// Struct represents a train station.
type station struct {
    Code        string  `json:"code"`
    Name        string  `json:"name"`
    Operator    string  `json:"operator"`
	Platforms   int     `json:"platforms"`
}

// Struct represents an exit at a station
type exit struct {
	Platform	int		`json:"platform"`
	Carriage	int		`json:"carriage"`
	Door		string	`json:"door"`
	ExitType	string	`json:"exitType"`
}