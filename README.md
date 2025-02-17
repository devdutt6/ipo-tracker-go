# IPO tracker Go

Pull IPO allotment status from 4 registrars MAASHITLA, BIGSHARE, LINKINTIME, and CAMEO for all you pans.

## Tools

Go, Fiber, Mongo

## How to start

Make sure you have Golang setup ready on your machine then **clone repo** then run `go mod tidy` to install all deps
the to start `go run ./cmd/main.go`

## APIs

- POST /register
- POST /login
- GET, POST, DELETE /pan
- GET checkAllotmentStatus
