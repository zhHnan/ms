chcp 65001
cd ms-user
docker build -t ms-user:latest .
cd ..
docker-compose up -dz