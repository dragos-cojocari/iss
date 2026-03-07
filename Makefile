.PHONY: setup

setup:
	@echo "Setting up development environment..."
	@echo "Installing pre-commit hooks..."
	pre-commit install
	@echo "Setup complete!"
