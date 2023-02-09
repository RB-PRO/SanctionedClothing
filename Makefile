all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/SanctionedClothing.git

pull:
	git pull git@github.com:RB-PRO/SanctionedClothing.git

pushW:
	git push https://github.com/RB-PRO/SanctionedClothing.git

pullW:
	git pull https://github.com/RB-PRO/SanctionedClothing.git

pushCar:
	scp main root@194.87.107.129:go/SanctionedClothing/

build-config:
	go env GOOS GOARCH

build-windows-to-linux:
	set GOARCH=amd64 set GOOS=linux go build .\cmd\main\main.go  

build-linux-to-windows:
	export GOARCH=amd64 export GOOS=windows go build .\cmd\main\main.go  