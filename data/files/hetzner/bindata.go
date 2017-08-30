// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package hetzner

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x98\xcd\x6e\xa3\x30\x10\xc7\xef\x79\x8a\x91\xcf\xdd\x2a\x76\x81\xcd\xe6\xda\x44\xbb\x97\xae\xaa\x65\x55\x21\xad\x72\xa0\x61\xd4\xa2\x24\x26\x32\x04\x89\x56\x79\xf7\x95\x21\xc5\xe5\x23\x38\x38\xa0\x5e\x42\xe2\xf1\xdf\xfe\xcd\x64\x06\x7f\xbc\x4f\x00\x08\xf7\x77\x48\xe6\x40\x5e\x31\x79\xe3\x28\xc8\x8d\x6c\x44\x9e\x92\x39\xfc\x9b\x00\x00\x90\x00\x53\x32\x01\x58\xe5\x16\x81\x2f\x61\xc4\xe3\xd2\xfa\x9e\x7f\x02\x90\x6d\xb4\xf6\x93\x30\xe2\x72\xac\x9f\x28\x76\x3e\xcf\xf2\xb1\x72\x63\xa1\x92\xa6\xc5\x52\xb5\xbe\x45\x1c\xd5\x48\xa7\xb9\xc8\xe9\xc7\x2a\x7f\x1e\xcb\x89\x43\x1e\x27\x3e\x5f\xe3\xdf\x6c\x8f\x2d\xd3\xc7\x9b\x83\x1c\x7e\xe9\x59\x53\x35\x41\x80\xf1\x5a\x84\xfb\x0f\xac\x05\x06\xe1\xda\x4f\x30\x80\x3f\x51\x94\x80\x8b\x22\x45\x01\x55\x89\xb4\xbf\x44\x22\x93\xfd\x1f\x22\x9e\xbc\x6e\x33\x50\xba\x42\xf2\xa9\xf7\x5e\xce\x6a\x29\x3f\xfd\x1d\x99\xc3\x1d\x53\x04\x61\xbc\x21\x73\x60\xd3\xe9\xb4\xf0\xe7\xe6\x2c\xf6\x37\xd7\x5d\x18\xa0\x57\x65\x23\xe1\x5b\x9d\xf4\x8f\xde\x0f\xda\x93\xbc\x2a\x31\xa1\x76\x6a\xd4\x94\xcd\xfa\x46\x5d\x42\x18\x44\xbd\x29\x1b\x8b\x5f\x17\x76\xca\xfa\xc7\xbd\xa2\x19\x02\x9c\xd9\x4e\x1d\xdc\xd2\x06\x9e\x32\xb3\xc8\x37\x74\x63\x79\x30\xd3\xd7\xeb\xaf\xec\x59\x84\x81\x49\xc9\xd6\x95\x5f\x52\xb5\x4b\xcf\xea\x9b\x3d\x55\xc9\x38\xd4\xba\xdc\x91\x0c\x66\xaf\xca\x01\x32\x47\x8f\x6f\xeb\xe9\x5d\x03\x74\x77\x58\x6e\xc7\x32\x58\xa1\xa8\x6b\x18\x77\x77\xf8\xc0\xb7\x38\x60\x6b\xf8\xed\xfe\xd9\x6e\x0f\x9c\xed\x4d\x6a\x7d\xb6\xdb\x66\xd9\x5e\x97\x8d\x83\xaf\xc9\xf6\x47\xcf\xe9\xbf\x40\x39\x5f\x1e\x74\xc9\x60\xb4\x3a\xd5\x65\x23\xe1\x77\xaf\x4d\x39\xc6\xef\xa7\x07\x34\xc1\xaf\xea\x46\x4a\x1a\xca\xba\xf8\xef\x3d\x7a\x7e\x0b\x9f\x9e\x60\xab\x9d\xda\x38\x9f\x42\x91\x1c\xfc\xed\x19\x4a\x5a\xdf\x7d\x35\x5f\x27\xdd\x8c\xec\x12\x46\x76\x15\x23\xab\x6f\x53\x9a\xd5\xd7\xcd\x78\x77\x09\xe3\xdd\xa0\x8c\x8d\x3f\x9b\x76\x97\xda\x7d\xd7\x79\x4d\x41\xea\x4e\x68\xfd\x20\xdb\x0e\x0a\xdd\x90\xf6\x25\x90\xf6\x55\x90\xf5\xba\xa1\x6d\xbb\xea\x6e\x4a\xe7\x12\x4a\xe7\x2a\xca\x99\x76\x03\xe4\x7c\x50\x96\x27\xf7\xcd\xe1\x19\x05\xc7\x24\x3f\xb6\x17\xd4\x24\x45\x11\x57\xee\x11\x94\x3f\xca\x2a\xb1\xe8\xed\xf7\x5b\xab\xa4\x68\x3a\xd6\xb0\x17\x57\x17\x6a\xb0\xd3\x05\xc6\x1c\x12\x71\xc0\xb2\xf5\x38\xf9\xfc\x5c\x4d\xe4\xb7\xe3\xff\x00\x00\x00\xff\xff\xa8\x28\x6e\x73\x10\x11\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 4368, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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
