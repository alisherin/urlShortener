``
docker-compose up -d --build
``

``
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
``

``
export POSTGRESQL_URL='postgres://shortener:123456@localhost:5432/shortener?sslmode=disable'
``

``
migrate -database ${POSTGRESQL_URL} -path database/migrations up
``

``
go run main.go
``