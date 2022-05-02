docker-compose down
git pull origin main
docker-compose build
docker-compose up -d redis gateway_service
docker-compose up -d user farm contract_service governance