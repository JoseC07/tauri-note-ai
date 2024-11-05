# Tauri Notes AI - Architecture Documentation

## System Architecture Overview

### Core Components
1. **Frontend (React + Tauri)**
   - React for UI components
   - Tauri for native desktop capabilities
   - Communicates with Go backend via HTTP

2. **Backend (Go)**
   - RESTful API server
   - Handles data persistence
   - Manages audio processing
   - Coordinates with AI processing

3. **AI Processing (Python)**
   - Text summarization
   - Classification
   - Speech-to-text processing

## Technology Choices & Alternatives

### Frontend Framework
**Current Choice: React + Tauri**
- Pros:
  - Strong ecosystem
  - Cross-platform capabilities
  - Native performance
  - Secure by default
- Alternatives:
  - Electron (Rejected due to higher resource usage)
  - Qt (Rejected due to steeper learning curve)
  - Flutter (Viable alternative, but less desktop-focused)

### Backend Framework
**Current Choice: Gin (Go)**
- Pros:
  - High performance
  - Simple routing
  - Good middleware support
  - Built-in input validation
- Alternatives:
  - Echo (Similar features, less community support)
  - Chi (More minimal, but requires more manual setup)
  - Fiber (Fast but newer, less battle-tested)

### Database
**Current Choice: SQLite + GORM**
- Pros:
  - Self-contained
  - Zero configuration
  - Perfect for desktop apps
  - GORM provides clean ORM interface
- Alternatives:
  - Embedded PostgreSQL (More features but more complex)
  - LevelDB (Key-value store, simpler but less querying capability)
  - Badger (Pure Go but less SQL-like interface)

### Audio Processing
**Current Choice: PortAudio**
- Pros:
  - Cross-platform
  - Low-level control
  - Mature library
- Alternatives:
  - SDL Audio (More complex, game-focused)
  - ALSA (Linux-only)
  - WebRTC (Browser-based, more complex)

## Data Flow

### Note Creation Flow
1. User creates note in UI
2. React component sends HTTP POST to Go backend
3. Go backend:
   - Validates input
   - Stores in SQLite
   - Triggers AI processing if needed
4. Response returned to UI

### Audio Processing Flow
1. Audio captured via PortAudio
2. Streamed to Go backend
3. Go backend:
   - Saves audio file
   - Triggers Python processing
   - Updates note with transcription
4. UI updated with results

## Design Decisions

### Why Separate Go Backend?
1. **Separation of Concerns**
   - UI logic stays in frontend
   - Data processing in dedicated service
   - Easier testing and maintenance

2. **Performance**
   - Go excels at concurrent operations
   - Efficient database operations
   - Good for audio processing

3. **Future Scalability**
   - Could become network service
   - Easy to add new features
   - Could support multiple frontends

### Why SQLite?
1. **Simplicity**
   - No separate server
   - Easy backups
   - Simple deployment

2. **Performance**
   - Fast for single-user scenarios
   - Good for document storage
   - Reliable ACID compliance

## Potential Future Improvements

### Short Term
1. **Caching Layer**
   - Add Redis for frequently accessed notes
   - Cache AI processing results

2. **Better Error Handling**
   - Structured error responses
   - Retry mechanisms for AI processing
   - Better error recovery

### Long Term
1. **Cloud Sync**
   - Optional remote backup
   - Multi-device sync
   - Collaborative features

2. **Advanced AI Features**
   - Real-time suggestions
   - Better categorization
   - Custom AI models

## Security Considerations

1. **Data Storage**
   - SQLite file encryption
   - Secure file permissions
   - Safe storage locations

2. **API Security**
   - Input validation
   - Rate limiting
   - CORS configuration

3. **Audio Data**
   - Secure storage of recordings
   - Proper cleanup of temporary files
   - Permission handling

## Development Guidelines

### Code Organization
- Follow standard Go project layout
- Keep controllers thin
- Use services for business logic
- Separate interface from implementation

### Testing Strategy
- Unit tests for business logic
- Integration tests for API endpoints
- E2E tests for critical flows
- Separate test database

### Error Handling
- Use custom error types
- Consistent error responses
- Proper logging
- Recovery middleware

## Monitoring and Debugging

### Logging
- Use structured logging (logrus)
- Different log levels
- Rotation policy
- Debug information

### Performance Monitoring
- Database query monitoring
- API response times
- Resource usage tracking
- Error rate monitoring

## Build and Deployment

### Development
- Hot reload for frontend
- Watch mode for Go
- Local development environment

### Production
- Binary compilation
- Resource bundling
- Update mechanism
- Installation process

## Known Technical Debt

1. **Audio Processing**
   - Basic implementation
   - Need better error handling
   - Limited format support

2. **Database**
   - Simple schema
   - No migrations yet
   - Basic indexing

3. **AI Integration**
   - Basic Python integration
   - No model management
   - Limited error handling

## References

### Documentation
- [Tauri Docs](https://tauri.app/docs/get-started/intro)
- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Guide](https://gorm.io/docs/)
- [PortAudio Documentation](http://www.portaudio.com/docs/)

### Best Practices
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [React Best Practices](https://reactjs.org/docs/thinking-in-react.html)
- [SQLite Performance](https://www.sqlite.org/performance.html) 