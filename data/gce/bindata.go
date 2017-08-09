// Code generated by go-bindata.
// sources:
// cloud.json
// credential.json
// DO NOT EDIT!

package gce

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9b\xdf\x8f\xa3\x36\x10\xc7\xdf\xf3\x57\x58\x3c\xdd\x49\x09\x0a\x04\xc2\x76\xdf\xba\xdb\x53\x7f\x48\x55\xab\xee\x9e\xfa\x50\xad\x22\x87\x4c\x02\x5a\x02\x9c\x6d\xb2\x97\x9e\xf2\xbf\x57\x10\x42\x88\x81\xd8\xee\xa1\xdd\xe3\xe4\x97\x5d\xb0\x07\xfb\xeb\x0f\x33\x8c\x3d\x52\xbe\x8c\x10\x32\x62\xbc\x05\xe3\x16\x19\x1b\x1f\x8c\x71\xde\x00\xf1\x8e\x1a\xb7\xe8\x9f\x11\x42\x08\x19\x2b\xd8\x15\xcd\x08\x19\x9f\xf0\xe9\x2a\x25\xc9\xea\x74\x9d\xc4\xb0\x4c\x3e\x1b\x23\x84\x9e\x8a\xc7\x09\x6c\xc2\x24\x3e\x8f\xf0\xa5\xf8\x5b\x75\xe4\x53\x7d\xc0\x94\x01\x89\xd1\xc7\x87\x72\x10\x84\x8c\x28\xf1\x31\x2b\xfb\xef\x80\x3c\x43\x04\x7b\x74\x9f\x64\x31\xdb\x8f\xd1\x43\x92\xb1\x00\xdd\x63\x92\x44\x61\x8c\xcf\x0f\xfd\x9b\xc4\x70\x9e\xa9\x68\xca\xe8\x04\x30\x65\xd6\x64\x59\x99\xd5\x5b\xfd\xd6\xd6\x95\x51\x36\x3e\x15\xff\x0f\xe3\x2e\xe5\x7f\xc3\x75\xe5\x8f\x01\xa0\x9f\x70\x14\x01\x1d\xa3\x3f\xfe\x12\x09\x7d\x81\x7c\x72\xcc\x49\x3a\xb6\x2e\x65\x25\xdd\x43\xcc\x08\x8e\x3a\x25\xe5\x0c\xfd\x30\x42\x77\x51\xb6\x5e\xd3\x31\xfa\x35\x79\x11\x12\xf4\x8f\x63\x36\xb5\x55\x1d\x3c\xdd\xaa\x83\x07\x5c\x75\xac\x55\x19\x7f\xc8\x48\x92\x42\xfb\xa2\x1e\x98\x89\x7e\x0e\x42\x1a\xe1\x30\x1e\xa3\x3b\x88\x36\x61\xb6\xbd\xba\x2a\x28\x46\xab\xe0\x8e\x3b\x7a\xfc\xce\x1e\x69\x1f\xf9\x91\x86\x18\xbd\x7b\xc4\xe1\x0b\x8e\xdf\x77\xbc\x93\x00\xc7\x9b\x20\xc3\x95\x83\x1f\xad\xaf\x2e\x00\xd3\x10\x97\xee\x7a\xf1\x56\x6a\xed\xcb\x8e\x76\x5f\x4d\xfa\x6f\x38\xed\x54\xfe\x98\x3c\xef\x93\x31\x2a\x4c\xc4\x72\xe3\x84\xb0\xa0\x4b\x73\xad\xb3\x29\xbc\xd6\xc9\xab\xaf\xbe\x34\x61\x4c\x19\x8e\x7d\x58\xb0\x7d\x0a\x2d\x1f\x1c\xf8\x9c\xfb\x11\x8e\x16\xf4\x39\x2b\xbe\x70\xd6\x84\x6e\x71\x14\x9d\x75\xaf\x80\xfa\x24\x4c\x4f\x8b\x6b\x1a\xf8\x98\xc1\x26\x21\xfb\xc2\xe7\xb8\xae\x34\x1f\xd4\xaa\xee\x09\xde\xe6\xf7\xa6\xd7\xca\x98\xd7\x12\x5b\x93\x5c\xfc\x0a\x93\xd5\xc4\xea\x14\xd4\x61\x75\xa1\xaa\xec\x17\x09\x9b\x99\x9e\xab\xac\xcc\x96\x52\x66\xab\x29\xb3\x39\x65\x9e\xa9\x2e\xcc\x91\x12\xe6\xa8\x09\x73\xf8\x77\xa9\xae\xeb\x46\x4a\xd7\x8d\x9a\xae\x1b\xfe\x55\x4e\xd5\x5d\x6c\x2e\xe7\x63\x73\x45\x27\x9b\x73\xd2\xe6\xea\xd2\x66\xb6\x9c\xb6\x0b\x3b\x19\x71\x33\xde\xd1\x2c\x5b\x5a\x5d\x10\x6e\x82\x2d\x6c\xaf\x07\x40\x8b\x51\x5d\xd7\x2f\xe1\x26\x40\xbf\xc3\x36\xbf\x15\x84\x80\x35\x53\x15\x76\x35\x00\x5a\x8c\x24\x85\xf1\x21\x60\xcf\x55\x85\x5d\x8d\x80\x16\x23\x49\x61\x7c\x0c\xb8\xb6\xaa\xb0\xeb\x21\xd0\x66\x25\x29\xad\x11\x04\xd6\xd4\x51\x15\x27\x08\x82\x56\x33\x49\x79\x8d\x30\xb0\xa7\x37\x2a\xf2\xfc\x34\x13\x87\x01\x67\xd4\x90\x76\xff\xe7\x47\x61\x0c\x98\xca\xb2\x84\x41\xc0\x19\xc9\xc8\xe2\x23\x60\x66\x2a\x85\x40\x3e\xa3\x30\x04\x38\x23\x19\x59\xbc\xff\x7b\xa6\x52\x00\xe4\x33\x8a\x03\x80\xb7\x92\x11\xd6\xf4\x7e\xc7\x54\x72\xff\x7c\x52\x09\xf7\x6f\x98\xc9\x88\x6b\xfa\xfe\xcd\xc9\xcb\xaa\x5d\xe4\x73\xb6\x04\x12\x03\x2b\x76\x90\x47\xb1\xc6\x0a\xd6\x38\x8b\xd8\x82\x02\xcb\xd2\xaa\x39\x3f\x2c\x27\x2b\xa0\xc0\x6a\x4d\x88\xdf\x0b\xdd\x22\xbb\xec\x3a\xd4\x31\x18\x3b\x20\x34\x3f\x18\x2f\x96\xfb\x05\xc4\xbb\xfa\xa8\xf9\x19\xbb\xbe\x79\x3e\x8f\x7d\x7e\x2e\x5f\xa6\x65\xda\xe6\xbc\xb6\x55\x6e\xc2\x2a\x2c\xd0\xbb\x25\x30\xfc\xfe\xd2\x10\xa7\x29\xbd\x90\x5d\xb4\x9e\x17\x3f\xa1\x40\x76\x40\xda\xa7\xe1\x2c\x71\xc4\x2a\xbb\x09\xf6\x79\x53\xca\x30\x61\x93\x0b\xae\xc6\xd4\xb4\x78\xb3\x20\xa1\x6c\x8d\x7d\x46\x8f\x63\xb9\xe6\xdc\xa8\x19\x1c\xb8\x65\xa6\x04\xf2\x37\xbe\x32\x6e\x11\x23\x19\x8c\x5a\xec\xae\x61\xf3\x84\xd8\xbc\x3e\xb0\x79\x92\xd8\xbc\x41\x60\x73\xcc\xa9\x00\x9b\x63\x4e\xbf\x1e\x1b\x3f\x4d\x17\x36\xc7\x9c\x4a\x63\xb3\x4d\xab\x31\xe8\xab\x92\xb3\x84\xe4\xac\x3e\xc8\x35\xdc\xa3\x83\x9c\x35\x1c\x72\x8e\x90\x9c\xd3\x07\x39\x47\x92\x9c\x33\x1c\x72\xae\x90\x9c\xdb\x07\x39\x57\x92\x9c\x3b\x1c\x72\xa2\xac\xea\xf4\x91\x55\xf9\x69\xba\xc9\xc9\x67\xd5\xb9\x69\xff\xf0\x86\xe4\x5c\xe1\x77\xce\xed\xe3\x3b\xc7\x4f\xd3\x45\xce\x55\xf8\xce\xbd\x39\x39\x5b\x48\xce\xee\x83\x9c\x2d\x49\xce\x1e\x0e\xb9\x99\x90\xdc\xac\x0f\x72\x33\x49\x72\xb3\xe1\x90\x13\xe5\x56\xb7\x8f\xdc\xca\x4f\xd3\x4d\x4e\x3e\xb7\xbe\x39\x39\x51\x6e\x75\xfb\xc8\xad\xfc\x34\xdd\xe4\xe4\x73\xeb\x9b\x93\x13\xe5\x56\xb7\x8f\xdc\xca\x4f\xd3\x4d\x6e\x40\xb9\x55\x74\x68\x6d\x58\x28\x22\x93\x3b\xad\xba\x0a\xa7\x55\x29\x64\xde\x05\x32\x35\x28\x9e\xf0\x48\xda\xb0\x50\x82\xc2\x3f\xdd\x05\xc5\x53\x38\x8b\x7e\x15\x94\xf2\xea\xa9\x2a\x64\x7d\xc2\xba\x56\xa4\x6b\x45\xdf\x1c\x36\x5d\x2b\xd2\xb5\xa2\xd7\x27\xa7\x6b\x45\xba\x56\xa4\x6b\x45\x43\xd9\xcf\xea\x5a\x91\xae\x15\xe9\x5a\xd1\x50\xc8\xe9\x5a\x91\xae\x15\xe9\x5a\xd1\x50\xc8\xe9\x5a\x91\xae\x15\x95\x57\xe7\x5a\x51\xf1\x3b\x1d\xd9\x6a\x91\x44\xd9\x43\xd7\x3b\xf4\xd9\x53\x9f\x3d\xf5\xce\x42\xef\x2c\xf4\xce\xe2\x7b\xdd\x59\x34\x92\x68\xf9\x03\x57\x9d\x46\x4f\x84\x74\x1a\xfd\x16\x92\x81\x4e\xa3\xba\x10\xf9\xfa\xe4\x74\x21\x52\x17\x22\x75\x21\x72\x28\xe4\xf4\x71\x41\x1f\x17\xf4\x71\xa1\x17\x64\xba\x10\xf9\x3f\xce\x50\xa3\xd3\xdd\x61\x74\xf8\x2f\x00\x00\xff\xff\xdf\xa4\x43\x0a\x6d\x48\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 18541, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _credentialJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\xbd\x8a\xc3\x30\x10\x04\xe0\x5e\x4f\x31\x6c\x7d\x4f\xe0\xee\xb8\xe2\xda\x83\x2b\x83\x31\x6b\x69\x6d\x94\xc8\x92\xd1\x8f\x49\x30\x7e\xf7\x20\xb9\x08\x49\xb3\xc5\xcc\xb7\xcc\xae\x00\x5a\x63\xd8\xac\x91\x48\x1d\x68\xd6\x42\x5f\x35\x34\x36\xad\x8e\x1f\x35\xfb\x0d\x61\x76\x82\x1f\x17\x8a\xc1\x9f\xe3\x3c\x85\xb8\xbc\xa9\xa1\x26\x9c\x2b\xbe\xa6\xe0\xcf\x6e\xb2\xe2\x4c\xa2\x0e\x17\x05\x00\x7b\xbb\x00\x79\x5e\xa4\xca\x24\x71\xb3\x5a\x06\xd6\x3a\x14\x9f\xdb\x53\x03\x8e\x47\x71\x55\xfc\x9f\x02\xdf\x9f\xc2\xfa\xb5\xb4\xb5\x2c\xf7\xcc\x51\x98\x5a\x73\x28\xa0\x6f\xdb\x99\xe7\xd7\x32\xdd\xca\x28\xd1\x4b\x96\x54\x5d\xaf\x8e\x67\x00\x00\x00\xff\xff\xc5\x8e\x97\xc8\xf7\x00\x00\x00")

func credentialJsonBytes() ([]byte, error) {
	return bindataRead(
		_credentialJson,
		"credential.json",
	)
}

func credentialJson() (*asset, error) {
	bytes, err := credentialJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "credential.json", size: 247, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"credential.json": credentialJson,
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
	"cloud.json": &bintree{cloudJson, map[string]*bintree{}},
	"credential.json": &bintree{credentialJson, map[string]*bintree{}},
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
