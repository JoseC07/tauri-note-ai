# Stage 1: Build the React front-end
FROM node:16-alpine AS frontend-builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the source code
COPY . .

# Build the React app
RUN npm run build

# Stage 2: Build the Go back-end
FROM golang:1.18-alpine AS backend-builder

# Install build dependencies
RUN apk add --no-cache \
    gcc \
    musl-dev \
    pkgconfig \
    portaudio-dev

# Set working directory
WORKDIR /app

# Create proper Go module structure
RUN mkdir -p tauri-notes-ai

# Copy Go source code maintaining module structure
COPY src-tauri/go-backend/ tauri-notes-ai/

# Build the Go application with static linking
WORKDIR /app/tauri-notes-ai
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-linkmode=external -extldflags=-static" -o ../go-backend

# Stage 3: Final stage - create a minimal image
FROM python:3.9-alpine

# Install necessary packages
RUN apk add --no-cache \
    ca-certificates \
    portaudio \
    gcc \
    musl-dev \
    python3-dev \
    linux-headers \
    openblas-dev

# Set working directory
WORKDIR /app

# Create and activate virtual environment
RUN python -m venv /app/venv
ENV PATH="/app/venv/bin:$PATH"

# Upgrade pip and install build tools
RUN pip install --no-cache-dir --upgrade pip wheel setuptools

# Install Python dependencies in smaller groups to better handle failures
RUN pip install --no-cache-dir \
    numpy>=1.24.3 \
    scipy>=1.11.3 \
    tqdm>=4.66.1 \
    python-dotenv>=1.0.0

# Install PyTorch and related packages
RUN pip install --no-cache-dir \
    --extra-index-url https://download.pytorch.org/whl/cpu \
    torch>=2.1.0 \
    torchaudio>=2.1.0 \
    torchvision>=0.16.0

# Install remaining packages
RUN pip install --no-cache-dir \
    transformers>=4.36.0 \
    sounddevice>=0.4.6 \
    librosa>=0.10.1 \
    nltk>=3.8.1

# Copy Python scripts
COPY src-tauri/python/ ./python/

# Copy the front-end build from Stage 1
COPY --from=frontend-builder /app/dist ./frontend

# Copy the Go binary from Stage 2
COPY --from=backend-builder /app/go-backend .

# Expose ports if necessary
EXPOSE 8080

# Create entrypoint script
RUN echo '#!/bin/sh' > /entrypoint.sh && \
    echo 'source /app/venv/bin/activate' >> /entrypoint.sh && \
    echo 'exec "$@"' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["./go-backend"]