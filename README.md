
#Testing Go HTTP Endpoints

Running the API server

#Make sure that you set your GOPATH correctly
go run src/main.go

#Check root path
curl -i localhost:9000/

#Check stats path
curl -i localhost:9000/stats

#Check favicon.ico path
curl localhost:9000/favicon.ico
