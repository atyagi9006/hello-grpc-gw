`hello-grpc-gw` is a hello world grpc-gw app.

Steps For generating cert :
1. openssl genrsa -out pkg/cert/server.key 2048
2. openssl req -new -x509 -sha256 -key pkg/cert/server.key -out pkg/cert/server.crt -days 3650
3. openssl req -new -sha256 -key pkg/cert/server.key -out pkg/cert/server.csr
4. openssl x509 -req -sha256 -in pkg/cert/server.csr -signkey pkg/cert/server.key -out pkg/cert/server.crt -days 3650

e.g 
➜$ openssl req -new -sha256 -key pkg/cert/server.key -out pkg/cert/server.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:in
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:localhost
Email Address []:

repeat same for crt.

project tree
.
├── Makefile
├── Readme.md
├── bin
│   ├── client
│   └── server
├── client
│   └── client.go
├── go.mod
├── go.sum
├── pkg
│   ├── api
│   │   └── handler.go
│   ├── cert
│   │   ├── server.crt
│   │   ├── server.csr
│   │   └── server.key
│   ├── proto
│   │   ├── api.pb.go
│   │   ├── api.pb.gw.go
│   │   ├── api.proto
│   │   └── api.swagger.json
│   └── start
│       ├── auth.go
│       └── start.go
└── server
    └── server.go

build server : make build-server
build client : make build-client

run server : bin/server
run client : bin/client

sample http req:
curl -H "login:amit" -H "password:tyagi" -X POST -d '{"greeting":"foo"}' 'http://localhost:7778/1/ping'      

Response : {"greeting":"Milgya greeting...bas rehn de.... gw"}