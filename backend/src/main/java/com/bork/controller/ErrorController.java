package com.bork.controller;

import jakarta.servlet.RequestDispatcher;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Custom error controller to handle errors without exposing internal details
 * Replaces the default Spring Boot Whitelabel Error Page
 */
@RestController
public class ErrorController implements org.springframework.boot.web.servlet.error.ErrorController {

    @RequestMapping("/error")
    public ResponseEntity<Map<String, Object>> handleError(HttpServletRequest request) {
        Object status = request.getAttribute(RequestDispatcher.ERROR_STATUS_CODE);

        Map<String, Object> errorResponse = new HashMap<>();
        errorResponse.put("timestamp", LocalDateTime.now().toString());

        if (status != null) {
            int statusCode = Integer.parseInt(status.toString());
            HttpStatus httpStatus = HttpStatus.valueOf(statusCode);

            errorResponse.put("status", statusCode);
            errorResponse.put("error", httpStatus.getReasonPhrase());

            // Generic messages without exposing internal details
            switch (statusCode) {
                case 404:
                    errorResponse.put("message", "The requested resource was not found");
                    break;
                case 400:
                    errorResponse.put("message", "Bad request");
                    break;
                case 401:
                    errorResponse.put("message", "Unauthorized");
                    break;
                case 403:
                    errorResponse.put("message", "Forbidden");
                    break;
                case 500:
                    errorResponse.put("message", "Internal server error");
                    break;
                default:
                    errorResponse.put("message", "An error occurred");
            }

            return ResponseEntity.status(httpStatus).body(errorResponse);
        }

        // Default error response
        errorResponse.put("status", 500);
        errorResponse.put("error", "Internal Server Error");
        errorResponse.put("message", "An unexpected error occurred");

        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(errorResponse);
    }
}
