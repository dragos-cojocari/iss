package com.bork.repository;

import com.bork.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

/**
 * Repository interface for User entity
 * Provides CRUD operations and custom queries for user management
 */
@Repository
public interface UserRepository extends JpaRepository<User, UUID> {

    /**
     * Find user by username (case-sensitive)
     * Used for authentication
     */
    Optional<User> findByUsername(String username);

    /**
     * Find user by email
     * Used for validation and account recovery
     */
    Optional<User> findByEmail(String email);

    /**
     * Check if username exists
     * Used for registration validation
     */
    boolean existsByUsername(String username);

    /**
     * Check if email exists
     * Used for registration validation
     */
    boolean existsByEmail(String email);

    /**
     * Find all locked accounts
     * Used for administrative monitoring
     */
    java.util.List<User> findByIsLockedTrue();
}
