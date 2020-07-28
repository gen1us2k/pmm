// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/sql/01_init.down.sql (20B)
// migrations/sql/01_init.up.sql (10.025kB)
// migrations/sql/02_postgresql_columns.down.sql (979B)
// migrations/sql/02_postgresql_columns.up.sql (2.04kB)
// migrations/sql/03_add_agent_type.down.sql (233B)
// migrations/sql/03_add_agent_type.up.sql (270B)
// migrations/sql/04_add_tables_column.down.sql (43B)
// migrations/sql/04_add_tables_column.up.sql (56B)
// migrations/sql/05_add_more_std_labels.down.sql (186B)
// migrations/sql/05_add_more_std_labels.up.sql (588B)
// migrations/sql/06_change_agent_type.down.sql (259B)
// migrations/sql/06_change_agent_type.up.sql (289B)

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __01_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x5a\x4d\x73\xdb\x38\x12\xbd\xfb\x57\xf4\x2d\x76\x95\x9d\xda\x49\xa6\xb6\x92\x49\xe5\x20\xdb\x4a\x46\x5b\x96\xec\x91\x34\x3b\xeb\x13\x0d\x91\x2d\x11\x6b\x12\x60\x00\x30\xb2\xf2\xeb\xa7\xf0\x41\x11\x24\x01\x49\xf1\x9c\x12\xab\x5f\xbf\xc6\x6b\x34\x80\x06\xa4\x9b\xf9\x78\xb4\x1c\xc3\x72\x74\x7d\x37\x86\x12\x95\xa0\xa9\x84\xf3\x33\x80\xab\x2b\x98\x12\xca\x20\xa3\x25\x32\x49\x39\x93\x67\x00\x4f\xdf\x6a\x14\x3b\x9a\x3d\xc1\x1d\xdf\xde\x10\x91\x51\x46\x0a\xaa\x76\xe7\x0b\x25\x28\xdb\x5c\xc0\xcd\xfd\x74\x3a\x9e\x2d\xe1\x4d\x4e\x64\x0e\x7c\x0d\xc6\x01\xd6\x94\x6d\x50\x54\x82\x32\xf5\xe6\x52\xf3\x48\x14\xdf\x69\x8a\x09\x23\x25\x1e\x27\x9b\x91\x12\x35\x99\xf3\x82\xf3\xc9\x03\x70\x01\x39\x97\x8a\x39\xd3\xed\xb5\xb1\xa2\x80\xd5\x0e\x32\x5c\x93\xba\x50\x17\x36\x56\x46\x14\x59\x11\x79\x42\x9c\x07\x2e\xd5\x46\xe0\xe2\x8f\xbb\xdf\xa0\xf1\x72\xe3\x4d\x73\x2c\xc9\x71\x86\xe9\xae\xe3\xfc\x09\x7c\x4a\x4b\x62\x09\x6b\x89\xe2\x34\xf1\x69\x41\x91\x29\xd0\x0e\xa0\x3d\xac\xbf\xfd\x34\xd1\x29\x38\x99\xa2\x9b\x35\xc3\x73\x75\x05\xb0\x50\x84\x65\x44\x64\x50\x90\x15\x16\x66\x9a\x05\x56\x05\x4d\x89\xa2\x9c\x25\x12\x4f\x88\xd0\xcc\x90\xe7\x08\x12\x55\x33\xd6\x5a\x2a\x14\xc7\x59\x6e\x2c\xd0\x93\xd9\xd4\x89\xda\x55\x27\xa4\x6a\xb9\xab\xfc\x3a\xb1\x14\xc8\xbe\x53\xc1\x59\x89\xec\x04\x1d\xe3\x16\xec\x8d\x82\xfc\x38\xee\x39\xfa\x4e\x68\x41\x56\x54\x5b\xe1\x07\x67\xce\x55\xe0\x86\x72\x76\xdc\x7d\x6e\x70\x5e\x4c\xc6\x33\x4c\x4a\x9e\x61\x71\x42\xf6\x79\x86\x60\xb0\x2e\xe1\x9c\x29\x42\x19\x8a\x13\xd7\xd7\x4d\x83\x07\xbf\x30\x6e\x6a\xa9\x78\xe9\x55\x85\xfd\xdf\xdb\x67\xdc\x3d\xc1\x48\x08\xb2\x3b\x0f\x13\xfb\xcc\x3e\x87\x61\x97\x76\x88\x8e\xeb\x3b\x29\x6a\x7c\x25\x9b\xf1\x75\x74\x64\xa3\x57\xc3\x29\x1b\xd3\x24\x43\xa6\xe8\x9a\xa2\xd0\x95\x62\xfc\x40\xe5\x44\x41\xca\x8b\x02\x53\x05\x84\x65\x20\x91\x65\xcd\x6e\xe8\x07\xb0\x65\x38\x66\x75\xf9\x41\xef\x91\x00\x6f\xda\xcf\x13\xca\xbe\x93\x82\x66\x6f\xe0\x33\xfc\xeb\xd2\x5a\xcb\x9d\xfc\x56\x5c\x55\x28\xd6\x6e\xed\xc3\x67\xf8\xa5\x63\x93\x05\xdf\x16\x7c\xa3\x0d\xef\x1a\x03\x67\x1b\x9e\xad\xae\x2a\xc1\xd7\xb4\x40\xa1\x6d\xef\xcf\x00\xfc\x62\x33\xc3\x36\xe5\xde\x19\x3b\x5f\x37\xa3\xfe\x0d\x1c\xf3\x25\xe8\xf0\x6e\xef\xb9\x04\x54\xe9\x5b\xab\xa8\x42\x41\x79\x96\x48\x45\x84\x7a\x82\x5b\xa2\x70\x49\x4b\xf4\x16\x93\xfe\x6b\x9b\x23\x6b\xd8\x75\x7d\xf2\x35\xac\xea\xf4\x19\x15\x18\x3f\xcc\x3a\x5c\x05\xb2\x8d\xca\x9f\xe0\xcf\x09\x53\xef\xdf\xb5\x54\xb7\xb5\x20\x8d\xbb\x47\x66\x99\x2c\x83\x77\x42\x1c\x9f\x43\x93\x3a\xc8\xe8\x06\xa5\x4a\x14\xbe\xa8\x4f\xee\x9c\xd9\x52\x95\xf3\x5a\x99\xfd\xd7\xad\xfe\x17\x52\x56\x05\x3e\x81\x25\x69\x39\xee\x19\xb6\xe7\x93\x43\xc1\x5a\xf0\x52\xef\x5b\xb0\xe6\x35\xcb\x80\x76\xc7\xe8\x50\xc9\x9a\x8b\x92\xa8\x6e\x1d\x8c\xff\x37\x9a\x3e\xdc\x8d\x93\x2f\xf7\xf3\xe9\x68\x99\x4c\x66\xff\x1d\xdd\x4d\x6e\xfd\x5a\x70\x08\xbf\x04\xbe\x4c\x66\x5f\xc7\xf3\x87\xf9\x64\xb6\x34\x05\xd0\x9d\xe4\x09\xcb\xf4\x56\x8a\xb2\x3b\xc7\x02\x49\xd1\x1d\xb6\x04\x2a\xa1\x12\x3c\xa7\x2b\xba\x9f\x12\x2a\x13\x25\x6a\xa6\x09\x32\x3b\x23\x1f\x42\xd4\x74\x1d\xe0\x52\x9c\x43\xc1\xd9\xc6\xac\x85\x2d\x91\xb0\x67\xea\x26\x62\xb8\x1c\x9a\x34\x2c\x1f\x1f\xc6\xa1\x24\xcc\x47\xb3\xdb\xfb\xa9\x9f\x83\xc5\xdd\xfd\x5f\xe3\xc5\xd2\x5f\x00\x5f\x46\x8b\xa5\xfb\xe8\xbd\xfb\xe8\xaf\xc9\xf2\xf7\x64\x3c\x9f\xdf\xcf\xf5\xa7\xbf\xc6\x12\xb5\xd5\x89\xea\x4e\xa9\x1e\x7e\x45\xd3\x67\xcc\xa0\xae\xba\xc3\x77\x4b\x65\x58\x1b\x53\xd7\x07\x0d\xea\x83\x32\xf8\xcf\xe2\x7e\x06\xb6\x02\xdc\x42\x62\x75\x99\x68\x18\x45\x99\xe8\x02\x4c\xb6\x44\x30\xca\x36\xf2\x09\xbe\x14\x9c\x74\x56\xc2\xef\x7c\x0b\x25\x61\x3b\x70\x0e\x66\x78\xda\x09\x1a\xa7\x7e\xd1\x35\x9f\xbf\x4d\x79\xb6\xdf\x29\xed\x02\xf3\x52\x70\x47\xa5\x59\xfd\x0d\x7a\xe0\x5b\xeb\x75\x65\x9d\xdd\xa0\x3a\xfb\x7f\xcd\x8c\x3b\x92\x34\x3e\x92\x81\x4e\x14\x82\x8b\x9f\x53\x69\x5c\x06\xeb\xca\xf0\x0c\x04\xfe\xfb\xd7\x80\xc0\x3b\x22\x95\x0e\xcc\x78\xcf\xd7\x13\x38\x70\xee\xea\x6b\x29\x0e\x28\x0c\xa8\x1a\x95\x86\xa6\xd1\x44\x19\xa8\x9c\x4a\x9f\x40\xf7\xcc\xb6\x74\x34\x9b\xe5\xda\x25\x8a\x96\x98\xa4\x7a\x78\x03\xca\x65\x8e\x7a\x23\x55\x68\x1a\x0e\x7c\xc1\xb4\x36\x3b\xa3\x76\xd1\x01\x24\xa6\x9c\x65\x36\x7f\x25\x36\x05\xd7\x21\x96\x75\xf9\x7a\xe2\x10\x61\x49\x59\x80\x70\x51\x92\xa2\x40\xa9\xec\xa9\xbb\x5f\x19\x49\x43\xe8\xa7\xb1\x4b\x47\x5e\x02\x74\xd7\x74\xb3\x79\x0d\x5b\xf5\xf1\x63\x80\xed\xe3\x47\x7d\xc8\xa5\xfa\x60\x2f\x0c\xdb\x49\xb4\x05\x4f\x9f\x87\x93\x33\x30\xc6\x13\x6c\x68\x15\x07\x92\x7e\xab\xa9\x40\xd0\x3e\x32\x94\xdd\x96\xcc\x4f\xee\xd0\xe8\xa5\x6a\x60\xf4\x95\x3b\xa3\xe0\x5b\x99\x48\xdd\x7f\x04\x04\xb4\xc6\xb8\x00\x56\x97\x2b\xdb\x07\x69\xb4\xee\x7a\x94\xd6\xa3\x72\x04\x7b\x65\xd8\x2b\x68\xd9\x02\x0a\x3c\xe3\x50\x41\x6b\x8c\x29\xd0\x9b\x2b\x65\x98\x45\x55\xec\x01\x61\x25\xb3\x9e\x8a\x94\x30\x86\x19\x5c\xc1\x62\x7c\x37\xbe\x59\x76\x35\xec\xb9\x62\x3a\x5a\x40\x44\xcb\x1e\x10\xd3\x43\xd6\x6b\x4c\xd5\x01\x3d\x7b\xc0\x49\x7a\xd2\x9c\xb0\x8d\xd1\xf3\xe7\xc3\xed\x68\x39\xbe\x84\xdb\xf1\xdd\x58\xff\x3b\x99\x2d\xc6\xf3\x9e\xbe\x3d\x77\x4c\x5f\x0b\x88\xe8\xdb\x03\x62\xfa\x04\x92\xb8\x36\x63\x3c\xb9\xe2\x34\xda\x36\x5d\x8a\xac\x0a\x94\x5d\x2d\x86\x2b\xa6\xc3\x1a\x23\x1a\x8c\x31\x30\xfe\x12\xc5\x06\x93\x8a\x48\x89\x32\x24\xa1\x63\x3f\x45\x85\x71\x00\xeb\x60\xdb\x34\xbd\x78\x24\x17\x0a\x48\xb1\xe1\x82\xaa\xbc\x84\x9c\x48\xc8\x49\xa6\x97\x56\xc6\xf7\x12\x3b\xb1\x02\x2a\xbb\xf6\xa1\xd0\x8e\x3d\xa0\x95\x32\xc6\xb3\x55\x42\x79\x22\x12\x5e\x05\xe5\xf6\x21\x61\xc5\xe6\xf4\x94\x46\x58\x2b\xbc\x22\x1b\xb4\xd3\xc7\x2b\xb4\xdd\xbd\x34\x97\x8c\xac\x2e\x30\xdb\x8b\xec\x47\x08\xe8\x1c\x40\x86\x52\xfb\x90\x23\x6a\x57\x3b\x15\x9e\xde\x21\x28\xac\x78\x41\x4b\x5a\x10\xa1\x27\xcc\xb9\x4c\xee\x6d\xe8\x4b\x58\xd5\x76\x8e\x6b\x46\x95\xee\x95\x0d\x4f\x50\xaf\x8d\x70\x44\xb1\x03\x1d\xd6\x6c\x41\x47\x54\x6f\x09\x0d\x1e\x04\x03\x4c\x44\x73\xae\x17\x64\xce\xb7\xb6\xf3\x3f\x6f\x8f\xb1\x0b\xa0\xfa\x58\xe0\xcf\x30\x61\x8c\xdf\x5e\xdb\x23\x4f\xd5\xa4\x28\x76\xb6\x04\x74\x3e\xf4\x95\xcb\xdd\x9f\x14\x17\x64\x83\xc1\x9c\x98\x01\x1c\x49\x89\xc5\x1c\xce\x88\xc1\xc4\x13\x22\x30\xb5\x47\xe7\x91\xa4\x74\x71\x3f\x9f\x18\x2d\xdc\xdd\x3c\x89\xbe\x7a\xe9\x8b\x81\xde\xd9\x6c\x2b\xd0\xcf\x40\x37\x5a\x3c\x0b\x3d\x5c\x34\x13\x5d\x5c\x3c\x1b\xdf\x6a\xac\xf1\x58\x2a\x3c\xd0\x3f\xc9\x83\xac\x4c\xc3\x49\x55\x8e\xc2\x24\x45\x5f\xac\x14\x07\x64\x0a\x85\xc1\xb9\x22\x32\xf1\xc0\x5c\x06\x24\xcd\xdc\x3b\x86\xfd\xb0\x71\xd3\xc9\xdc\xb7\xae\xfd\x64\x7a\xe3\x8d\x67\xd2\x07\x45\xd3\xe8\x81\xe2\x39\xd4\xfb\x9d\x4c\x32\x2a\x15\x65\xe9\xa1\x3c\xf6\x80\x07\xb7\x54\x52\x55\x82\xbf\xd0\x92\x28\x2c\x76\xbd\x0d\xb6\x66\xf4\x5b\x8d\x66\x9f\x95\x5e\x7e\x49\x9a\xa2\x94\xc3\x1d\xb6\x17\x36\x9e\x92\x3e\x30\x9a\x96\x1e\x30\x90\x1a\xdb\x5c\xdb\x47\x9f\x50\x46\x3a\xf6\x93\x8a\xaa\xd5\x49\xfb\x77\x13\x47\x13\x10\xd6\xb5\x0f\xf5\x74\xec\x01\x19\xee\x24\x88\xf4\xd1\x9e\xf5\x94\x86\xc0\xc0\xf7\x9d\x34\x29\x0a\xd7\x49\xb7\x72\x3c\xc2\x80\x18\xdf\x3a\x94\xe2\x59\x03\x42\x54\x59\x25\xb6\x8f\x0a\x09\xf1\xac\xc7\xfa\x4e\x85\x65\xc5\x05\x11\x3b\xd7\x96\x41\x2a\x90\xe8\xed\x8d\x33\x28\xb1\xe4\x62\x67\xd6\xe6\x7e\xb6\xf6\xda\xbc\x18\x01\x6d\xbe\x75\xa8\xcd\xb3\x46\xb4\x65\x54\x3e\x1f\x11\xe8\x43\xfe\x89\x4a\xcd\x73\x40\xa3\x1f\x26\x22\xb4\x03\x09\xab\xf5\x21\x87\xa6\x33\x91\xf4\xc7\x91\x39\x75\x90\x48\x85\x72\x45\x0a\x58\xd0\x1f\xf6\x12\x6c\x2a\x54\x6b\xd3\xd5\x39\x48\x42\x2d\x31\xb3\x6f\x1b\xd1\xc9\x75\xc1\x0e\xcd\x70\x03\x39\x30\xcd\x0e\xd2\x17\x7e\x75\x05\xd7\x9c\x17\x48\x58\xf3\x94\xed\x56\x71\x9a\xe4\xe1\x03\xcc\x59\xc2\xda\xff\x30\xbb\xc9\x0d\x49\x73\x84\x9c\x7a\xcb\x70\x5d\x17\x45\xa2\xef\x8a\x21\xca\xd6\x18\x5f\xf3\x76\x9f\xaa\x50\xac\xb9\x28\x31\x03\x02\xda\xcb\xa6\xd1\x5c\x42\xbb\xa1\xfe\xcf\x69\x3c\x94\x31\xfe\x74\x28\xed\x05\xe7\xc4\xfe\xdb\x3c\x7d\x53\x96\xe1\x0b\xca\x8b\xc0\xb4\x1d\x2e\xa0\x23\xd1\x9b\xc5\x41\x18\xd0\xb2\x2a\x68\xaa\xdb\x5f\x7d\xa8\x33\x32\xa8\xa2\x40\x6c\xce\x6c\xb9\x1f\x1c\x43\x03\x3a\x32\x16\xd9\x8f\x67\xde\xc7\x74\xdb\xd9\x2e\xde\x36\xf9\xb4\x40\x7d\x1f\x0b\xe6\xbe\xb1\x1d\x13\x6f\x16\x05\x81\x06\x3f\x24\x3f\x20\x6f\x80\x89\x07\x6b\xa0\xf6\xd1\x7a\x3f\xdd\x7d\x49\x12\x0b\x4c\x95\xad\x1c\x41\xd8\x06\xa3\xc5\x15\x41\x9e\x72\x90\x69\xa0\xbb\xd2\x3a\xf5\x86\x00\x24\x12\x91\xe6\x7a\x4c\x04\x04\xae\x51\x20\x4b\xb1\x37\xed\x2e\xac\x8d\x18\x1f\x96\xb5\xbf\x66\x30\xc6\x53\xea\x41\x28\x93\x36\x21\xd5\xc1\x21\xe4\x98\x06\xe7\x26\x80\x3a\x7d\x38\xcd\x8a\x7b\xc6\x5d\xf3\x15\x8d\x66\x30\x5b\xeb\x33\xea\xa2\xd1\x77\x64\xb2\xd6\x8d\xaf\x79\x78\x16\x7c\xdb\x8e\x4f\x57\x44\x3c\x41\xad\xf5\x94\xf1\x68\xb8\x1b\xc2\x16\x05\x42\xc6\x19\x42\x2d\x75\x0f\x6d\x33\xd5\x0b\xcb\xb7\xc1\xd3\xa4\x35\x9e\x1a\x54\x4f\x05\xdf\xf6\xe8\x63\x9b\x6a\x6b\x7c\xb5\xa6\xd5\xce\x3e\xef\x51\xd7\x2b\x76\xe7\x9c\xf1\xc4\xec\x7e\x89\xae\x91\xd0\x08\xba\x80\x53\x46\xb1\xff\x02\xc3\xdf\x5d\xfd\x80\x1b\xce\xb3\xe3\x51\xfb\xa8\xd7\x84\xd6\x1c\x7e\xfc\xab\x2b\x30\xdf\x14\x77\x8f\xc9\x8c\xa7\x32\x11\xa8\x6a\x11\x79\x51\xed\x02\x4e\x7a\xa9\x73\x60\xc8\x78\x5a\x97\x9d\x5e\xb6\x4b\x16\x68\x08\x7a\x80\x61\x3b\xd0\x05\x84\x9e\x1c\x51\x56\x9c\x49\x3c\x70\xcf\xe8\x43\xe2\xa2\x1a\x24\x58\xa4\xe9\x02\xf7\x9b\xbc\x40\x59\x17\x6a\xdf\x22\xb5\x6f\x91\x3d\xfa\xd0\x8b\x64\x1f\x12\x78\x97\xec\x41\x02\x52\x4d\x2e\xdc\xf3\x75\x74\xea\x1a\xfb\x49\x6b\xc8\x3d\x85\x47\x26\xae\xa1\x8a\xcd\xdb\xde\x1e\x99\xb6\xc6\xee\x4b\x39\xbb\x80\xf1\xec\xeb\x64\x36\x86\xcf\x30\x45\xb1\xc1\xa5\x40\x84\x87\xd1\x7c\x39\x59\x4e\xee\x67\x70\xfd\x08\x8a\x3f\x3e\x3e\x3e\x4e\xa7\xb7\xb7\xe7\xfe\x6f\x10\x2e\xce\xee\xe7\xb7\xe3\x39\x5c\x3f\x9e\x01\xd8\xaf\x93\xdd\xaf\xcd\x2e\xcd\x1f\xfe\x4f\xc6\xec\x27\xcd\xaf\xac\x9c\xdd\xfe\xc2\xc1\xfc\xbf\xf9\x75\x95\xfd\xcb\xfb\xad\x94\xfd\xc0\x0f\x6b\xbe\x4a\x5e\x8c\x97\xcb\xc9\xec\xeb\xc2\xae\xaf\x64\x23\x08\xab\x0b\x22\xa8\xda\xc1\x67\xf8\xf0\xcb\xc7\x77\x9f\xce\xfe\x0e\x00\x00\xff\xff\x92\x74\x17\xdb\x29\x27\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x65, 0x78, 0x4f, 0x3a, 0x62, 0x59, 0x56, 0xef, 0x30, 0x98, 0x8a, 0x1c, 0x8f, 0xd6, 0x41, 0xcc, 0x84, 0xc2, 0x87, 0x2f, 0x24, 0x2f, 0x77, 0x57, 0xc3, 0x69, 0x3d, 0xb7, 0xcf, 0x1c, 0x7, 0x17}}
	return a, nil
}

