# Proxy Checker in Go

This is a simple proxy checker written in Go. It checks the links in the `links.json` file and returns a report.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)

### Installing

```bash
git clone https://github.com/danieljancar/go-proxy-request-checker.git
cd go-proxy-checker
```

### Configuration

The `config/links.json` file contains the proxies to be checked. The format is as follows, make sure to create a `links.json` file in the `config` folder if it does not exist.

```json
[
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  },
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  },
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 403
  }
]
```

## Usage

### Source Code

```bash
go run cmd/main.go
```

### Binary

```bash
go build -o proxy-checker cmd/main.go
./proxy-checker
```

> Recommended: Run the binary on the project root directory to avoid any issues.

## Output

```bash
$ go run main.go
2023/12/21 18:30:12 Requesting https://www.google.com
2023/12/21 18:30:13 Request to https://www.google.com successful with expected status code 200
2023/12/21 18:30:13 Requesting https://www.slack.com
2023/12/21 18:30:13 Request to https://www.slack.com successful with expected status code 200
Do you wish to export the results? (y/n): y
Enter file path to save the report (default: report.json): test-output
2023/12/21 18:30:25 Report successfully saved to test-output
```

# Known Issues

- [ ] If running the binary, the report is not saved to the specified file path, also the links.json file is not found.

# Contributing

I welcome contributions from everyone. Before you get started, please take a moment to review the [guidelines](.github/CONTRIBUTING.md).

## Branches

This project uses the following branches:

- `master` - The main branch. This branch is used for production releases.
- `develop` - The development branch. This branch is used for development and testing.

## Commit Messages

When you commit, make sure to follow the [Contributing Guide](.github/CONTRIBUTING.md) for commit messages. This will help keep the commit messages clean and consistent across the project. If you don't follow the guide, your pull request will be rejected.

# License

This project is licensed under the [MIT](LICENSE) license. By contributing to this project, you agree that your contributions will be released under the same license. Also, you agree to the [Contributor Covenant Code of Conduct](.github/CODE_OF_CONDUCT.md) and [Developer Certificate of Origin](.github/DCO.md).

## Disclaimer

This project is for educational purposes only. I am not responsible for any misuse of this project
