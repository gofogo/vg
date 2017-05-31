// Code generated by go-bindata.
// sources:
// data/fish
// data/sh
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

var _dataFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5f\x4f\xe3\x38\x10\x7f\xef\xa7\x18\x4a\xa5\xb6\x48\x6e\x04\x8f\x20\x4e\x02\x15\x41\x25\xa0\x15\xc7\x1d\x0f\xa8\x17\xb9\x89\x93\x58\x24\x76\xce\x76\xd2\xd5\xaa\x1f\x7e\x65\xa7\xa9\xe3\xd4\x65\x5f\xd6\x12\x8a\x98\xf9\xcd\x3f\xff\x66\xc6\x3d\x3f\x0b\x36\x94\x05\x09\x95\x59\x30\x18\x24\x15\x8b\x14\xe5\x0c\xea\x74\x00\x00\x40\x13\x60\x5c\x41\xc4\x2b\xa6\x60\x84\x45\x5a\xc3\x5f\x10\xc4\xa4\x0e\x58\x95\xe7\x06\xa2\x4f\xc4\x8b\x02\xb3\xb8\xb5\xd2\x47\x10\x55\x09\x66\xfe\x25\x2c\x36\xdf\x73\x20\x51\xc6\x1b\x37\x9f\x97\x6b\x23\x93\x5b\xaa\xa2\xcc\x95\x19\x8f\x58\x12\xc0\x91\xa2\x35\x56\xe4\x20\xd5\x27\xac\xd3\xb0\x55\xb4\x76\xb3\x19\xea\xdb\xc6\xe4\xa4\xb5\x47\x65\x4c\xc6\x17\x63\x07\x6b\xab\x6a\xe2\x1c\x8a\xd1\x7f\xf6\xaa\xba\x09\xb5\x97\xf6\x09\x88\xc1\x70\xf4\xef\xe2\xed\xfd\x9f\xbb\xe7\xc7\xe5\x10\x6c\x7a\x9e\x14\x8c\xc3\x83\xe9\xa4\x73\xdd\x53\x40\xb9\x82\xab\x8e\xb9\x24\x0a\x4c\xd5\x57\x6b\x98\x6c\xb0\x24\x0c\x17\x04\x46\xab\x8f\xf9\xd4\xf5\xa5\x81\x28\xfd\x01\x87\x24\xf6\xb7\x75\xb5\xf6\xab\xc3\xd5\xdd\xfb\x13\x8c\x9e\x96\x2f\x0f\xc1\xac\xa6\x42\x55\x38\x4f\x79\x60\x8b\x70\xdd\x86\xd6\x70\xf9\x3c\x7f\x5c\x36\xd6\xcd\xf7\x5b\xe0\xfd\xe2\x55\xe3\xee\x17\xaf\xdf\xc0\x1a\x6f\xc6\x97\x83\x6a\xe3\xb8\x59\x5f\xfb\xe2\xee\x23\xb9\x48\xdd\xec\x8d\xc7\xe2\x2b\xa6\x02\x50\x79\x84\x90\x22\x0a\x7a\x88\x26\x5b\xc7\x7b\x5b\xae\x89\x61\xf3\x6c\xb8\xff\xd9\xe5\x3e\x9c\x2f\xfe\xbe\xbb\x7f\x7e\x08\x57\x6f\xcb\x97\xd5\x7b\xb7\x15\xda\x1e\x92\x80\x22\xd0\x23\x18\x96\x82\x17\xa5\x82\x90\xe7\x71\xd8\x11\x1c\x59\x80\x4f\x69\xe8\xd7\x23\xa6\x9b\x6f\x62\x33\x98\xc2\xd0\x1d\x82\x53\xde\xdb\x49\xf5\x37\x79\xaf\x69\x3d\xa5\x76\x6b\x33\x89\x3c\x08\xc1\xc5\x35\xbc\x72\x48\x39\x6c\xb9\xf8\x92\x25\x8e\x08\x50\xd9\xcc\x36\x01\xac\x40\x65\x04\x0a\x5e\x10\x66\xf3\xa8\x53\xc8\x48\x5e\xf6\x96\x09\x5c\x3a\xeb\xc4\x30\x71\xa0\xd9\xd3\x65\x2e\xaa\x21\xcc\xd7\xb5\x5d\x9c\x0f\xf5\x07\xd8\x25\x5e\xbe\x1c\xfa\xfb\xa4\xfc\xce\x82\xf8\x69\x74\xc7\x9f\xd8\xf1\xf6\x0a\x43\xb7\x7e\xe2\x99\xc2\xd3\xda\xde\x25\x1f\xab\x8d\x79\xaf\x8f\xcc\xb6\xac\x14\xb7\x3b\x1c\x21\xce\x50\x8d\x05\xc5\x9b\x9c\xc0\xea\x63\xde\x5e\xb5\x54\x58\x55\x12\x10\xa2\x12\xed\x77\x31\x92\xd5\x46\x2a\xaa\x2a\xed\xec\xd4\x5b\xb3\x7f\x6c\x12\xca\x62\xc0\x0c\x74\x34\xb4\x8f\xd6\x4c\x4e\x4e\x6c\xd6\x39\x58\x55\x28\x38\x57\x66\x95\x1a\xfd\x36\xa3\x39\x01\x45\xa4\xde\xc5\x3d\xd0\xd9\x2d\x0c\xed\x50\xd1\xa4\x81\xa1\x04\x86\x7d\x68\x67\x97\xba\x53\xa8\xc3\x33\xb2\x0d\x0f\xea\xd0\x6c\xf2\x49\x84\xd5\xf7\x5e\xa6\x8e\x1b\x9a\xc0\x86\xa4\x94\x99\x87\xda\x94\xf4\xbf\xe5\xf7\x06\xb8\xd8\x57\xe0\x09\x75\x76\xdb\x59\x7e\x37\x87\xc1\xea\x9e\x3a\x05\xfb\xd8\x1e\xbb\x70\x57\x4f\xcf\xbe\x43\x4b\x5f\x7d\x0e\x2a\xa3\x12\xa4\x12\xb4\x94\x66\x03\xe4\x58\x2a\x28\xb1\xca\xf4\xb3\x5b\x72\x46\x98\x82\x44\xf0\xc2\x28\xb5\x7c\xe6\xbe\x80\x3d\x3e\x26\xfb\x9f\x16\x3d\xf1\x0e\x24\x89\x61\x2c\x77\xc1\xe7\x7f\xc1\xfa\x62\xb4\xdb\x8d\xa7\xee\x86\x3b\x6e\xc8\xc1\xaf\x00\x00\x00\xff\xff\xfc\x2d\x16\xc0\x13\x09\x00\x00")

