// Code generated by go-bindata.
// sources:
// templates/.envrc
// templates/.gitignore
// templates/Dockerfile
// templates/Makefile
// templates/config.go
// templates/k8s.yaml
// templates/main.go
// templates/otel.go
// templates/runner.conf
// templates/tracing.go
// DO NOT EDIT!

package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesEnvrc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf7\x37\x34\x34\xf4\xf5\x77\x09\xf5\x71\xb5\xcd\xcf\xe3\x02\x04\x00\x00\xff\xff\x23\xba\x3e\x29\x0f\x00\x00\x00")

func templatesEnvrcBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnvrc,
		"templates/.envrc",
	)
}

func templatesEnvrc() (*asset, error) {
	bytes, err := templatesEnvrcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/.envrc", size: 15, mode: os.FileMode(420), modTime: time.Unix(1564034532, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x2c\x28\xe0\x2a\xc9\x2d\xe0\x02\x04\x00\x00\xff\xff\x57\x45\x21\xbf\x08\x00\x00\x00")

func templatesGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_templatesGitignore,
		"templates/.gitignore",
	)
}

func templatesGitignore() (*asset, error) {
	bytes, err := templatesGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/.gitignore", size: 8, mode: os.FileMode(420), modTime: time.Unix(1551887816, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0b\xf2\xf7\x55\x48\xcf\xcf\x49\xcc\x4b\xb7\x4a\xcc\x29\xc8\xcc\x4b\xe5\x72\xf5\x0b\x09\x8a\x0c\xf0\xf7\xf4\x0b\x51\xd0\x4f\x2c\x28\xe0\x0a\xf7\x0f\xf2\x76\xf1\x0c\x52\xd0\x4f\xcf\xd7\x2f\x2e\x4a\x06\x8b\x71\x39\xba\xb8\x28\x24\x16\x14\x28\xe8\x73\x01\x02\x00\x00\xff\xff\xcd\x75\x7e\xe4\x42\x00\x00\x00")

func templatesDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerfile,
		"templates/Dockerfile",
	)
}

func templatesDockerfile() (*asset, error) {
	bytes, err := templatesDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Dockerfile", size: 66, mode: os.FileMode(420), modTime: time.Unix(1636764845, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\xc1\x8e\x9b\x30\x10\x86\xcf\xcc\x53\x8c\x56\x3e\xd8\x52\x71\xb6\xb7\x0a\x89\x83\x1b\x28\xbb\x6a\x37\x50\xd8\x7b\x42\xc0\x21\x28\x06\xa3\xd8\x48\x89\x10\xef\x5e\xc5\xa1\xb4\x27\xd0\x7c\xf6\x37\xff\x8c\x45\x96\x85\xd3\x84\x7c\x57\x76\x12\xe7\x19\xde\x3f\x44\x12\xbb\xca\x7b\x57\x36\xae\x14\xa5\xdb\x9f\x71\xbe\xcf\xd3\xf4\xd3\x81\x48\x57\x17\x79\xcd\xb5\xb6\x0f\xba\x13\x1f\x71\x91\x89\x6d\xbc\x6a\xcc\x50\x56\xee\x22\xfc\xf8\xfd\x29\x92\x90\xd0\xff\x0c\x6c\x43\xa8\xeb\xc1\x00\x8a\x37\x11\x12\x6a\xce\x52\x29\xac\x9d\x14\xdb\xde\x0c\xb2\xb2\xe8\xfb\x27\x7d\xed\x4a\x8b\x2f\xd3\x74\x98\x26\x6c\xfb\x5a\xde\x90\xe7\x72\xd0\x51\xdb\x48\x63\x0d\xbe\xe2\x3c\x1f\xe6\xf9\x05\x09\xfd\xca\x18\x80\x95\xc6\x06\xe0\x35\x1a\x1f\x7f\xc8\x37\x9c\x73\x80\x46\x07\xe0\x6d\x93\x74\x1f\xef\xc4\xf7\x5f\x71\x14\xbe\x62\x92\xa6\x45\xa8\xda\x7e\xbc\x61\xa3\xf1\x38\xb6\xaa\x46\x5f\x63\x39\x0c\x00\xcf\x18\x01\x2e\x16\xf0\x96\x5c\xcb\x29\x8b\x84\xba\xa1\x18\x72\x5c\xe1\x30\x9a\xf3\x0a\x00\x6a\x39\x28\x7d\x0f\x96\x91\xc0\xbb\x8c\x47\x59\x59\xf5\x68\xa0\xee\xe8\x9f\xf0\xf2\xcd\xf0\x7b\xd9\xa9\x7f\xc8\xf7\xfb\x75\x71\x84\xae\x2b\x65\x68\xa4\xc5\xd6\x3d\xc4\xd3\xda\xc9\xde\x6e\x08\x15\x59\xc6\xf0\xf9\x09\x09\xad\x4a\xa5\xb0\x78\x13\x5f\xfe\x66\x60\xf0\x27\x00\x00\xff\xff\x51\xf2\x28\x5d\xd7\x01\x00\x00")