var __02_postgresql_columnsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x4d\x0e\x82\x30\x10\xc5\xf1\xbd\xa7\xe8\x01\xbc\x81\x2b\x54\x76\x28\x86\xe0\xba\x7c\x4d\x42\x43\x8b\xa6\x1d\xe3\xf5\x0d\x1a\x12\xad\x94\xd7\xfd\x3f\xbf\x69\x5f\x92\x95\x69\x21\xca\x64\x9f\xa5\xc2\x10\x5b\xd5\xba\x8d\x10\xc7\x22\xbf\x88\x43\x9e\x5d\x4f\x67\x51\x19\xe9\xfa\xda\x52\x27\x1b\x3d\x38\xd9\x2b\x96\xed\xc8\xd5\x16\x67\xee\x61\x50\x66\xa9\xee\x62\xb8\x77\x17\xe1\x75\xca\xb2\xa2\x28\x72\x4e\x23\xd4\xa7\x55\xcc\x34\xc6\xa8\x73\x1a\x50\xf5\xad\xad\x35\x1c\xd2\xab\xb0\xb5\x36\xa3\x9f\x61\x0d\x8c\xb8\x50\x62\x13\x4c\xb8\x50\x06\x4c\x26\x73\xc7\x9f\xf6\x2a\x68\x81\xe7\xfd\x87\x01\xb1\xd1\xc3\xe7\x22\x2b\x43\x21\xed\x37\x5a\x91\xa6\x63\x04\xa9\xaf\x6a\xb2\x76\xaf\x00\x00\x00\xff\xff\x92\x15\x3e\x17\xd3\x03\x00\x00")

