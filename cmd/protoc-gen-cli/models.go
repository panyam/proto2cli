package main

/**
 * Our proto plugin generates a bunch of CLI commands either one for
 * a group of services or a group of methods or in any other order
 * the proto author chooses.
 *
 * These CLI models can be thought of as an intermediate format before
 * they are rendered into the final CLI code for particular language
 * or framework targets (eg python with click, golang with cobra etc)
 * via custom templates
 *
 * One of the goals for the CLIs is also to ensure that CLIs can be
 * extended by the user to either add their own custom commands or
 * perform other things like adding custom flags to existing commands
 * or even adding custom help text to existing commands.  One example
 * of this could be via including generated code in a user's larger
 * CLI codebase.
 */
type CLI struct {
	// The executable name of the CLI
	Name string

	// help text describing the CLI
	HelpText string

	// What commands does a CLI support
	// Commands here are hierarchical and can be used to
	// compose groups of commands
	Commands []*Command
}

/**
 * A CommandGroup is a group of commands that are related to each other and allow
 * subcommands to be added to them along with containing things like help strings
 * and other metadata.
 */
type Command struct {
	// (Short) Name of the command
	Name string

	// Title for the command
	Title string

	// Flags passed to the command
	Flags []*Flag

	// Sub commands supported by this command once all
	// args and flags have been processed
	SubCommands []*Command

	// Help text for the command
	HelpText string

	// Any help text that should be shown *before* describing the flags
	PreAmble []string

	// Notes/PostAmble to be shown after describing flags
	Notes []string
}

type Flag struct {
	// Name of the flag (eg --my-flag-name)
	Name string

	// (Short) Name of the flag (eg -m)
	ShortName string

	// Title of the flag
	Title string

	// Help text for the flag
	HelpText string

	// Is the flag required
	Required bool

	// Where the default values comes from
	// nil => no default value
	// env:<ENV_VAR> => default value comes from environment variable named ENV_VAR
	// file:<FILE_PATH> => default value comes from file at FILE_PATH
	// value:<VALUE> => default value is VALUE
	DefaultFrom *string
}
