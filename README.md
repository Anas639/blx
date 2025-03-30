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

#### Find Task

`find` searches for matches in the task name or project name.
Consider The following tasks:

```sh
$ blx ls -a      
+---+-----------------------------------------------+--------+----------+---------+
| # | NAME                                          | STATUS | DURATION | PROJECT |
+---+-----------------------------------------------+--------+----------+---------+
| 1 | this is an example                            | new    |       0s | N/A     |
| 2 | more examples                                 | new    |       0s | N/A     |
| 3 | this won't be selected                        | paused |      47s | N/A     |
| 4 | this will be selected because of project name | new    |       0s | example |
+---+-----------------------------------------------+--------+----------+---------+

```

We will find tasks that contain the keyword *"exam"* in the task name or project name.

```sh 

$ blx find "exam"
+---+-----------------------------------------------+--------+----------+---------+
| # | NAME                                          | STATUS | DURATION | PROJECT |
+---+-----------------------------------------------+--------+----------+---------+
| 1 | this is an example                            | new    |       0s | N/A     |
| 2 | more examples                                 | new    |       0s | N/A     |
| 4 | this will be selected because of project name | new    |       0s | example |
+---+-----------------------------------------------+--------+----------+---------+
```

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
    blx end <task_id>
    ```

5. List tasks:

   ```sh
   blx ls
   ```

By default `ls` will list ongoing tasks only. If you wish to list all the tasks 
use `--all` or `-a`.

   ```sh
   blx ls -a
   ```

You can also filter by statuses (new, ongoing, paused, ended) using `--status` or `-s`

   ```sh
   blx ls -s new # list new tasks only
   blx ls -s ongoing,paused # list ongoing and paused tasks
   ```

## Polybar Integration

To integrate *blx* into your *polybar* configuration:

### Add *blx* module to your `modules.ini`

Your `modules.ini` file is probably located at `~/.config/polybar/`

  ```ini 
  [module/blx]
  type = custom/script
  exec = ~/.config/polybar/scripts/blx.sh
  label-foreground = ${colors.white}
  label-background = ${colors.base}
  label=󱎫 %output%
  tail=true
  ```

### Create a custom script

  ```sh
  mkdir ~/.config/polybar/scripts 
  touch ~/.config/polybar/scripts/blx.sh
  chmod +x ~/.config/polybar/scripts/blx.sh
  ```


Then call `blx watch` from your custom script

  ```sh 
  #!/bin/zsh

  BLX="$HOME/go/bin/blx"
  $BLX watch 
  
  ```

### Add the *blx* module to your `config.ini`

Your `config.ini` is probably located at `~/.config/polybar/`

```ini 
modules-right = blx
```

*blx* will now appear on the right side of your bar

![Polybar No Task](./assets/polybar_no_task.png "No Task")

  ```sh 
  blx start <task_id>
  ```

![Polybar Task Started](./assets/polybar_task_start.png "Task Started")

  ```sh 
  blx pause <task_id>
  ```

![Polybar Task Paused](./assets/polybar_task_pause.png "Task Paused")

## License

MIT License

## Contributing

Feel free to submit issues and pull requests!

