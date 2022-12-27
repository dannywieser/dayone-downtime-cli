package export

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

const exportWorkingDir = "output"

func Unzip(archivePath string) error {
	archive, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, archiveFile := range archive.File {
		err := processArchiveFile(archiveFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func processArchiveFile(archiveFile *zip.File) error {
	filePath := filepath.Join(exportWorkingDir, archiveFile.Name)

	if archiveFile.FileInfo().IsDir() {
		os.MkdirAll(filePath, os.ModePerm)
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, archiveFile.Mode())
	if err != nil {
		return err
	}

	fileInArchive, err := archiveFile.Open()
	if err != nil {
		return err
	}

	if _, err := io.Copy(dstFile, fileInArchive); err != nil {
		return err
	}

	dstFile.Close()
	fileInArchive.Close()
	return nil
}
