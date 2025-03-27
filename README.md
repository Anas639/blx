# BLX CLI

A simple and efficient CLI tool for tracking time spent on tasks and projects.

## Features

- Create, update, and delete tasks.
- Assign tasks to projects.
- Track task statuses: `new`, `ongoing`, `paused`, `ended`.
- Start, pause, and end tasks.
- Automatically log task sessions when pausing or ending a task.
- List tasks with filtering options.
- Manage projects (create, update, delete).

## Installation

Ensure you have Go installed, then run:

```sh
go install github.com/anas639/blx@latest
```

Or clone and build manually:

```sh
git clone https://github.com/anas639/blx.git
cd blx 
go build
go install
```

## Usage

### Task Commands

#### Create a new task
```sh
blx new "Task Name"
```

#### List all tasks
```sh
blx ls
```

**Filters:**

- `--all` – List all tasks. 

#### Start a task

```sh
blx start <task_id>
```

#### Pause a task

```sh
blx pause <task_id>
```

#### End a task

```sh
blx end <task_id>
```

#### Update a task

```sh
blx update <task_id> --name "New Task Name"
```

#### Assign a task to a project

```sh
blx update <task_id> --project <project_id>
```

#### Delete a task

```sh
blx delete <task_id>
```

#### Check Elapsed Time

The `time` command displays the elapsed time of an ongoing task.

If you don't you provide a `task_id`, the `time` command will fetch the last active task

```sh
blx time <task_id>
```



#### Watch tasks

The `watch` command listens for incoming task events and prints them.

When you start a task the `watch` command displays the elapsed time. If you pause or stop a task, it prints
the current status. If there's an older ongoing task, `watch` automatically switches to tracking that task.

```sh
blx watch 
```

> The `watch` command is different from the `time` command. `time` will only print elapsed time of a single task and won't listen on events


### Project Commands

#### Create a new project

```sh
blx project new "Project Name"
```

#### List all projects

```sh
blx project ls
```

#### Rename a project

```sh
blx project update <project_id> --name "New Project Name"
```

#### Delete a project

```sh
blx project delete <project_id>
```

## Example Workflow

1. Create a project:

    ```sh
    blx project new "Web Development"
    ```

2. Create a task and assign it to the project:

    ```sh
    blx new "Implement login system"
    blx update <task_id> --project <project_id>
    ```

3. Start tracking time:

    ```sh
    blx start <task_id>
    ```

4. Pause or end the task:

    ```sh
    blx pause <task_id>
    blx task end <task_id>
    ```

5. List tasks:

   ```sh
   blx ls
   ```

## License

MIT License

## Contributing

Feel free to submit issues and pull requests!

