serve-docs:
	uv sync
	uv run mkdocs serve
fmt:
	go fmt ./...