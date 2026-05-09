package com.bork.dto;

import io.swagger.v3.oas.annotations.media.Schema;
import jakarta.validation.constraints.NotBlank;

/**
 * DTO for login request
 */
@Schema(description = "Login request with username and password")
public class LoginRequest {

    @Schema(description = "Username", example = "student1", required = true)
    @NotBlank(message = "Username is required")
    private String username;

    @Schema(description = "Password", example = "Test123!", required = true)
    @NotBlank(message = "Password is required")
    private String password;

    public LoginRequest() {
    }

    public LoginRequest(String username, String password) {
        this.username = username;
        this.password = password;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}
