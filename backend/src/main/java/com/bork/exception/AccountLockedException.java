package com.bork.exception;

/**
 * Exception thrown when attempting to login to a locked account
 */
public class AccountLockedException extends AuthenticationException {

    public AccountLockedException() {
        super("Account is locked due to multiple failed login attempts");
    }

    public AccountLockedException(String message) {
        super(message);
    }
}
