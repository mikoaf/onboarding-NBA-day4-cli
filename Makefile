# @abarobotics 2025
# Makefile automated build and run

# include .env

# Variables
NAME = project
BUILD_PATH = ./build
MAIN_DIR = ./cmd/app/
MAIN_FILE = main.go

# Target: build
.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BUILD_PATH)/$(NAME) $(MAIN_DIR)*go
	@echo "Build completed. Output at $(BUILD_PATH)/ with name '$(NAME)'"

# Target: run without build
.PHONY: run
run:
	@echo "Running the application without building..."
	@cd $(BUILD_PATH) && go run .$(MAIN_DIR)*go
	@echo "Run completed."

# Target: run prebuild program
.PHONY: start
start:
	@cd $(BUILD_PATH) && ./$(NAME)

.PHONY: build-run
build-run:
	@echo "Building the application..."
	@go build -o $(BUILD_PATH)/$(NAME) $(MAIN_DIR)$(MAIN_FILE)
	@echo "Build completed. Output at $(BUILD_PATH)/$(NAME) with name $(NAME)"
	@echo "Starting application"
	@cd $(BUILD_PATH) && ./$(NAME)
	@echo "Build and run completed"	
	