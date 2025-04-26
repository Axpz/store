curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"username":"testuser","password":"testpass123","email":"test@example.com"}'
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"username":"testuser","password":"testpass123","email":"test2@example.com"}'
curl -X GET http://localhost:8080/api/users/VVAvQNyLfHaYgLEIdKvJ0A
curl -X GET http://localhost:8080/api/users/Q7BfOU1WEcVKGp6OILruIQ
curl -X DELETE http://localhost:8080/api/users/VVAvQNyLfHaYgLEIdKvJ0A
curl -X DELETE http://localhost:8080/api/users/Q7BfOU1WEcVKGp6OILruIQ
# curl -X POST http://localhost:8080/api/auth/login -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"testpass123"}'

# curl -X GET http://localhost:8080/api/users/VVAvQNyLfHaYgLEIdKvJ0A
# curl -X DELETE http://localhost:8080/api/users/VVAvQNyLfHaYgLEIdKvJ0A