package structs

const apiVersion string = "v1/"
const baseURL string = "https://api.userede.com.br/desenvolvedores/" + apiVersion + "transactions/"

// APIBaseURL return the string containing the base url to the rede's api
func APIBaseURL() string {
	return baseURL
}
