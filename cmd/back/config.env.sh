#SERVER
export HOST=0000
export GRPC="8080"
export METRIC="8081"
export GW="8082"
export FILES="8084"

export AUTH_KEY="super-duper-secret-key-qwertyuio"
export DEV_MODE=true

#FILE_STORE SETTINGS
export SECURE=false
export ENDPOINT="minio:9000"
export ACCESS_KEY="test_svc"
export SECRET_KEY="test_pass"

#DB
export MIGRATE_DIR="migrate"
export DRIVER="postgres"
export POSTGRES="postgres://user_svc:user_pass@postgres:5432/user_db?sslmode=disable"