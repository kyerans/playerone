run:
	go run cmd/playerone/main.go

register:
	curl -X POST http://localhost:8080/license/register \
		-H "Content-Type: application/json" \
		-d '{"kid": "b1a2c3d4", "key": "0123456789abcdef"}'

license:
	curl -X POST http://localhost:8080/license \
		-H "Content-Type: application/json" \
		-d '{"kids": ["b1a2c3d4"], "type": "org.w3.clearkey"}'