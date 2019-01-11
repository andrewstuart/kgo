// Code generated by go-bindata.
// sources:
// templates/Dockerfile
// templates/Makefile
// templates/config.go
// templates/k8s.yaml
// templates/main.go
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

var _templatesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0b\xf2\xf7\x55\x48\xcf\xcf\x49\xcc\x4b\xe7\xe2\x0a\xf7\x0f\xf2\x76\xf1\x0c\x52\x50\x71\xf7\x0f\x70\x0c\xf1\xd0\x4f\x2c\x28\xe0\xe2\x72\x74\x71\x51\xd0\x53\xd0\xe3\xe2\x0a\x0a\xf5\x53\x48\xcf\x57\x48\x4f\x2d\x51\x50\x53\x03\xb1\x92\x4a\x33\x73\x52\x14\x74\xf3\x15\xc0\x0a\x01\x01\x00\x00\xff\xff\xac\xe3\x8a\xb6\x4a\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/Dockerfile", size: 74, mode: os.FileMode(420), modTime: time.Unix(1545379654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x51\x6b\xf2\x30\x14\x86\xaf\x3d\xbf\xe2\x20\xb9\x48\xe0\x6b\xfc\x76\x37\x84\x5e\x04\xed\x54\x86\xda\x55\xef\x67\x6c\xa3\x16\xd3\x26\x98\x08\x93\x92\xff\x3e\x8c\xae\xdb\x55\xcb\xfb\x90\xe7\xbc\xbc\x22\xcf\xd3\xae\x43\xbe\x92\x8d\xc2\x10\x60\xb1\x14\xb3\x2c\x26\x8b\x46\x1e\x63\x34\x5d\x4f\xde\xb3\xe2\xb3\x58\xaf\xb7\x11\x4c\x4d\x79\x56\x97\xc2\x18\x7f\xa7\x2b\xb1\xcc\x36\xb9\x98\x64\xbd\xc6\x59\x59\xc6\x87\xf0\xf6\xb1\x15\xb3\x94\xd0\x3f\x06\x36\x22\x34\xde\x60\x00\x9b\xb9\x48\x09\x75\x27\xa5\x35\x56\x51\x8a\x75\xeb\xac\x2a\x3d\x26\xc9\xc1\x5c\x1a\xe9\x71\xd8\x75\xbb\xae\xc3\xba\xad\xd4\x17\xf2\x42\x59\x33\xad\x8f\xca\x79\x87\xff\x31\x84\x5d\x08\x43\x24\xf4\x85\x31\x00\xaf\x9c\x1f\xc3\xe0\x68\xf0\xfe\x87\x7c\xc4\x39\x07\x78\x78\xc7\x31\x83\xc1\xf3\xca\xfe\x5a\xeb\x0a\x13\x8f\x84\xc6\x8a\x0c\x39\xf6\xd0\x5e\xdd\xa9\x07\x00\x95\xb2\xda\xdc\xc6\xcf\x82\x30\x38\x5f\xf7\xaa\xf4\x1a\xa5\xb5\xfa\x86\xc9\x01\xcf\xaf\x8e\xdf\x64\xa3\x7f\x51\x92\xb4\xfd\x0c\x84\xf6\x03\x31\x74\xca\x63\x1d\x67\x7d\x58\x1b\xd5\xfa\x11\xa1\x22\xcf\x19\x3e\x3e\x29\xa1\xa5\xd4\x1a\x37\x73\xf1\xef\xa7\x03\x83\xef\x00\x00\x00\xff\xff\x0d\x5a\x42\xe7\xa5\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/Makefile", size: 421, mode: os.FileMode(420), modTime: time.Unix(1541574082, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6a\xdc\x30\x10\x3d\x5b\x5f\x31\x15\x3d\xc8\x50\x64\x4a\x6f\x2d\x7b\x08\xc1\xd0\x52\x9a\x84\xe4\xb0\xc7\xa0\xca\x63\x59\xd4\x96\xc4\x58\x72\x1b\xb6\xfe\xf7\x22\xdb\x1b\x76\x97\xed\x21\x37\xf1\xde\x9b\x37\xcf\x6f\x1c\x94\xfe\xa5\x0c\xc2\xa0\xac\x63\xcc\x0e\xc1\x53\x04\xc1\x0a\x3e\x46\xb2\xce\x8c\x9c\xb1\x82\x1b\x1b\xbb\xf4\x53\x6a\x3f\x54\xed\xe8\x7c\xb4\xed\xcb\xeb\x83\x9f\xf3\xa3\xa5\x14\x46\x74\x55\xef\x0d\xa5\xf1\x92\x0d\xed\xc7\x4f\xd5\x64\x03\x12\x67\x25\x63\x93\xa2\xbc\x4c\xb7\x06\x76\xb0\xc0\xf2\x0e\x7f\x8b\x92\x15\xbd\x01\xd8\xc1\x6a\xb2\x61\x25\x63\x6d\x72\x1a\x46\x8c\x29\xdc\x7a\xd7\x5a\x23\x4a\x38\x2c\xe3\xf2\xa6\x69\x56\xe8\x41\xc5\x4e\x70\xc9\xcb\xab\xf8\xfb\xaf\xf7\x3f\xea\xea\x70\x00\x79\xa7\x06\x84\x79\xfe\x8f\xae\xc2\xa8\x2f\x64\xab\xee\x09\xe3\xaa\xcb\x84\xe0\x57\x9c\x9e\x30\xd6\x6e\x7a\x20\x6c\xed\x9f\x13\xc1\x5f\x48\x21\x20\x9d\x7a\xdd\xa4\xe8\x07\x15\xad\xae\xdd\x24\xce\xa6\xbf\xe3\xcb\x23\x86\x5e\x69\x24\xb1\x1d\x22\x97\xf0\x8a\x71\xc9\x3f\x00\x7f\xe6\x65\xb6\xb2\x2d\x20\x11\x7c\xde\x41\x36\x78\x44\xd5\x7c\x73\xc7\x7a\xbe\x2c\xd4\xbb\x1d\x38\xdb\xe7\xaa\x8a\xde\xc8\xbd\x8d\x5d\x4d\xe4\x49\x20\x51\x29\xd7\x27\xd7\x3e\xf5\x0d\x38\x1f\x81\x50\x35\x60\x9d\x8d\x56\xf5\xa0\x17\xa3\xfc\x6d\xf3\x96\xfa\x7e\x33\xbf\xed\x94\x33\x28\xf2\x49\xc4\x33\x1c\x7f\x07\x59\x4f\xe8\xe2\x72\x96\xb7\x05\xbb\x92\x6c\xaf\xc8\x9d\x07\xeb\xbd\x6a\x4e\x22\x15\x33\x2b\xe6\x5c\x81\xf1\xcb\x8a\xbd\x8a\xba\x3b\x6e\x60\x33\xfb\x17\x00\x00\xff\xff\x1b\x8f\x84\x25\xdd\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/config.go", size: 733, mode: os.FileMode(420), modTime: time.Unix(1544858246, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesK8sYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xc1\x8a\xdb\x30\x10\xbd\xfb\x2b\x1e\xa1\x0b\xed\xc1\x89\xd3\xd3\x22\xe8\xa1\xb0\x97\x3d\x34\x84\x16\x7a\x9f\xc8\x93\x44\x44\x96\x54\x69\x9c\x36\x18\xff\x7b\x51\x62\xbb\xce\x6e\xf7\xb6\x3a\x49\xf3\x66\x46\x4f\xf3\x9e\x28\x98\x9f\x1c\x93\xf1\x4e\x81\xff\x08\xbb\xbc\x4d\xab\xf3\x7a\xc7\x42\xeb\xe2\x64\x5c\xad\xf0\xc4\xc1\xfa\x4b\xc3\x4e\x8a\x86\x85\x6a\x12\x52\x05\xe0\xa8\x61\x85\xae\xc3\x72\x43\x0d\xa3\xef\x87\x58\x0a\xa4\x67\xc0\xf5\x98\xd1\x14\x58\xe7\xba\xc8\xc1\x1a\x4d\x49\x61\x5d\x00\xc2\x4d\xb0\x24\x9c\x11\x60\xde\x3f\x2f\x4b\x3b\xb6\x69\x3c\x01\x14\xc2\xcb\x2b\x81\xb1\x71\x5e\xda\x3b\x21\xe3\x38\x4e\x45\xe5\x7f\x89\xde\x96\x69\xe8\x30\x40\x4f\x5e\x9f\x38\x7e\xf7\x5e\xd0\xf7\xab\x1c\x79\xce\xe0\xab\xec\x6d\x6b\xed\xd6\x5b\xa3\x2f\x0a\x5f\xed\x6f\xba\xa4\x09\x8f\x9c\x7c\x1b\x35\xcf\xf8\xe6\xe0\xaf\x96\x93\xdc\xc5\x00\x1d\x5a\x85\x75\x55\x35\x77\xd1\x86\x1b\x1f\x2f\x0a\x9f\xab\xea\x9b\x99\x90\xe0\xe3\xbc\xbc\xfc\xf7\xc6\xad\x8f\xa2\xf0\x58\x3d\x56\x45\x59\x96\xc5\x5c\xcc\xf3\x28\xde\x0f\x8e\x67\xa3\xf9\x5d\x94\x9b\x98\x94\xd7\x6d\xbe\xfa\x4a\x4b\x28\x1e\x58\x66\x6c\x80\xc4\x96\xb5\xf8\x78\xe3\xfd\x4a\xb6\xae\x83\xd9\x63\xf9\xec\x0e\x91\x53\x42\xdf\x77\x1d\x3e\x18\x77\x80\xfa\x82\x10\x8d\x93\x3d\x16\x0f\x69\xf9\x90\x16\x43\xcd\xc7\x9a\xf7\xd4\x5a\xc9\x42\x35\x64\x1c\x16\x94\xa4\xa5\x28\x4b\xed\x17\x9f\x72\xcb\x97\x13\x78\xd3\xce\xc3\xa5\xef\x32\x11\xb1\xc3\x3c\x12\xeb\xc8\xb2\x19\xdb\x5c\xdf\xd2\xf7\x4b\xb1\x37\x7b\x1c\xfd\x64\x81\x72\x86\xe7\xdf\xd0\x5a\x1e\x7a\xe4\x24\x75\x8f\x02\x47\x91\x30\x8a\x1f\x48\x8e\x33\x5f\xe7\xa3\xc2\x6a\x72\xc6\x8e\xf4\x89\x5d\x3d\x77\x5a\xba\xa9\xbf\x79\xf3\x07\x4c\x29\xa3\x7a\x59\x1a\x76\x75\xce\xf9\x1b\x00\x00\xff\xff\x87\x7f\xb8\x09\x1d\x04\x00\x00")

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

	info := bindataFileInfo{name: "templates/k8s.yaml", size: 1053, mode: os.FileMode(420), modTime: time.Unix(1547179257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xce\xc1\x6a\xe3\x30\x10\xc6\xf1\xb3\xe6\x29\xc4\x9c\xe4\xc5\xc8\xd9\x5b\x08\xec\x61\x59\x08\x5b\x28\x25\x24\x87\x9e\x15\x67\x6c\x8b\xca\x92\x3a\x1e\xc5\x81\x92\x77\x2f\x76\x03\x7d\x82\x5e\x3f\x98\xf9\xff\xb2\x6b\xdf\x5c\x4f\x7a\x74\x3e\x02\xf8\x31\x27\x16\x6d\x40\x61\x37\x0a\x82\xc2\x48\xd2\x0c\x22\x19\x01\x14\xf6\x5e\x86\x72\xb6\x6d\x1a\x9b\x3e\xb1\x0f\xc1\x35\x63\xb9\x21\x54\x00\x5d\x89\xed\xfa\xc4\x54\xfa\x03\xd4\x44\x52\xf2\xbf\x14\x3b\xdf\x9b\x0a\x54\xe8\xed\x53\xec\x92\xc1\x49\x1c\x0b\x5d\xb0\x02\x50\xac\x77\x7f\xf4\x58\x6e\xf6\x85\xe6\x63\x2a\x42\x6c\xd6\xd9\x1e\x9c\x0c\x06\x9b\x81\x5c\x90\x01\x2b\xfb\xdf\xc5\x4b\x20\xde\x97\xd8\x9a\xa5\x63\x66\xbd\x90\xec\x91\xa6\x9c\xe2\x44\xaf\xec\x85\xb8\xd6\xac\x7f\x3d\xf6\xf7\x42\x93\xac\x10\xd5\x8d\x62\xf7\x99\x7d\x94\xce\xcc\xb5\xc6\x81\x42\x48\x7a\x4e\x1c\x16\x85\xba\x2f\x49\x97\xfd\x62\xf9\x2a\x1f\x98\x3a\x7f\x33\xd8\xb8\xec\x9b\xeb\x6f\xac\xec\xa9\x9c\xf9\xdb\xe7\xb2\x7f\x88\x56\x10\x36\x58\xeb\x1f\x51\xad\x57\xcf\x7e\x12\x8a\x7f\xe3\xe5\x44\x7c\x25\x83\xbb\xed\x66\xbb\xc1\x5a\x73\x05\x77\xf8\x0c\x00\x00\xff\xff\xf4\xa2\x25\x1d\xbd\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/main.go", size: 445, mode: os.FileMode(420), modTime: time.Unix(1547178186, 0)}
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
	"templates/config.go": templatesConfigGo,
	"templates/k8s.yaml": templatesK8sYaml,
	"templates/main.go": templatesMainGo,
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
		"config.go": &bintree{templatesConfigGo, map[string]*bintree{}},
		"k8s.yaml": &bintree{templatesK8sYaml, map[string]*bintree{}},
		"main.go": &bintree{templatesMainGo, map[string]*bintree{}},
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

