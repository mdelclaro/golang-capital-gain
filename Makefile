# Install dependencies
install:
	go install gotest.tools/gotestsum@latest && go mod tidy
	
# Test the application
test: 
	gotestsum --format pkgname ./...

# Run the application
run:
	go run cmd/main.go < $(case)
	
# Build the application with Docker
docker-build:
	docker build -t golang-capital-gain . 

# Run the application with Docker
docker-run:
	docker run -i golang-capital-gain < $(case)