func _02_postgresql_columnsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_postgresql_columnsDownSql,
		"02_postgresql_columns.down.sql",
	)
}

func _02_postgresql_columnsDownSql() (*asset, error) {
	bytes, err := _02_postgresql_columnsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_postgresql_columns.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xba, 0xff, 0x8a, 0x34, 0x78, 0x29, 0x20, 0x64, 0xd4, 0x9c, 0x10, 0x61, 0x92, 0xf3, 0x20, 0xc, 0x39, 0xb0, 0xf0, 0x75, 0x33, 0x2e, 0x61, 0x44, 0x22, 0x32, 0x16, 0x53, 0x75, 0xb5, 0x32, 0xd4}}
	return a, nil
}

var __02_postgresql_columnsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x3b\x4f\xc3\x30\x14\x85\xf7\xfe\x8a\xbb\x15\xa4\x8a\x01\x46\xa6\xd2\x96\xa9\x0f\x09\x85\x39\x75\x9c\x5b\x72\x15\x3f\x2a\xfb\x56\x15\xfc\x7a\xe4\x3e\x20\x7d\xd0\xd8\x11\x4b\xa6\xf3\xe9\x7c\xbe\x39\xc3\x69\x36\x79\x83\x6c\xf8\x32\x9d\x80\x46\x76\x24\x7d\x0f\x60\x38\x1e\xc3\x68\x31\x7d\x9f\xcd\x61\xa9\x73\x5f\x09\x87\x65\x5e\xa8\xda\xe7\x15\x71\x2e\x0d\x2f\xe1\x55\x59\xc1\x4f\x8f\x83\xd6\xb4\xdf\xe8\x9f\x34\x8c\x16\xb3\xd9\x64\x9e\x41\x3f\xb3\x2c\x14\x98\x8d\x2e\xd0\x81\x5d\xc1\x9e\x82\x42\x59\x59\x7b\x90\x42\x56\x08\x15\xb1\x87\xe2\x13\xb8\x42\xf0\x2c\x18\x35\x1a\xee\xdf\x6e\x74\x28\xca\x04\xc1\x5d\xbc\x8b\x61\x00\x2f\xdc\x1e\x5a\xe4\x4a\x72\x4c\x98\xe2\x77\x24\xba\x28\x1e\xd8\x64\xcb\xad\x23\x66\x34\x09\x96\x47\xa2\x8b\xe5\x81\x8d\xb2\x54\x56\x0a\x15\xbb\xc3\xb3\x70\x9c\xdb\x0e\xda\xab\x25\x8e\xb0\x51\x17\xb1\xc1\xf3\x74\xb2\x5d\xc2\x02\x1b\x5d\x71\x03\xbc\x02\x74\xf0\x4b\x99\x5f\xa3\x31\x6e\x7d\x57\x80\x0e\x8a\x29\xdb\x63\xd4\xeb\xe8\xdf\x7b\x16\x8e\x53\x0b\x50\xf2\xcf\xfd\x6d\x8a\x3b\xdc\x65\x3e\x5d\x2e\xe5\x6c\x85\xaa\xf7\x37\x60\xd2\xd8\xe2\x76\x9a\xbd\xe5\x15\x02\xa7\xdd\xe0\xd7\xe1\x1b\x70\x32\x1f\x07\xd3\x01\x90\x01\x4d\x4a\x91\x47\x69\x4d\xe9\xe1\x8e\x56\xc0\x4e\xc8\x3a\x27\x1b\x6a\x42\x96\x3c\xa0\x11\x85\xc2\x72\x00\x96\x2b\x74\x5b\xf2\x08\x5f\xe8\xec\xfd\x5f\x2f\x0a\x17\xc0\xd8\x27\x35\xc2\x9d\xde\x14\xf8\x7f\x7d\x53\x0f\xe0\xf9\x3b\x00\x00\xff\xff\xb5\x67\xf1\x34\xf8\x07\x00\x00")

