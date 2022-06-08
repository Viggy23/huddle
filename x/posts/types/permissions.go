package types

// DONTCOVER

import (
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var (
	// PermissionWrite identifies users that can create content inside the subspace
	PermissionWrite = subspacestypes.RegisterPermission("write content")

	// PermissionInteractWithContent allows users to interact with content inside the subspace (eg. polls)
	PermissionInteractWithContent = subspacestypes.RegisterPermission("interact with content")

	// PermissionEditOwnContent allows users to edit their own content inside the subspace
	PermissionEditOwnContent = subspacestypes.RegisterPermission("edit own content")

	// PermissionModerateContent allows users to moderate other user's content
	PermissionModerateContent = subspacestypes.RegisterPermission("moderate content")
)