func dataFishBytes() ([]byte, error) {
	return bindataRead(
		_dataFish,
		"data/fish",
	)
}

func dataFish() (*asset, error) {
	bytes, err := dataFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fish", size: 2323, mode: os.FileMode(420), modTime: time.Unix(1496224695, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\xdd\x6e\xe2\x3c\x10\xbd\xcf\x53\x4c\x5d\x4b\x85\x4a\x14\xf1\x5d\x52\x45\x2a\xfd\x40\x2d\x52\x81\xa8\xed\xb6\x17\x55\x37\x72\x13\x43\xac\x12\x3b\x8a\x0d\xac\xb6\xf4\xdd\x57\x76\x20\xbf\x26\xcd\x0d\x8a\x7d\x66\x4e\xe6\xcc\xf1\x98\xf3\xb3\xfe\x07\xe3\x7d\x19\x39\x8e\xb3\x5d\x75\xba\xf0\xe5\x00\x00\x04\x44\x52\x40\x78\x80\x80\x71\xb3\xa0\x1f\x12\x28\xb6\x25\x8a\x76\xf3\x15\xfd\xf8\xdb\x95\x7f\xdc\x01\x84\xbf\x6e\x86\xff\x7d\xa3\x0a\xe2\xfa\x3a\x7f\x0d\xe9\xe9\x24\xc5\xde\x8f\x69\x2e\xab\xd1\x34\x88\x04\x04\x64\xbd\x66\x7c\xa5\xbf\x72\x43\xd6\xb0\x5d\x41\x20\xe2\x98\xf0\xb0\x02\x3d\xac\xe9\x6d\x7c\x63\x76\xa8\x24\x81\xe3\x7c\x3b\x4e\xb9\x90\x5c\x08\xb6\x84\x37\xe8\x71\x40\xf8\x65\xfa\xf8\xfc\x6b\xf4\x70\xb7\x40\xf0\x7e\x0d\x2a\xa2\x85\x30\xdb\x55\xa9\x30\xb3\xba\x64\x4e\x96\xfc\x4f\x22\x52\x05\x79\xac\x8b\x07\xa5\xb4\x7f\xdb\xd3\x36\x83\x3b\x1f\x44\x52\x4e\x62\x0a\xd8\x7b\x1d\x77\x5b\xa9\x7c\x6f\xf4\x7c\xef\xe2\xfb\xc5\x6c\xd2\xbf\xda\xb2\x54\xab\xb2\x12\xfd\x82\xb0\x12\xe6\x17\x71\x8b\x87\xb1\xf7\x34\x70\xb1\xf7\x34\x38\x8d\xb8\x5b\x64\xe9\xb3\xdf\x36\xdc\xed\x74\xae\x61\xb7\xd3\x79\x0b\x9f\xc9\x65\x32\x95\x41\x47\x92\x6a\x4d\x43\x0b\xe9\x81\xa6\x0a\xd4\xce\x2e\x83\x4c\x32\x94\x7d\xcb\xd0\xb0\xa1\x53\xcd\xf0\xc7\xd3\xa7\xd1\xed\xc3\xc4\xf7\x1e\x17\x33\xef\xf9\x64\x6f\xb4\x50\xa8\x53\xc4\x75\x41\xcb\x86\x2a\x8d\x89\x3f\x43\x96\x42\x2f\x81\xfa\xe7\xc9\x34\xe8\xd7\x10\x99\x4e\x07\x33\x16\x9e\xaa\xd9\xf1\x27\xdf\xe8\xe3\x30\x49\x53\x91\x0e\x61\x2e\x60\x25\x60\x27\xd2\x4f\x99\x90\x80\x02\x93\xd9\x31\xa6\x40\x94\x8e\x82\x58\xc4\x94\xab\xb2\x95\x23\xba\x4e\xf2\xf7\x94\xaa\x4d\xca\x61\x70\xac\xa8\x56\x3b\x6e\xf8\xc6\xd2\x15\x8b\x23\x6c\x5d\xb6\x19\xac\xd1\xbe\x3a\x61\x6e\x99\x0d\x97\xb4\xe4\xfe\xa6\xa1\x2d\x96\xb3\x79\xd5\xea\xf3\xda\xa9\x32\x0d\xf2\x1e\x27\x2f\x25\xac\xf7\x3a\x76\x11\x6a\x2e\x1b\xff\xf8\xff\x2f\x66\xb3\xd1\x7c\xec\xe2\xea\xbb\xe3\xeb\x3e\x27\xa9\x88\x13\xe5\x1f\x46\x53\xde\xeb\x73\x58\x32\x1e\x02\xe1\x40\x36\x4a\xf4\x0e\x5e\x60\x82\xc3\x92\xad\x69\x61\x07\x84\xe0\xcc\x05\x84\xdb\xa9\x9b\x36\xf9\x21\xa0\xe2\xe1\x8c\x48\x0f\x1d\x04\x56\xae\x57\x0b\x41\xe6\x9d\xb2\x73\xac\xa2\x99\xac\x19\x8d\x91\x43\x57\x9b\x0f\x62\x2d\x75\x4d\x43\x8b\x66\x8e\xd3\x8c\xcc\x75\x2c\x84\xf3\x53\x21\x94\xab\xf9\xcc\xc6\x2e\x62\x6b\x6a\xea\xaa\x41\x32\x41\x4d\x41\xa1\xc8\xcb\xc9\xce\xde\xb2\x09\x2f\x8d\xd7\xa6\x08\xfa\xe1\x74\xe7\xe7\x10\x5f\xcf\x6f\x17\x77\x02\xa2\xda\x53\x55\xef\xb8\x63\xab\xdd\xea\xd9\xef\x09\x40\xb8\x99\xff\x60\x89\x96\x21\x51\x3a\xf0\xf9\xb5\x6b\x49\x54\xc1\x1f\xba\x68\xe9\x6f\x6d\xf7\x1c\x54\xc4\x24\x48\x95\xb2\x44\x9a\x31\xb3\x26\x52\x41\x42\x54\xa4\x6f\xe0\x44\x70\xca\x15\x2c\x53\x11\x9b\x4d\xbd\x7e\x55\xff\x9b\x51\xf4\xab\x63\x06\x5a\x5d\x2a\xd8\x83\xa4\x21\x5c\xc8\x7d\xff\xed\x77\xff\xfd\x12\xef\xf7\x17\x99\x64\xa1\xe0\xc6\x37\x41\x94\xec\x42\x7f\xb9\xe1\x81\x0e\x92\x6e\x07\x7f\xd5\x96\xde\x6e\xde\xbf\x01\x35\xbd\x83\xba\x36\x47\x39\xff\x02\x00\x00\xff\xff\x75\x7e\xa5\xeb\x2d\x09\x00\x00")

func dataShBytes() ([]byte, error) {
	return bindataRead(
		_dataSh,
		"data/sh",
	)
}

func dataSh() (*asset, error) {
	bytes, err := dataShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sh", size: 2349, mode: os.FileMode(420), modTime: time.Unix(1496223361, 0)}
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
	"data/fish": dataFish,
	"data/sh": dataSh,
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
	"data": &bintree{nil, map[string]*bintree{
		"fish": &bintree{dataFish, map[string]*bintree{}},
		"sh": &bintree{dataSh, map[string]*bintree{}},
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
