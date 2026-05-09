# Ingineria Sistemelor de Software - Project BORK

## Documentation

All functional specifications and technical documentation can be found in [docs](./docs/index.md).

## Lab Assignments

| Date       | Description                                                                             | Files                                                                                             |
| ---------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| 2026-05-09 | Lab 3 & 4: add class diagram and sequence diagrams for 2 usecases. Add the UI prototype | `docs/class_diagram.md`<br>`docs/sd_login.md`<br>`docs/sd_rentbook.md`<br>`docs/tui_prototype.md` |
| 2026-03-21 | Lab 2: Use Cases and Use Case Diagram                                                   | `docs/[ISS 2025-2026] Use case template.docx`<br>`docs/nfr.md`<br>`docs/usecase_diagram.md`       |
| 2026-03-07 | Lab 1: Flow diagram                                                                     | `docs/flow_diagram.md`                                                                            |
| 2026-03-07 | Lab 1: Project goals                                                                    | `Makefile`<br>`docs/specification.md`                                                             |

## Development Environment Setup

To set up your development environment, run:

```bash
make setup
```

See the [Makefile](./Makefile) for available commands.

## Contribution Guidelines

- **Pull Requests**: All contributions must be made via pull requests
- **Commit Format**: Follow [Conventional Commits](https://www.conventionalcommits.org/) specification
- **Signed Commits**: All commits must be signed with GPG/SSH keys
- **Pre-commit Hooks**: Run pre-commit hooks before submitting (`pre-commit run --all-files`)
- **Squash Commits**: Squash all commits into a single commit before merging
- **Stay Updated**: Keep your branch up to date with the main branch
