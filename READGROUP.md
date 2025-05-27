# 🚀 LogAnalyzer - Distributed Log Analysis Tool

[![Go Version](https://img.shields.io/badge/Go-1.24.3-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](.)

A powerful command-line tool for analyzing log files from multiple sources in parallel with robust error handling and concurrent processing.

## 👥 Development Team

- **Mathieu** - Lead Developer & Project Architect
- **Yoann** - the junior core Developer 
- **Swann** - Core Developer

---

## 📋 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Commands](#-commands)
- [Configuration](#-configuration)
- [Examples](#-examples)
- [Project Structure](#-project-structure)
- [Error Handling](#-error-handling)
- [Contributing](#-contributing)

---

## ✨ Features

- **🔄 Concurrent Processing**: Analyze multiple log files simultaneously using goroutines
- **🛡️ Robust Error Handling**: Custom error types with detailed reporting
- **📊 Multiple Output Formats**: Console display and JSON export
- **🕒 Timestamped Reports**: Automatic date stamping for organized reporting
- **🔍 Status Filtering**: Filter results by success/failure status
- **📁 Auto Directory Creation**: Automatically creates output directories
- **⚡ Performance Optimized**: Simulated analysis with configurable timing

---

## 🚀 Installation

### Prerequisites

- Go 1.24.3 or higher
- Git

### Build from Source

```bash
# Clone the repository
git clone https://github.com/Mathieu-ai/loganizer.git
cd loganizer

# Initialize Go module and install dependencies
go mod tidy

# Build the application
go build -o loganalyzer main.go

# (Optional) Install globally
go install
```

---

## ⚡ Quick Start

1. **Create a configuration file** (`config.json`):
```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  }
]
```

2. **Run the analysis**:
```bash
./loganalyzer analyze --config config.json --output results.json
```

---

## 📖 Commands

### Root Command

```bash
loganalyzer
```

**Description**: A CLI tool for analyzing log files
**Usage**: Base command that shows help and available subcommands

---

### `analyze` Command

Analyze log files based on configuration with concurrent processing.

#### Syntax
```bash
loganalyzer analyze [flags]
```

#### Required Flags

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--config` | `-c` | Path to JSON configuration file | `--config config.json` |

#### Optional Flags

| Flag | Description | Default | Example |
|------|-------------|---------|---------|
| `--output` | `-o` | Export results to JSON file | None | `--output report.json` |
| `--status` | Filter results by status (OK/FAILED) | None | `--status FAILED` |
| `--timestamp` | Add timestamp to output filename | false | `--timestamp` |

#### Examples

**Basic Analysis:**
```bash
./loganalyzer analyze --config config.json
```

**Analysis with JSON Export:**
```bash
./loganalyzer analyze --config config.json --output reports/analysis.json
```

**Filter Failed Logs Only:**
```bash
./loganalyzer analyze --config config.json --status FAILED
```

**Timestamped Output:**
```bash
./loganalyzer analyze --config config.json --output report.json --timestamp
# Creates: 250127_report.json (for January 27, 2025)
```

**Complete Example with All Features:**
```bash
./loganalyzer analyze \
  --config config.json \
  --output reports/2024/analysis.json \
  --status OK \
  --timestamp
```

---

## 🔧 Configuration

### Configuration File Structure

The configuration file must be a valid JSON array containing log configuration objects:

```json
[
  {
    "id": "unique-identifier",
    "path": "path/to/logfile.log",
    "type": "log-type"
  }
]
```

#### Configuration Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | ✅ | Unique identifier for the log |
| `path` | string | ✅ | File path (absolute or relative) |
| `type` | string | ✅ | Log type classification |

#### Supported Log Types

- `nginx-access` - Nginx access logs
- `mysql-error` - MySQL error logs  
- `custom-app` - Custom application logs
- `generic` - Generic log format

### Sample Configuration Files

**Basic Configuration:**
```json
[
  {
    "id": "web-server",
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "database",
    "path": "/var/log/mysql/error.log", 
    "type": "mysql-error"
  }
]
```

**Development Configuration:**
```json
[
  {
    "id": "app-logs",
    "path": "logs/application.log",
    "type": "custom-app"
  },
  {
    "id": "debug-logs",
    "path": "logs/debug.log",
    "type": "generic"
  }
]
```

---

## 📊 Output Format

### Console Output

```
=== Log Analysis Results ===
ID: web-server-1
Path: test_logs/access.log
Status: OK
Message: Analyse terminée avec succès.
---
ID: invalid-path
Path: /non/existent/log.log
Status: FAILED
Message: Fichier introuvable.
Error: file not found: /non/existent/log.log
---

=== Summary ===
Total logs analyzed: 2
Successful: 1
Failed: 1

Failed logs breakdown:
- invalid-path: Fichier introuvable.
```

### JSON Output

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "file not found: /non/existent/log.log"
  }
]
```

---

## 💡 Examples

### Example 1: Basic Log Analysis

```bash
# Create test configuration
cat > basic_config.json << EOF
[
  {
    "id": "app-server",
    "path": "logs/app.log",
    "type": "custom-app"
  }
]
EOF

# Run analysis
./loganalyzer analyze --config basic_config.json
```

### Example 2: Batch Analysis with Export

```bash
# Analyze multiple logs and export to organized directory
./loganalyzer analyze \
  --config production_config.json \
  --output reports/prod/$(date +%Y%m%d)_analysis.json
```

### Example 3: Error Investigation

```bash
# Filter only failed analyses for troubleshooting
./loganalyzer analyze \
  --config config.json \
  --status FAILED \
  --output failed_logs.json
```

### Example 4: Daily Report Generation

```bash
# Generate timestamped daily report
./loganalyzer analyze \
  --config daily_config.json \
  --output reports/daily_report.json \
  --timestamp
```

---

## 📁 Project Structure

```
loganizer/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command definition
│   └── analyze.go         # Analyze command implementation
├── internal/              # Internal packages
│   ├── analyser/          # Core analysis logic
│   │   ├── analyzer.go    # Main analyzer with goroutines
│   │   └── errors.go      # Custom error types
│   ├── config/            # Configuration management
│   │   └── config.go      # JSON config loading
│   └── reporter/          # Result reporting
│       └── reporter.go    # Console & JSON output
├── test_logs/             # Sample log files
├── reports/               # Generated reports
├── config.json           # Sample configuration
├── main.go               # Application entry point
├── go.mod                # Go module definition
└── README.md             # This file
```

---

## ⚠️ Error Handling

### Custom Error Types

The application implements robust error handling with custom error types:

#### FileNotFoundError
- **Trigger**: When log files don't exist or are inaccessible
- **Message**: "Fichier introuvable."
- **Details**: Full system error message

#### ParsingError  
- **Trigger**: Simulated parsing failures (10% random chance)
- **Message**: "Erreur de parsing."
- **Details**: Specific parsing error information

### Error Detection Functions

```go
// Check error types
if analyzer.IsFileNotFoundError(err) {
    // Handle file not found
}

if analyzer.IsParsingError(err) {
    // Handle parsing error  
}

// Extract error details
if fnfErr, ok := analyzer.GetFileNotFoundError(err); ok {
    fmt.Printf("Missing file: %s\n", fnfErr.Path)
}
```

---

## 🏗️ Technical Details

### Concurrency Model

- **Goroutines**: One per log file for parallel processing
- **WaitGroup**: Synchronization of concurrent operations  
- **Channels**: Safe result collection from goroutines
- **Timing**: Random analysis duration (50-200ms) for realistic simulation

### Performance Characteristics

- **Scalability**: Handles hundreds of log files concurrently
- **Memory**: Efficient channel-based result collection
- **Error Rate**: 10% simulated parsing failures for testing
- **File I/O**: Minimal file system operations

---

## 🔍 Bonus Features

### ✅ Implemented Bonus Features

1. **📁 Automatic Directory Creation**
   - Creates output directories automatically
   - Uses `os.MkdirAll()` for nested paths

2. **🕒 Timestamped Output Files**
   - `--timestamp` flag adds YYMMDD prefix
   - Format: `250127_report.json`

3. **🔍 Status Filtering**
   - `--status` flag filters by OK/FAILED
   - Useful for error investigation

---

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Setup

```bash
# Clone your fork
git clone https://github.com/yourusername/loganizer.git
cd loganizer

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build and test
go build -o loganalyzer main.go
./loganalyzer analyze --config config.json
```

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 📞 Support

For support and questions:

- **GitHub Issues**: [Create an issue](https://github.com/Mathieu-ai/loganizer/issues)
- **Team Contact**: Reach out to Mathieu, Yoann, or Swann

---

**Made with ❤️ by the LogAnalyzer Team**
