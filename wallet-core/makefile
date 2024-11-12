create_migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq init

migrate_up:
	migrate -path=internal/database/migrations -database "mysql://root:root@tcp(localhost:3306)/wallet" -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database "mysql://root:root@tcp(localhost:3306)/wallet" -verbose down

migrate_down_one:
	migrate -path=internal/database/migrations -database "mysql://root:root@tcp(localhost:3306)/wallet" -verbose down 1

migrate_version:
	migrate -path=internal/database/migrations -database "mysql://root:root@tcp(localhost:3306)/wallet" version

migrate_force:
	migrate -path=internal/database/migrations -database "mysql://root:root@tcp(localhost:3306)/wallet" force 3

.PHONY: create_migration migrate_up migrate_down migrate_down_one migrate_version migrate_force
