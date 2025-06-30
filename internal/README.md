# internal/

This directory contains the core internal logic of the application. It is structured by responsibilities, following a modular design to promote separation of concerns and maintainability.

## 📁 Folder Structure

```
internal/
├── api/                  # Handles API interfaces (client & server)
│   ├── client/           # Outgoing API calls to external services
│   ├── server/           # Incoming API logic, handlers, routes
│   └── api.go            # Shared API types, structs, or setup
│
├── config/               # Configuration loading and parsing
│   └── config.go         # Handles app configuration (e.g., from YAML/ENV)
│
├── hardware/             # Low-level hardware or device interaction logic
│   └── hardware.go       # Interacts with hardware interfaces
│
├── models/               # Core data models and domain entities
│   └── models.go         # Structs used throughout the app
├── sensor/               # Handles sensor data acquisition
│   └── sysinfo.go        # Sensor data acquisition function
│
└── utils/                # Utility helpers and shared functions
```
