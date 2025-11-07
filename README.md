# raks-coin-exchange

# Integratoin Explanation
- Please see ./internal/transaction/transaction_handle.go for reference.

## Library Import
- the library package "github.com/paweenwatkwanja/transaction-broadcasting" is imported with naming specifically for this local package.

## Integration
- I wrap the broadcast object instantiation inside initBroadcastService() method.
- In this method, I assign two custom requests. One for transaction a retry machanism and one for a HTTP request's retry mechanism and other configurations.
- There we have an object of broadcastService.

## Usage
### Prepare a Broadcast Request
- I create an object of BroadcastRequest with details of transaction. These details come from the context of local service.

### Prepare a Url for Broadcasting
- this forms a whole Url  for braodcasting a transaction.

### Call the BroadcastTransaction()
- this requires an endpoint and a BroadcastRequest.
- it returns the txHash and an error.

### Proccess any local busiess logic (Optional)

### Prepare a Url for Monitoring
- this forms a whole Url for monitoring a transaction by concatenating the endpoint, the path, and the parameter (txHash).

### Call the MonitorTransaction()
- this requires only the Url.
- it returns a txStatus and an error.

### Proccess any local busiess logic (Optional)

### Call the HandleStatus()
- this requires the same Url as MonitorTransaction()'s and the txStatus.

### End of transaction broadcasting

### Proccess any local busiess logic (Optional)

# How to Run this Program
## Commands
### start service with docker compose
- `docker compose -f docker-compose-yaml up --build`
- *** please provide endpoint in docker-compose.yaml (the endpoint under Environment)
- *** no `/` at the end of the endpoint

### down service with docker compose
- `docker compose -f docker-compose-yaml down`

### start with Go command
- `APP_HOST=<app_host> APP_PORT=8080 ENDPOINT=<endpoint> go run main.go`

### POSTMAN
- Method: POST
- URL: http://<app_host>:8080/v1/transactions
- request 
    - {
    "symbol": "ETH",
    "price": 100000,
    "timestamp": 1678912345,
    "retry_attempt": 10,
    "retry_duration": 10,
    }
- other fields with type int
    - "retry_attempt_http"
    - "retry_duration_http"
    - "timeout"