# Go-Chassis Example

This is a example illustrates creating a simple micro service using go-chassis.


#### Running the Client and Server 
Client and Server in this example can be run on the same machine
- Start the Service-Center
```go
docker run -d -p 30100:30100 servicecomb/service-center
```  
- Get Chassis
```go
go get github.com/ServiceComb/go-chassis
```
- Download this example
- Start the Client
```go
cd client
go build
./client
```
- Start the Server
```go
cd server
go build
./server
```
- Verify the Communication between client and server
```go
curl -X GET   http://127.0.0.1:8081/saymessage/EngineeringDay
Welcome to : Engineeringday
```
