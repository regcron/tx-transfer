Assumptions:
POST /accounts
- Account id must be a positive integer
- Initial balance must be > 0
GET /accounts/:accountId
- 

Steps to run:
- Requirements: Make sure you have Docker installed and have started Docker Engine
- To build project: make build
- To start services: make run
- To tail tx service logs: make logs
- To stop all services: make stop
- To run unit tests: make unit-tests