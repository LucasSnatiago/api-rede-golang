# Unnoficial go version of [Rede's API](https://www.userede.com.br/desenvolvedores/pt/produto/e-Rede#tutorial).

This repository contains the basic methods to use the Rede's API.

## Usage

Add this line to the imports of your project:
```go
import github.com/lusantisuper/api-rede-golang
```

## Usage

Edit the [login.json](https://github.com/lusantisuper/api-rede-golang/blob/main/login.json) file and add your keys or use the method: ```login.ReadLogin(PV, IntegrationKey)```

#
The request struct contains all data needed to make the payment request.
```go
struct structs.Request
```
#
Use the method Pay to make the payment:
```go
func Pay(r *structs.Request, login *login.Login) error
```
The function returns a error if the payments didn't work, else return nil.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache V2](https://choosealicense.com/licenses/apache-2.0/)