package graph

type SharedFileSystemer interface {
	// CreateUser creates a new user.
	//
	// # Arguments
	//  - name: The name of the user to create.
	//
	// # Errors
	//	- `ErrConflict` if a user with the same name already exists.
	CreateUser(name string) (User, error)

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

	// GetFolderbyID fetches a folder by its id.
	//
	// # Arguments
	//   - user: The user who is fetching the folder.
	//   - id: The id of the folder to fetch.
	//
	// # Errors
	//  - `ErrNotFound` if `folder`'s parent is not found.
	//  - `ErrUnauthorized` if the user does not have read access to the folder.
	GetFolderByID(user User, id string) (Folder, error)

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

	GetUserByID(id string) (User, error)
}
