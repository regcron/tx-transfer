Base URL: localhost:8000/api/v1

Assumptions:
- Account id must be a positive integer and unique
- Initial balance must be > 0
- Account balance have limited number of decimals (<12)
- Transfer amount must not be <=0
- Allow self-transfer

APIs:
POST /accounts
GET /accounts/:accountId
POST /transactions

Steps to run:
- Requirements: Make sure you have Docker installed and have started Docker Engine
- To build project: make build
- To start services: make run.  localhost:8000/api/v1/accounts or localhost:8000/api/v1/transactions
- To view tx service logs: make logs
- To stop all services: make stop
- To run unit tests: make unit-tests