func _02_postgresql_columnsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_postgresql_columnsUpSql,
		"02_postgresql_columns.up.sql",
	)
}

func _02_postgresql_columnsUpSql() (*asset, error) {
	bytes, err := _02_postgresql_columnsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_postgresql_columns.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe2, 0x1d, 0xed, 0x3f, 0xb, 0xb3, 0x8d, 0x0, 0x6e, 0x61, 0xd9, 0x9f, 0x1, 0x3e, 0xeb, 0xe5, 0x5c, 0x54, 0x66, 0xbf, 0xcf, 0xd5, 0xc0, 0x15, 0xcf, 0xf, 0x81, 0xc6, 0x35, 0x31, 0x5a, 0x93}}
	return a, nil
}

var __03_add_agent_typeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xcd\x4a\xc4\x30\x18\x45\xf7\xf3\x14\x77\x17\x85\x8e\xf8\xb3\x11\xc5\x45\x1d\x23\x08\xcd\x14\x24\x2e\x5c\xcd\xc4\xf4\x6b\x1b\xc8\x9f\x49\x54\xfa\xf6\x12\xaa\xe8\xf6\x1c\xb8\xe7\xb6\x9d\xe4\xcf\x90\xed\x7d\xc7\xe1\xa8\x24\xa3\x33\x44\xff\xf0\xf4\xf8\x8a\x5d\xdf\xbd\x88\x3d\x8e\x6a\x22\x5f\x0e\x65\x89\x74\x04\xf7\x1f\xee\xfa\x64\x03\xb0\x3f\x7a\x30\xfe\x53\x59\x33\x30\xdc\xe1\xbc\xa9\xce\x2d\xf9\xdd\x6e\x23\xa5\x31\xeb\x99\x9c\xaa\xe6\xe2\x9f\xc9\x36\x7c\xd9\x30\x55\x7c\xb9\xe2\xe0\xa7\x30\xbc\x6d\x63\x0a\xa3\xb1\x94\xaa\xb9\xda\x00\xa7\xd8\xf5\x42\xf0\xbd\x04\x6b\x6b\x0f\x72\x89\x84\x32\xab\x02\x1d\xac\x25\x5d\x10\xc6\xdf\xdf\x37\xf8\xd9\x6d\x50\xd3\x58\xdb\x0d\xa8\xe8\x33\x76\xfb\x1d\x00\x00\xff\xff\xf5\x7c\x8f\xa8\xe9\x00\x00\x00")

