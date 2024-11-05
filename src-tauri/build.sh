# Add to existing build.sh
echo "Building Go backend..."
cd go-backend
go build -o ../backend
cd .. 