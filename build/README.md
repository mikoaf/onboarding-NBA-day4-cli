# 🛠️ Project Configuration & Build

This directory contains essential setup files and configuration needed to build and run the project.

## 📁 Folder Structure

```bash
build/                  # Build-related scripts and assets (binary, packaging, etc.)
├── config/              # Centralized configuration directory
│   └── config.yaml      # Project-specific configuration file (YAML)
├── project                 # Project executable binary
└── README.md               # Documentation for the config and build setup
```

## 📦 Description

- `build/`: This folder typically contains build scripts, assets, and setup instructions for compiling or packaging the project for different environments.

- `config/config.yaml`: Used to store all configuration files. These can include environment variables, service endpoints, ports, hardware pin mapping, etc.

- `project`: A main binary file. You can rename it with a relevant name. Specify executable name in the `Makefile` file.

## ✅ Usage Notes

Make sure to define your config values inside `config/config.yaml` before building or running the app.

This folder is commonly ignored from version control (.gitignore) if it contains sensitive or environment-specific configurations.

## 💡 Tips

Use `NewConfig` struct from `config.go` to load a configuration file into your application.
