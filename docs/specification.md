# Library Management System - Project Specification

**Project Code:** BORK (Book Organization & Rental Kiosk)

## Overview

A simplified library management system designed for university use, allowing students to browse, filter, and rent books. The system does not include administrative interfaces; instead, data is managed through file imports.

## Scope

This is a university project focused on core library functionality without complex administrative features. The system emphasizes user interactions with the book catalog and rental management.

## Features

For a visual representation of how features interact and the user flow, see the [Feature Flow Diagram](./flow_diagram.md).

### F1. Import Functionality

- **Books and Users**: All book inventory and user accounts are created and updated via import from CSV/JSON files
- **No Admin Interface**: There is no librarian or administrator role; all data management happens through file imports
- **Batch Updates**: Changes to book stock or user information are applied by importing updated files

### F2. Authentication

- **Login**: Users authenticate with their credentials to access the system
- **Logout**: Users can securely log out of their session

### F3. Book Browsing

- **View Available Books**: Users can see a complete list of books available in the library
- **Filter Options**: Books can be filtered by:
  - Book name/title
  - Author
  - Category/genre

### F4. Rental Display

- **Current Rentals**: At the top of the book list, users always see:
  - Books they currently have rented
  - The date since when each book has been rented
  - Visual indication of rental status

### F5. Book Rental

- **Shopping Cart Flow**: The rental process follows a shopping cart pattern similar to an online store:
  - Users can add available books to their rental cart
  - Users can review their cart before finalizing the rental
  - Users can remove books from the cart before checkout
  - A "checkout" action finalizes the rental of all books in the cart
- **Rental Limit**: Users can rent a maximum of 3 books at any given time (including books already rented and books in cart)
- **Availability Check**: Books can only be added to cart if they are currently available (not rented by another user)
- **Rental Period**: Books can be rented for a maximum of 1 month (30 days)
- **Return Process**: Users can return books to make them available again

### F6. Notifications

- **Overdue Alerts**: When users log in, they are notified if any of their rented books are past the return date
- **Rental Status**: Clear indication of rental deadlines and overdue status

## Non-Functional Requirements

- **Usability**: Simple and intuitive interface for students
- **Data Integrity**: Proper validation of imported data
- **Availability Tracking**: Accurate real-time tracking of book availability
- **Date Management**: Reliable tracking of rental periods and overdue notifications
- **Testing**: Unit tests must be created to validate core functionality and business logic
- **Optional Enhancements**:
  - Run tests in a CI/CD pipeline for automated validation
  - Containerization using Docker/Docker Compose for simplified deployment and development environment setup

## Technical Constraints

- Simple architecture suitable for a university project
- File-based data import (CSV/JSON) instead of complex admin panels
- Focus on core functionality: authentication, browsing, filtering, and rental management
- Clear business rules enforcement (3-book limit, 1-month rental period)
- **UML Modeling**: A [UML (Unified Modeling Language)](https://www.omg.org/spec/UML/) model must be created before implementation
  - The model must be created using a tool that has a [UML metamodel](https://www.omg.org/spec/UML/2.5.1/About-UML/) behind it
  - The model should include class diagrams, sequence diagrams, and other relevant UML diagrams
  - Recommended tools: Enterprise Architect, Visual Paradigm, StarUML, or similar UML-compliant tools
- **ORM Usage**: Implementation must use an Object-Relational Mapping (ORM) framework for database interactions
- **Layered Architecture**: Application must be structured with separate layers:
  - Presentation layer (UI/interface)
  - Business layer (business logic and rules)
  - Data layer (database access and persistence)
- **OOP Language**: Implementation must use an Object-Oriented Programming language
