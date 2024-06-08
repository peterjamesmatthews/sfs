package graph

type SharedFileSystemer interface {
	// Authenticate determines the requesting user.
	//
	// # Arguments
	//	- string: The Authorization header from the request.
	//
	// # Returns
	//  - The user who is making the request.
	//
	// # Errors
	//  - `ErrUnauthorized` if an authenticated user cannot be determined.
	Authenticate(string) (User, error)

	// CreateUser creates a new user.
	//
	// # Arguments
	//  - name: The name of the user to create.
	//
	// # Errors
	//	- `ErrConflict` if a user with the same name already exists.
	CreateUser(name string, password string) (User, error)

	// GetTokens fetches a user's access and refresh tokens, creating them if necessary.
	//
	// # Arguments
	//  - name: The name of the user to fetch tokens for.
	//  - password: The password of the user.
	//
	// # Errors
	//	- `ErrUnauthorized` if the user is not found or the password is incorrect.
	GetTokens(name string, password string) (Tokens, error)

	// RefreshTokens refreshes a user's access and refresh tokens.
	//
	// # Arguments
	//  - refresh: The user's refresh token.
	//
	// # Errors
	//	- `ErrForbidden` if the refresh token is expired.
	RefreshTokens(refresh string) (Tokens, error)

	// CreateFolder creates a new folder.
	//
	// # Arguments:
	//  - user: The user who is creating the folder.
	//  - parentID: The id of the parent folder.
	//  - name: The name of the folder to create.
	//
	// # Errors
	//	- `ErrNotFound` if the parent folder is not found.
	//	- `ErrUnauthorized` if the user does not have write access to the parent folder.
	//	- `ErrConflict` if a folder owned by `user` named `name` already exists in the parent folder.
	CreateFolder(user User, parentID *string, name string) (Folder, error)

	// GetNodeByURI fetches a node by its uri.
	//
	// # Arguments
	//  - user: The user who is fetching the node.
	//  - uri: a '/' separated path of node names.
	//
	// # Errors
	//  - `ErrNotFound` if the node is not found.
	//  - `ErrUnauthorized` if `user` does not have read access to any of the nodes in the uri.
	GetNodeByURI(user User, uri string) (Node, error)

	// RenameNode renames a node.
	//
	// # Arguments
	//  - user: The user who is renaming the node.
	//  - id: The id of the node to rename.
	//  - name: The new name of the node.
	//
	// # Errors
	//  - `ErrNotFound` if the node is not found.
	//  - `ErrUnauthorized` if `user` does not have write access to the node.
	RenameNode(user User, id string, name string) (Node, error)

	// MoveNode moves a node to a new parent.
	//
	// # Arguments
	//  - user: The user who is moving the node.
	//  - id: The id of the node to move.
	//  - dstID: The id of the new parent.
	//
	// # Errors
	//  - `ErrNotFound` if the node is not found.
	//  - `ErrUnauthorized` if the user does not own the node or if the user does not have write access to the destination parent.
	MoveNode(user User, id string, dstID *string) (Node, error)

	// GetRoot fetches the root folder of the shared file system.
	//
	// # Arguments
	//  - user: The user who is fetching the root folder.
	GetRoot(user User) (Folder, error)

	// InsertFolder inserts a folder into the shared file system.
	//
	// # Arguments
	//  - user: The user who is inserting the folder.
	//  - folder: The folder to insert.
	//
	// # Errors
	//  - `ErrNotFound` if `folder`'s parent is not found.
	//  - `ErrUnauthorized` if the user does not write access to `folder`'s parent.
	InsertFolder(user User, folder Folder) (Folder, error)

	// InsertFile inserts a file into the shared file system.
	//
	// # Arguments
	//  - user: The user who is inserting the file.
	//  - file: The file to insert.
	//
	// # Errors
	//  - `ErrNotFound` if `file`'s parent is not found.
	//  - `ErrUnauthorized` if `user` does not have write access to `file`'s parent
	InsertFile(user User, file File) (File, error)

	// WriteFile writes content to a file.
	//
	// # Arguments
	//  - user: The user who is writing to the file.
	//  - fileID: The id of the file to write to.
	//  - content: The content to write to the file.
	//
	// # Errors
	// - `ErrNotFound` if `file`'s parent is not found.
	// - `ErrUnauthorized` if `user` does not have write access to the file
	WriteFile(user User, fileID string, content string) (File, error)
}
