# MediCity Backend

MediCity is a Hospital Management and Online Consultation Platform designed to streamline the interaction between patients, doctors, and administrators. This backend provides secure RESTful APIs for authentication, appointment management, medical records, doctor verification, and more.

---

## Features

### Authentication & Authorization
- JWT-based authentication
- Role-based access control
- Secure password hashing
- User login and registration

### Admin
- Manage patients
- Verify doctor profile and grant them access to the system
- Block/Unblock users(both patients and doctors)
- Manage departments
- View dashboard statistics
- manage payment to the doctors

### Doctor
- Manage profile
- Add qualifications
- Manage available time slots
- View appointments
- cancel appoinments
- consult patients through google meet
- Issue prescriptions
- check patients profile, vital data, uploaded medical reports, previous prescriptions
- check wallet and request for withdrawal
- view feedback from patients
- send message to admin
- check notifications

### Patient
- Register and manage profile
- search for doctors
- view doctor profile
- Book, cancel appointments
- View appointment history
- Access prescriptions
- View, upload medical reports
- consult doctor through google meet
- give and view feedback and rating for doctors

### Appointment Management
- Book appointments
- Cancel appointments
- Manage appointment status

---

## Tech Stack

- Go (Golang)
- Gin Web Framework
- GORM
- PostgreSQL
- JWT Authentication
- Clean Architecture
- Swagger / OpenAPI

---

## Project Structure

```
mediCityHealthCare/
│
├── cmd/
│   └── main.go
│
├── config/
│
├── internal/
│   ├── domain/
│   │     ├── entity/
│   │     └── repository/
│   │
│   ├── usecase/
│   │
│   ├── handler/
│   │
│   ├── infrastructure/
│   │     ├── database/
│   │     ├── repository/
│   │     └── jwt/
│   │
│   ├── middleware/
│   └── routes/
│
├── pkg/
├── .env
├── go.mod
└── README.md
```

---

## Prerequisites

Before running the project, ensure you have installed:

- Go (version 1.xx or later)
- PostgreSQL
- Git

---

## Installation

Clone the repository

```bash
git clone https://github.com/<your-username>/medicity-backend.git
```

Navigate to the project

```bash
cd medicity-backend
```

Install dependencies

```bash
go mod tidy
```

Create a `.env` file and configure your environment variables.

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=medicity
JWT_SECRET=your_secret_key
SERVER_PORT=8080
```

Run the application

```bash
go run cmd/main.go
```

---

## API Documentation

Swagger/OpenAPI documentation is available after running the application.

```
http://localhost:8080/swagger/index.html
```

---

## Authentication

Protected endpoints require a JWT access token.

```
Authorization: Bearer <access_token>
```

---

## User Roles

- Admin
- Doctor
- Patient

---

## Database

Database: PostgreSQL

ORM: GORM

The project follows a normalized relational database design.

---

## Architecture

The project follows **Clean Architecture**, separating responsibilities into:

- Domain
- Routes
- Repository
- Handler
- Middleware
- usecase

This improves maintainability, scalability, and testability.

---

## API Modules

- Authentication
- Users
- Doctors
- Patients
- Departments
- Appointments
- Medical Records
- Notifications

---

## Testing

API testing can be performed using:

- Swagger UI

---

## Future Enhancements

- Email notifications
- SMS notifications
- Analytics dashboard

---

## Author

**Rinsiya K M**

Backend Developer (Go)

---

## License

This project is developed for educational purposes as part of a backend development training program.
