#!/bin/bash

set -e  # stop script if anything fails

echo "🛠️  Building image..."
make docker-build-prod

echo "📦 Saving Docker image..."
docker save -o scan-fi.tar scan-fi:latest

echo "📂 Creating remote folder and uploading files..."
ssh ubuntu@urban-things.com "mkdir -p ~/scan-fi"
scp scan-fi.tar .env ubuntu@urban-things.com:~/scan-fi/

echo "🚀 Deploying on remote server..."
ssh ubuntu@urban-things.com << 'EOF'
  cd ~/scan-fi
  sudo docker load -i scan-fi.tar
  sudo docker stop scan-fi-app || true
  sudo docker rm scan-fi-app || true
  sudo docker run -d \
    --name scan-fi-app \
    --env-file .env \
    -e ENV=production \
    -p 8000:8000 \
    scan-fi:latest
EOF

echo "✅ Deployment complete. App should be live at: http://scan-fi.urban-things.com"
