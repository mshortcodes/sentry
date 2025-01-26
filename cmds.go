package main

type cmds map[string]cmd

type cmd struct {
	name        string
	description string
	callback    func(*state) error
}

func getCmds() cmds {
	return cmds{
		"help": {
			name:        "help",
			description: "lists available commands",
			callback:    cmdHelp,
		},
		"add": {
			name:        "add",
			description: "adds a new password",
			callback:    cmdAdd,
		},
		"create": {
			name:        "create",
			description: "creates a new user",
			callback:    cmdCreate,
		},
		"get": {
			name:        "get",
			description: "retrieves passwords",
			callback:    cmdGet,
		},
		"login": {
			name:        "login",
			description: "logs a user in",
			callback:    cmdLogin,
		},
		"reset": {
			name:        "reset",
			description: "resets the database",
			callback:    cmdReset,
		},
	}
}
