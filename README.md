# README

## 📌 About

This CLI monitor shows you about current temperature in your city, CPU usage, RAM usage (with total RAM in bytes), and disk usage (with total disk space in bytes) on your computer. If you'd like to contribute on this project, feel free to contact me 😊.

### Structure

```
cmd/
  app/
    app.go          # Main application struct
    main.go         # The entry point of the application
    README.md       # Documentation for this CLI app layer
internal/           # Internal package library
  api/              # API Client and Server
    client/         # API Client logic
    server/         # API Server logic
    api.go          # API Client and Server wrapper (optional)
  config/           # Application configuration library
  env/              # Environment variable parser
  hardware/         # Hardware library interface
  models/           # Data structure models for communication
  sensor/           # Handles sensor data acquisition
  services/         # Application service routine
  utils/            # Package available for all services
services/           # Application business logic and use cases
  controller.go     # High-level orchestration (controller layer)
  service.go        # Service layer (contains core app logic)
Makefile            # Makefile build automation
README.md           # You're here!
```

### Purpose

This folder is intended to provide a CLI structure for launching the golang application. It follows the Go convention for placing app entry points in `cmd/{appname}` folders, keeping the core logic modular and reusable.

### `main.go`

- Initializes necessary services.
- Parses the configuration and/or environment variables.
- Starts the core backend logic using the packages from `internal/`.

### Usage

Install make on your computer. Run the app from the root directory:

```bash
make run
```

build a binary:

```bash
make build
```

Or build and run a binary:

```bash
make build-run
```
This project still running on localhost server. After you run this project, open http://localhost:"YOUR_PORT" on your browser. Add `/weather` to check about temperature in your city, `/sys/cpu` for CPU usage, `/sys/memory` for RAM usage (with total RAM in bytes), and `/sys/disk` for disk usage (with total disk space in bytes).

### Environment Variables

You may define environment-specific settings in a `.env` file or `config.yaml` (not included in version control) in `build/config`. Choose your preferred one.

---

For further details, refer to the root-level `README.md` or relevant documentation inside the `internal/` directory.

## Application Configuration

Create your own config file at:

> ./config/config.yaml

Note: The config/ folder is ignored via .gitignore, so your local configs remain private.

🔑 Example `config/config.yaml`

```yaml
sub_config_a:
  field_a: "http://localhost:8080"
  field_b: 100
  field_c: 1.23
  field_d: true

sub_config_b:
  field_a: "http://localhost:8081"
  field_b: 200
  field_c: 2.34
  field_d: false

sub_config_c:
  field_a: "http://localhost:8082"
  field_b: 300
  field_c: 3.45
  field_d: true
```

You can define multiple sub-configs and load them using custom Go logic under `internal/config/config.go`.

## 👨‍💻 Author

##### 📧 mikoaf