func templatesMakefileBytes() ([]byte, error) {
	return bindataRead(
		_templatesMakefile,
		"templates/Makefile",
	)
}

func templatesMakefile() (*asset, error) {
	bytes, err := templatesMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Makefile", size: 471, mode: os.FileMode(420), modTime: time.Unix(1636764852, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6a\xdc\x30\x10\x3d\x5b\x5f\x31\x15\x3d\xc8\xa5\x68\x29\xbd\xb5\xec\x21\xa4\x86\x96\xd2\x24\x24\xd0\x3d\x06\x55\x1e\xcb\xa2\xb6\x64\xc6\x23\xb7\x61\xeb\x7f\x2f\xb2\x77\x43\x36\x6c\x0b\xb9\x18\x31\xef\xcd\x7b\xcf\x4f\x1a\x8c\xfd\x69\x1c\x42\x6f\x7c\x10\xc2\xf7\x43\x24\x06\x25\x0a\x39\x32\xf9\xe0\x46\x29\x44\x21\x9d\xe7\x36\xfd\xd0\x36\xf6\x9b\x66\x0c\x91\x7d\xf3\xf0\x78\x90\xa7\xf8\xe8\x29\x0d\x23\x86\x4d\x17\x1d\xa5\xf1\x39\x3a\x34\xef\xde\x6f\x26\x3f\x20\x49\x51\x0a\x31\x19\xca\x66\x9d\x83\x2d\xac\x0b\xfa\x0a\x7f\xa9\x32\x63\x4d\x0a\x16\x46\xe4\x34\x5c\xc6\xd0\x78\xa7\x4a\x78\xb3\x6c\xea\xef\xf9\x0b\x7b\x51\xd8\xc6\xc1\x87\x2d\xac\xd3\x75\x31\xcf\xf4\x45\x5d\xaf\x3b\x37\x86\x5b\x25\xb5\x3c\x3f\x7f\xfd\xf9\xfa\x5b\xb5\xd9\xef\x41\x5f\x99\x1e\x61\x9e\xff\xc1\xdb\x20\xdb\x67\xb4\x95\x77\x87\xbc\xf2\x32\xa0\xe4\x19\xa5\x3b\xe4\x2a\x4c\x37\x84\x8d\xff\xfd\x84\xf0\x07\xd2\x90\xff\xe1\x54\xeb\x13\x36\x26\x75\xac\x64\x30\x3d\xca\xb7\xf0\x5f\xfe\x45\xe2\xd8\x1b\xf6\xb6\x0a\x93\x3a\x71\xfb\x8a\x0f\xb7\x38\x74\xc6\x22\xa9\xc3\x2d\xe6\x72\x1e\x67\x52\x67\xed\x7b\x59\x66\x29\xdf\x00\x12\xe5\x16\xb3\xc0\x2d\x9a\xfa\x4b\x38\xf6\xfd\x71\x81\x5e\x6d\x21\xf8\x2e\xd7\x5d\x74\x4e\xef\x3c\xb7\x15\x51\x24\x85\x44\xa5\x5e\x8f\xd2\xc6\xd4\xd5\x10\x22\x03\xa1\xa9\xc1\x07\xcf\xde\x74\x60\x17\xa1\xdc\xc5\x7c\x48\x7d\x7d\x10\xbf\x6c\x4d\x70\xa8\xf2\x1d\xab\x7b\x38\xbe\x25\x5d\x4d\x18\xb8\x5c\xbc\x5e\x14\xec\x4c\xb2\x9d\xa1\x70\x1a\xac\x8b\xa6\x7e\x12\xa9\x98\x45\x31\xe7\x0a\x5c\x5c\x2c\x76\x86\x6d\x7b\x74\x10\xa2\x20\xe4\x44\x21\x43\x62\x16\x7f\x03\x00\x00\xff\xff\xf2\xeb\x80\xb3\x27\x03\x00\x00")

func templatesConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigGo,
		"templates/config.go",
	)
}

func templatesConfigGo() (*asset, error) {
	bytes, err := templatesConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.go", size: 807, mode: os.FileMode(420), modTime: time.Unix(1647024160, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesK8sYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x0c\xa1\x0b\xed\xc1\x8e\xd3\xd3\x22\xe8\xa1\xb0\x97\x85\x36\x84\x16\x7a\x9f\xc8\x93\x44\x44\x5f\x95\xc6\xdb\x06\xe3\xff\x5e\x26\xb1\x5d\x27\xbb\x7b\x8b\x4e\xd2\xbc\x99\xd1\xd3\xbc\x27\x8c\xe6\x17\xa5\x6c\x82\x57\x80\x31\xe6\xe5\xcb\xaa\x38\x1a\xdf\x28\x78\xa2\x68\xc3\xc9\x91\xe7\xc2\x11\x63\x83\x8c\xaa\x00\xf0\xe8\x48\x41\xd7\x41\xb5\x46\x47\xd0\xf7\x43\x2c\x47\xd4\x33\xe0\x7c\x14\x34\x47\xd2\x52\x97\x28\x5a\xa3\x31\x2b\x58\x15\x00\x99\x2c\x69\x0e\x49\x10\x00\x87\xac\x0f\xdf\x70\x4b\x36\x5f\x02\x20\x5c\x6e\x6f\x61\x72\xd1\x22\xd3\x50\x33\xe3\x24\xcb\x5e\x95\xbf\xd9\x00\x60\x24\x23\x4b\x07\xcf\x68\x3c\xa5\xa9\xa8\x7c\xf3\x71\x97\x65\x1c\xee\x07\xe8\x29\xe8\x23\xa5\x1f\x21\x30\xf4\xfd\x52\x22\xcf\x02\xbe\xca\xde\xb4\xd6\x6e\x82\x35\xfa\xa4\xe0\xab\xfd\x83\xa7\x3c\xe1\x89\x72\x68\x93\xa6\x19\x5f\x09\xfe\x6e\x29\xf3\x55\x0c\x40\xc7\x56\xc1\xaa\xae\xdd\x55\xd4\x91\x0b\xe9\xa4\xe0\x73\x5d\x7f\x37\x13\x12\x43\x9a\x97\x97\xff\xdf\xb8\x09\x89\x15\x3c\xd6\x8f\x75\x51\x96\x65\x31\x57\x7d\x12\xfc\x27\xa5\x17\xa3\xe9\x2e\x6a\x4f\x4c\xca\xf3\x56\xae\x3e\xd3\x62\x4c\x7b\xe2\x19\x9b\x5b\x2b\xbc\x92\xad\xeb\xc0\xec\xa0\x7a\xf6\xfb\x44\x39\x43\xdf\x77\x1d\x7c\x30\x7e\x0f\xea\x0b\xc4\x64\x3c\xef\x60\xf1\x90\xab\x87\xbc\x18\x6a\x3e\x36\xb4\xc3\xd6\xb2\x08\xe5\xd0\x78\x58\x60\xe6\x16\x13\x57\x3a\x2c\x3e\x49\xcb\xdb\x09\xd0\x5f\x26\x2f\x5b\x71\xff\x96\x18\xc7\x89\x0c\x97\xde\x65\x22\x6c\x87\x79\x64\xd2\x89\x78\x3d\xb6\x39\xbf\xa5\xef\x2b\xb6\x17\x7b\x1c\xc2\x64\x81\x72\x86\xcb\x0f\x6a\x2d\x0d\x3d\x24\x49\x5d\xa3\x00\x07\xe6\x38\x8a\x1f\x91\x0f\x33\x5f\xcb\x51\xc1\x72\x72\xc6\x16\xf5\x91\x7c\x33\x77\x5a\xbe\xa8\xbf\x7e\xf7\x07\x4c\x29\xa3\x7a\x22\x0d\xf9\x46\x72\xfe\x05\x00\x00\xff\xff\x62\xda\x84\x0d\x46\x04\x00\x00")

