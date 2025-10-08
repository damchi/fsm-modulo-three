# Makefile for FSM Modulo Three project

# Variables
APP_NAME=fsm-modulo-three
DOC_DIR=doc

# Run the API server
run:
	@echo "Starting $(APP_NAME)..."
	go run ./cmd/app

# Run all tests with verbose output
test:
	@echo "Running tests..."
	go test ./... -v

# Generate Swagger documentation
doc:
	@echo "Generating Swagger API documentation..."
	swag init -g ./cmd/app/main.go -o ./doc
	@echo "Documentation generated in docs/"

# Clean generated files
clean:
	@echo "🧼 Cleaning up..."
	rm -rf $(DOC_DIR)
