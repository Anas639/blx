# Changelog

## 0.4.0 - [2025-03-30]

### Added

- Implement `find` command to run a  full-text search on task and project names
- Implement `project find` to run a full-text search on project names

### Changed

- Skip DB initialization if the DB file exists to speed up startup.


## 0.3.0 - [2025-03-29]

### Added

- Assign task to an existing project during creation using the `--project` flag
- Automatically start a task after creating it using the `--start` flag
- Add status filter using the `--status` flag when listing tasks

### Fixed

- Fix nil reference error caused by udpListener using nil UDPConn 

## 0.2.0 - [2025-03-27]

### Added

- Introduce the `time` command to track the elapsed time in real-time
- Implement a UDP-based local event system
- Introduce the `watch` command to listen on task events locally (useful for third parties like Polybar Modules)

## 0.0.1 - [2025-03-25]

### Added
- Create task
- Update task name
- Assign task to project
- Start task
- Pause task
- End task
- List all tasks
- Delete task
- Create project
- Update project name 
- list all projects
- delete project

