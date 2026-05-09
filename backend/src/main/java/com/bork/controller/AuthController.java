package com.bork.controller;

import com.bork.dto.LoginRequest;
import com.bork.dto.LoginResponse;
import com.bork.model.Session;
import com.bork.model.User;
import com.bork.service.AuthenticationService;
import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * REST controller for authentication operations
 * Handles login, logout, and session management
 */
@RestController
@RequestMapping("/api/auth")
public class AuthController {

    private static final String SESSION_COOKIE_NAME = "BORK_SESSION";
    private static final int COOKIE_MAX_AGE = 30 * 60;

    @Autowired
    private AuthenticationService authenticationService;

    /**
     * Login endpoint
     * POST /api/auth/login
     *
     * @param loginRequest Username and password
     * @param response HTTP response to set cookie
     * @return Login response with session ID and user info
     */
    @PostMapping("/login")
    public ResponseEntity<LoginResponse> login(
            @Valid @RequestBody LoginRequest loginRequest,
            HttpServletResponse response) {

        Session session = authenticationService.login(
                loginRequest.getUsername(),
                loginRequest.getPassword()
        );

        Cookie sessionCookie = new Cookie(SESSION_COOKIE_NAME, session.getSessionId().toString());
        sessionCookie.setHttpOnly(true);
        sessionCookie.setPath("/");
        sessionCookie.setMaxAge(COOKIE_MAX_AGE);
        response.addCookie(sessionCookie);

        LoginResponse loginResponse = new LoginResponse(session.getSessionId(), session.getUser());
        return ResponseEntity.ok(loginResponse);
    }

    /**
     * Logout endpoint
     * POST /api/auth/logout
     *
     * @param request HTTP request to get session cookie
     * @param response HTTP response to clear cookie
     * @return Success message
     */
    @PostMapping("/logout")
    public ResponseEntity<Map<String, String>> logout(
            HttpServletRequest request,
            HttpServletResponse response) {

        UUID sessionId = getSessionIdFromRequest(request);
        if (sessionId != null) {
            authenticationService.logout(sessionId);
        }

        Cookie sessionCookie = new Cookie(SESSION_COOKIE_NAME, "");
        sessionCookie.setHttpOnly(true);
        sessionCookie.setPath("/");
        sessionCookie.setMaxAge(0);
        response.addCookie(sessionCookie);

        Map<String, String> responseBody = new HashMap<>();
        responseBody.put("message", "Logged out successfully");
        return ResponseEntity.ok(responseBody);
    }

    /**
     * Get current user info
     * GET /api/auth/me
     *
     * @param request HTTP request to get session cookie
     * @return Current user information
     */
    @GetMapping("/me")
    public ResponseEntity<LoginResponse.UserInfo> getCurrentUser(HttpServletRequest request) {
        UUID sessionId = getSessionIdFromRequest(request);
        if (sessionId == null) {
            return ResponseEntity.status(401).build();
        }

        User user = authenticationService.getUserFromSession(sessionId);
        return ResponseEntity.ok(new LoginResponse.UserInfo(user));
    }

    /**
     * Extract session ID from cookie
     */
    private UUID getSessionIdFromRequest(HttpServletRequest request) {
        Cookie[] cookies = request.getCookies();
        if (cookies != null) {
            for (Cookie cookie : cookies) {
                if (SESSION_COOKIE_NAME.equals(cookie.getName())) {
                    try {
                        return UUID.fromString(cookie.getValue());
                    } catch (IllegalArgumentException e) {
                        return null;
                    }
                }
            }
        }
        return null;
    }
}
