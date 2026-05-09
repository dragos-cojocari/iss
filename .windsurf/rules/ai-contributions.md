---
description: Log all AI contributions to AI_Contributions.md
alwaysApply: true
---

# AI Contributions Logging Rule

**Incrementally** update `AI_Contributions.md` in the project root after each completed task, not just at the end of the conversation. Follow the existing format exactly:

## Session Entry Format

```markdown
---

# Session N: <Short Descriptive Title>

**Date**: <current date>
**Duration**: ~<estimated duration>
**Focus**: <brief focus area>

## Task Log

### 1. <Task Title>

**Request**: <what the user asked>

**Complexity**: <Low | Medium | High>
**Follow-up**: <Yes/No (explanation)>
**Analysis**:

- <bullet points describing what was done>

**Missing/Suggestions**:

- <what could have been clearer, or "None">
```

## Rules

1. **Increment the session number** based on the last session in the file.
2. **Log every task** performed during the conversation, numbered sequentially.
3. **Include a Session Summary** at the end with total tasks, complexity breakdown, and impact.
4. **Update the "Overall Summary: All Sessions" section** at the bottom of the file with revised cumulative statistics (total sessions, tasks, metrics).
5. **Be honest** about complexity and suggestions — this is a learning tool.
6. **Do not skip this step** even if the session was short or trivial.
7. If a task involved fixing issues found by Snyk or other tools, note that in the analysis.