func _03_add_agent_typeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_agent_typeDownSql,
		"03_add_agent_type.down.sql",
	)
}

func _03_add_agent_typeDownSql() (*asset, error) {
	bytes, err := _03_add_agent_typeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_agent_type.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x22, 0xa4, 0x18, 0x3c, 0x12, 0x1b, 0x61, 0xa3, 0x28, 0x96, 0x96, 0x58, 0xc5, 0x18, 0xa0, 0x95, 0xdb, 0xf8, 0xfc, 0xb5, 0x86, 0xad, 0x3d, 0xdc, 0xb4, 0xf9, 0xac, 0x96, 0x9e, 0xdf, 0xbb, 0x8f}}
	return a, nil
}

var __03_add_agent_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x4f\x4b\xc3\x40\x10\x47\xef\xfd\x14\xbf\x5b\x14\x52\xf1\xdf\x41\x14\x0f\xb1\x46\x10\x92\x06\x24\x1e\x3c\xb5\xeb\x76\xb2\x0d\xec\x3f\x77\x47\x25\xdf\x5e\xa6\x51\xec\x61\x2f\xef\x2d\x33\x6f\xaa\xa6\xaf\x5f\xd0\x57\x0f\x4d\x0d\x47\x9c\x46\x9d\xd1\x76\x8f\xcf\x4f\x6f\x58\x75\xcd\x6b\xbb\xc6\x56\x19\xf2\xbc\xe1\x29\xd2\x16\xb5\xff\x74\x37\x27\x0b\xa0\xf8\xa7\x9b\xd1\x7f\x29\x3b\xee\x0a\xdc\xe3\xbc\x14\xe7\xa6\xfc\x61\x97\x91\xd2\x90\xf5\x9e\x9c\x12\x73\x71\x64\xb2\x0d\xdf\x36\x18\xc1\x97\x33\x0e\xde\x84\xdd\xfb\x32\xa6\x30\x8c\x96\x92\x98\xab\x83\x89\x21\xb3\x49\x74\x98\x67\x32\x2b\x96\x47\x8e\x3c\x67\xf9\x74\xbd\x00\x4e\xb1\xea\xda\xb6\x5e\xf7\x28\x2a\x89\x42\x3f\x45\x02\xef\x15\x43\x07\x6b\x49\x33\xc2\xf0\x77\xdc\x2d\x7e\x97\x97\x90\x3e\xcc\x81\x25\x88\xf5\x59\x71\xf7\x13\x00\x00\xff\xff\x6e\xa0\x3b\x29\x0e\x01\x00\x00")

