package com.bork.repository;

import com.bork.model.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

/**
 * Repository interface for Category entity
 * Provides CRUD operations and custom queries for category management
 */
@Repository
public interface CategoryRepository extends JpaRepository<Category, UUID> {

    /**
     * Find category by name (case-sensitive)
     * Used for filtering and validation
     */
    Optional<Category> findByName(String name);

    /**
     * Check if category name exists
     * Used for validation during imports
     */
    boolean existsByName(String name);

    /**
     * Find category by name (case-insensitive)
     * Used for flexible searching
     */
    Optional<Category> findByNameIgnoreCase(String name);
}
