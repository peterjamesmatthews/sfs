package graph

import (
	"pjm.dev/sfs/graph/model"
)

type SharedFileSystemer interface {
	// GetNodeByID fetches a node by its id.
	//
	// # Arguments
	//  - user: The user who is fetching the node.
	//  - id: The id of the node to fetch.
	//
	// # Errors
	//  - `ErrNotFound` if the node is not found.
	//  - `ErrUnauthorized` if `user` does not have read access to the node.
	GetNodeByID(user model.User, id string) (model.Node, error)

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
	RenameNode(user model.User, id string, name string) (model.Node, error)

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
	MoveNode(user model.User, id string, dstID string) (model.Node, error)

	// GetRoot fetches the root folder of the shared file system.
	//
	// # Arguments
	//  - user: The user who is fetching the root folder.
	//
	// # Errors
	//  - `ErrNotFound` if the shared file system doesn't have a root.
	//  - `ErrUnauthorized` if the user does not have read access to the root.
	GetRoot(user model.User) (model.Folder, error)

	// InsertFolder inserts a folder into the shared file system.
	//
	// # Arguments
	//  - user: The user who is inserting the folder.
	//  - folder: The folder to insert.
	//
	// # Errors
	//  - `ErrNotFound` if `folder`'s parent is not found.
	//  - `ErrUnauthorized` if the user does not write access to `folder`'s parent.
	InsertFolder(user model.User, folder model.Folder) (model.Folder, error)

	// GetFolderbyID fetches a folder by its id.
	//
	// # Arguments
	//   - user: The user who is fetching the folder.
	//   - id: The id of the folder to fetch.
	//
	// # Errors
	//  - `ErrNotFound` if `folder`'s parent is not found.
	//  - `ErrUnauthorized` if the user does not have read access to the folder.
	GetFolderByID(user model.User, id string) (model.Folder, error)

	// InsertFile inserts a file into the shared file system.
	//
	// # Arguments
	//  - user: The user who is inserting the file.
	//  - file: The file to insert.
	//
	// # Errors
	//  - `ErrNotFound` if `file`'s parent is not found.
	//  - `ErrUnauthorized` if `user` does not have write access to `file`'s parent
	InsertFile(user model.User, file model.File) (model.File, error)

	// GetFileByID fetches a file by its id.
	//
	// # Arguments
	//  - user: The user who is fetching the file.
	//  - id: The id of the file to fetch.
	//
	// # Errors
	// - `ErrNotFound` if `file`'s parent is not found.
	// - `ErrUnauthorized` if `user` does not have read access to the file
	GetFileByID(user model.User, id string) (model.File, error)

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
	WriteFile(user model.User, fileID string, content string) (model.File, error)
}
