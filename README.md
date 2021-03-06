# RESTFULL API Go Lang

Buil Simple RESTFULL API GO LANG Using gorilla mux http library

## Usage
``` bash
git clone https://github.com/wildanicardi/restapi-go.git 
cd restapi-go
go build
go run main.go

```

## API Endpoints
### All Article
 * Path : ``` /api/articles```
 * Method : ```GET```
 * Response : ``` 200```

### Create Single Article
 * Path : ``` /api/articles```
 * Method : ```POST```
 * Field : ```Title, Description, User ```
 * Response : ``` 201```

### Details Article
 * Path : ``` /api/article/{id}```
 * Method : ```GET```
 * Response : ``` 200```

### Update Article
 * Path : ``` /api/article/{id}```
 * Method : ```PUT/PATCH```
 * Field : ```Title, Description, User ```
 * Response : ``` 200```

### Delete Article
 * Path : ``` /api/article/{id}```
 * Method : ```DELETE```
 * Response : ``` 200```


## License
[MIT](https://choosealicense.com/licenses/mit/) wildanicardi