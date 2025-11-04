run:
	go run cmd/playerone/main.go

register:
	curl -X POST http://localhost:8080/license/register \
		-H "Content-Type: application/json" \
		-d '{"kid": "b1a2c3d4e5f60718293a4b5c6d7e8f90", "key": "0123456789abcdeffedcba9876543210"}'

license:
	curl -X POST http://localhost:8080/license \
		-H "Content-Type: application/json" \
		-d '{"kids": ["b1a2c3d4e5f60718293a4b5c6d7e8f90"], "type": "org.w3.clearkey"}'