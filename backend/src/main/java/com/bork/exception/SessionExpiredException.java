package com.bork.exception;

/**
 * Exception thrown when session is expired or invalid
 */
public class SessionExpiredException extends AuthenticationException {

    public SessionExpiredException() {
        super("Session has expired or is invalid");
    }

    public SessionExpiredException(String message) {
        super(message);
    }
}
