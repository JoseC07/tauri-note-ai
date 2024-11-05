# Getting Started with Tauri Notes AI

## Prerequisites

Before you begin, ensure you have the following installed:
- Node.js (v16 or later)
- Go (v1.19 or later)
- Python (v3.8 or later)
- Rust (latest stable version)

### System-specific Requirements

#### Windows
- Visual Studio Build Tools
- Windows 10 or later
- WebView2

#### macOS
- Xcode Command Line Tools
- macOS 10.15 or later

#### Linux
- Development packages:
  ```bash
  # Ubuntu/Debian
  sudo apt install libwebkit2gtk-4.0-dev \
      build-essential \
      curl \
      wget \
      file \
      libssl-dev \
      libgtk-3-dev \
      libayatana-appindicator3-dev \
      librsvg2-dev
      python3-dev \ 
      portaudio19-dev
  ```

## Quick Start

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/tauri-notes-ai
   cd tauri-notes-ai
   ```

2. **Install Dependencies**
   ```bash
   # Install npm dependencies
   npm install

   # Install Python dependencies
   pip install -r src-tauri/python/requirements.txt

   # Install Go dependencies
   cd src-tauri/go-backend
   go mod download
   cd ../..
   ```

3. **Initialize Application**
   ```bash
   # Run the initialization script
   chmod +x src-tauri/init.sh
   ./src-tauri/init.sh
   ```

4. **Start Development Environment**
   ```bash
   # Start the Go backend
   cd src-tauri/go-backend
   go run main.go &
   cd ../..

   # Start the Tauri application
   npm run tauri dev
   ```

## Development Setup

### Environment Configuration

1. **Create Environment File**
   Create a `.env` file in the `src-tauri/go-backend` directory:
   ```env
   APP_ENV=development
   PORT=5000
   ```

2. **Configure Python Environment**
   ```bash
   # Create and activate virtual environment (optional but recommended)
   python -m venv venv
   source venv/bin/activate  # On Windows: .\venv\Scripts\activate
   
   # Install required packages
   pip install -r src-tauri/python/requirements.txt
   ```

### Directory Structure 