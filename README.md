# Blockchain Supply Service
A Go-based REST API that retrieves the total supply of NGL tokens from an external service and returns it as JSON. Configurable via environment variables or config files, it runs an HTTP server with a /getTotalSupply endpoint to fetch the latest token data

### Running the Service in Docker

1. **Run the service using Docker Compose:**
   If Docker and Docker Compose are installed, you can start the service easily. Simply run:
   ```bash
   docker-compose up --build
   ```

2. **Access the API:**
   Once the service is up, you can interact with the API via `curl` commands. Here are a few examples:

   - Basic GET request:
     ```bash
     curl http://localhost:8081/getTotalSupply
     ```

   - Verbose mode to check details:
     ```bash
     curl -v http://localhost:8081/getTotalSupply
     ```

   **Note:** If `8081` doesn't work, it may be in use by your VPN service. Try updating the port to avoid this conflict.

3. **Stopping the service:**
   To stop the service, use:
   ```bash
   docker-compose down
   ```

---

### Running the Application Without Containerization

1. **Check configuration:**
   Ensure the `config.json` file exists in the root directory with the appropriate configuration, such as:
   ```json
   {
     "port": "8081",
     "url": "https://cosmos.entangle.fi/cosmos/bank/v1beta1/supply"
   }
   ```

2. **Build and run the service:**
   To build and run the service manually, use the following commands:
   ```bash
   go build -o ./bin/blockchain-supply-service ./cmd/http/main.go
   ./bin/blockchain-supply-service
   ```

3. **Access the API:**
   Once the service is running, you can make requests to the API using `curl`. Here are some examples:

   - Basic request:
     ```bash
     curl http://localhost:8081/getTotalSupply
     ```

   - For verbose output:
     ```bash
     curl -v http://localhost:8081/getTotalSupply
     ```

---

**Reminder:** If `8081` doesn't work, it may be in use by your VPN service. Try updating the port to avoid this conflict.
