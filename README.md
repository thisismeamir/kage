# Project Kage

**K**ernel of **A**utonomous **G**raph-based **E**xecution

## Overview

Kage is a lightweight, modular automation system that enables users to create custom workflows and automations on their local machines. Unlike traditional system automation tools, Kage focuses on personal productivity workflows that can chain together APIs, AI services, file operations, and system commands seamlessly.

## Core Concepts

### Atoms
The fundamental building blocks of Kage. Atoms are:
- **Stateless**: Pure functions that produce consistent outputs given the same inputs
- **Self-contained**: Each atom defines its own interface and requirements
- **Composable**: Can be connected to other atoms through covalence bonds
- **Hot-reloadable**: Can be modified and updated without system restart

Examples of atoms:
- API fetchers
- LLM processors
- File readers/writers
- Git operations
- Logic gates (if-then, filters, switches)
- System monitors
- Event watchers

### Covalence Bonds (Mappings)
Connections between atoms that handle data transformation:
- **Pure data transformation**: Reshape, extract, and format data between atoms
- **No logic**: Complex conditional logic belongs in dedicated logic atoms
- **Custom per connection**: Each bond can have unique mapping rules

### Modules (Proteins)
Complete workflows composed of multiple atoms connected via covalence bonds:
- **Self-contained**: Include all necessary atoms and their connections
- **Shareable**: Can be exported as `.kmodule` files
- **Editable**: Can be modified and customized after import

## Architecture

### Resource Management
- **Lazy Loading**: Atoms are only loaded when their inputs are ready
- **Performance Awareness**: System resource monitoring enables intelligent scheduling
- **Event-Driven**: Lightweight watchers trigger heavier processing only when needed

### Execution Models
- **On-Demand**: Modules triggered by events (file changes, webhooks, etc.)
- **Scheduled**: Modules running on timers or schedules
- **Conditional**: Modules that only run when system resources allow

### File System
- **`.katom`**: Self-contained atom definitions for sharing individual components
- **`.kmodule`**: Complete workflow packages with internal dependencies
- **Configuration**: Human-readable config files define module behavior

## User Interface

### Visual Editor
- **Localhost Frontend**: Web-based drag-and-drop interface
- **Node-based**: Visual representation of atoms and their connections
- **Real-time**: Hot reload capabilities for immediate feedback
- **Atom Creation**: Built-in tools for defining new atom types

### Dashboard
- **Module Status**: View active, idle, and scheduled modules
- **System Monitoring**: Track resource usage and system health
- **Event Logs**: History of triggers, executions, and errors
- **Performance Metrics**: Insights into module efficiency and resource usage

## Technical Stack

### Backend
- **Language**: Go
- **Architecture**: Monolithic for simplicity and lightweight deployment
- **File Watching**: Built-in `fsnotify` for configuration changes
- **Concurrency**: Goroutines for module execution management

### Frontend
- **Framework**: React with React Flow for visual pipeline editing
- **Embedding**: SPA served by Go's built-in HTTP server
- **Interactivity**: Real-time updates for module status and system monitoring

## Example Workflows

### Automated Content Pipeline
```
API Fetch → LLM Processing → Document Generation → File Save → Git Commit → Git Push
```

### Smart Email Response
```
Email Watcher → Email Parser → Context Analysis → LLM Response → Email Sender
```

### System Maintenance
```
System Monitor → Performance Check → Cleanup Tasks → Report Generation → Notification
```

## Design Principles

### Modularity
Every component is designed to be replaceable and extensible. No hardcoded primitives force users into specific patterns.

### Resource Efficiency
Kage respects system resources and user workflow, only consuming CPU and memory when actually needed.

### Flexibility
Users can define any automation they need without being constrained by predetermined use cases or interfaces.

### Shareability
The `.katom` and `.kmodule` ecosystem enables community-driven development and sharing of automation components.

## Getting Started

1. **Install Kage**: Download and run the lightweight Go binary
2. **Access Interface**: Open the localhost web interface
3. **Import Primitives**: Load starter pack of basic `.katom` files
4. **Build Workflows**: Use the visual editor to create your first module
5. **Share & Extend**: Export your creations and import community modules

## Open Source Philosophy

Kage embraces open-source principles:
- **Transparency**: All modules are inspectable and editable
- **Community**: Share and improve workflows through the `.katom`/`.kmodule` ecosystem
- **Flexibility**: No vendor lock-in or prescribed ways of working
- **Extensibility**: Every user can contribute new atoms and workflows