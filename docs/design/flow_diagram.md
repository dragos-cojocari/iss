# BORK Feature Flow Diagram

This diagram illustrates the flow and interactions between features in the BORK (Book Organization & Rental Kiosk) system.

```mermaid
graph TD
    F1[F1: Import Functionality<br/>CSV/JSON Import]
    F2[F2: Authentication<br/>Login/Logout]
    F6[F6: Notifications<br/>Overdue Alerts]
    F3[F3: Book Browsing<br/>View & Filter Books]
    F4[F4: Rental Display<br/>Current Rentals]
    F5a[F5: Add to Cart]
    F5b[F5: Review Cart]
    F5c[F5: Checkout]
    F5d[F5: Return Books]

    F1 -.->|Initial Setup| F2
    F2 -->|Login Success| F6
    F6 -->|Check Overdue| F4
    F4 -->|Display Current| F3
    F3 -->|Browse Available| F5a
    F5a -->|Add Books| F5b
    F5b -->|Review| F5c
    F5b -->|Modify| F5a
    F5c -->|Finalize Rental| F4
    F4 -->|Manage Rentals| F5d
    F5d -->|Return Complete| F4
    F3 -.->|Apply Filters| F3
    F2 -.->|Logout| F2

    style F1 fill:#e1f5ff
    style F2 fill:#fff4e1
    style F3 fill:#e8f5e9
    style F4 fill:#f3e5f5
    style F5a fill:#ffe0b2
    style F5b fill:#ffe0b2
    style F5c fill:#ffe0b2
    style F5d fill:#ffe0b2
    style F6 fill:#ffebee
```

## Legend

- Solid arrows (→) indicate primary user flow
- Dashed arrows (-.→) indicate setup, optional, or cyclic actions
- Colors distinguish different feature categories

## Flow Description

1. **F1 (Import)**: Initial setup of books and users via CSV/JSON import
2. **F2 (Authentication)**: User login entry point
3. **F6 (Notifications)**: Immediate check for overdue books after login
4. **F4 (Rental Display)**: Shows current rentals at the top
5. **F3 (Book Browsing)**: Main interface for viewing and filtering available books
6. **F5 (Book Rental)**: Shopping cart flow with multiple steps:
   - Add books to cart
   - Review cart (can modify by going back to add/remove)
   - Checkout to finalize rental
   - Return books to make them available again
