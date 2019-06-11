// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/sql/01_init.down.sql (20B)
// migrations/sql/01_init.up.sql (9.644kB)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __01_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x6b\xf9\xb4\xa3\x14\x00\x00\x00")

func _01_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initDownSql,
		"01_init.down.sql",
	)
}

func _01_initDownSql() (*asset, error) {
	bytes, err := _01_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0xb9, 0x22, 0xf, 0xca, 0x89, 0xa5, 0x5, 0xf2, 0xed, 0x48, 0xd7, 0xe0, 0x9d, 0xec, 0x2, 0xab, 0x87, 0xf3, 0x9d, 0x3d, 0x5a, 0x71, 0x39, 0xe, 0x4a, 0x88, 0xde, 0x1d, 0x0, 0x1d, 0x4e}}
	return a, nil
}

var __01_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x5a\x5f\x73\xdb\xb6\xf2\x7d\xf7\xa7\xd8\x37\xdb\x33\x52\xe7\xf7\x6b\x3b\x77\x9a\x66\xf2\x20\xdb\x4a\xab\x3b\xb6\xec\x4a\xea\x6d\xfd\xc4\x40\xe4\x4a\xc4\x35\x09\x30\x00\x18\x59\xf9\xf4\x77\xf0\x87\x24\x48\x80\xb2\x92\xf6\x25\xb5\xf6\xec\x01\xce\xee\x02\x58\x40\xba\x5d\xcd\x67\x9b\x39\x6c\x66\x37\xf7\x73\x58\x7c\x84\xe5\xe3\x06\xe6\x7f\x2f\xd6\x9b\x35\x94\xa8\x04\x4d\xe5\xc5\xd5\x05\x00\xc0\x74\x0a\x0f\x84\x32\xc8\x68\x89\x4c\x51\xce\xa4\xf9\xf8\x73\x8d\xe2\x48\x33\x58\x2b\x41\xd9\x1e\x6e\x1f\x1f\x1e\xe6\xcb\x0d\x5c\xe6\x44\xe6\xc0\x77\xd6\x0e\x3b\xca\xf6\x28\x2a\x41\x99\xba\x9c\x18\x3f\x89\xe2\x0b\x8a\xc0\x6d\xf1\x04\x5c\x40\xce\xa5\x62\xa4\x44\x4d\x70\x77\xe3\xb0\xce\x31\x23\x8a\x6c\x89\xc4\xc0\xf5\x89\x4b\xb5\x17\xb8\xfe\xe3\xfe\xd7\x16\xd4\x0c\x96\xe6\x58\x92\xc0\xe3\xe1\xd8\x03\xbf\x07\x9f\xc2\xfa\x38\x82\x5a\xa2\x30\x13\x1a\x52\xa4\x05\x45\xa6\x8c\x1d\x34\xc0\xe1\xed\xc7\x89\xd6\x31\xe6\xd2\x57\xea\xfc\xa6\x53\x80\xb5\x22\x2c\x23\x22\x83\x82\x6c\xb1\xb0\x51\x16\x58\x15\x34\x25\x3a\xec\x89\xc4\x90\x73\xe9\x82\xe5\xe1\x40\xa2\x6a\x67\x53\x4b\x15\x09\xf6\xad\xfb\xdc\x9b\x80\x0e\x35\x4d\x31\x51\xc7\x2a\x14\xbb\xd1\x1f\xf2\x5d\x03\x72\x2e\xc8\xbe\x50\xc1\x99\xae\x8a\xc0\x63\xee\xd9\xbc\x51\xc8\xd7\x00\x39\xfb\x42\x68\x41\xb6\xb4\xa0\xea\x08\x5f\x39\x6b\xa0\x02\xf7\x5a\xcc\x10\xbe\xb2\x1f\x7b\x9c\x8c\x67\x98\x94\x3c\xc3\x22\x8c\x0e\xcf\x10\x8c\xa9\x09\x08\x67\x8a\x50\x86\x22\x89\x26\xf5\xb6\x31\x43\x3f\x35\xb7\xb5\x54\xbc\xf4\xf3\xf2\xc9\xfe\xff\x0f\x2f\x78\xfc\x04\x33\x21\xc8\xf1\xca\x92\x5d\x7b\x6c\xbe\x97\x61\x94\x8e\xb2\xf1\xfe\x42\x8a\x1a\xcf\xf4\x37\xd8\x86\x80\xec\x75\x95\x45\x56\xdf\x22\xd3\x6b\x74\x47\x51\xe8\x7c\x19\x18\xa8\x9c\x28\x48\x79\x51\x60\xaa\x80\xb0\x0c\x24\xb2\xac\x59\xe4\x3d\x42\x93\xfc\x39\xab\xcb\x5f\xec\xd2\xd7\xff\x5d\x76\xa6\x84\xb2\x2f\xa4\xa0\xd9\x25\x7c\x80\xff\x9b\x74\x88\xf2\x28\x3f\x17\xd3\x0a\xc5\xce\xad\x1d\xf8\x00\xff\x1f\xd8\x65\xc1\x0f\x05\xdf\x6b\xe3\x8f\xbe\x91\xb3\x3d\xcf\xb6\xd3\x4a\xf0\x1d\x2d\x50\x68\xfb\x4f\xc6\xec\x85\x62\x66\x84\x98\x32\xec\xa9\xe1\xbb\x46\xc7\xaf\xe0\xf8\x27\xa0\x27\xe2\x56\xf1\x04\x50\xa5\x3f\x38\x8d\x15\x0a\xca\xb3\x44\x2a\x22\x14\xdc\x11\x85\x1b\x5a\xa2\x57\xe4\xfa\xaf\x43\x8e\xac\x61\xd7\x75\xc6\x77\xb0\xad\xd3\x17\x54\x60\xdc\x30\xeb\x73\x15\xc8\xf6\x2a\x87\x3f\x17\x4c\xfd\xf4\x63\x47\x75\x57\x0b\xd2\xb8\x7b\x64\x96\xc9\x31\x78\x3b\x63\x90\x45\x13\x30\xc8\xe8\x1e\xa5\x4a\x14\xbe\xaa\xf7\x6e\x3b\x3d\x50\x95\xf3\x5a\x99\x9d\xab\x59\x85\xaf\xa4\xac\x8a\xb0\x96\x1f\x19\x76\xdb\x70\x03\xda\x09\x5e\xea\xfd\x01\x76\xbc\x66\x19\xd0\xc1\x9c\x1c\x2c\xd9\x71\x51\x12\xe5\x2a\xe1\x72\xfe\xf7\xec\xe1\xe9\x7e\x9e\x7c\x7c\x5c\x3d\xcc\x36\xc9\x62\xf9\x9f\xd9\xfd\xe2\xce\x16\x01\x34\x56\x9b\x73\xb8\xfc\xb8\x58\xfe\x36\x5f\x3d\xad\x16\xcb\x8d\xc9\xb4\x97\xc4\x05\xcb\xf4\x16\x85\xb2\x9f\x43\x81\xa4\xe8\x4f\x53\x02\x95\x50\x09\x9e\xd3\x2d\xed\x42\x4e\x65\xa2\x44\xcd\x34\x43\x66\x22\xfe\x4b\x8c\x9a\xee\x22\x5c\x8a\x73\x28\x38\xdb\x9b\xea\x3f\x10\x09\x2d\xd1\x40\xb8\xb7\x00\x5a\xd9\x9b\xe7\xa7\xf9\x50\xf4\x6a\xb6\xbc\x7b\x7c\x68\x34\xaf\xef\x1f\xff\x9a\xaf\xad\x5e\x1d\x82\xd9\x7a\xe3\xfe\xfc\x69\x02\x97\x7f\x2d\x36\xbf\x27\xf3\xd5\xea\x71\xa5\x3f\xf9\x39\x1a\x90\x83\x0e\x48\x3f\x55\x7a\x9a\x15\x4d\x5f\x30\x83\xba\x1a\x4c\xd3\xd5\x7c\x78\xae\xb9\xcf\x83\xbc\x53\x06\xff\x5e\x3f\x2e\xc1\x26\xb6\x59\x10\xac\x2e\x13\x8d\xa3\x28\x13\x5d\x59\xc9\x81\x08\x46\xd9\x5e\xc2\xc7\x82\x93\x5e\x45\xff\xce\x0f\x50\x12\x76\x04\x87\x37\xd3\xd3\x3e\xd0\xfa\x0c\x8b\xe9\x53\x63\xf9\x21\xe5\x59\xbb\xcb\xd9\xa5\xe2\x05\xe1\x9e\x4a\xb3\x8e\x1b\x74\xc4\xbb\x66\xaa\x71\x77\x13\xf3\x77\x49\x6d\xd6\x04\x48\xd2\x13\xb3\x09\xb4\xa2\x10\x5c\x7c\x93\x52\xe3\x11\xca\xb4\x44\x81\xc8\x7f\xfd\x1c\x11\x79\x4f\xa4\xd2\x23\x33\x1e\x78\x7b\x22\x03\xf7\xbe\xc6\x8e\xe4\x94\xca\x50\xd9\xac\x34\x34\x8d\x9d\x32\x50\x39\x95\x7d\x82\xe9\xb4\xed\xfe\xf4\x9f\x96\xed\x98\x28\x5a\x62\x92\x32\x15\x92\x6e\x72\xd4\x5b\xa3\x42\x73\xd4\xe3\x2b\xa6\xb5\xd9\xeb\xb4\x87\x1e\x42\x62\xca\x59\x66\xa3\x58\x62\x5b\x7a\x3d\x62\x59\x97\xdf\x4d\x1c\x25\x2c\x29\x0b\x09\xd7\x25\x29\x0a\x94\xca\x9e\xa4\xed\x1a\x49\x1a\xc2\x5e\x1c\xfa\x74\xe4\x35\xa4\xbb\xa1\xfb\xfd\x77\xb1\x55\xef\xde\x85\x6c\xef\xde\xe9\x23\x25\xd5\x87\x77\x61\xd8\xce\xa3\x2d\x78\xfa\x12\x24\x27\x34\x8e\x06\xd8\xd0\x2a\x0e\x24\xfd\x5c\x53\x81\xa0\x5d\x64\x34\xba\x1d\x99\x17\xdc\x88\xb1\x0b\x55\x68\xf4\x94\x37\x46\xc1\x0f\x32\x91\xba\xc1\x88\x08\xe8\x8c\xa3\x02\x58\x5d\x6e\x6d\xaf\xa3\xc1\xba\xb3\x51\x5a\x8f\xca\xd1\xb5\xe2\x9d\x82\x8e\x2d\xa2\xc0\x33\x86\x0a\x3a\xe3\x98\x02\xbd\xcf\x52\x86\xd9\xa8\x8a\x16\x10\x55\xb2\x1c\xa8\x48\x09\x63\x98\xc1\x14\xd6\xf3\xfb\xf9\xed\x66\xa0\xa1\xe5\x1a\xd3\xd1\x01\x46\xb4\xb4\x80\x31\x3d\x64\xb7\xc3\x54\x9d\xd0\xd3\x02\xce\xd1\x93\xe6\x84\xed\x8d\x9e\x3f\x9f\xee\x66\x9b\xf9\x04\xee\xe6\xf7\x73\xfd\xef\x62\xb9\x9e\xaf\x86\xfa\x5a\xee\x31\x7d\x1d\x60\x44\x5f\x0b\x18\xd3\x27\x90\x8c\x6b\x33\xc6\x73\x2b\x4e\x83\x6d\x5b\xa5\xc8\xb6\x40\x39\xd0\x62\xb8\xc6\x74\x58\xe3\x88\x06\x63\x8c\xcc\xbf\x44\xb1\xc7\xa4\x22\x52\xa2\x8c\x49\xe8\xd9\xcf\x50\x61\xf0\x60\xf1\xb6\x31\xd3\x8b\x47\x72\xa1\x80\x14\x7b\x2e\xa8\xca\x4b\xc8\x89\x84\x9c\x64\x7a\x69\x65\xbc\x93\xd8\x1b\x2b\xa2\xb2\x6f\x0f\x85\xf6\xec\x11\xad\x94\x31\x9e\x6d\x13\xca\x13\x91\xf0\x2a\x2a\x77\x08\x89\x2a\x36\xe7\xa7\x34\xc2\x3a\xe1\x15\xd9\xa3\x4d\x1f\xaf\xd0\xf6\xeb\xd2\x5c\x1b\xb2\xba\xc0\xac\x13\x39\x1c\x21\xa2\x33\x80\x84\x52\x87\x90\x37\xd4\x6e\x8f\x2a\x9e\xde\x10\x14\x55\xbc\xa6\x25\x2d\x88\xd0\x09\x73\x1e\x8b\x47\x3b\xf2\x04\xb6\xb5\xcd\x71\xcd\xa8\xd2\xdd\xb1\xa1\x89\xeb\xb5\x23\xbc\xa1\xd8\x81\x4e\x6b\xb6\xa0\x37\x54\x1f\x08\x8d\x1e\x04\x01\x26\xae\x39\xd7\x0b\x32\xe7\x07\xdb\xeb\x5f\x75\xc7\xd8\x35\x50\x7d\x2c\xf0\x17\x58\x30\xc6\xef\x6e\xec\x91\xa7\x6a\x52\x14\x47\x5b\x02\x3a\x1e\xfa\x52\xe5\x6e\x48\x8a\x0b\xb2\xc7\x78\x4c\xcc\x04\xde\x08\x89\xc5\x9c\x8e\x88\xc1\x8c\x07\x44\x60\x6a\x8f\xce\x37\x82\xd2\xc7\x7d\x73\x60\xb4\x70\x77\xb7\x24\xfa\xb2\xa5\xaf\x08\x7a\x67\xb3\xad\x40\x10\x81\xfe\x68\xe3\x51\x18\xe0\x46\x23\xd1\xc7\x8d\x47\xe3\x73\x8d\x35\xbe\x15\x0a\x0f\xf4\x0f\xe2\x20\x2b\xd3\x70\x52\x95\xa3\x30\x41\xd1\x57\x2c\xc5\x01\x99\x42\x61\x70\xae\x88\xcc\x70\x60\xae\x04\x92\x66\xee\x65\xc2\x7e\xd8\xb8\xe9\x60\xb6\xad\x6b\x10\x4c\x6f\xbe\xe3\x91\xf4\x41\xa3\x61\xf4\x40\xe3\x31\xd4\x1b\x9e\x4c\x32\x2a\x15\x65\xe9\xa9\x38\x0e\x80\xa7\xb6\x54\x52\x55\x82\xbf\xd2\x92\x28\x2c\x8e\x83\x0d\xb6\x66\xf4\x73\x8d\x66\x9f\x95\x5e\x7c\x49\x9a\xa2\x94\x91\x1d\x76\x30\xec\x78\x48\x86\xc0\xd1\xb0\x0c\x80\x91\xd0\xd8\xee\xda\x3e\xe3\xc4\x22\xd2\xb3\x9f\x53\x54\x9d\x4e\x1a\xdc\x4d\x1c\x4d\x44\x58\xdf\x1e\xea\xe9\xd9\x23\x32\xdc\x49\x30\xd2\x47\x7b\xd6\x33\x1a\x02\x83\x6e\x3b\x69\x52\x14\xae\x93\xf6\xe4\x78\x84\x11\x31\xbe\x35\x94\xe2\x59\x23\x42\x54\x59\x25\xb6\x91\x8a\x09\xf1\xac\x6f\xf4\x9d\x0a\xcb\x8a\x0b\x22\x8e\xae\x2d\x83\x54\xa0\x79\x38\xe2\x0c\x4a\x2c\xb9\x38\x9a\xb5\xd9\x66\xab\xd3\xe6\x8d\x11\xd1\xe6\x5b\x43\x6d\x9e\x75\x44\x5b\x46\xe5\xcb\x1b\x02\x7d\xc8\x3f\x50\xa9\x69\x4e\x69\xf4\x87\x19\x11\xda\x83\xc4\xd5\xfa\x90\x53\xe9\x4c\x24\xfd\xfa\x46\x4e\x1d\x24\x5e\xa1\x5c\x91\x02\xd6\xf4\xab\xbd\x04\x9b\x0a\xd5\xda\x74\x75\x06\x41\xa8\x25\x66\xf6\x75\x63\x3c\xb9\x6e\xb0\x53\x19\x6e\x20\x27\xd2\xec\x20\x81\xf0\xe9\x14\x6e\x38\x2f\x90\xb0\xe1\x6b\x4a\x9a\xe4\xf1\x23\xcc\x59\xa2\xea\xff\x30\xfb\xc9\x2d\x49\x73\x84\x9c\xfa\x0b\x71\x57\x17\x45\xa2\xaf\x8b\x31\xca\xce\x38\xba\xea\xed\x4e\x55\xa1\xd8\x71\x51\x62\x06\x04\xb4\x93\x0d\xa4\xb9\x86\x0e\x86\xfa\x2f\xa7\xe3\x43\x19\xe3\xb7\x0e\xa5\x9d\xe0\x8a\xd8\x7f\x9b\xe7\x6d\xca\x32\x7c\x45\x79\x1d\x4b\xdc\xe9\x12\x3a\x3d\x7a\xb3\x3c\x08\x03\x5a\x56\x05\x4d\x75\x03\xac\x8f\x75\x46\x82\x3a\x8a\x8d\xcd\x99\x2d\xf8\x93\x73\x68\x40\xa7\xe7\x22\x87\xe3\x99\x17\x32\xdd\x78\x76\xcb\xd7\x0b\x3e\x2d\x50\x5f\xc9\xa2\xb1\x6f\x6c\x6f\x88\x37\xcb\x82\x40\x03\x8f\x90\x9f\x90\x17\x60\x46\x07\x6b\x90\xf6\x01\xbb\x4d\x77\x20\x49\x62\x81\xa9\xb2\x95\x23\x08\xdb\xe3\x68\x71\x8d\x20\xcf\x38\xca\x34\xce\x5d\x6a\x9d\x7a\xe3\x0f\x12\x89\x48\x73\x3d\x27\x02\x02\x77\x28\x90\xa5\x38\x4c\xbb\x1b\xd6\x8e\x38\x3e\x2d\x6b\xff\x8e\xc9\x18\x47\xa9\x27\xa1\x4c\xd8\x84\x54\xa7\xa7\x90\x63\x1a\xcd\x4d\x04\x75\xf6\x74\x9a\x15\xf7\x82\xc7\xe6\x6b\x19\x4d\x60\x36\xd7\x17\xd4\x45\xa3\x6f\xc9\x64\xa7\x5b\x5f\xf3\xf8\x2c\xf8\xc1\x9b\x9f\xae\x88\xf1\x00\x75\xd6\x33\xe6\xa3\xd1\x6e\x0a\x07\x14\x08\x19\x67\x08\xb5\xd4\x5d\xb4\x8d\xd4\x70\x58\x7e\x88\x9e\x27\x9d\xf1\xcc\x41\x75\x2a\xf8\x61\x48\x3f\xb6\xa9\x76\xc6\xef\xd5\xb4\x3d\xda\x07\x3e\xea\xba\xc5\x41\xce\x19\x4f\xcc\xf6\x97\xe8\x22\x89\xcd\xa0\x0f\x38\x63\x16\xed\x17\x19\xfe\xee\xda\x1b\x70\xcf\x79\xf6\xf6\xa8\x43\xd4\x77\x0c\xad\x29\xfa\xe3\x4f\xa7\x60\xbe\x01\x1e\x1c\x94\x19\x4f\x65\x22\x50\xd5\x62\xe4\x55\xb5\x0f\x38\xe7\xb5\xce\x61\x21\xe3\x69\x5d\xf6\xfb\xd9\x3e\x59\xa4\x29\x18\x00\xc2\x96\xa0\x0f\x88\x3d\x3b\xa2\xac\x38\x93\x78\xe2\xae\x31\x84\x8c\x8a\x6a\x80\xe0\xbe\x7e\xd6\x9d\x60\xbb\xcd\x0b\x94\x75\xa1\xda\x36\xc9\x7b\x8f\x1c\xd0\xc7\x5e\x25\x87\x90\xc8\xdb\xe4\x00\x12\x91\x6a\x62\xe1\xde\xb0\x47\x53\xd7\xd8\xcf\x59\x45\xee\x39\x7c\x2c\x71\x0d\xd5\x58\xde\x5a\xfb\x48\xda\x1a\xbb\x27\xe5\xe2\xfa\x62\xbe\xfc\x6d\xb1\x9c\xc3\x07\x78\x40\xb1\xc7\x8d\x40\xbc\x78\x9a\xad\x36\x8b\xcd\xe2\x71\x09\x37\xcf\xa0\xf8\xf3\xf3\xf3\xf3\xc3\xc3\xdd\xdd\x95\xff\xc3\x82\xeb\x8b\xc7\xd5\xdd\x7c\xa5\x11\x57\xee\x27\x52\x13\xf7\x3b\xa6\x49\xfb\x83\xa3\x49\xfb\xf3\x84\xe6\xf7\x45\x13\xff\x97\x43\x93\xde\x4f\x15\xae\xdf\x5f\xfc\x2f\x00\x00\xff\xff\xf8\xd6\x21\x1e\xac\x25\x00\x00")

func _01_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initUpSql,
		"01_init.up.sql",
	)
}

func _01_initUpSql() (*asset, error) {
	bytes, err := _01_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0xc1, 0x9d, 0x48, 0xeb, 0x4f, 0xf8, 0xce, 0xb6, 0xc, 0x59, 0xe0, 0x33, 0xf9, 0x88, 0x74, 0xa3, 0x61, 0x4c, 0xb9, 0x51, 0x1f, 0x2a, 0xb0, 0xc5, 0xea, 0xc6, 0x4b, 0x5e, 0x86, 0x92, 0xaf}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"01_init.down.sql": _01_initDownSql,

	"01_init.up.sql": _01_initUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"01_init.down.sql": &bintree{_01_initDownSql, map[string]*bintree{}},
	"01_init.up.sql":   &bintree{_01_initUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
