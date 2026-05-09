# BORK - Conceptual Class Diagram

This document presents the conceptual model for the BORK (Book Organization & Rental Kiosk) system using a UML class diagram. The model represents the core domain entities, their attributes, relationships, and multiplicities.

## Class Diagram

```mermaid
classDiagram
    class User {
        -int userId
        -string username
        -string passwordHash
        -string email
        -string firstName
        -string lastName
        -DateTime createdAt
        -boolean isLocked
        -int failedLoginAttempts
        +login(username, password) boolean
        +logout() void
        +getCurrentRentals() List~Rental~
        +getOverdueRentals() List~Rental~
        +isRentalLimitReached() boolean
    }

    class Book {
        -int bookId
        -string title
        -string author
        -string isbn
        -Category category
        -boolean isAvailable
        -DateTime addedAt
        +getAvailabilityStatus() boolean
        +markAsRented() void
        +markAsAvailable() void
    }

    class Category {
        -int categoryId
        -string name
        -string description
        +getBooks() List~Book~
    }

    class Rental {
        -int rentalId
        -int userId
        -int bookId
        -DateTime rentalDate
        -DateTime dueDate
        -DateTime returnDate
        -RentalStatus status
        +calculateDueDate() DateTime
        +isOverdue() boolean
        +getDaysOverdue() int
        +returnBook() void
    }

    class RentalCart {
        -int cartId
        -int userId
        -DateTime createdAt
        -DateTime updatedAt
        +addBook(book) boolean
        +removeBook(book) void
        +getCartItems() List~CartItem~
        +getTotalItemCount() int
        +clear() void
        +checkout() List~Rental~
        +validateCartLimit() boolean
    }

    class CartItem {
        -int cartItemId
        -int cartId
        -int bookId
        -DateTime addedAt
    }

    class Session {
        -string sessionId
        -int userId
        -DateTime createdAt
        -DateTime lastAccessedAt
        -DateTime expiresAt
        -boolean isActive
        +isExpired() boolean
        +refresh() void
        +invalidate() void
    }

    class ImportLog {
        -int logId
        -ImportType type
        -string filename
        -DateTime importDate
        -ImportStatus status
        -int recordsProcessed
        -int recordsFailed
        -string errorDetails
        +logSuccess(records) void
        +logError(error) void
    }

    class RentalStatus {
        <<enumeration>>
        ACTIVE
        RETURNED
        OVERDUE
    }

    class ImportType {
        <<enumeration>>
        BOOKS
        USERS
    }

    class ImportStatus {
        <<enumeration>>
        SUCCESS
        PARTIAL
        FAILED
    }

    %% Relationships
    User "1" --> "0..*" Rental : has
    User "1" --> "0..1" RentalCart : owns
    User "1" --> "0..*" Session : has

    Book "1" --> "0..*" Rental : rented in
    Book "*" --> "1" Category : belongs to
    Book "1" --> "0..*" CartItem : added to

    Rental "*" --> "1" RentalStatus : has

    RentalCart "1" --> "0..*" CartItem : contains

    CartItem "*" --> "1" Book : references

    ImportLog "*" --> "1" ImportType : has
    ImportLog "*" --> "1" ImportStatus : has

    %% Notes
    note for User "Maximum 3 active rentals per user\nRental period: 30 days\nAccount locks after 3 failed login attempts"
    note for Rental "Due date = rental date + 30 days\nOverdue if current date > due date"
    note for RentalCart "Total items (cart + active rentals) ≤ 3"
```

## Entity Descriptions

### User

Represents a student who can browse books, manage rentals, and receive notifications. Users are authenticated and subject to rental limits.

**Key Business Rules:**

- Maximum 3 active rentals at any time
- Account locks after 3 failed login attempts
- Session timeout after 30 minutes of inactivity

### Book

Represents a physical book in the library inventory. Books can be available or rented, and belong to a specific category.

**Key Business Rules:**

- A book can only be rented by one user at a time
- Availability status is updated in real-time
- Books are managed through CSV/JSON imports

### Category

Represents a classification/genre for books (e.g., Fiction, Science, History). Helps users filter and browse books.

### Rental

Represents an active or historical rental transaction. Tracks when a book was rented, when it's due, and when it was returned.

**Key Business Rules:**

- Rental period is exactly 30 days
- Due date is automatically calculated as rental date + 30 days
- Status changes from ACTIVE to RETURNED when book is returned
- Status is OVERDUE if current date > due date and not yet returned

### RentalCart

Represents a user's shopping cart for books they intend to rent. Follows an e-commerce pattern where users can add/remove books before checkout.

**Key Business Rules:**

- Each user has at most one active cart
- Total items in cart + active rentals cannot exceed 3
- Cart is cleared after successful checkout
- Books in cart must still be available at checkout

### CartItem

Represents an individual book added to a rental cart. Links the cart to specific books.

### Session

Represents an authenticated user session. Manages login state and session expiration.

**Key Business Rules:**

- Sessions expire after 30 minutes of inactivity
- Only one active session per user (optional constraint)
- Sessions are invalidated on logout

### ImportLog

Tracks the history of CSV/JSON file imports for books and users. Provides audit trail and error tracking.

**Key Business Rules:**

- Records both successful and failed imports
- Maintains error details for troubleshooting
- Supports partial imports (some records succeed, some fail)

## Enumerations

### RentalStatus

- **ACTIVE**: Book is currently rented and not yet returned
- **RETURNED**: Book has been returned
- **OVERDUE**: Book is past due date and not yet returned

### ImportType

- **BOOKS**: Import of book inventory data
- **USERS**: Import of user account data

### ImportStatus

- **SUCCESS**: All records imported successfully
- **PARTIAL**: Some records imported, some failed
- **FAILED**: Import completely failed

## Key Relationships

### User ↔ Rental (1 to 0..\*)

- A user can have zero or more rentals (current and historical)
- Each rental belongs to exactly one user
- **Constraint**: Maximum 3 ACTIVE rentals per user

### User ↔ RentalCart (1 to 0..1)

- A user can have at most one rental cart
- A cart belongs to exactly one user

### Book ↔ Category (\* to 1)

- Each book belongs to exactly one category
- A category can have zero or more books

### Book ↔ Rental (1 to 0..\*)

- A book can have zero or more rental records (historical)
- Each rental is for exactly one book
- **Constraint**: Only one ACTIVE rental per book at a time

### RentalCart ↔ CartItem (1 to 0..\*)

- A cart contains zero or more items
- Each cart item belongs to exactly one cart

### CartItem ↔ Book (\* to 1)

- Each cart item references exactly one book
- A book can be in zero or more carts (but typically just one)

## Architectural Layers

This conceptual model will be implemented across three layers:

### Presentation Layer

- User interface components
- Session management
- Input validation and sanitization

### Business Layer

- Business logic for rental rules (3-book limit, 30-day period)
- Overdue calculation
- Cart management
- Authentication and authorization
- Import processing

### Data Layer

- ORM entities mapping to database tables
- Repository pattern for data access
- Transaction management
- Data validation

## Notes

- All date/time fields use UTC timezone
- Password storage uses bcrypt hashing (minimum cost factor 12)
- The model enforces referential integrity through foreign key relationships
- Audit fields (createdAt, updatedAt) support tracking and debugging
- The design supports future extensions (e.g., book reservations, notifications)
