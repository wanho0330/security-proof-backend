// Package file is a package for handling file processes.
package file

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Netflix/go-env"

	"security-proof/pkg/constants"
)

// fileManage struct is composed of a path.
type fileManager struct {
	Path string `env:"SECURITY_PROOF_FILE_PATH,default=/Users/wanho/GolandProjects/security-proof/savedproofs"`
}

// SaveFile function is returning a save file path and an error, accepting a file name and a data.
func SaveFile(fileName string, data []byte) (string, error) {
	config := fileManager{}
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return "", errors.Join(constants.ErrFileSave, err)
	}

	saveFilePath := filepath.Join(config.Path, fileName)
	safeFilePath, err := isSafePath(config.Path, saveFilePath)
	if err != nil {
		return "", err
	}

	file, err := os.Create(safeFilePath) // #nosec G304
	if err != nil {
		return "", errors.Join(constants.ErrFileSave, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Failed to close file: %v\n", closeErr)
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		return "", errors.Join(constants.ErrFileSave, err)
	}

	return saveFilePath, nil
}

// isSafePath function is returning a safe path and an error, accepting a base path and a requested path.
func isSafePath(basePath string, requestedPath string) (string, error) {
	absBasePath, err := filepath.Abs(filepath.Clean(basePath))
	if err != nil {
		return "", errors.Join(constants.ErrFilePath, err)
	}

	absRequestedPath, err := filepath.Abs(filepath.Clean(requestedPath))
	if err != nil {
		return "", errors.Join(constants.ErrFilePath, err)
	}

	if !strings.HasPrefix(absRequestedPath, absBasePath) {
		return "", errors.Join(constants.ErrFilePath, constants.ErrFilePathTraversal)
	}

	return absRequestedPath, nil
}

// ImageToHash function is returning a hex string and an error, accepting a file path string pointer.
// Does not return an error if the image path does not exist.
func ImageToHash(filePath *string) (string, error) {
	if filePath == nil {
		return "", nil
	}

	image, err := os.ReadFile(*filePath)
	if err != nil {
		return "", errors.Join(constants.ErrProofUpload, err)
	}

	hashHex := sha256.Sum256(image)
	return hex.EncodeToString(hashHex[:]), nil
}
