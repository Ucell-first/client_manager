# Client Manager

🚀 Technologies

**Backend:** Go 1.24.4
**Database:** PostgreSQL 13
**Frontend:** HTML Templates + Bulma CSS Framework
**Deployment:** Docker & Docker Compose


1. Run with Docker Compose:

```docker-compose up --build```

This command will:
1. Start the PostgreSQL container
2. Wait until the database is ready (health check)
3. Run the migration service (creates tables)
4. Wait for migrations to complete
5. Build and run the Go application
6. Open the app at http://localhost:8080

Method 2: Local Development
1. Run the application:
```go run main.go```
Then open http://localhost:8080 in your browser.

🌐 Usage

Main pages:

Home Page: http://localhost:8080 or http://localhost:8080/users
Displays all users
Button to add new user
Add New User: http://localhost:8080/user/new
MSISDN (phone number) – required
Name – required
Status (active/inactive) – optional
View User: http://localhost:8080/user/view?id=USER_ID
View detailed user information
Edit User: http://localhost:8080/user/edit?id=USER_ID
Update existing user data
Delete User: http://localhost:8080/user/delete?id=USER_ID
Completely remove a user (with confirmation)
