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

DATABASE_URL=postgresql://<your-database-connection-url>

Make sure your PostgreSQL server is running locally, or modify the settings to point to a remote PostgreSQL instance.

3. Run the Application Locally
Once everything is set up, you can run the application locally by using the following command:


go run main.go
This will start the server on http://localhost:8080.

4. Access Swagger Documentation
After the application starts, navigate to http://localhost:8080/swagger in your browser to view the Swagger UI, where you can test the CRUD APIs and view detailed API documentation.

5.Testing the API
Swagger UI to test the API endpoints.
Available Endpoints
GET /api/v1/blog-post: Get all blog posts.
POST /api/v1/blog-post: Create a new blog post.
GET /api/v1/blog-post/{id}: Get a specific blog post by ID.
DELETE /api/v1/blog-post/{id}: Delete a blog post by ID.
PATCH /api/v1/blog-post/{id}: Update a blog post by ID.
You can interact with these endpoints through Swagger UI at https://blog-crud-dmua.onrender.com/swagger/index.html#/.

License
This project is licensed under the MIT License - see the LICENSE file for details.

### Key Sections in the `README.md`:

1. **Installation Instructions**:
   - Step-by-step guide to clone, set up, and run the app locally.
2. **Database Configuration**:
   - Steps to set up PostgreSQL and configure the `.env` file.
3. **Testing the API**:
   - Basic guide on interacting with the API using Swagger UI.
  
Make sure to replace the placeholder in the repository URL (`https://github.com/your-username/crud-app.git`) with your actual GitHub username or repository link.

Let me know if you need further adjustments!