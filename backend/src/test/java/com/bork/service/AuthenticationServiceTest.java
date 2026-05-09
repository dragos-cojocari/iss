package com.bork.service;

import com.bork.exception.AccountLockedException;
import com.bork.exception.InvalidCredentialsException;
import com.bork.exception.SessionExpiredException;
import com.bork.model.Session;
import com.bork.model.User;
import com.bork.repository.SessionRepository;
import com.bork.repository.UserRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.time.LocalDateTime;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class AuthenticationServiceTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private SessionRepository sessionRepository;

    @Mock
    private PasswordEncoder passwordEncoder;

    @InjectMocks
    private AuthenticationService authenticationService;

    private User testUser;
    private Session testSession;

    @BeforeEach
    void setUp() {
        testUser = new User();
        testUser.setUserId(UUID.randomUUID());
        testUser.setUsername("testuser");
        testUser.setPasswordHash("hashedPassword");
        testUser.setEmail("test@example.com");
        testUser.setFirstName("Test");
        testUser.setLastName("User");
        testUser.setIsLocked(false);
        testUser.setFailedLoginAttempts(0);

        testSession = new Session(testUser);
    }

    @Test
    void login_WithValidCredentials_ShouldReturnSession() {
        when(userRepository.findByUsername("testuser")).thenReturn(Optional.of(testUser));
        when(passwordEncoder.matches("password", "hashedPassword")).thenReturn(true);
        when(sessionRepository.save(any(Session.class))).thenReturn(testSession);

        Session result = authenticationService.login("testuser", "password");

        assertNotNull(result);
        verify(userRepository).save(testUser);
        verify(sessionRepository).save(any(Session.class));
        assertEquals(0, testUser.getFailedLoginAttempts());
    }

    @Test
    void login_WithInvalidUsername_ShouldThrowInvalidCredentialsException() {
        when(userRepository.findByUsername("wronguser")).thenReturn(Optional.empty());

        assertThrows(InvalidCredentialsException.class, () -> {
            authenticationService.login("wronguser", "password");
        });

        verify(passwordEncoder, never()).matches(anyString(), anyString());
        verify(sessionRepository, never()).save(any(Session.class));
    }

    @Test
    void login_WithInvalidPassword_ShouldThrowInvalidCredentialsException() {
        when(userRepository.findByUsername("testuser")).thenReturn(Optional.of(testUser));
        when(passwordEncoder.matches("wrongpassword", "hashedPassword")).thenReturn(false);

        assertThrows(InvalidCredentialsException.class, () -> {
            authenticationService.login("testuser", "wrongpassword");
        });

        verify(userRepository).save(testUser);
        assertEquals(1, testUser.getFailedLoginAttempts());
        verify(sessionRepository, never()).save(any(Session.class));
    }

    @Test
    void login_WithLockedAccount_ShouldThrowAccountLockedException() {
        testUser.setIsLocked(true);
        when(userRepository.findByUsername("testuser")).thenReturn(Optional.of(testUser));

        assertThrows(AccountLockedException.class, () -> {
            authenticationService.login("testuser", "password");
        });

        verify(passwordEncoder, never()).matches(anyString(), anyString());
        verify(sessionRepository, never()).save(any(Session.class));
    }

    @Test
    void login_WithMaxFailedAttempts_ShouldLockAccountAndThrow() {
        testUser.setFailedLoginAttempts(4);
        when(userRepository.findByUsername("testuser")).thenReturn(Optional.of(testUser));
        when(passwordEncoder.matches("wrongpassword", "hashedPassword")).thenReturn(false);

        assertThrows(AccountLockedException.class, () -> {
            authenticationService.login("testuser", "wrongpassword");
        });

        assertEquals(5, testUser.getFailedLoginAttempts());
        assertEquals(true, testUser.getIsLocked());
        verify(userRepository).save(testUser);
    }

    @Test
    void validateSession_WithValidSession_ShouldReturnRefreshedSession() {
        UUID sessionId = UUID.randomUUID();
        when(sessionRepository.findActiveSessionWithUser(any(UUID.class), any(LocalDateTime.class)))
                .thenReturn(Optional.of(testSession));
        when(sessionRepository.save(any(Session.class))).thenReturn(testSession);

        Session result = authenticationService.validateSession(sessionId);

        assertNotNull(result);
        verify(sessionRepository).save(testSession);
    }

    @Test
    void validateSession_WithInvalidSession_ShouldThrowSessionExpiredException() {
        UUID sessionId = UUID.randomUUID();
        when(sessionRepository.findActiveSessionWithUser(any(UUID.class), any(LocalDateTime.class)))
                .thenReturn(Optional.empty());

        assertThrows(SessionExpiredException.class, () -> {
            authenticationService.validateSession(sessionId);
        });

        verify(sessionRepository, never()).save(any(Session.class));
    }

    @Test
    void logout_WithValidSession_ShouldInvalidateSession() {
        UUID sessionId = UUID.randomUUID();
        testSession.setSessionId(sessionId);
        when(sessionRepository.findById(sessionId)).thenReturn(Optional.of(testSession));

        authenticationService.logout(sessionId);

        verify(sessionRepository).save(testSession);
        assertEquals(false, testSession.getIsActive());
    }

    @Test
    void logout_WithInvalidSession_ShouldNotThrow() {
        UUID sessionId = UUID.randomUUID();
        when(sessionRepository.findById(sessionId)).thenReturn(Optional.empty());

        authenticationService.logout(sessionId);

        verify(sessionRepository, never()).save(any(Session.class));
    }

    @Test
    void getUserFromSession_WithValidSession_ShouldReturnUser() {
        UUID sessionId = UUID.randomUUID();
        when(sessionRepository.findActiveSessionWithUser(any(UUID.class), any(LocalDateTime.class)))
                .thenReturn(Optional.of(testSession));

        User result = authenticationService.getUserFromSession(sessionId);

        assertNotNull(result);
        assertEquals(testUser.getUserId(), result.getUserId());
        assertEquals(testUser.getUsername(), result.getUsername());
    }

    @Test
    void getUserFromSession_WithInvalidSession_ShouldThrowSessionExpiredException() {
        UUID sessionId = UUID.randomUUID();
        when(sessionRepository.findActiveSessionWithUser(any(UUID.class), any(LocalDateTime.class)))
                .thenReturn(Optional.empty());

        assertThrows(SessionExpiredException.class, () -> {
            authenticationService.getUserFromSession(sessionId);
        });
    }

    @Test
    void invalidateAllUserSessions_ShouldCallRepository() {
        authenticationService.invalidateAllUserSessions(testUser);

        verify(sessionRepository).invalidateAllUserSessions(testUser);
    }

    @Test
    void cleanupExpiredSessions_ShouldCallRepository() {
        authenticationService.cleanupExpiredSessions();

        verify(sessionRepository).deleteExpiredSessions(any(LocalDateTime.class));
    }
}
