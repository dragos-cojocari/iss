---
description: Security best practices for all generated and modified code
alwaysApply: true
---

# Security Rule

1. **Never document tokens or secrets** in README files or any other documentation. Use placeholder values (e.g., `<YOUR_API_KEY>`) and reference environment variables instead.
2. **Alert immediately** if any modified or reviewed file contains hardcoded secrets, API keys, passwords, tokens, or credentials. Flag the exact line and suggest remediation (environment variables, secret managers, `.env` files in `.gitignore`).
3. **Always use secure defaults** in all generated code:
   - HTTPS over HTTP
   - Parameterized queries over string concatenation
   - Least-privilege permissions
   - Input validation and output encoding
   - Secure headers (CORS, CSP, etc.)
   - Non-root users in containers
   - Read-only filesystems where possible
