package com.bork.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.sql.DataSource;
import java.sql.Connection;
import java.util.HashMap;
import java.util.Map;

/**
 * Health check endpoint for monitoring application and database status
 */
@RestController
@RequestMapping("/api/health")
public class HealthController {

    @Autowired
    private DataSource dataSource;

    @GetMapping
    public ResponseEntity<Map<String, Object>> health() {
        Map<String, Object> health = new HashMap<>();
        health.put("status", "UP");
        health.put("application", "BORK Backend");
        health.put("version", "0.1.0-SNAPSHOT");

        // Check database connection
        try (Connection connection = dataSource.getConnection()) {
            health.put("database", "Connected");
            health.put("databaseProductName", connection.getMetaData().getDatabaseProductName());
            health.put("databaseProductVersion", connection.getMetaData().getDatabaseProductVersion());
        } catch (Exception e) {
            health.put("database", "Disconnected");
            health.put("databaseError", e.getMessage());
            return ResponseEntity.status(503).body(health);
        }

        return ResponseEntity.ok(health);
    }
}
