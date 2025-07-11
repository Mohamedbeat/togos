# Togo - Terminal Todo App

Togo is a modern, interactive terminal-based todo list application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) in Go. It features a beautiful TUI, modal dialogs, keyboard navigation, and persistent storage.

## Features

- **Add, edit, and delete todos** with title and description
- **Mark todos as done/undone**
- **View creation and completion dates**
- **Centered modal dialogs** for add, edit, save, and delete actions
- **Table view** for todos with scrolling and keyboard navigation
- **Persistent storage** in your home directory (`~/togos.json`)
- **Fully keyboard-driven**

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/togo.git
   cd togo
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Run the app:**
   ```sh
   go run .
   ```

## Usage

- **Navigate** the list with <kbd>j</kbd>/<kbd>k</kbd> or <kbd>up</kbd>/<kbd>down</kbd>
- **Add a todo:** <kbd>a</kbd>
- **Edit a todo:** <kbd>e</kbd>
- **Delete a todo:** <kbd>d</kbd>
- **Mark as done/undone:** <kbd>space</kbd> or <kbd>enter</kbd>
- **Save todos:** <kbd>w</kbd> or <kbd>ctrl+w</kbd>
- **Quit:** <kbd>q</kbd> or <kbd>ctrl+c</kbd>
- **In modals:**
  - <kbd>esc</kbd> to cancel/close
  - <kbd>tab</kbd>/<kbd>shift+tab</kbd> or <kbd>up</kbd>/<kbd>down</kbd> to move between fields
  - <kbd>enter</kbd> to submit (on last field)

## Data Storage

- Todos are saved in a JSON file at `~/togos.json`.
- Data persists between runs.

## Development

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Bubbles](https://github.com/charmbracelet/bubbles)
- Modular code: each modal/view is in its own file for clarity
- To add features, edit the corresponding view/model file

## Screenshots



## License

MIT 