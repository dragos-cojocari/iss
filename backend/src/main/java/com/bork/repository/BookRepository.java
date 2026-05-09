package com.bork.repository;

import com.bork.model.Book;
import com.bork.model.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository interface for Book entity
 * Provides CRUD operations and custom queries for book management
 */
@Repository
public interface BookRepository extends JpaRepository<Book, UUID> {

    /**
     * Find all available books
     * Used for browsing available inventory
     */
    List<Book> findByIsAvailableTrue();

    /**
     * Find all unavailable (rented) books
     * Used for rental tracking
     */
    List<Book> findByIsAvailableFalse();

    /**
     * Find books by category
     * Used for category filtering
     */
    List<Book> findByCategory(Category category);

    /**
     * Find available books by category
     * Used for filtered browsing
     */
    List<Book> findByCategoryAndIsAvailableTrue(Category category);

    /**
     * Find book by ISBN
     * Used for validation and duplicate detection
     */
    Optional<Book> findByIsbn(String isbn);

    /**
     * Check if ISBN exists
     * Used for validation during imports
     */
    boolean existsByIsbn(String isbn);

    /**
     * Search books by title (case-insensitive, partial match)
     * Used for book search functionality
     */
    @Query("SELECT b FROM Book b WHERE LOWER(b.title) LIKE LOWER(CONCAT('%', :searchTerm, '%'))")
    List<Book> searchByTitle(@Param("searchTerm") String searchTerm);

    /**
     * Search books by author (case-insensitive, partial match)
     * Used for book search functionality
     */
    @Query("SELECT b FROM Book b WHERE LOWER(b.author) LIKE LOWER(CONCAT('%', :searchTerm, '%'))")
    List<Book> searchByAuthor(@Param("searchTerm") String searchTerm);

    /**
     * Search books by title or author (case-insensitive, partial match)
     * Used for general book search
     */
    @Query("SELECT b FROM Book b WHERE LOWER(b.title) LIKE LOWER(CONCAT('%', :searchTerm, '%')) " +
           "OR LOWER(b.author) LIKE LOWER(CONCAT('%', :searchTerm, '%'))")
    List<Book> searchByTitleOrAuthor(@Param("searchTerm") String searchTerm);

    /**
     * Search available books by title or author
     * Used for filtered search
     */
    @Query("SELECT b FROM Book b WHERE b.isAvailable = true AND " +
           "(LOWER(b.title) LIKE LOWER(CONCAT('%', :searchTerm, '%')) " +
           "OR LOWER(b.author) LIKE LOWER(CONCAT('%', :searchTerm, '%')))")
    List<Book> searchAvailableByTitleOrAuthor(@Param("searchTerm") String searchTerm);

    /**
     * Find books by category and search term
     * Used for combined filtering
     */
    @Query("SELECT b FROM Book b WHERE b.category = :category AND " +
           "(LOWER(b.title) LIKE LOWER(CONCAT('%', :searchTerm, '%')) " +
           "OR LOWER(b.author) LIKE LOWER(CONCAT('%', :searchTerm, '%')))")
    List<Book> searchByCategoryAndTitleOrAuthor(@Param("category") Category category,
                                                 @Param("searchTerm") String searchTerm);
}
