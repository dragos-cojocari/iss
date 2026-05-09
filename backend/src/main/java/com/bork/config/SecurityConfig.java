package com.bork.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

/**
 * Security configuration for password encoding
 *
 * Configures BCrypt password encoder with strength 12 for secure password hashing
 */
@Configuration
public class SecurityConfig {

    /**
     * BCrypt password encoder bean
     *
     * Strength 12 provides a good balance between security and performance
     * Each increment doubles the hashing time
     *
     * @return BCryptPasswordEncoder with strength 12
     */
    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder(12);
    }
}
