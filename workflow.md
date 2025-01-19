User Registration Flow:

User provides email or phone
System checks for existing user in Redis
Creates new user record if none exists
Sends verification OTP
Logs registration event

Login Initiation:

User requests login
System performs rate limit check
Generates and stores OTP
Sends OTP via Twilio
Logs login attempt

OTP Verification:

User submits OTP
System validates OTP against Redis
Manages attempt counting
Creates session on success
Handles blocking on too many failures

Session Management:

Validates active sessions
Updates last activity
Handles session expiration
Logs session activities

Rate Limiting:

Checks limits for each action
Manages counters in Redis
Blocks excessive attempts
Logs rate limit events

## Security

Security Measures:

JWT validation on every request
Multi-level rate limiting (IP, User, Action)
Session validation and updates
Account blocking after max attempts
OTP expiration and attempt tracking
Input validation and sanitization

Error Handling:

HTTP 400: Bad Request (Invalid input)
HTTP 401: Unauthorized (Invalid credentials)
HTTP 403: Forbidden (Blocked account)
HTTP 404: Not Found (User doesn't exist)
HTTP 409: Conflict (Duplicate registration)
HTTP 429: Too Many Requests (Rate limit)
HTTP 500: Internal Server Error (System errors)

Data Flow Between Services:

Redis → User data, sessions, OTPs, rate limits
Elasticsearch → Logs, metrics, errors, audit trails
Twilio → Email and SMS communication
Application → Orchestrates all services

Monitoring Points:

Authentication attempts
OTP delivery status
Rate limit triggers
Session activities
Error occurrences
Security events