func _03_add_agent_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_agent_typeUpSql,
		"03_add_agent_type.up.sql",
	)
}

func _03_add_agent_typeUpSql() (*asset, error) {
	bytes, err := _03_add_agent_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_agent_type.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0xfb, 0x74, 0x75, 0xaa, 0xed, 0x12, 0xc4, 0xfd, 0xf4, 0x7, 0x2a, 0x33, 0xa, 0xd0, 0x6b, 0xea, 0x23, 0x4e, 0x77, 0x74, 0x40, 0xec, 0xc1, 0x17, 0xcc, 0x65, 0x30, 0xf, 0x7b, 0x2b, 0xa5}}
	return a, nil
}

var __04_add_tables_columnDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x28\x49\x4c\xca\x49\x2d\x4e\xb0\x06\x04\x00\x00\xff\xff\x3d\x3a\x24\xe4\x2b\x00\x00\x00")

func _04_add_tables_columnDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__04_add_tables_columnDownSql,
		"04_add_tables_column.down.sql",
	)
}

func _04_add_tables_columnDownSql() (*asset, error) {
	bytes, err := _04_add_tables_columnDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "04_add_tables_column.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0x49, 0xd1, 0x9f, 0x81, 0xac, 0x26, 0x7, 0x57, 0xdf, 0x33, 0x19, 0x94, 0xd2, 0x85, 0x4b, 0xaa, 0x8a, 0x46, 0xd5, 0xae, 0xaa, 0x54, 0xd, 0x85, 0x68, 0xc4, 0xbf, 0x1a, 0x56, 0xdf, 0x8a}}
	return a, nil
}

var __04_add_tables_columnUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xe6\x52\x50\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x28\x49\x4c\xca\x49\x2d\x4e\x50\x70\x2c\x2a\x4a\xac\xd4\x08\x2e\x29\xca\xcc\x4b\xd7\xb4\x06\x04\x00\x00\xff\xff\x58\x41\xc6\xd2\x38\x00\x00\x00")

func _04_add_tables_columnUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__04_add_tables_columnUpSql,
		"04_add_tables_column.up.sql",
	)
}

