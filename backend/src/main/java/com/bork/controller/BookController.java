package com.bork.controller;

import com.bork.model.Book;
import com.bork.repository.BookRepository;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

/**
 * REST controller for book operations
 * Provides endpoints for browsing and searching books
 */
@RestController
@RequestMapping(value = "/api/books", produces = "application/json")
@Tag(name = "Books", description = "Book browsing and search endpoints")
@SecurityRequirement(name = "cookieAuth")
public class BookController {

    @Autowired
    private BookRepository bookRepository;

    /**
     * Get all books
     */
    @Operation(summary = "Get all books", description = "Retrieve all books in the library")
    @ApiResponse(responseCode = "200", description = "Books retrieved successfully")
    @GetMapping
    public ResponseEntity<List<Book>> getAllBooks() {
        List<Book> books = bookRepository.findAll();
        return ResponseEntity.ok(books);
    }

    /**
     * Get all available books
     */
    @Operation(summary = "Get available books", description = "Retrieve only books that are currently available for rental")
    @ApiResponse(responseCode = "200", description = "Available books retrieved successfully")
    @GetMapping("/available")
    public ResponseEntity<List<Book>> getAvailableBooks() {
        List<Book> books = bookRepository.findByIsAvailableTrue();
        return ResponseEntity.ok(books);
    }

    /**
     * Get book by ID
     */
    @Operation(summary = "Get book by ID", description = "Retrieve a specific book by its UUID")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Book found"),
            @ApiResponse(responseCode = "404", description = "Book not found")
    })
    @GetMapping("/{id}")
    public ResponseEntity<Book> getBookById(
            @Parameter(description = "Book UUID") @PathVariable UUID id) {
        return bookRepository.findById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    /**
     * Search books by title or author
     */
    @Operation(summary = "Search books", description = "Search for books by title or author (case-insensitive, partial match)")
    @ApiResponse(responseCode = "200", description = "Search results retrieved")
    @GetMapping("/search")
    public ResponseEntity<List<Book>> searchBooks(
            @Parameter(description = "Search term") @RequestParam String q) {
        List<Book> books = bookRepository.searchByTitleOrAuthor(q);
        return ResponseEntity.ok(books);
    }
}
