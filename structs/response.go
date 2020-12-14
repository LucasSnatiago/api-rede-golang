package structs

// Response Struct to return the response of the Rede's API
type Response struct {
	// Return code
	ReturnCode string `json:"ReturnCode"`
	// Return message
	ReturnMessage string `json:"ReturnMessage"`
	// Reference to the Request.Reference code
	Reference string `json:"Reference"`
	// Unique payment indentifier
	Tid string `json:"Tid"`
	// Sequencial number returned by Rede
	Nsu string `json:"Nsu"`
	// Authorization code by the card's manufactory
	AuthorizationCode string `json:"AuthorizationCode"`
	// Date Time

}
