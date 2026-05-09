package com.bork.dto;

import com.bork.model.User;

import java.util.UUID;

/**
 * DTO for login response
 */
public class LoginResponse {

    private UUID sessionId;
    private UserInfo user;

    public LoginResponse() {
    }

    public LoginResponse(UUID sessionId, User user) {
        this.sessionId = sessionId;
        this.user = new UserInfo(user);
    }

    public UUID getSessionId() {
        return sessionId;
    }

    public void setSessionId(UUID sessionId) {
        this.sessionId = sessionId;
    }

    public UserInfo getUser() {
        return user;
    }

    public void setUser(UserInfo user) {
        this.user = user;
    }

    public static class UserInfo {
        private UUID userId;
        private String username;
        private String email;
        private String firstName;
        private String lastName;

        public UserInfo() {
        }

        public UserInfo(User user) {
            this.userId = user.getUserId();
            this.username = user.getUsername();
            this.email = user.getEmail();
            this.firstName = user.getFirstName();
            this.lastName = user.getLastName();
        }

        public UUID getUserId() {
            return userId;
        }

        public void setUserId(UUID userId) {
            this.userId = userId;
        }

        public String getUsername() {
            return username;
        }

        public void setUsername(String username) {
            this.username = username;
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public String getFirstName() {
            return firstName;
        }

        public void setFirstName(String firstName) {
            this.firstName = firstName;
        }

        public String getLastName() {
            return lastName;
        }

        public void setLastName(String lastName) {
            this.lastName = lastName;
        }
    }
}
