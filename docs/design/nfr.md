# Non-Functional Requirements

This document outlines the non-functional requirements (NFRs) for the BORK (Book Organization & Rental Kiosk) system. These requirements define the quality attributes, constraints, and operational characteristics of the system.

## Table of Contents

- [Performance Requirements](#performance-requirements)
- [Usability Requirements](#usability-requirements)
- [Reliability Requirements](#reliability-requirements)
- [Security Requirements](#security-requirements)
- [Maintainability Requirements](#maintainability-requirements)
- [Portability Requirements](#portability-requirements)
- [Data Integrity Requirements](#data-integrity-requirements)
- [Testing Requirements](#testing-requirements)
- [Optional Enhancements](#optional-enhancements)

## Performance Requirements

### NFR-1: Response Time

- **Description**: The system shall respond to user actions within acceptable time limits
- **Criteria**:
  - Page load time: ≤ 2 seconds for book list display
  - Search/filter operations: ≤ 1 second
  - Login authentication: ≤ 1 second
  - Cart operations (add/remove): ≤ 500ms
  - Checkout process: ≤ 3 seconds
- **Priority**: High
- **Rationale**: Users expect responsive interfaces for a smooth browsing and rental experience

### NFR-2: Concurrent Users

- **Description**: The system shall support multiple concurrent users without performance degradation
- **Criteria**:
  - Support at least 50 concurrent active sessions
  - No significant performance degradation with up to 100 concurrent users
- **Priority**: Medium
- **Rationale**: University library systems need to handle multiple students accessing the system simultaneously

### NFR-3: Database Query Performance

- **Description**: Database queries shall execute efficiently
- **Criteria**:
  - Book catalog queries: ≤ 500ms for up to 10,000 books
  - Rental status checks: ≤ 200ms
  - User authentication queries: ≤ 300ms
- **Priority**: Medium
- **Rationale**: Efficient database access ensures overall system responsiveness

## Usability Requirements

### NFR-4: User Interface Simplicity

- **Description**: The system shall provide a simple and intuitive interface for students
- **Criteria**:
  - New users should be able to browse and rent books without training
  - Maximum 3 clicks to complete common tasks (browse, filter, add to cart)
  - Clear visual feedback for all user actions
  - Consistent navigation across all pages
- **Priority**: High
- **Rationale**: Students need to quickly find and rent books without extensive learning curve

### NFR-5: Accessibility

- **Description**: The system shall be accessible to users with disabilities
- **Criteria**:
  - WCAG 2.1 Level AA compliance for web interfaces
  - Keyboard navigation support for all features
  - Screen reader compatibility
  - Sufficient color contrast (minimum 4.5:1 ratio)
- **Priority**: Medium
- **Rationale**: University systems should be accessible to all students

### NFR-6: Error Messages

- **Description**: The system shall provide clear and helpful error messages
- **Criteria**:
  - Error messages explain what went wrong in plain language
  - Error messages suggest corrective actions
  - No technical jargon or stack traces visible to users
- **Priority**: High
- **Rationale**: Clear error messages reduce user frustration and support burden

## Reliability Requirements

### NFR-7: System Availability

- **Description**: The system shall be available during library operating hours
- **Criteria**:
  - 99% uptime during operating hours (approximately 8 hours downtime per year)
  - Planned maintenance during off-hours only
  - Graceful degradation if components fail
- **Priority**: High
- **Rationale**: Students need reliable access to the library system during operating hours

### NFR-8: Data Backup

- **Description**: The system shall maintain regular backups of all data
- **Criteria**:
  - Daily automated backups of database
  - Backup retention for at least 30 days
  - Backup verification and restoration testing monthly
- **Priority**: High
- **Rationale**: Protect against data loss from hardware failures or errors

### NFR-9: Error Recovery

- **Description**: The system shall recover gracefully from errors
- **Criteria**:
  - Transaction rollback on failed operations
  - Session preservation during temporary network issues
  - Clear recovery instructions for users
- **Priority**: Medium
- **Rationale**: Minimize impact of errors on user experience

## Security Requirements

### NFR-10: Authentication

- **Description**: The system shall securely authenticate users
- **Criteria**:
  - Password minimum length: 8 characters
  - Password complexity requirements (letters, numbers, special characters)
  - Account lockout after 3 failed login attempts (15-minute lockout)
  - Session timeout after 30 minutes of inactivity
- **Priority**: High
- **Rationale**: Protect user accounts from unauthorized access

### NFR-11: Password Storage

- **Description**: The system shall securely store user credentials
- **Criteria**:
  - Passwords hashed using bcrypt or similar strong algorithm
  - Minimum cost factor of 12 for bcrypt
  - No plaintext password storage
- **Priority**: Critical
- **Rationale**: Protect user credentials in case of database breach

### NFR-12: Authorization

- **Description**: The system shall enforce proper authorization
- **Criteria**:
  - Users can only view and modify their own rentals
  - Users cannot access other users' data
  - All privileged operations require authentication
- **Priority**: High
- **Rationale**: Ensure data privacy and prevent unauthorized actions

### NFR-13: Data Transmission Security

- **Description**: The system shall protect data in transit
- **Criteria**:
  - HTTPS/TLS 1.2 or higher for all communications
  - Secure cookie flags (HttpOnly, Secure, SameSite)
  - No sensitive data in URL parameters
- **Priority**: High
- **Rationale**: Protect user data from interception

### NFR-14: Input Validation

- **Description**: The system shall validate all user inputs
- **Criteria**:
  - Server-side validation for all inputs
  - Protection against SQL injection
  - Protection against XSS attacks
  - Sanitization of file upload contents
- **Priority**: Critical
- **Rationale**: Prevent security vulnerabilities from malicious input

## Maintainability Requirements

### NFR-15: Code Quality

- **Description**: The system code shall be maintainable and well-documented
- **Criteria**:
  - Code follows language-specific style guides
  - Functions/methods have clear, single responsibilities
  - Maximum cyclomatic complexity of 10 per function
  - Inline comments for complex logic
- **Priority**: Medium
- **Rationale**: Facilitate future maintenance and enhancements

### NFR-16: Documentation

- **Description**: The system shall have comprehensive documentation
- **Criteria**:
  - API documentation for all public interfaces
  - Database schema documentation
  - Deployment and configuration guides
  - User manual for students
- **Priority**: Medium
- **Rationale**: Enable future developers and administrators to understand the system

### NFR-17: Logging

- **Description**: The system shall maintain comprehensive logs
- **Criteria**:
  - Log all authentication attempts (success and failure)
  - Log all rental transactions
  - Log all errors with stack traces
  - Log rotation to prevent disk space issues
  - Logs retained for at least 90 days
- **Priority**: Medium
- **Rationale**: Support debugging, auditing, and security monitoring

### NFR-18: Modularity

- **Description**: The system shall be organized in modular layers
- **Criteria**:
  - Clear separation between presentation, business, and data layers
  - Loose coupling between modules
  - High cohesion within modules
  - Dependency injection for testability
- **Priority**: High
- **Rationale**: Facilitate testing, maintenance, and future enhancements

## Portability Requirements

### NFR-19: Browser Compatibility

- **Description**: The system shall work across modern web browsers
- **Criteria**:
  - Support for Chrome (last 2 versions)
  - Support for Firefox (last 2 versions)
  - Support for Safari (last 2 versions)
  - Support for Edge (last 2 versions)
- **Priority**: High
- **Rationale**: Users should be able to access the system from their preferred browser

### NFR-20: Database Portability

- **Description**: The system shall use standard SQL and ORM abstractions
- **Criteria**:
  - ORM framework abstracts database-specific features
  - Standard SQL for queries where possible
  - Database migrations for schema changes
- **Priority**: Medium
- **Rationale**: Allow flexibility in database choice and facilitate testing

## Data Integrity Requirements

### NFR-21: Data Validation

- **Description**: The system shall ensure data integrity through validation
- **Criteria**:
  - Validation of imported CSV/JSON data before database insertion
  - Foreign key constraints enforced at database level
  - Check constraints for business rules (e.g., rental limit)
  - Atomic transactions for multi-step operations
- **Priority**: High
- **Rationale**: Prevent invalid data from entering the system

### NFR-22: Rental Tracking Accuracy

- **Description**: The system shall accurately track book availability and rental status
- **Criteria**:
  - Real-time availability updates
  - Accurate rental date and due date calculations
  - Correct overdue status determination
  - No double-booking of books
- **Priority**: Critical
- **Rationale**: Core business requirement for library management

### NFR-23: Data Consistency

- **Description**: The system shall maintain data consistency
- **Criteria**:
  - ACID properties for all database transactions
  - Referential integrity maintained at all times
  - Consistent state after failed operations (rollback)
- **Priority**: High
- **Rationale**: Ensure reliable system state

## Testing Requirements

### NFR-24: Unit Test Coverage

- **Description**: The system shall have comprehensive unit tests
- **Criteria**:
  - Minimum 70% code coverage for business logic layer
  - All critical business rules tested
  - Tests for edge cases and error conditions
  - Automated test execution
- **Priority**: High
- **Rationale**: Ensure code quality and facilitate refactoring

### NFR-25: Integration Testing

- **Description**: The system shall have integration tests
- **Criteria**:
  - Tests for database interactions
  - Tests for authentication flow
  - Tests for complete rental workflows
  - Tests run in isolated test environment
- **Priority**: Medium
- **Rationale**: Verify components work together correctly

### NFR-26: Test Automation

- **Description**: Tests shall be automated and repeatable
- **Criteria**:
  - Tests executable via command line
  - Tests produce consistent results
  - Test data setup and teardown automated
  - Fast test execution (< 5 minutes for full suite)
- **Priority**: Medium
- **Rationale**: Enable continuous verification of system quality

## Optional Enhancements

### NFR-27: CI/CD Pipeline (Optional)

- **Description**: Automated testing and deployment pipeline
- **Criteria**:
  - Automated test execution on code commits
  - Automated code quality checks
  - Automated deployment to staging environment
  - Build status notifications
- **Priority**: Low (Optional)
- **Rationale**: Improve development workflow and code quality

### NFR-28: Containerization (Optional)

- **Description**: Docker/Docker Compose for simplified deployment
- **Criteria**:
  - Dockerfile for application containerization
  - Docker Compose for multi-container orchestration
  - Environment-specific configuration via environment variables
  - Volume mounts for persistent data
- **Priority**: Low (Optional)
- **Rationale**: Simplify deployment and development environment setup

### NFR-29: Performance Monitoring (Optional)

- **Description**: System performance monitoring and alerting
- **Criteria**:
  - Response time monitoring
  - Error rate tracking
  - Resource utilization metrics
  - Alerting for performance degradation
- **Priority**: Low (Optional)
- **Rationale**: Proactive identification of performance issues

## NFR Traceability Matrix

| NFR ID | Category        | Priority | Related Use Cases        |
| ------ | --------------- | -------- | ------------------------ |
| NFR-1  | Performance     | High     | All                      |
| NFR-2  | Performance     | Medium   | All                      |
| NFR-3  | Performance     | Medium   | UC-4, UC-5, UC-6         |
| NFR-4  | Usability       | High     | All                      |
| NFR-5  | Usability       | Medium   | All                      |
| NFR-6  | Usability       | High     | All                      |
| NFR-7  | Reliability     | High     | All                      |
| NFR-8  | Reliability     | High     | All                      |
| NFR-9  | Reliability     | Medium   | All                      |
| NFR-10 | Security        | High     | UC-2, UC-3               |
| NFR-11 | Security        | Critical | UC-2                     |
| NFR-12 | Security        | High     | All Student UCs          |
| NFR-13 | Security        | High     | All                      |
| NFR-14 | Security        | Critical | UC-1, UC-7, UC-10        |
| NFR-15 | Maintainability | Medium   | All                      |
| NFR-16 | Maintainability | Medium   | All                      |
| NFR-17 | Maintainability | Medium   | All                      |
| NFR-18 | Maintainability | High     | All                      |
| NFR-19 | Portability     | High     | All                      |
| NFR-20 | Portability     | Medium   | All                      |
| NFR-21 | Data Integrity  | High     | UC-1, UC-7, UC-10        |
| NFR-22 | Data Integrity  | Critical | UC-6, UC-7, UC-10, UC-11 |
| NFR-23 | Data Integrity  | High     | All                      |
| NFR-24 | Testing         | High     | All                      |
| NFR-25 | Testing         | Medium   | All                      |
| NFR-26 | Testing         | Medium   | All                      |
| NFR-27 | Optional        | Low      | All                      |
| NFR-28 | Optional        | Low      | All                      |
| NFR-29 | Optional        | Low      | All                      |
