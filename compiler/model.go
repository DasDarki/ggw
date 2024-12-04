package compiler

// Module represents a module.
type Module struct {
	Namespace string   // The global namespace name of the module. This is the "path" in import statements.
	Name      string   // The name of the module.
	Scripts   []Script // The scripts.
}

// Script represents a complete chunk of code.
type Script struct {
	Module      string   // The module name. If not specified, the string is empty.
	Imports     []Import // The imports.
	AvailableOn Side     // The side where the script is available.
}

// Import represents an import statement.
type Import struct {
	Location Location // The location of the import statement in the source code.
	Module   string   // The module name.
	Alias    string   // The alias name. If not specified, the string is empty.
	Path     string   // The path to the module.
}

// VariableKind represents the kind of a variable.
type VariableKind string

const (
	// Const represents a constant variable. The variable is immutable.
	Const VariableKind = "const"
	// Let represents a variable that can be reassigned.
	Let VariableKind = "let"
)

type Variable struct {
	Location Location     // The location of the variable in the source code.
	Name     string       // The name of the variable.
	Kind     VariableKind // The kind of the variable.
}

// Side represents the side of the web app.
type Side string

const (
	// Frontend represents the client/browser side.
	Frontend Side = "frontend"
	// Backend represents the server side.
	Backend Side = "backend"
	// Shared represents the shared side. Which means that the code can be used in both sides - backend and frontend.
	Shared Side = "shared"
)

// WireProtocol represents a wire protocol.
type WireProtocol string

const (
	// HTTP represents the HTTP wire protocol. The default wire protocol.
	HTTP WireProtocol = "http"
	// WebSocket represents the WebSocket wire protocol.
	WebSocket WireProtocol = "websocket"
)

// Wire defines a wiring instruction.
type Wire struct {
	Protocol WireProtocol // The wire protocol.
	IsPublic bool         // Whether the wire is public or not. If not, the call must be done authenticated.
	Roles    []string     // The roles that can access the wire. The user must have ALL the roles to access the wire.
}

// GrantStatement is the statement that grants access to a wire.
type GrantStatement struct {
	Location Location // The location of the grant statement in the source code.
	Roles    []string // The roles that are given the session when grant is called.
}
