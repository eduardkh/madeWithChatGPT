# Echo with SwagGo PoC Setup Guide

This guide outlines the steps to set up a Proof of Concept (PoC) using Echo and SwagGo.

## Steps

1. **Create Echo Project**

   - Initialize a new Echo project.

     ```bash
     go mod init Echo_SwagGo
     ```

2. **Install Dependencies**

   - Install Echo and SwagGo:

     ```bash
     go get -u github.com/labstack/echo/v4
     go get -u github.com/swaggo/echo-swagger
     go get -u github.com/swaggo/swag/cmd/swag
     go install github.com/swaggo/swag/cmd/swag@latest # install the CLI
     ```

3. **Define Handlers**

   - Write CRUD operation handler functions in Echo.

4. **Add SwagGo Annotations**

   - Annotate the handlers with SwagGo comments for Swagger documentation.

5. **Generate Swagger Documentation**

   - Run `swag init` to generate Swagger documentation.

6. **Import Generated Docs**

   - Import the generated `docs` in the main Go file.
     ```go
     import _ "Echo_SwagGo/docs"
     ```

7. **Configure Echo Server**

   - Set up Echo routes and the Swagger UI endpoint.

8. **Run the Server**
   - Start the Echo server.
   - Access the Swagger UI at `http://localhost:<port>/swagger/index.html`.
