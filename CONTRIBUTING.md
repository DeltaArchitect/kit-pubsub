# Contributing to kit-pubsub

Thank you for your interest in contributing to `kit-pubsub`! We welcome contributions from the community.

## Getting Started

1.  **Fork the repository** on GitHub.
2.  **Clone your fork** locally:
    ```bash
    git clone https://github.com/YOUR_USERNAME/kit-pubsub.git
    cd kit-pubsub
    ```
3.  **Create a new branch** for your feature or bugfix:
    ```bash
    git checkout -b feature/my-new-feature
    ```

## Development Workflow

### Prerequisites
- Go 1.26 or later

### Running Tests
Please ensure all tests pass before submitting a Pull Request.
```bash
go test ./... -v -race
```

### Coding Style
We follow standard Go coding conventions.
- Run `gofmt` to format your code.
- Run `go vet` to check for common errors.
- Ensure your code is linted (e.g., using `golangci-lint` if available).

## Submitting a Pull Request

1.  Push your branch to your fork.
2.  Open a Pull Request against the `main` branch of the original repository.
3.  Fill out the Pull Request Template.
4.  Wait for code review and address any feedback.

## License
By contributing, you agree that your contributions will be licensed under the MIT License.
