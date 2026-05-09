package com.bork.model;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class UserTest {

    private User user;

    @BeforeEach
    void setUp() {
        user = new User();
        user.setUserId(UUID.randomUUID());
        user.setUsername("testuser");
        user.setEmail("test@example.com");
        user.setPasswordHash("hashedPassword");
        user.setFirstName("Test");
        user.setLastName("User");
        user.setIsLocked(false);
        user.setFailedLoginAttempts(0);
    }

    @Test
    void incrementFailedLoginAttempts_ShouldIncreaseCounter() {
        assertEquals(0, user.getFailedLoginAttempts());

        user.incrementFailedLoginAttempts();
        assertEquals(1, user.getFailedLoginAttempts());

        user.incrementFailedLoginAttempts();
        assertEquals(2, user.getFailedLoginAttempts());
    }

    @Test
    void incrementFailedLoginAttempts_ShouldLockAccountAfterFiveAttempts() {
        assertFalse(user.getIsLocked());

        for (int i = 0; i < 5; i++) {
            user.incrementFailedLoginAttempts();
        }

        assertEquals(5, user.getFailedLoginAttempts());
        assertTrue(user.getIsLocked());
    }

    @Test
    void resetFailedLoginAttempts_ShouldResetCounter() {
        user.setFailedLoginAttempts(3);
        assertEquals(3, user.getFailedLoginAttempts());

        user.resetFailedLoginAttempts();
        assertEquals(0, user.getFailedLoginAttempts());
    }

    @Test
    void userCreation_ShouldHaveDefaultValues() {
        User newUser = new User();

        assertNotNull(newUser);
        assertEquals(0, newUser.getFailedLoginAttempts());
        assertFalse(newUser.getIsLocked());
    }

    @Test
    void userSettersAndGetters_ShouldWorkCorrectly() {
        UUID userId = UUID.randomUUID();
        user.setUserId(userId);
        user.setUsername("newusername");
        user.setEmail("newemail@example.com");
        user.setFirstName("New");
        user.setLastName("Name");

        assertEquals(userId, user.getUserId());
        assertEquals("newusername", user.getUsername());
        assertEquals("newemail@example.com", user.getEmail());
        assertEquals("New", user.getFirstName());
        assertEquals("Name", user.getLastName());
    }
}
