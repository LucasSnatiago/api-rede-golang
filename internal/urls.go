package internal

const apiVersion string = "v1/"
const baseURL string = "https://api.userede.com.br/erede/" + apiVersion + "transactions/"
const baseURLTest string = "https://api.userede.com.br/desenvolvedores/" + apiVersion + "transactions/"

// APIBaseURL return the string containing the base url to the rede's api
func APIBaseURL() string {
	return baseURL
}

// APIBaseURLTest return the string containing the base url to the rede's api
func APIBaseURLTest() string {
	return baseURLTest
}