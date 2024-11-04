// Package constants is a package for handling constants.
package constants

import "errors"

// Defines errors related to the database.
var (
	ErrDBisOpen = errors.New("db open error")
	ErrNewDB    = errors.New("new db error")
)

// Defines errors related to the redis.
var (
	ErrRedisOpen = errors.New("redis open error")
	ErrNewRedis  = errors.New("new redis error")
)

// Defines errors related to the elastic.
var (
	ErrElasticCountExist = errors.New("elastic count exist")
	ErrElasticCountAll   = errors.New("elastic count all")
)

// Defines errors related to the transaction in database.
var (
	ErrBegin     = errors.New("begin tx error")
	ErrCommit    = errors.New("commit error")
	ErrRollback  = errors.New("rollback error")
	ErrExecute   = errors.New("execute error")
	ErrRowResult = errors.New("row result error")
	ErrQuery     = errors.New("query error")
)

// Defines errors related to the repository.
var (
	ErrItemNotFound      = errors.New("item not found")
	ErrRepositoryUnknown = errors.New("repository unknown")
)

// Defines errors related to the user service.
var (
	ErrUserIDDuplicate = errors.New("user id duplicate")
	ErrUserCreate      = errors.New("create user error")
	ErrUserUpdate      = errors.New("update user error")
	ErrUserDelete      = errors.New("delete user error")
	ErrUserSignIn      = errors.New("user sign in error")
	ErrUserRead        = errors.New("read user error")
	ErrUsersList       = errors.New("user list error")
	ErrUserToken       = errors.New("user token error")
)

// Defines errors related to the token.
var (
	ErrTokenSaveRefresh  = errors.New("save refresh token error")
	ErrTokenCreate       = errors.New("create token error")
	ErrTokenRead         = errors.New("read token error")
	ErrTokenValidate     = errors.New("validate token error")
	ErrTokenParse        = errors.New("parse token error")
	ErrTokenRotation     = errors.New("rotation token error")
	ErrTokenDoesNotMatch = errors.New("refresh token does not match")
	ErrTokenDelete       = errors.New("delete token error")
	ErrTokenRoleMissing  = errors.New("role missing")
	ErrTokenRoleAuth     = errors.New("role auth error")
)

// Defines errors related to the proof service.
var (
	ErrProofCreate          = errors.New("create proof error")
	ErrProofUpdate          = errors.New("update proof error")
	ErrProofDelete          = errors.New("delete proof error")
	ErrProofUpload          = errors.New("upload proof error")
	ErrProofRead            = errors.New("read proof error")
	ErrProofList            = errors.New("list proof error")
	ErrProofReadFirstImage  = errors.New("read first image error")
	ErrProofReadSecondImage = errors.New("read second image error")
	ErrProofReadLog         = errors.New("read log error")
	ErrProofConfirm         = errors.New("confirm proof error")
	ErrProofUpdateConfirm   = errors.New("confirm update proof error")
)

// Defines errors related to the dashboard service.
var (
	ErrDashboardRead    = errors.New("dashboard read error")
	ErrDashboardUnknown = errors.New("unknown error")
)

// Defines errors related to the file.
var (
	ErrFileSave          = errors.New("save file error")
	ErrFilePath          = errors.New("file path error")
	ErrFilePathTraversal = errors.New("file path traversal error")
)
