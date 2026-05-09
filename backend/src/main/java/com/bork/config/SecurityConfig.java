package com.bork.config;

import com.bork.filter.SessionAuthenticationFilter;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

/**
 * Security configuration for password encoding and session authentication
 *
 * Configures BCrypt password encoder and session validation filter
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

    /**
     * Register session authentication filter
     * This filter validates sessions for protected endpoints
     *
     * @param filter SessionAuthenticationFilter instance
     * @return FilterRegistrationBean configuration
     */
    @Bean
    public FilterRegistrationBean<SessionAuthenticationFilter> sessionAuthFilter(SessionAuthenticationFilter filter) {
        FilterRegistrationBean<SessionAuthenticationFilter> registration = new FilterRegistrationBean<>();
        registration.setFilter(filter);
        registration.addUrlPatterns("/api/*");
        registration.setOrder(1);
        return registration;
    }
}
