package fileStorage

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)


type LocalDiskOpts struct {
	rootDir string
	baseUrl string
	urlPrefix string
}

func NewLocalDiskOpts(rootDir string, baseUrl string, urlPrefix string) *LocalDiskOpts {
	return &LocalDiskOpts{rootDir: rootDir, baseUrl: baseUrl, urlPrefix: urlPrefix}
}

func (opts LocalDiskOpts) getRootDir() string {
	return strings.TrimSuffix(opts.rootDir, "/") + "/"
}

func (opts LocalDiskOpts) getBaseUrl() string {
	return strings.TrimSuffix(opts.baseUrl, "/") + "/"
}

func (opts LocalDiskOpts) UrlPrefix() string {
	return strings.Trim(opts.urlPrefix, "/")+ "/"
}

type localDisk struct {
	opts *LocalDiskOpts
}

func (d localDisk) CopyToRoot(file *multipart.FileHeader) (string, error) {
	return d.CopyToDir(file, "")
}

func (d localDisk) CopyToDir(file *multipart.FileHeader, dir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	path := dir + strconv.FormatInt(time.Now().UnixNano(), 10) + file.Filename
	dst, err := os.Create(d.opts.getRootDir() + d.trimPathSlashPrefix(path))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (d localDisk) Url(path string) string {
	return d.opts.getBaseUrl() + d.opts.UrlPrefix() + d.trimPathSlashPrefix(path)
}

func (d localDisk) trimPathSlashPrefix(path string) string {
	return strings.TrimPrefix(path, "/")
}

func NewPublicDisk(opts *LocalDiskOpts) *localDisk {
	return &localDisk{
		opts: opts,
	}
}