func _04_add_tables_columnUpSql() (*asset, error) {
	bytes, err := _04_add_tables_columnUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "04_add_tables_column.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa0, 0xad, 0x9, 0xc0, 0x4f, 0x33, 0x6d, 0xf1, 0x16, 0x4e, 0xe3, 0x73, 0xff, 0xd3, 0x5d, 0x85, 0x8b, 0xce, 0x6b, 0x1b, 0x96, 0x57, 0x33, 0x90, 0xb4, 0x6e, 0x8, 0x1f, 0x93, 0x3e, 0xb9, 0xae}}
	return a, nil
}

var __05_add_more_std_labelsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xc8\xcb\x4f\x49\x8d\xcf\x4c\x49\xd0\xc1\x2a\x91\x97\x98\x9b\x8a\x43\xaa\xa4\xb2\x00\x8b\x54\x6e\x62\x72\x46\x66\x1e\x76\x13\x93\xf3\xf3\x4a\x12\x33\xf3\x52\x8b\xb0\xca\x16\xa7\x16\x95\x65\x26\x83\x75\x72\x59\x73\x01\x02\x00\x00\xff\xff\x1d\x12\xd5\x34\xba\x00\x00\x00")

func _05_add_more_std_labelsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__05_add_more_std_labelsDownSql,
		"05_add_more_std_labels.down.sql",
	)
}

func _05_add_more_std_labelsDownSql() (*asset, error) {
	bytes, err := _05_add_more_std_labelsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "05_add_more_std_labels.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4c, 0xd5, 0x2, 0x6e, 0xa6, 0x74, 0xa2, 0xf4, 0x47, 0x69, 0xb8, 0xd9, 0x70, 0x9d, 0x50, 0x5e, 0x1f, 0x7b, 0x77, 0x3f, 0xe7, 0x3b, 0xfc, 0xb, 0x17, 0x86, 0xc, 0x2a, 0xc2, 0x41, 0xb6, 0xf1}}
	return a, nil
}

var __05_add_more_std_labelsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\x4a\xc6\x30\x10\x05\xe0\x7d\x4f\x31\xbb\x5f\xc1\x1b\xb8\x8a\x69\x5d\x25\x2d\xd8\xba\xb6\x21\x19\x75\xa0\x99\x48\x1a\x94\xde\x5e\xda\x52\xa4\x25\x42\xf6\x8f\xef\xcd\x1b\xa1\x86\xe6\x05\x06\xf1\xa4\x1a\xf0\x98\x22\xd9\xb9\x02\x10\x75\x0d\xb2\x53\xaf\xba\x85\x91\x83\xc3\x37\x72\x23\xa8\xf0\x23\x4d\x74\xc4\x66\xa2\xb4\xdc\xf5\x29\x12\x7f\xdc\x83\xec\xb4\x6e\xda\x01\x6e\x6d\x70\x08\xe4\x90\x13\xbd\x13\xc6\x1b\x88\xe7\xd5\xde\x01\x1f\x1c\x4e\xe3\x43\xce\x66\xe3\xb1\x50\x5f\xa3\x67\x97\x5c\x1e\x4d\xcb\x57\x29\xba\x46\xcf\xe8\x76\xd1\x95\xf5\xc6\x7e\x12\x97\xbd\x42\xef\xd9\x7f\xbf\xb1\x5d\x77\x2d\xb0\x81\x93\x21\xc6\x58\x54\x21\x8f\x74\xae\xe4\x8f\xca\x4e\x99\x31\x7e\x93\x2d\x9b\xd2\xef\xd9\x5c\xcb\xc1\x6c\x6b\xaa\xc7\xea\x37\x00\x00\xff\xff\x27\xe2\xd8\x13\x4c\x02\x00\x00")

func _05_add_more_std_labelsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__05_add_more_std_labelsUpSql,
		"05_add_more_std_labels.up.sql",
	)
}

func _05_add_more_std_labelsUpSql() (*asset, error) {
	bytes, err := _05_add_more_std_labelsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "05_add_more_std_labels.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x66, 0x48, 0x59, 0xed, 0x41, 0xa1, 0x7, 0xb, 0xf5, 0x5e, 0x8c, 0x93, 0x98, 0x3e, 0xf1, 0x9e, 0xd6, 0x9b, 0x50, 0x7f, 0x2d, 0xdf, 0xf9, 0xf2, 0x5e, 0x57, 0x20, 0x80, 0x5, 0xeb, 0x46, 0xe2}}
	return a, nil
}

var __06_change_agent_typeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcf\x4d\x4b\xc4\x30\x10\xc6\xf1\xbb\x9f\xe2\xb9\x45\xa1\x2b\xbe\x1d\x44\xf1\x50\xd7\x0a\x42\xbb\x05\xa9\x07\x4f\xdb\x98\x9d\x66\x0b\x79\x33\x19\x95\x7e\x7b\x19\x5d\xf4\x90\xcb\x8f\x30\xf3\x9f\xba\x1d\x9a\x67\x0c\xf5\x7d\xdb\x60\xf4\xc4\x79\x36\x65\x44\xd7\x3f\x3c\x3d\xbe\x62\xdd\xb7\x2f\xdd\x06\xa3\xb6\x14\x78\xcb\x4b\xa2\x11\x4d\xf8\xf0\xd7\xc7\xea\x9f\xb6\x73\xf8\xd4\x6e\xde\x29\xdc\xe1\xac\x82\xf2\x4b\x79\x77\xab\x44\x79\x2a\x66\x4f\x5e\x8b\x9f\xff\x79\x71\xf1\xcb\x45\x2b\x78\x21\x18\x83\x8d\xbb\xb7\x55\xca\x71\x9a\x1d\x65\xf1\xcb\x0a\x2a\xc5\xc2\x36\xd3\xcf\x24\x5b\x58\xb3\x3c\xf2\x14\xb8\xc8\x97\xab\x13\xac\xfb\xae\x6b\x36\x03\x54\x2d\x25\x18\x96\x44\xe0\xbd\x66\x98\xe8\x1c\x19\x46\x9c\x70\xb8\xe7\x06\x87\xad\x15\x24\x0b\xbf\x5d\x15\x88\xcd\xa9\xba\x3d\xfa\x0e\x00\x00\xff\xff\x03\xb2\xa6\x30\x03\x01\x00\x00")

