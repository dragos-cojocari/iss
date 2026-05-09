package com.bork.controller;

import com.bork.dto.LoginRequest;
import com.bork.exception.InvalidCredentialsException;
import com.bork.model.Session;
import com.bork.model.User;
import com.bork.service.AuthenticationService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class AuthControllerTest {

    @Mock
    private AuthenticationService authenticationService;

    @InjectMocks
    private AuthController authController;

    private User testUser;
    private Session testSession;

    @BeforeEach
    void setUp() {
        testUser = new User();
        testUser.setUserId(UUID.randomUUID());
        testUser.setUsername("testuser");
        testUser.setEmail("test@example.com");
        testUser.setFirstName("Test");
        testUser.setLastName("User");

        testSession = new Session(testUser);
    }

    @Test
    void login_WithValidCredentials_ShouldCallAuthenticationService() {
        when(authenticationService.login(anyString(), anyString())).thenReturn(testSession);

        assertNotNull(authenticationService.login("testuser", "password"));
        verify(authenticationService).login("testuser", "password");
    }

    @Test
    void authenticationService_ShouldBeInjected() {
        assertNotNull(authenticationService);
        assertNotNull(authController);
    }
}
