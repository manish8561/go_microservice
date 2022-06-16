docker-compose down
# git pull origin main
git pull origin user_change_password
# git pull origin stage-v2
docker-compose build
docker-compose up -d redis gateway_service
docker-compose up -d user farm contract_service event_service