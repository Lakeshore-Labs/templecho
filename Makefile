# Makefile for templecho project

# Variables
BUILD_DIR := build
GO := go
GOFLAGS := -v

# Help target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  templ-generate    - Generate templ files"
	@echo "  build-demo        - Build demo application"
	@echo "  dev-demo          - Start demo with live reload (tailwind + templ + air)"
	@echo "  demo-tailwind-clean - Clean build demo CSS"
	@echo "  demo-tailwind-watch - Watch demo CSS changes"
	@echo "  demo-templ        - Watch templ files for demo"
	@echo "  demo-air          - Run demo with air live reload"
	@echo "  clean             - Remove build artifacts"

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Generate Templ templates
.PHONY: templ-generate
templ-generate:
	@echo "Generating Templ templates..."
	@if ! command -v templ >/dev/null 2>&1; then \
		echo "Templ not found. Installing templ..."; \
		go install github.com/a-h/templ/cmd/templ@latest; \
	fi
	templ generate

# Build demo web app
.PHONY: build-demo
build-demo: templ-generate | $(BUILD_DIR)
	@echo "Building demo (CSS will be loaded from Tailwind CDN)..."
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/demo ./cmd/demo

# Development server for demo with live reload
.PHONY: dev-demo
dev-demo:
	@echo "🎨 Starting demo development server with live reload..."
	@echo "📦 Server will run on http://localhost:8081"
	@$(MAKE) demo-tailwind-clean
	@$(MAKE) -j3 demo-tailwind-watch demo-templ demo-air

# Watch templ files for demo
.PHONY: demo-templ
demo-templ:
	templ generate --watch

# Air server for demo
.PHONY: demo-air
demo-air:
	PORT=8081 air -c cmd/demo/.air.toml

# Build demo CSS with tailwind (clean output)
.PHONY: demo-tailwind-clean
demo-tailwind-clean:
	./tailwindcss -i ./cmd/demo/static/css/src/input.css -o ./cmd/demo/static/css/output.css --minify

# Watch demo CSS with tailwind
.PHONY: demo-tailwind-watch
demo-tailwind-watch:
	./tailwindcss -i ./cmd/demo/static/css/src/input.css -o ./cmd/demo/static/css/output.css --watch

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf tmp
	rm -f cmd/demo/static/css/output.css
