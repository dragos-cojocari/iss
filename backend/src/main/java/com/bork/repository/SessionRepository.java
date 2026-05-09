package com.bork.repository;

import com.bork.model.Session;
import com.bork.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository interface for Session entity
 * Provides CRUD operations and custom queries for session management
 */
@Repository
public interface SessionRepository extends JpaRepository<Session, UUID> {

    /**
     * Find session by ID and check if active and not expired
     */
    @Query("SELECT s FROM Session s WHERE s.sessionId = :sessionId AND s.isActive = true AND s.expiresAt > :now")
    Optional<Session> findActiveSession(@Param("sessionId") UUID sessionId, @Param("now") LocalDateTime now);

    /**
     * Find session by ID with user eagerly fetched
     */
    @Query("SELECT s FROM Session s JOIN FETCH s.user " +
           "WHERE s.sessionId = :sessionId AND s.isActive = true " +
           "AND s.expiresAt > :now")
    Optional<Session> findActiveSessionWithUser(
            @Param("sessionId") UUID sessionId,
            @Param("now") LocalDateTime now);

    /**
     * Find all active sessions for a user
     */
    @Query("SELECT s FROM Session s WHERE s.user = :user AND s.isActive = true AND s.expiresAt > :now")
    List<Session> findActiveSessionsByUser(@Param("user") User user, @Param("now") LocalDateTime now);

    /**
     * Invalidate all active sessions for a user
     * Used when user changes password or for security purposes
     */
    @Modifying
    @Query("UPDATE Session s SET s.isActive = false WHERE s.user = :user AND s.isActive = true")
    void invalidateAllUserSessions(@Param("user") User user);

    /**
     * Delete expired sessions (cleanup task)
     */
    @Modifying
    @Query("DELETE FROM Session s WHERE s.expiresAt < :now")
    void deleteExpiredSessions(@Param("now") LocalDateTime now);

    /**
     * Find all sessions for a user (for monitoring)
     */
    List<Session> findByUser(User user);
}
