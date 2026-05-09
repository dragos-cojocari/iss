# BORK API Guide

This guide provides examples for interacting with the BORK API using curl.

## 📖 Interactive Documentation (Swagger UI)

**The easiest way to explore and test the API is through the interactive Swagger UI:**

**URL:** http://localhost:8080/swagger-ui.html

### Features

- 🔍 **Browse all endpoints** - See all available API operations organized by category
- 🧪 **Test in browser** - Execute requests directly from the UI
- 📝 **View schemas** - See request/response models with examples
- 🔐 **Authenticate** - Login and test protected endpoints with session cookies
- 📥 **Export spec** - Download OpenAPI JSON/YAML for client generation

### How to Use Swagger UI

1. **Start the backend**: `make dev-backend`
2. **Open Swagger UI**: http://localhost:8080/swagger-ui.html
3. **Login**: Use the `POST /api/auth/login` endpoint with test credentials
4. **Test endpoints**: Click "Try it out" on any endpoint to test it

---

## Table of Contents

- [Interactive Documentation](#interactive-documentation-swagger-ui)
- [Authentication](#authentication)
- [Books](#books)
- [Users](#users)
- [Session Management](#session-management)
- [Error Handling](#error-handling)

---

## Authentication

### Login

Authenticate with username and password to receive a session cookie.

**Endpoint:** `POST /api/auth/login`

**Request:**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"student1","password":"Test123!"}' \
  -c cookies.txt
```

**Response (200 OK):**

```json
{
  "sessionId": "217b8d51-4551-45b8-8cd4-4db6269b4c30",
  "user": {
    "userId": "aefce4e9-2eee-4119-a237-6de7b2935b4b",
    "username": "student1",
    "email": "student1@university.edu",
    "firstName": "Alice",
    "lastName": "Johnson"
  }
}
```

**Error Responses:**

- `401 Unauthorized` - Invalid credentials
- `403 Forbidden` - Account locked (3 failed attempts)

---

### Get Current User

Get information about the currently authenticated user.

**Endpoint:** `GET /api/auth/me`

**Request:**

```bash
curl http://localhost:8080/api/auth/me \
  -b cookies.txt
```

**Response (200 OK):**

```json
{
  "userId": "aefce4e9-2eee-4119-a237-6de7b2935b4b",
  "username": "student1",
  "email": "student1@university.edu",
  "firstName": "Alice",
  "lastName": "Johnson"
}
```

---

### Logout

Invalidate the current session and clear the session cookie.

**Endpoint:** `POST /api/auth/logout`

**Request:**

```bash
curl -X POST http://localhost:8080/api/auth/logout \
  -b cookies.txt
```

**Response (200 OK):**

```json
{
  "message": "Logged out successfully"
}
```

---

## Books

All book endpoints require authentication (session cookie).

### Get All Books

Retrieve all books in the library.

**Endpoint:** `GET /api/books`

**Request:**

```bash
curl http://localhost:8080/api/books \
  -b cookies.txt
```

**Response (200 OK):**

```json
[
  {
    "bookId": "09c8a648-e7da-4552-9730-cedf1b4b14c3",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "isbn": "9780743273565",
    "category": {
      "categoryId": "92034bc4-2f97-4533-9ac6-4b0c4bfcde9f",
      "name": "Fiction",
      "description": "Literary works of imaginative narration"
    },
    "isAvailable": true,
    "addedAt": "2026-05-09T16:51:28.02391"
  }
]
```

---

### Get Available Books

Retrieve only books that are currently available for rental.

**Endpoint:** `GET /api/books/available`

**Request:**

```bash
curl http://localhost:8080/api/books/available \
  -b cookies.txt
```

**Response:** Same format as "Get All Books", but filtered to `isAvailable: true`

---

### Get Book by ID

Retrieve a specific book by its UUID.

**Endpoint:** `GET /api/books/{id}`

**Request:**

```bash
curl http://localhost:8080/api/books/09c8a648-e7da-4552-9730-cedf1b4b14c3 \
  -b cookies.txt
```

**Response (200 OK):**

```json
{
  "bookId": "09c8a648-e7da-4552-9730-cedf1b4b14c3",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "isbn": "9780743273565",
  "category": {
    "categoryId": "92034bc4-2f97-4533-9ac6-4b0c4bfcde9f",
    "name": "Fiction",
    "description": "Literary works of imaginative narration"
  },
  "isAvailable": true,
  "addedAt": "2026-05-09T16:51:28.02391"
}
```

**Error Response:**

- `404 Not Found` - Book does not exist

---

### Search Books

Search for books by title or author (case-insensitive, partial match).

**Endpoint:** `GET /api/books/search?q={searchTerm}`

**Request:**

```bash
# Search for books with "clean" in title or author
curl "http://localhost:8080/api/books/search?q=clean" \
  -b cookies.txt

# Search for books by author
curl "http://localhost:8080/api/books/search?q=fitzgerald" \
  -b cookies.txt
```

**Response:** Same format as "Get All Books", filtered by search term

---

## Users

User endpoints require authentication. These are primarily for testing and administrative purposes.

### Get All Users

**Endpoint:** `GET /api/users`

**Request:**

```bash
curl http://localhost:8080/api/users \
  -b cookies.txt
```

**Response (200 OK):**

```json
[
  {
    "userId": "aefce4e9-2eee-4119-a237-6de7b2935b4b",
    "username": "student1",
    "email": "student1@university.edu",
    "firstName": "Alice",
    "lastName": "Johnson",
    "createdAt": "2026-05-09T16:51:28.019589",
    "isLocked": false,
    "failedLoginAttempts": 0
  }
]
```

---

### Get User by ID

**Endpoint:** `GET /api/users/{id}`

**Request:**

```bash
curl http://localhost:8080/api/users/aefce4e9-2eee-4119-a237-6de7b2935b4b \
  -b cookies.txt
```

---

### Get User by Username

**Endpoint:** `GET /api/users/username/{username}`

**Request:**

```bash
curl http://localhost:8080/api/users/username/student1 \
  -b cookies.txt
```

---

## Session Management

### Session Cookie

The session is managed via an HTTP-only cookie named `BORK_SESSION`. The cookie:

- Contains the session UUID
- Expires after 30 minutes of inactivity
- Is automatically refreshed on each authenticated request
- Is cleared on logout

### Using Sessions with curl

**Save session cookie:**

```bash
curl ... -c cookies.txt
```

**Use saved session cookie:**

```bash
curl ... -b cookies.txt
```

**View cookie contents:**

```bash
cat cookies.txt
```

---

## Error Handling

All errors return a consistent JSON format:

```json
{
  "timestamp": "2026-05-09T19:47:58.455544",
  "status": 401,
  "error": "Unauthorized",
  "message": "Invalid username or password"
}
```

### Common HTTP Status Codes

- **200 OK** - Request succeeded
- **401 Unauthorized** - Authentication required or invalid credentials
- **403 Forbidden** - Account locked or insufficient permissions
- **404 Not Found** - Resource does not exist
- **500 Internal Server Error** - Server error

### Authentication Errors

**Invalid Credentials (401):**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"student1","password":"wrong"}'
```

**Account Locked (403):**
After 3 failed login attempts, the account is automatically locked.

**Session Expired (401):**

```bash
# Try to access protected endpoint without session
curl http://localhost:8080/api/books
```

Response:

```json
{
  "timestamp": "2026-05-09T19:51:53.513874",
  "status": 401,
  "error": "Unauthorized",
  "message": "No session found"
}
```

---

## Complete Workflow Example

Here's a complete example of authenticating and accessing protected resources:

```bash
# 1. Login and save session cookie
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"student1","password":"Test123!"}' \
  -c cookies.txt

# 2. Get current user info
curl http://localhost:8080/api/auth/me -b cookies.txt

# 3. Browse available books
curl http://localhost:8080/api/books/available -b cookies.txt

# 4. Search for a specific book
curl "http://localhost:8080/api/books/search?q=gatsby" -b cookies.txt

# 5. Get book details
curl http://localhost:8080/api/books/{bookId} -b cookies.txt

# 6. Logout
curl -X POST http://localhost:8080/api/auth/logout -b cookies.txt

# 7. Verify session is invalid (should return 401)
curl http://localhost:8080/api/books -b cookies.txt
```

---

## Test Credentials

All test users have the password: **`Test123!`**

| Username | Email                   | Name           |
| -------- | ----------------------- | -------------- |
| student1 | student1@university.edu | Alice Johnson  |
| student2 | student2@university.edu | Bob Smith      |
| student3 | student3@university.edu | Carol Williams |
| jdoe     | john.doe@university.edu | John Doe       |
| testuser | test@university.edu     | Test User      |

---

## Tips

1. **Pretty Print JSON:** Use `jq` for readable output

   ```bash
   curl http://localhost:8080/api/books -b cookies.txt | jq
   ```

2. **Verbose Output:** Add `-v` flag to see headers and debug info

   ```bash
   curl -v http://localhost:8080/api/auth/login ...
   ```

3. **Silent Mode:** Use `-s` to suppress progress output

   ```bash
   curl -s http://localhost:8080/api/books -b cookies.txt
   ```

4. **Save Response:** Redirect output to a file
   ```bash
   curl http://localhost:8080/api/books -b cookies.txt > books.json
   ```

---

## Security Notes

- **HTTPS in Production:** Always use HTTPS in production to protect session cookies
- **Cookie Security:** Session cookies are HttpOnly (not accessible via JavaScript)
- **Session Expiration:** Sessions expire after 30 minutes of inactivity
- **Account Locking:** Accounts lock after 3 failed login attempts
- **Password Hashing:** Passwords are hashed with BCrypt (cost factor 12)

---

For more information, see:

- [Quick Start Guide](./quickstart.md)
- [Development Guide](./development.md)
- [Backend README](../backend/README.md)
- [Database README](../database/README.md)
