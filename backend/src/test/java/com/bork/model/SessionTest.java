package com.bork.model;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class SessionTest {

    private User testUser;
    private Session session;

    @BeforeEach
    void setUp() {
        testUser = new User();
        testUser.setUserId(UUID.randomUUID());
        testUser.setUsername("testuser");
        testUser.setEmail("test@example.com");

        session = new Session(testUser);
    }

    @Test
    void sessionCreation_ShouldInitializeWithDefaults() {
        assertEquals(testUser, session.getUser());
        assertTrue(session.getIsActive());
        assertNotNull(session.getExpiresAt());
        assertNotNull(session.getLastAccessedAt());
    }

    @Test
    void sessionCreation_ShouldExpireIn30Minutes() {
        LocalDateTime now = LocalDateTime.now();
        LocalDateTime expiresAt = session.getExpiresAt();

        long minutesDiff = java.time.Duration.between(now, expiresAt).toMinutes();
        assertTrue(minutesDiff >= 29 && minutesDiff <= 30);
    }

    @Test
    void refresh_ShouldUpdateExpiryAndLastAccessed() throws InterruptedException {
        LocalDateTime originalExpiry = session.getExpiresAt();
        LocalDateTime originalLastAccessed = session.getLastAccessedAt();

        Thread.sleep(10);

        session.refresh();

        assertTrue(session.getExpiresAt().isAfter(originalExpiry));
        assertTrue(session.getLastAccessedAt().isAfter(originalLastAccessed));
    }

    @Test
    void invalidate_ShouldSetIsActiveToFalse() {
        assertTrue(session.getIsActive());

        session.invalidate();

        assertFalse(session.getIsActive());
    }

    @Test
    void isExpired_WithFutureExpiry_ShouldReturnFalse() {
        assertFalse(session.isExpired());
    }

    @Test
    void isExpired_WithPastExpiry_ShouldReturnTrue() {
        session.setExpiresAt(LocalDateTime.now().minusMinutes(1));

        assertTrue(session.isExpired());
    }

    @Test
    void sessionSettersAndGetters_ShouldWorkCorrectly() {
        UUID sessionId = UUID.randomUUID();
        LocalDateTime now = LocalDateTime.now();

        session.setSessionId(sessionId);
        session.setIsActive(false);
        session.setCreatedAt(now);
        session.setExpiresAt(now.plusMinutes(30));
        session.setLastAccessedAt(now);

        assertEquals(sessionId, session.getSessionId());
        assertFalse(session.getIsActive());
        assertEquals(now, session.getCreatedAt());
        assertEquals(now.plusMinutes(30), session.getExpiresAt());
        assertEquals(now, session.getLastAccessedAt());
    }
}
