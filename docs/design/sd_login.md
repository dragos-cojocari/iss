# BORK - Login Sequence Diagram

This document presents the sequence diagram for the login process in the BORK (Book Organization & Rental Kiosk) system, illustrating the interactions between the user, presentation layer, business layer, and data layer.

## Related Use Cases

- **UC-2: Login** - Primary use case
- **UC-12: View Overdue Notifications** - Included in login flow

## Sequence Diagram

```mermaid
sequenceDiagram
    actor Student
    participant UI as Presentation Layer
    participant AuthService as Business Layer<br/>(AuthService)
    participant RentalService as Business Layer<br/>(RentalService)
    participant SessionMgr as Business Layer<br/>(SessionManager)
    participant UserRepo as Data Layer<br/>(UserRepository)
    participant RentalRepo as Data Layer<br/>(RentalRepository)
    participant DB as Database

    Student->>UI: Navigate to login page
    activate UI
    UI-->>Student: Display login form
    deactivate UI

    Student->>UI: Enter username and password
    activate UI
    UI->>AuthService: authenticate(username, password)
    activate AuthService

    AuthService->>UserRepo: findByUsername(username)
    activate UserRepo
    UserRepo->>DB: SELECT * FROM users WHERE username = ?
    activate DB
    DB-->>UserRepo: User record
    deactivate DB
    UserRepo-->>AuthService: User object
    deactivate UserRepo

    alt User not found
        AuthService-->>UI: AuthenticationError: Invalid credentials
        UI-->>Student: Display error message
        Note over Student,UI: Return to login form
    else User account locked
        AuthService->>AuthService: Check isLocked flag
        AuthService-->>UI: AuthenticationError: Account locked
        UI-->>Student: Display account locked message
        Note over Student,UI: Terminate use case
    else Valid user found
        AuthService->>AuthService: validatePassword(password, passwordHash)

        alt Password invalid
            AuthService->>UserRepo: incrementFailedAttempts(userId)
            activate UserRepo
            UserRepo->>DB: UPDATE users SET failedLoginAttempts = failedLoginAttempts + 1
            activate DB
            DB-->>UserRepo: Success
            deactivate DB
            deactivate UserRepo

            AuthService->>AuthService: Check if failedAttempts >= 3

            alt Lock account
                AuthService->>UserRepo: lockAccount(userId)
                activate UserRepo
                UserRepo->>DB: UPDATE users SET isLocked = true
                activate DB
                DB-->>UserRepo: Success
                deactivate DB
                deactivate UserRepo
            end

            AuthService-->>UI: AuthenticationError: Invalid credentials
            UI-->>Student: Display error message
            Note over Student,UI: Return to login form
        else Password valid
            AuthService->>UserRepo: resetFailedAttempts(userId)
            activate UserRepo
            UserRepo->>DB: UPDATE users SET failedLoginAttempts = 0
            activate DB
            DB-->>UserRepo: Success
            deactivate DB
            deactivate UserRepo

            AuthService->>SessionMgr: createSession(userId)
            activate SessionMgr
            SessionMgr->>SessionMgr: Generate sessionId
            SessionMgr->>SessionMgr: Calculate expiresAt (now + 30 min)
            SessionMgr->>DB: INSERT INTO sessions (sessionId, userId, createdAt, expiresAt)
            activate DB
            DB-->>SessionMgr: Success
            deactivate DB
            SessionMgr-->>AuthService: Session object
            deactivate SessionMgr

            AuthService-->>UI: AuthenticationSuccess(sessionId, user)

            UI->>RentalService: checkOverdueRentals(userId)
            activate RentalService
            RentalService->>RentalRepo: getActiveRentals(userId)
            activate RentalRepo
            RentalRepo->>DB: SELECT * FROM rentals WHERE userId = ? AND status = 'ACTIVE'
            activate DB
            DB-->>RentalRepo: Rental records
            deactivate DB
            RentalRepo-->>RentalService: List of Rental objects
            deactivate RentalRepo

            RentalService->>RentalService: Calculate due dates and check overdue
            RentalService->>RentalService: Filter overdue rentals (currentDate > dueDate)
            RentalService-->>UI: List of overdue rentals
            deactivate RentalService

            alt Has overdue rentals
                UI-->>Student: Display overdue notifications
                Student->>UI: Acknowledge notifications
            end

            UI-->>Student: Display main dashboard
            deactivate UI
        end
    end
    deactivate AuthService
```

## Sequence Description

### Normal Flow

1. **Student navigates to login page**

   - Student requests access to the BORK system
   - Presentation layer displays the login form

2. **Student enters credentials**

   - Student inputs username and password
   - Presentation layer sends credentials to AuthService

3. **User lookup**

   - AuthService requests UserRepository to find user by username
   - UserRepository queries database for user record
   - Database returns user data (if found)

4. **Credential validation**

   - AuthService validates the password against stored hash using bcrypt
   - If valid, failed login attempts are reset to 0

5. **Session creation**

   - AuthService requests SessionManager to create new session
   - SessionManager generates unique sessionId
   - Session expiration time is set to 30 minutes from creation
   - Session record is persisted to database

6. **Overdue rental check**

   - Presentation layer requests RentalService to check for overdue rentals
   - RentalService retrieves all active rentals for the user
   - RentalService calculates due dates (rental date + 30 days)
   - RentalService identifies rentals where current date > due date

7. **Display results**
   - If overdue rentals exist, notifications are displayed
   - Student acknowledges notifications
   - Main dashboard is displayed

### Exception Flows

#### E1: Invalid Credentials

- User not found or password doesn't match
- Failed login attempts counter is incremented
- If attempts reach 3, account is locked
- Error message displayed to student

#### E2: Account Locked

- System detects `isLocked` flag is true
- Account locked message displayed
- Use case terminates

## Key Components

### Presentation Layer

- **UI**: Handles user interface, form display, and user interactions

### Business Layer

- **AuthService**: Manages authentication logic, password validation, and account locking
- **RentalService**: Handles rental business logic including overdue calculations
- **SessionManager**: Manages session creation, validation, and expiration

### Data Layer

- **UserRepository**: Data access for user entities
- **RentalRepository**: Data access for rental entities
- **Database**: Persistent storage

## Business Rules Enforced

1. **Password Validation**: Passwords are hashed using bcrypt (minimum cost factor 12)
2. **Account Lockout**: Account locks after 3 consecutive failed login attempts
3. **Session Expiration**: Sessions expire after 30 minutes of inactivity
4. **Overdue Calculation**: Rentals are overdue if current date > (rental date + 30 days)
5. **Automatic Notification**: Overdue rentals are automatically checked and displayed on login

## Security Considerations

- Passwords are never stored in plaintext
- Password validation uses secure hashing (bcrypt)
- Failed login attempts are tracked to prevent brute force attacks
- Sessions have automatic expiration
- Generic error messages prevent username enumeration
