docker compose -p multiple-docker-compose \
    -f frontend-service/docker-compose.yml \
    -f backend-service/docker-compose.yml \
    --env-file ./.env \
    up --build
