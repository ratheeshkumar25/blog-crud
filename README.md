# Blog-CRUD App with Go

A simple Blog-CRUD application built with Go, designed to handle basic Create, Read, Update, and Delete operations. This app is also integrated with Swagger for API documentation.

## Features

- **Create**: Add new records.
- **Read**: Fetch existing records.
- **Update**: Modify existing records.
- **Delete**: Remove records.
- **Swagger Documentation**: Access API documentation via Swagger UI.

## Prerequisites

- Go (go version go1.23.1)
- Docker (for containerization)
- PostgreSQL (for data storage)
- Swagger (for API documentation)

## Clone the Repository

## To get started, clone the repository to your local machine:

```bash
git clone https://github.com/your-username/crud-app.git
cd crud-app
Setup and Installation
1. Install Dependencies
Ensure you have Go installed on your machine. If not, follow the installation guide for Go here.

Run the following command to install the required Go packages:

go mod tidy
2. Configure Database
You need to set up a PostgreSQL database for the application. Create a new PostgreSQL database named crud_app. Update the database connection settings in the .env file with your database credentials:


DB_HOST=localhost
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=crud_app
Make sure your PostgreSQL server is running locally, or modify the settings to point to a remote PostgreSQL instance.

3. Run the Application Locally
Once everything is set up, you can run the application locally by using the following command:


go run main.go
This will start the server on http://localhost:8080.

4. Access Swagger Documentation
After the application starts, navigate to http://localhost:8080/swagger in your browser to view the Swagger UI, where you can test the CRUD APIs and view detailed API documentation.

Docker Setup (Optional)
If you prefer to run the application with Docker, follow these steps:

Build the Docker image:

First, ensure that your Dockerfile is set up correctly. Then, build the Docker image using the following command:

docker build -t crud-app .
Run the container:

Run the Docker container with the following command, ensuring you pass the .env file for database configuration:

docker run -p 8080:8080 --env-file .env crud-app
This will run the app inside a Docker container and expose the app on http://localhost:8080.

Testing the API
Once the application is up and running, you can use the Swagger UI to test the API endpoints.

Available Endpoints
POST /posts: Create a new blog post.
GET /posts: Get all blog posts.
GET /posts/{id}: Get a specific blog post by ID.
PUT /posts/{id}: Update a specific blog post by ID.
DELETE /posts/{id}: Delete a specific blog post by ID.
You can interact with these endpoints through Swagger UI at http://localhost:8080/swagger.

License
This project is licensed under the MIT License - see the LICENSE file for details.

### Key Sections in the `README.md`:

1. **Installation Instructions**:
   - Step-by-step guide to clone, set up, and run the app locally.
2. **Database Configuration**:
   - Steps to set up PostgreSQL and configure the `.env` file.
3. **Docker Setup**:
   - Optional Docker setup for running the application in a container.
4. **Testing the API**:
   - Basic guide on interacting with the API using Swagger UI.
  
Make sure to replace the placeholder in the repository URL (`https://github.com/your-username/crud-app.git`) with your actual GitHub username or repository link.

Let me know if you need further adjustments!