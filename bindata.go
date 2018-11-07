// Code generated by go-bindata.
// sources:
// templates/Dockerfile
// templates/Makefile
// templates/k8s.yaml
// DO NOT EDIT!

package main

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

var _templatesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0b\xf2\xf7\x55\x48\xcf\xcf\x49\xcc\x4b\xb7\xca\xcf\x4b\x2a\xcd\xcc\x49\xe1\x02\x04\x00\x00\xff\xff\xa0\x4d\x60\x6e\x14\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/Dockerfile", size: 20, mode: os.FileMode(420), modTime: time.Unix(1541561220, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xb1\x6a\xc3\x30\x10\x86\xe7\xdc\x53\xdc\xa0\x41\x1a\xa4\xac\xc5\xe0\x41\x24\x6e\x08\x25\xb1\xea\x66\x2f\x8a\xa5\xa6\xc5\x72\x25\x6a\x79\x30\x21\xef\x5e\x72\x01\xb5\xdb\xdd\xff\x71\xdf\xf1\x6b\x63\xea\xeb\x15\xd5\xd1\x8e\x1e\x6f\x37\xd8\x1f\xf4\xae\xa1\x64\x3f\xda\x0b\x45\xdb\x76\xf3\xd2\x74\xef\x5d\xdb\x9e\x08\x6c\x63\x3f\xf8\x9f\x2e\xc6\x7c\xa7\x47\x7d\x68\xde\x8c\xde\x34\x45\x33\x25\xdb\xd3\x21\x3c\xbf\x9e\xf4\xae\x66\xfc\x9f\x41\xac\x19\xa7\x1f\x02\x20\xfb\x29\x57\xb0\xba\x44\xbc\x4f\xa8\xd6\x4a\x29\x00\x47\xfa\x8a\x32\x58\x3d\x36\x3c\xcf\x5f\xc1\xa1\xcc\xc8\x38\x49\x05\x2a\x2c\x30\xcd\xd3\x67\x01\x00\xce\xa7\x10\x97\x0a\x1f\x14\x56\xc3\x7c\xf6\x7d\x0e\x68\x53\x0a\x0b\xca\x0f\x1c\x9e\x26\xb5\xd8\x31\xfc\x21\xe7\x83\xcf\x1e\x53\x74\x28\xe5\x77\xe9\xc0\x78\x69\x27\x50\x92\xa1\x66\x5c\x1b\x23\xe0\x37\x00\x00\xff\xff\xc7\xe6\x0a\xde\x38\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/Makefile", size: 312, mode: os.FileMode(420), modTime: time.Unix(1541562737, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesK8sYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xbb\x6e\xf3\x30\x0c\x85\x77\x3f\x05\x5f\xc0\x7f\xec\x7f\x0a\xb4\x15\xc8\xd2\xa1\x45\xd0\x02\xdd\x19\xe5\x20\x10\xa2\x5b\x25\x3a\xad\x11\xf8\xdd\x0b\xe5\xe2\x2a\x4d\xc7\x72\x12\xcf\x67\x92\xc7\x38\x1c\xcd\x1b\x52\x36\xc1\x2b\xc2\xa7\xc0\x97\x67\x5e\x1c\xfa\x0d\x84\xfb\x66\x6f\xfc\x56\xd1\x0a\xd1\x86\xd1\xc1\x4b\xe3\x20\xbc\x65\x61\xd5\x10\x79\x76\x50\x74\x3c\xd2\xbf\x67\x76\xa0\x69\xba\x68\x39\xb2\xae\xc0\xa9\x2d\x34\x47\xe8\x32\x97\x10\xad\xd1\x9c\x15\xf5\x0d\x91\xc0\x45\xcb\x82\x42\x88\xea\xfd\xa5\x2c\x6f\x60\xf3\xb5\x23\xe2\x18\x7f\x9e\x24\xba\x2e\x2e\xa5\x83\x17\x36\x1e\x69\x1e\x6a\x7f\x35\x7a\x2e\xe3\x78\x77\x41\xab\xa0\xf7\x48\x2f\x21\x08\x4d\xd3\xa2\x28\x8f\x05\xde\x7d\xbd\x1e\xac\x5d\x07\x6b\xf4\xa8\xe8\xc1\x7e\xf0\x98\x67\x9e\x90\xc3\x90\x34\x2a\xbf\x45\x7c\x1f\x90\xe5\x46\x23\xd2\x71\x50\xd4\x77\x9d\xbb\x51\x1d\x5c\x48\xa3\xa2\xff\x5d\xf7\x64\x66\x12\x43\xaa\xc7\xdb\xef\x7f\x5c\x87\x24\x8a\x96\xdd\xb2\x6b\xda\xb6\x6d\xea\x30\x0f\xd7\xf0\x5e\x91\x0e\x46\xe3\x4f\x92\x9b\x9d\xb4\xa7\x67\x39\x7d\xb2\x25\x9c\x76\x90\xca\x0d\x51\x86\x85\x96\x90\xce\xbe\xef\x62\xfb\x0a\x00\x00\xff\xff\xf3\x46\xe3\x63\x78\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/k8s.yaml", size: 632, mode: os.FileMode(420), modTime: time.Unix(1541562291, 0)}
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
	"templates/Dockerfile": templatesDockerfile,
	"templates/Makefile": templatesMakefile,
	"templates/k8s.yaml": templatesK8sYaml,
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
		"Dockerfile": &bintree{templatesDockerfile, map[string]*bintree{}},
		"Makefile": &bintree{templatesMakefile, map[string]*bintree{}},
		"k8s.yaml": &bintree{templatesK8sYaml, map[string]*bintree{}},
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
