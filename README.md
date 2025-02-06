# Sentry 🛡️

## Table of Contents

- [About](#about)
- [Commands](#commands)
  - [Add](#add)
  - [Create](#create)
  - [Delete](#delete)
  - [Edit](#edit)
  - [Exit](#exit)
  - [Get](#get)
  - [Help](#help)
  - [Login](#login)
  - [Logout](#logout)
  - [Reset](#reset)
  - [Wipe](#wipe)
- [Installation](#installation)
  - [Database Setup](#database-setup)

## About

Sentry is a terminal-based password manager. I enjoy CLI tools and working in the terminal in general. I also use password managers regularly and recently completed the cryptography course on Boot.dev so I thought this would be a good project to work on.

Sentry is a fast REPL, uses a SQLite database, supports multiple users, encrypts passwords, and implements a cache. This was the first project I had done completely in Neovim which was an additional fun challenge.

Key concepts:

- Encryption
- Hashing
- Authentication
- Caching
- Writing SQL queries
- CRUD operations
- Input parsing/validation

## Commands

### add 🔒

Add a new password.

![add](images/add.png)

---

### `create` 🔓

Create a new user.

![create](images/create.png)

---

### `delete` 🔒

Delete a password.

![delete](images/delete.png)

---

### `edit` 🔒

Edit a password and its name.

![edit](images/edit.png)

---

### `exit` 🔓

Exit the program.

![exit](images/exit.png)

---

### `get` 🔓

Get passwords.

![get](images/get.png)

---

### `help` 🔓

List available commands.

![help](images/help.png)

---

### `login` 🔓

Log a user in.

![login](images/login.png)

---

### `logout` 🔒

Log a user out.

![logout](images/logout.png)

---

### `reset` 🔓

---

### `wipe` 🔒

Wipe all passwords from current user.

![wipe](images/wipe.png)

---

## Installation

Sentry only works on Linux/Mac and requires Golang. If on Windows, use WSL.

1. Install Go 1.22 or later

```bash
curl -sS https://webi.sh/golang | sh
```

2. Install Sentry

```bash
go install github.com/mshortcodes/sentry
```

### Database Setup