func _06_change_agent_typeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__06_change_agent_typeDownSql,
		"06_change_agent_type.down.sql",
	)
}

func _06_change_agent_typeDownSql() (*asset, error) {
	bytes, err := _06_change_agent_typeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "06_change_agent_type.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x6e, 0xec, 0x68, 0x15, 0xaa, 0xa4, 0x3b, 0x44, 0x6e, 0x3c, 0xc0, 0x27, 0xe6, 0x59, 0x82, 0x85, 0x7b, 0x3, 0xb, 0xa, 0x28, 0xb7, 0x73, 0x75, 0x86, 0x2b, 0x2d, 0xa, 0x59, 0x7, 0xb}}
	return a, nil
}

var __06_change_agent_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcf\xbd\x4a\xc4\x40\x14\x05\xe0\xde\xa7\x38\x5d\x14\x32\xe2\x5f\x21\xca\x16\x71\x8d\x20\x24\x1b\x90\x58\x58\x99\x71\xf6\x66\x36\x30\x7f\x3b\x73\x55\xf2\xf6\x92\x6c\x52\xd8\x9e\xef\xc0\x3d\xb7\xa8\xda\xf2\x0d\x6d\xf1\x54\x95\xe8\x2c\x71\x1c\x54\xea\x50\x37\xcf\xaf\x2f\x1f\xd8\x36\xd5\x7b\xbd\x43\x27\x35\x39\xfe\xe4\x31\x50\x87\xd2\x7d\xdb\xfb\xf3\xec\x28\x9d\x98\x63\x31\xc5\x62\x70\x3f\xd2\x0c\xfb\x6c\x73\x95\x63\x36\x3b\xa6\xa3\x11\x81\x62\x9f\xd4\x81\xac\x3c\x95\xb3\xcd\xf5\x3f\x4f\xc6\xff\x1a\xaf\x57\xbc\x59\xd1\x3b\xed\xf7\x5f\x22\x44\xdf\x0f\x86\xe2\xea\xb7\x8b\x07\x9f\x58\x47\x9a\x2f\xe8\xc4\x92\xc9\x92\xe3\xb4\xd6\xee\x2e\xb0\x6d\xea\xba\xdc\xb5\xc8\x8a\x29\x42\x3b\x06\x02\x1f\x24\x43\x79\x63\x48\x31\x7c\x8f\xe5\xdb\x07\x2c\x2b\x72\x4c\x73\x71\xda\x9b\x83\x58\x5d\x66\x8f\x67\x7f\x01\x00\x00\xff\xff\x95\x60\x5f\x66\x21\x01\x00\x00")

func _06_change_agent_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__06_change_agent_typeUpSql,
		"06_change_agent_type.up.sql",
	)
}

func _06_change_agent_typeUpSql() (*asset, error) {
	bytes, err := _06_change_agent_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "06_change_agent_type.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0xe5, 0xb8, 0xa7, 0xd6, 0x2d, 0x10, 0x46, 0x87, 0x61, 0xa, 0x82, 0x26, 0x7f, 0x4d, 0xdf, 0x3c, 0xaa, 0xf8, 0xf6, 0xe2, 0xe4, 0x9b, 0x3d, 0x9e, 0x95, 0xc, 0xa8, 0xfa, 0x77, 0xdd, 0x97}}
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
	"01_init.down.sql":                _01_initDownSql,
	"01_init.up.sql":                  _01_initUpSql,
	"02_postgresql_columns.down.sql":  _02_postgresql_columnsDownSql,
	"02_postgresql_columns.up.sql":    _02_postgresql_columnsUpSql,
	"03_add_agent_type.down.sql":      _03_add_agent_typeDownSql,
	"03_add_agent_type.up.sql":        _03_add_agent_typeUpSql,
	"04_add_tables_column.down.sql":   _04_add_tables_columnDownSql,
	"04_add_tables_column.up.sql":     _04_add_tables_columnUpSql,
	"05_add_more_std_labels.down.sql": _05_add_more_std_labelsDownSql,
	"05_add_more_std_labels.up.sql":   _05_add_more_std_labelsUpSql,
	"06_change_agent_type.down.sql":   _06_change_agent_typeDownSql,
	"06_change_agent_type.up.sql":     _06_change_agent_typeUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"01_init.down.sql":                {_01_initDownSql, map[string]*bintree{}},
	"01_init.up.sql":                  {_01_initUpSql, map[string]*bintree{}},
	"02_postgresql_columns.down.sql":  {_02_postgresql_columnsDownSql, map[string]*bintree{}},
	"02_postgresql_columns.up.sql":    {_02_postgresql_columnsUpSql, map[string]*bintree{}},
	"03_add_agent_type.down.sql":      {_03_add_agent_typeDownSql, map[string]*bintree{}},
	"03_add_agent_type.up.sql":        {_03_add_agent_typeUpSql, map[string]*bintree{}},
	"04_add_tables_column.down.sql":   {_04_add_tables_columnDownSql, map[string]*bintree{}},
	"04_add_tables_column.up.sql":     {_04_add_tables_columnUpSql, map[string]*bintree{}},
	"05_add_more_std_labels.down.sql": {_05_add_more_std_labelsDownSql, map[string]*bintree{}},
	"05_add_more_std_labels.up.sql":   {_05_add_more_std_labelsUpSql, map[string]*bintree{}},
	"06_change_agent_type.down.sql":   {_06_change_agent_typeDownSql, map[string]*bintree{}},
	"06_change_agent_type.up.sql":     {_06_change_agent_typeUpSql, map[string]*bintree{}},
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
