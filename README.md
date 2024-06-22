
Assumptions:
- Account id must be a positive integer and unique
- Initial balance must be > 0
- Account balance have limited number of decimals (<12)
- Transfer amount must not be <=0
- Allow self-transfer

APIs:<br/>
Base URL: localhost:8000/api/v1
- POST /accounts
- GET /accounts/:accountId
- POST /transactions

Steps to run:
- Requirements: Make sure you have Docker installed and have started Docker Engine. All the following commands to be run from project root folder.
- Build: `make build`
- Start services: `make run`
- View service logs: `make logs`
- Stop all services: `make stop`
- Run unit tests: `make unit-tests`
