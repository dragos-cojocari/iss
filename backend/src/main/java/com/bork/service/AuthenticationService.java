package com.bork.service;

import com.bork.exception.AccountLockedException;
import com.bork.exception.InvalidCredentialsException;
import com.bork.exception.SessionExpiredException;
import com.bork.model.Session;
import com.bork.model.User;
import com.bork.repository.SessionRepository;
import com.bork.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.UUID;

/**
 * Service for handling authentication operations
 * Manages login, logout, and session validation
 */
@Service
public class AuthenticationService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private SessionRepository sessionRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    /**
     * Authenticate user and create session
     *
     * @param username User's username
     * @param password User's plain text password
     * @return Session object with session ID
     * @throws InvalidCredentialsException if credentials are invalid
     * @throws AccountLockedException if account is locked
     */
    @Transactional
    public Session login(String username, String password) {
        User user = userRepository.findByUsername(username)
                .orElseThrow(() -> new InvalidCredentialsException());

        if (user.getIsLocked()) {
            throw new AccountLockedException();
        }

        if (!passwordEncoder.matches(password, user.getPasswordHash())) {
            user.incrementFailedLoginAttempts();
            userRepository.save(user);

            if (user.getIsLocked()) {
                throw new AccountLockedException();
            }

            throw new InvalidCredentialsException();
        }

        user.resetFailedLoginAttempts();
        userRepository.save(user);

        Session session = new Session(user);
        return sessionRepository.save(session);
    }

    /**
     * Validate and refresh session
     *
     * @param sessionId Session UUID
     * @return Session object if valid
     * @throws SessionExpiredException if session is invalid or expired
     */
    @Transactional
    public Session validateSession(UUID sessionId) {
        Session session = sessionRepository.findActiveSessionWithUser(sessionId, LocalDateTime.now())
                .orElseThrow(() -> new SessionExpiredException());

        session.refresh();
        return sessionRepository.save(session);
    }

    /**
     * Logout user and invalidate session
     *
     * @param sessionId Session UUID to invalidate
     */
    @Transactional
    public void logout(UUID sessionId) {
        sessionRepository.findById(sessionId).ifPresent(session -> {
            session.invalidate();
            sessionRepository.save(session);
        });
    }

    /**
     * Get user from session
     *
     * @param sessionId Session UUID
     * @return User object
     * @throws SessionExpiredException if session is invalid
     */
    @Transactional(readOnly = true)
    public User getUserFromSession(UUID sessionId) {
        Session session = sessionRepository.findActiveSessionWithUser(sessionId, LocalDateTime.now())
                .orElseThrow(() -> new SessionExpiredException());

        return session.getUser();
    }

    /**
     * Invalidate all sessions for a user
     * Used for security purposes (e.g., password change)
     *
     * @param user User whose sessions to invalidate
     */
    @Transactional
    public void invalidateAllUserSessions(User user) {
        sessionRepository.invalidateAllUserSessions(user);
    }

    /**
     * Cleanup expired sessions (scheduled task)
     */
    @Transactional
    public void cleanupExpiredSessions() {
        sessionRepository.deleteExpiredSessions(LocalDateTime.now());
    }
}
