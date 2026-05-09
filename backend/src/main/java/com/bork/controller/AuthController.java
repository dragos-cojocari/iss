package com.bork.controller;

import com.bork.dto.LoginRequest;
import com.bork.dto.LoginResponse;
import com.bork.model.Session;
import com.bork.model.User;
import com.bork.service.AuthenticationService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * REST controller for authentication operations
 * Handles login, logout, and session management
 */
@RestController
@RequestMapping(value = "/api/auth", produces = "application/json")
@Tag(name = "Authentication", description = "User authentication and session management endpoints")
public class AuthController {

    private static final Logger logger = LoggerFactory.getLogger(AuthController.class);
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
    @Operation(
            summary = "User login",
            description = "Authenticate with username and password. Returns session cookie and user information."
    )
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Login successful",
                    content = @Content(schema = @Schema(implementation = LoginResponse.class))),
            @ApiResponse(responseCode = "401", description = "Invalid credentials"),
            @ApiResponse(responseCode = "403", description = "Account locked (3 failed attempts)")
    })
    @PostMapping(value = "/login", consumes = "application/json")
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
    @Operation(
            summary = "User logout",
            description = "Invalidate current session and clear session cookie."
    )
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Logout successful"),
            @ApiResponse(responseCode = "401", description = "No valid session")
    })
    @SecurityRequirement(name = "cookieAuth")
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
    @Operation(
            summary = "Get current user",
            description = "Get information about the currently authenticated user."
    )
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "User information retrieved",
                    content = @Content(schema = @Schema(implementation = LoginResponse.UserInfo.class))),
            @ApiResponse(responseCode = "401", description = "Not authenticated")
    })
    @SecurityRequirement(name = "cookieAuth")
    @GetMapping("/me")
    public ResponseEntity<LoginResponse.UserInfo> getCurrentUser(HttpServletRequest request) {
        // User is already validated and set by SessionAuthenticationFilter
        User user = (User) request.getAttribute("user");

        if (user == null) {
            logger.warn("GET /api/auth/me - No user in request attributes");
            return ResponseEntity.status(401).build();
        }

        logger.info("GET /api/auth/me - User: {}", user.getUsername());
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