func templatesK8sYamlBytes() ([]byte, error) {
	return bindataRead(
		_templatesK8sYaml,
		"templates/k8s.yaml",
	)
}

func templatesK8sYaml() (*asset, error) {
	bytes, err := templatesK8sYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/k8s.yaml", size: 1094, mode: os.FileMode(420), modTime: time.Unix(1565288430, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x6f\xdc\x36\x10\x3d\x8b\xbf\x62\xc0\x13\x99\x6e\xa8\x14\x28\x8a\xc0\x40\x0e\x41\x5a\xd7\x01\x12\xdb\xd8\xb5\x91\x43\x51\x04\xb4\x34\x92\x08\x53\xa4\x3a\x1c\x79\x37\xdd\xea\xbf\x17\xa4\x76\xf3\xd1\x8f\x43\x8f\x7a\xc3\xc7\x79\xef\xcd\x50\x93\x6d\x1e\x6d\x8f\x30\x5a\x17\x84\x70\xe3\x14\x89\x41\x89\x4a\x76\x23\x4b\x51\xc9\x80\x5c\x0f\xcc\x93\x14\xa2\x1a\xd0\x7a\x1e\x40\xda\xc4\xb3\x25\x36\x4d\xac\xfb\xf8\x7c\x45\x9b\x01\x9b\xc7\x4c\xe8\x1d\x0f\xf3\x83\x69\xe2\x58\xf7\x91\x9c\xf7\xb6\x1e\x6c\x68\x3d\x52\xfa\x8f\xf2\x38\x1f\xfe\x51\x99\xe6\xae\xab\xc7\x48\xd8\xc4\xc0\x78\x60\x29\xb4\x10\x4f\x96\xb2\xb4\x11\x19\x09\x5e\x41\xef\xe3\x83\xf5\xe6\x7d\xfe\x54\xf2\x78\x04\x73\x6d\x47\x84\x3f\x81\xa3\x8f\x7b\x24\x58\x16\xb9\x81\x11\x99\x5c\x63\x3e\x38\x1e\x7e\xc2\xd4\x90\x9b\xd8\xc5\xa0\x24\xbb\x11\x81\xed\x23\x06\xe8\x22\xc1\xe4\x1a\x68\xe3\x3e\xf8\x68\x5b\xa9\xb5\xa8\x1a\x26\x00\x78\x05\xa5\x9b\xb9\xc6\xfd\xa5\x8f\x96\x7f\xfc\xe1\xca\x25\x8e\x3d\xd9\x51\xc9\x01\xbd\x8f\x1f\xf7\x91\x7c\xfb\xd1\x5b\xc6\xd0\x7c\x92\x1b\x70\x21\x31\xcd\x23\x06\x2e\x4d\xef\x83\x63\x35\x07\xc7\x66\x97\xcd\xb4\x5a\x67\x2b\xdd\x1c\x9a\x92\xb9\xd2\x70\xcc\xcd\x0e\x70\xf1\x0a\xbe\x32\x6c\x2e\x23\xed\x5c\x1f\xac\x4f\x4a\x8b\x2a\x21\xcf\xd3\x0d\xa3\x57\x0d\x1f\xb2\xba\xae\xcf\x84\x02\xbf\x89\xa1\x73\x7d\x3e\xe5\x7b\xf3\x36\x74\x51\xc9\xc4\x96\x18\x5b\xa9\x85\xa8\x08\xcb\xd1\x75\x4c\xd9\xc8\x16\x7b\x97\x98\x3e\x65\x06\x95\xb6\xf3\xa1\xe0\x71\xce\x49\x66\xd4\xdc\x5a\x1e\x94\xac\x57\x92\xd4\xe6\x6a\x1d\xa1\x22\xec\xf3\x9d\x76\x72\x99\xb8\x9e\xbb\x25\xec\xdc\x41\xc9\xda\x4e\xae\x7e\xfa\x5e\x6a\xb3\x9b\x1f\xe8\x7c\x99\xa8\xea\x1a\xee\x13\xc2\xe8\xda\xd6\xe3\xde\x12\xc2\xdb\xed\xbb\x72\x87\xb9\x4f\xa8\x72\xf7\xf7\x9f\x6b\x97\x73\x68\x54\x4e\x47\x0d\x90\xf7\xee\xdc\x59\x7f\xf3\x95\x33\xab\x08\x79\xa6\xf0\x0d\xfe\x85\xbd\x5f\xf1\x2d\xa6\x29\x86\x84\x1f\xc8\x31\xd2\x06\x08\x9e\x9d\xf0\xdf\x67\x4c\x5c\xc2\xaf\xaa\x12\x57\x36\x94\x97\xc2\x5c\xc7\x7d\x0e\xa1\xaa\x5a\xec\x90\xa0\x5c\x77\x3a\x98\xb7\xc2\x6c\xb1\x89\xd4\xe6\x41\x6c\x56\xc2\xce\x85\x06\x55\xb9\x44\x9f\xa6\x9c\x94\xde\x80\x65\x26\xf7\x30\x33\x9a\x1d\x93\x0b\xbd\x92\x93\xe5\x41\x6e\x80\xcc\xfd\xf6\x5d\xc9\x4e\x97\x46\xcb\xda\x6f\x30\x3b\xa4\x27\xbc\xba\xbb\xbb\x55\xfb\x0d\x50\x06\x17\x2d\xaa\x45\x9f\x32\x3f\x8f\xe5\xcb\x40\x3e\x1b\x86\xff\xe5\xb8\x68\x4f\x53\xb6\x1c\x19\xbd\xb9\x23\xdb\xe4\x47\x94\x37\xd2\xf4\x31\x8f\x30\xbb\x59\x3d\x16\x54\x66\x31\x6b\x20\x69\x32\x3f\x87\xb6\x8c\xb6\xea\x46\x36\x97\x13\xb9\xc0\x5d\x96\xbc\x3e\x09\x90\xf0\x1d\x34\x5d\x6f\x7e\x41\x3e\x3b\x0f\x76\xc4\xf2\xae\x96\xcc\x4b\xf4\x54\x76\x32\x0b\x2b\x9e\x29\xab\x7a\xdd\xb6\x74\x01\xf2\xe2\xe5\x8b\x97\x2f\xe4\x46\x54\xd5\xf1\x08\xae\x03\xf3\xe6\x66\xbb\x83\xe7\xcb\x22\xaa\xea\x64\xfb\x02\xce\xff\x94\x52\x54\x25\xbd\x33\xf2\xda\xe7\xe7\xdf\x5e\xa1\x6d\x91\x92\xfa\xf5\xb7\x54\x34\x1c\xe5\xeb\x99\x87\x48\xee\x0f\x9b\xff\x00\x72\xd1\x9b\x7f\xa3\xdd\x90\xeb\x5d\xf8\x9a\xf6\xec\x74\x54\x2b\xd2\x27\x55\xe8\x13\xfe\x5d\x11\x9d\x6b\xa1\x5d\x4b\x8b\x10\x8b\xf8\x2b\x00\x00\xff\xff\x29\xc6\xe0\x18\x60\x05\x00\x00")

func templatesMainGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainGo,
		"templates/main.go",
	)
}

func templatesMainGo() (*asset, error) {
	bytes, err := templatesMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.go", size: 1376, mode: os.FileMode(420), modTime: time.Unix(1647025072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOtelGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x5b\x6b\x2b\x37\x10\x7e\x96\x7e\x85\x2a\x28\xec\x16\x77\xb7\xa7\x2d\x1c\x08\x9c\x87\x36\xa7\xa1\xa5\x75\x1a\xe2\x5e\x9e\x15\xed\x78\x2d\xa2\x95\xc4\x68\xd6\x71\x30\xfe\xef\x45\xda\x5d\xdf\x48\x9c\xb4\xf4\xc5\x36\x73\xfd\xf4\xcd\x37\xe3\xa0\xf4\xa3\x6a\x41\x74\xca\x38\xce\x4d\x17\x3c\x92\x28\x38\x93\xda\x3b\x82\x0d\x49\xce\xa4\xf5\x6d\xfa\x22\xd3\x81\xe4\x9c\xc9\xd6\x57\x3e\x80\x23\xb0\xd0\x01\xe1\x73\x65\x7c\xed\x09\xac\xbc\xe0\xab\x15\x11\x9a\x87\x9e\xe0\x62\x14\x6c\x12\x00\xc0\x58\x7b\xb2\x21\x7f\xa4\x00\xa3\xff\x5b\xd6\xd1\xcf\x16\xc3\xbf\x2e\x42\xa8\x34\x1c\x7e\xbd\x59\x62\x6c\xda\x5a\xff\xa0\x2e\xb3\x11\xd0\x07\xd5\x2a\x32\xde\x49\xce\x12\xd7\xe8\xad\x05\x14\xaf\xa7\xc4\xe6\x71\x6a\x70\x88\xaf\x1f\x54\xcc\xe4\x04\xf4\x1a\x62\xf4\xef\xac\xb0\x0f\xdf\x17\x78\x57\x5a\x04\x0b\x9a\x3c\xd6\xd1\x74\xc1\x5e\x1e\x65\x4a\x43\x88\xbe\x47\xfd\x76\x60\xe6\x57\x72\x16\xa1\xd3\xde\xad\x2f\x3d\x62\x88\xa8\xd7\x1f\xaa\x8f\xd5\x37\x92\x97\x9c\x2f\x7b\xa7\x45\x04\xea\xc3\xef\x04\xb6\xd0\xb4\x11\xa3\x7a\xab\xeb\xe1\xbb\x14\x5b\xce\xea\x5a\xfc\x91\xda\x44\xce\x08\x67\x02\x10\xc5\xd5\x27\x71\x32\xdd\xea\x16\x9e\x52\xfe\xec\xcc\xfc\xb7\xa1\xd5\x4f\xae\x09\xde\x38\x2a\xe4\x09\xb0\xaf\x75\x1a\x44\x22\xa5\x7a\xc5\x7e\xf5\xfd\x77\x1f\x3e\xca\xf2\xa5\x9a\xbf\xb8\x08\xba\x47\x28\xca\x92\x33\xb3\xcc\x98\xbe\xf8\x24\x9c\xb1\x09\x30\xb3\xbe\xad\x6e\x14\x29\x5b\x00\x62\xc9\xd9\x8e\x33\x84\xb8\x87\x3e\xb1\x5b\xcd\x01\x5b\x28\x38\x63\x7b\xcb\x67\x58\xaa\xde\x52\x51\xce\x8e\xad\xb7\xf0\x94\xba\xfe\x30\xad\x62\x4c\x39\x13\xe7\xd5\x42\xaf\xa0\x53\x7f\xde\xff\x36\x3b\xb1\x02\xae\x8d\x86\x5b\xd5\xc1\xaf\xf0\x5c\x2d\x08\x8d\x6b\x0b\xb9\xdd\x8a\x2a\xd9\xc4\x6e\x27\xcb\x97\x12\xfe\x02\x8c\xc6\xbb\xe3\x9c\x06\xd6\x60\x7d\x18\xe3\xf7\x07\x61\xef\x07\xb7\x96\x33\x21\x03\xfa\xa6\xd7\x79\x35\x72\x64\xfa\x78\x2f\x3d\x49\x22\xd5\x02\x28\x0f\x1a\xef\xd0\xaf\x4d\x03\x58\x64\xd6\xd3\xf3\xcf\xec\x9c\xb1\xc1\x95\x68\xb9\x1f\x69\x2a\x10\x62\x6e\x7c\x70\xfd\xa8\x48\xaf\x72\x9d\x0c\xa6\x3c\x6a\x04\x1b\x9a\xab\x70\x37\xee\xb3\xc7\xe2\x68\xb5\xab\xdc\x6e\xd4\xe0\x76\x57\xf2\x2c\xc2\x79\xde\xa5\xc8\x99\xb6\x06\x1c\x4d\x22\x3c\x9c\xa9\x04\xf4\x3a\xfb\x8a\x33\xc7\xff\xa7\xc3\xb3\xa2\x27\x42\x84\x4d\x38\x59\x8f\x21\xf6\xb0\x1b\x03\xee\xcb\x23\x59\x16\xf2\x46\x19\x0b\x8d\x20\x2f\x34\x82\x22\x10\xb4\x02\xb1\x47\x23\xa6\x7b\x7b\x25\xbe\x4c\x63\x9f\x46\xc8\x59\xe8\xe3\x0a\x72\xf3\xc3\x9d\xcb\xcd\x39\x3b\x9c\xb9\x64\xb8\x51\xa9\xd0\xf3\x20\xe2\x7c\x90\x26\x85\xff\x6c\x22\xf9\x16\x55\xf7\xd9\xc4\x41\x65\xc6\xbb\x61\x19\xf2\xeb\x46\x55\x1d\x1d\xde\x81\xda\x11\x52\x01\x9b\xf0\x92\xff\x7a\x00\x7f\x07\x68\x7c\x53\x7c\xfb\x55\xfa\x47\xac\x16\xa0\xbd\x6b\x46\x91\x0e\xa7\x3f\x09\x63\x0e\x74\x24\xb4\xe1\x4d\x25\xdf\xf1\x7f\x02\x00\x00\xff\xff\x34\xed\xbd\x10\x71\x07\x00\x00")

func templatesOtelGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesOtelGo,
		"templates/otel.go",
	)
}

func templatesOtelGo() (*asset, error) {
	bytes, err := templatesOtelGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/otel.go", size: 1905, mode: os.FileMode(420), modTime: time.Unix(1647025132, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRunnerConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x48\x2b\x4a\x2d\xce\x50\x48\xce\xcf\x4b\xcb\x4c\xe7\xe2\xca\x4c\xcf\xcb\x2f\x4a\x4d\xb1\x52\x28\x4f\x4d\xd2\x51\x28\xc9\x2d\xd0\x51\x48\x2c\x2e\x4e\x2d\x29\xe6\x02\x04\x00\x00\xff\xff\x3c\xbc\xcd\x0e\x2a\x00\x00\x00")

func templatesRunnerConfBytes() ([]byte, error) {
	return bindataRead(
		_templatesRunnerConf,
		"templates/runner.conf",
	)
}

func templatesRunnerConf() (*asset, error) {
	bytes, err := templatesRunnerConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/runner.conf", size: 42, mode: os.FileMode(420), modTime: time.Unix(1551887851, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTracingGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x4f\x6b\x1b\x3f\x10\x3d\x4b\x9f\x42\x3f\x9d\xb4\x60\xe4\x5f\x4f\x85\x42\x0e\x6d\xda\x42\x69\xe3\x86\xb8\x7f\xce\xaa\x3c\x5e\x8b\x68\x35\x62\x76\x76\xe3\x60\xf6\xbb\x17\xed\xda\x8e\x13\xc2\xd2\x8b\x10\xef\xbd\xd1\x9b\x79\x9a\xec\xfc\xbd\xab\x41\x35\x2e\x24\x29\x43\x93\x91\x58\x19\x29\x74\xc4\x5a\x4b\x29\x74\x8d\x16\x33\x24\x86\x08\x0d\x30\x3d\xda\x80\x4b\x64\x88\x7a\x86\x5b\x3a\x66\x0a\x7f\x3a\x86\x59\x15\xec\x8b\x1b\x50\xbb\x44\x8e\x79\x3c\x98\x9c\x87\xa7\x5b\x4d\xd9\xcf\x3e\x91\x09\xb3\xab\x1d\x07\x4c\xb3\xba\x76\x73\xbf\x24\x68\xb1\x23\x3f\xdf\x53\x11\x8e\xd6\x5a\x8a\x16\x1a\x8f\xa9\x57\x33\xea\x49\xb1\xec\xdf\xd8\xb7\xf6\x7f\x2d\x2b\x29\xb7\x5d\xf2\x2a\xa4\xc0\xa6\x52\x07\x29\x98\x16\x0a\x88\xd4\xbb\x2b\xf5\x6c\x2a\xbb\x82\x07\xe3\x79\xbf\x78\x01\xff\x0e\xbc\xfb\x94\x36\x19\x43\x62\xa3\x0f\x07\x65\xbf\x33\x44\x7b\x8d\x31\x82\x67\xa4\x13\xa7\x86\x41\x57\xaf\x15\x7f\x49\x2d\xf8\x8e\xc0\x54\x95\x14\x61\x3b\x9a\xff\x77\xa5\x52\x88\xa5\x1d\x11\xb1\xb6\x9f\x1d\xbb\x68\x80\xa8\x92\x62\x90\x82\xa0\x3d\xf7\x78\xca\xc8\xde\x00\xd5\x60\xa4\x10\x67\xe4\x23\x6c\x5d\x17\xd9\x54\x8b\x4b\x74\x05\x0f\xc5\xf5\xfd\xe9\xc3\xdb\x52\x73\x4a\xce\xae\xfd\x0e\x1a\xf7\xf3\xee\xdb\xe2\x19\x0a\xd4\x07\x0f\x2b\xd7\xc0\x57\x78\xb4\x6b\xa6\x90\xea\x69\xd8\x82\x4d\xa3\xbd\x52\xf0\x0b\xa8\x0d\x98\x2e\x6b\x36\xd0\x43\xc4\x7c\xd4\x9f\xd7\xee\xcc\x43\xea\xf5\x42\xe9\x4c\xb8\x99\x34\xe5\xf8\xd7\x60\xca\x17\xdb\x35\xf0\x8f\x12\x30\xdd\x12\xf6\x61\x03\x64\xc6\xbc\xcb\xe0\x2f\x70\x29\xc4\x44\x95\x40\xee\x8e\x01\x19\x82\x76\x34\x7e\xa2\x3e\x38\xf6\xbb\xf1\x9d\xb1\x99\xea\xc2\x08\xf6\x7c\xe3\xf2\xed\x71\xa9\x91\xcc\xc5\x7e\xdb\xd1\xee\x1a\x13\xc3\x9e\x0f\x43\x25\x07\xf9\x37\x00\x00\xff\xff\x54\xa9\x24\x00\xbd\x03\x00\x00")

func templatesTracingGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesTracingGo,
		"templates/tracing.go",
	)
}

func templatesTracingGo() (*asset, error) {
	bytes, err := templatesTracingGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tracing.go", size: 957, mode: os.FileMode(420), modTime: time.Unix(1637593097, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/.envrc": templatesEnvrc,
	"templates/.gitignore": templatesGitignore,
	"templates/Dockerfile": templatesDockerfile,
	"templates/Makefile": templatesMakefile,
	"templates/config.go": templatesConfigGo,
	"templates/k8s.yaml": templatesK8sYaml,
	"templates/main.go": templatesMainGo,
	"templates/otel.go": templatesOtelGo,
	"templates/runner.conf": templatesRunnerConf,
	"templates/tracing.go": templatesTracingGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		".envrc": &bintree{templatesEnvrc, map[string]*bintree{}},
		".gitignore": &bintree{templatesGitignore, map[string]*bintree{}},
		"Dockerfile": &bintree{templatesDockerfile, map[string]*bintree{}},
		"Makefile": &bintree{templatesMakefile, map[string]*bintree{}},
		"config.go": &bintree{templatesConfigGo, map[string]*bintree{}},
		"k8s.yaml": &bintree{templatesK8sYaml, map[string]*bintree{}},
		"main.go": &bintree{templatesMainGo, map[string]*bintree{}},
		"otel.go": &bintree{templatesOtelGo, map[string]*bintree{}},
		"runner.conf": &bintree{templatesRunnerConf, map[string]*bintree{}},
		"tracing.go": &bintree{templatesTracingGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

