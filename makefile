all:
	@echo "starting all"
	docker-compose up --build -d

prod:
	docker-compose -f docker-compose-prod.yaml up --build -d

stop:
	docker-compose down

gen:
	go mod download gopkg.in/yaml.v3 &&  go get github.com/99designs/gqlgen@v0.17.35 && go run github.com/99designs/gqlgen
make_migrate:
	go get ariga.io/atlas-provider-gorm/gormschema@v0.1.0 && atlas migrate diff --env gorm