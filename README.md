# **BookApp - GinGonic & Gorm**

A RESTful API built using the Gin framework in Go, designed to manage users and books. The application allows users to sign up, log in, and perform CRUD operations on books. JWT-based authentication is used to secure sensitive routes.

## Features

- **User Authentication**:  
  Users can sign up and log in.  
  Authentication is handled using JWT tokens.

- **Book Management**:  
  Users can create, update, and delete books.  
  Books are associated with the user who created them.  
  Users can list books by title, author, or ID.

- **Database**:  
  PostgreSQL is used as the database, managed using Docker and Docker Compose.  
  GORM is used for ORM (Object-Relational Mapping).

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/hurairaz/BookApp-GinGonic-Gorm.git
cd BookApp-GinGonic-Gorm
```

### 2. Set Up JWT Secret Key

Create a `.env` file in the root directory and add your JWT secret key:

```bash
JWT_SECRET_KEY=your-secret-key
```

### 3. Database Setup

The PostgreSQL database is configured via Docker Compose. It runs on port `5433` and uses the following credentials:

- `POSTGRES_USER=myuser`
- `POSTGRES_PASSWORD=mypassword`
- `POSTGRES_DB=bookappdb`

Run the following command to start the database:

```bash
docker-compose up -d
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Application

```bash
go run main.go
```

The server will be available at `http://localhost:8080`.

### 6. Stop the Application

To stop and remove the database container, run:

```bash
docker-compose down
```

## Project Structure

```
├── auth             
│   └── jwt.go
├── config            
│   └── db.go
├── controllers       
│   ├── book.go
│   └── user.go
├── models            
│   ├── book.go
│   └── user.go
├── services          
│   ├── book.go
│   └── user.go
├── docker-compose.yml 
├── go.mod             
├── main.go            
└── .gitignore         
```

## Technologies Used

- **Go**: Core programming language
- **Gin**: Web framework for building APIs
- **GORM**: Object Relational Mapper (ORM) for database management
- **PostgreSQL**: Relational database management system
- **Docker**: Containerization for the PostgreSQL database

