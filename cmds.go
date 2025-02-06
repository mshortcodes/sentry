package main

type cmds map[string]cmd

type cmd struct {
	name        string
	description string
	callback    func(*state) error
}

func getCmds() cmds {
	return cmds{
		"add": {
			name:        "add",
			description: "add a new password",
			callback:    cmdAdd,
		},
		"create": {
			name:        "create",
			description: "create a new user",
			callback:    cmdCreate,
		},
		"delete": {
			name:        "delete",
			description: "delete a password",
			callback:    cmdDelete,
		},
		"edit": {
			name:        "edit",
			description: "edit a password and its name",
			callback:    cmdEdit,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    cmdExit,
		},
		"get": {
			name:        "get",
			description: "get passwords",
			callback:    cmdGet,
		},
		"help": {
			name:        "help",
			description: "list available commands",
			callback:    cmdHelp,
		},
		"login": {
			name:        "login",
			description: "log a user in",
			callback:    cmdLogin,
		},
		"logout": {
			name:        "logout",
			description: "log a user out",
			callback:    cmdLogout,
		},
		"reset": {
			name:        "reset",
			description: "delete all users and passwords",
			callback:    cmdReset,
		},
		"wipe": {
			name:        "wipe",
			description: "delete all passwords from the current user",
			callback:    cmdWipe,
		},
	}
}
