## Cocom
### Backend service for Soda Vending Machine written in Go

### Running Cocom

1 Compile the code:
```bash
make dependencies
make build
```
2 Run the app:
```bash
./cocom
```

### Running the project with Docker
```bash
docker build -t cocom . && docker run -p 8080:8080 -it cocom
```

### REST API
```http request
GET  http://localhost:8080/api/v1/sodas
POST http://localhost:8080/api/v1/coins  {"coins" : 100}
POST http://localhost:8080/api/v1/soda/{sodaID}/purchase

Admin access

Authorization: Bearer SuperSecretToken

POST http://localhost:8080/api/v1/admin/soda/{sodaID}/inventory  {"amount" : 100}
POST http://localhost:8080/api/v1/admin//soda/{sodaID}/price  {"price" : 100}
POST http://localhost:8080/api/v1/admin/sodas  {"id: "soda-id", "name": "SodaName", "description":"Soda Description", "price": 100, "quantity": 0}
```

### Note
All prices are in cents $1.00 = 100 cents
