---
description: Keep documentation in sync with code changes
alwaysApply: true
---

# Documentation Sync Rule

After every code or configuration change, review the project documentation and **propose updates** to keep it accurate.

## Rules

1. **Review after each change**: After implementing a requested change, scan `docs/` and relevant `README.md` files for content that is now outdated or incomplete.
2. **Propose before writing**: Present the suggested documentation changes to the user for approval before editing.
3. **Brevity over verbosity**: Documentation is technical. Prefer concise tables, short bullet points, and minimal prose. Avoid restating what is obvious from the code.
4. **Scope**: Only update documentation directly affected by the change. Do not rewrite unrelated sections.
5. **Files to check**: `docs/`, root `README.md`, `backend/README.md`, `frontend/README.md`, `database/README.md`, and any file referenced by the change.
