docker-compose down
# git pull origin main
git pull origin stake_event
# git pull origin stage-v2
docker-compose build
docker-compose up -d redis gateway_service
docker-compose up -d user farm contract_service governance event_service