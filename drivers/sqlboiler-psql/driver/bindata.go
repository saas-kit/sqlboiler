// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates_test/singleton/psql_main_test.tpl
// override/templates_test/singleton/psql_suites_test.tpl
// override/templates_test/singleton/psql_upsert.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5d\x6f\xdb\xb8\x12\x7d\x96\x7e\xc5\x34\xb8\x40\xa4\x0b\x47\xb9\xcf\xbd\xf0\x43\x93\x7e\xdc\xe2\x6e\x53\x6f\xd3\x6e\x81\x2d\x8a\x80\x96\x46\x36\x11\x9a\x54\x29\x2a\xa9\x57\xd5\x7f\x5f\xcc\x50\xb2\x64\x5b\x6d\x9c\x66\x3f\xba\x4f\x89\x38\xe4\xcc\xe1\x9c\x43\xce\xd0\x75\x7d\x02\xff\x72\x62\xae\xf0\x42\xac\xf0\x52\xea\x45\xa5\x84\x85\xc7\x53\x48\xde\xd2\x68\x42\xc3\xf0\x05\xca\xce\xf2\x05\x9c\x74\x0a\xcf\x45\x89\x70\xd2\x34\x21\x3b\xb8\x11\xf6\xe0\xe5\xa9\x58\xa1\xda\x5e\x5e\xa6\x4b\x5c\x09\x5e\xb0\xbf\x34\xb9\xec\xad\xbc\x40\xe6\x90\x3c\xc9\xb2\x17\xca\xcc\x85\x62\x27\xa7\xa7\xf0\xae\x28\xd1\xba\x17\x20\x9c\xc3\x55\xe1\x4a\x10\x1a\xa4\xa6\xb1\x09\x08\x9d\x41\x66\x90\xc7\xaa\x22\x13\x0e\xc1\x58\x90\x0b\x6d\x2c\x82\xd1\x90\x1a\x9d\x2b\x99\xba\x24\xcc\x2b\x9d\x42\x64\xe0\xdf\x75\xbd\x9f\x94\xa6\x89\xbb\x30\x11\xa3\xd0\xc6\x41\x72\x61\xce\x8d\x76\xf8\xd9\x35\x4d\xea\x3e\x93\x2f\xfa\x48\xda\xc1\x09\xd4\x35\xea\x8c\x50\xb6\xa1\x5f\xeb\xf3\x36\x1c\xcc\x8d\x51\x93\x4d\xf4\x73\xa3\xaa\x95\x2e\xe1\xc3\xc7\xd2\x59\xa9\x17\x93\x76\xc1\xfe\xf8\xed\x52\x3a\x54\xb2\x74\x90\x24\x89\x1f\x8c\x01\xad\x35\x16\xea\x30\xb0\xe8\x2a\xab\xc1\x24\x1e\xab\x87\x3a\x84\x39\x37\x52\x25\x2f\xd0\x3d\x3d\x8b\xe2\xba\x46\x55\x22\x43\x9f\x40\x67\x68\x67\xb6\x76\x9d\x35\xcd\x64\x0f\xfc\x1e\xee\x1d\xb8\x03\x94\x49\x92\xc4\x61\x13\x86\x9b\x4c\x84\x9e\x44\xa2\x65\x40\x24\xfd\x3b\x13\x5a\xa6\x3b\x94\xce\x1e\xc6\x29\xb0\xcf\x92\xc6\x38\x45\x87\x93\x3c\xfb\x01\x59\xae\xc3\x40\xe6\xb4\x11\x3a\x27\x3f\x16\xc5\xff\x65\x58\x8f\xa6\xa0\xa5\x22\x9c\x41\x41\x89\x8f\x38\xe2\x7b\x2b\x8a\x67\xd6\x46\x68\x6d\x1c\x87\x41\x33\x26\x87\xaf\xf0\x3f\x46\x3f\x54\x74\x9f\xd0\x37\x7e\xc6\xb4\x72\xc6\xde\xe7\x90\x0f\x5c\x17\xdf\xa9\x8d\xd9\x7e\xca\x09\x89\x4f\xef\xb3\x16\xd3\x20\xf1\xfb\x82\xe9\xa7\xb7\x43\x83\x55\xe3\x74\xfc\x65\x42\x1a\x11\xfc\x50\xe0\x84\xfc\xef\x14\xcb\x86\xbe\x3f\x5c\x18\x07\x92\xff\x8f\xe7\x7e\x53\x2a\x64\x0e\x06\xa6\x3d\x09\x6d\xe9\x60\x7b\x99\x5c\xe0\x6d\x74\x54\xd7\xc9\xec\x7a\x41\x89\x68\x9a\xc7\xa0\x0d\xd4\xf5\xa0\x3c\x37\x0d\x14\xd6\xdc\xc8\x0c\x33\xc8\x8d\x85\x8a\x13\x74\xc4\xac\x85\x01\xd5\x76\x62\x48\x51\xce\x8f\x9c\x5c\x61\xe9\xc4\xaa\xb8\xf2\xb3\xae\x96\xa8\x0a\xb4\x47\x90\x40\xe3\x67\xf7\xda\xfb\x9f\x31\xd7\x25\xd3\xbd\xa5\xd2\xcc\x9c\x61\x6e\x2c\x7a\x1e\x78\xd2\xc1\x92\xdd\x97\x5c\xbf\x5b\x82\xcb\x68\x39\xfd\x61\x18\xe8\xdf\x9e\x62\x2e\x2a\xe5\x4a\x0a\xfc\xa9\x42\x2b\xb1\x4c\x2e\x8c\xfe\x15\xad\x69\x4d\x97\x48\x4a\xd8\x6d\x7d\x9a\xa6\x25\xe1\xbd\x74\xcb\x76\xe6\x04\x4c\x1c\x86\xc1\xe9\x29\x9c\x55\x52\x65\x90\x8a\x74\x89\x70\x8d\x6b\x90\xfa\x44\x49\x8d\x50\x2d\x94\x54\x6b\x38\x81\xd5\xba\xfc\xa4\xe0\xa6\x84\x82\xfe\x16\xd6\xcc\x15\xae\xca\x30\x98\x57\x39\x21\x29\x9d\x5d\x09\xbd\x50\x48\x77\xf8\x59\x95\xe7\x68\xa3\x98\x73\xb4\xa7\x18\xda\xe1\xbc\xca\x93\xf7\x56\x3a\x3c\x5b\x3b\x8c\x8e\xdd\x31\x11\x03\xa4\xcc\x31\x73\xce\xe6\x70\x77\x38\xa1\x61\x22\xf7\x6a\x02\x29\x81\xb0\x42\x2f\x70\x4f\x8b\x5b\x0e\x2f\x59\x69\x51\x7a\x1f\x87\xdb\x12\x7e\xb0\xbb\x5e\xf9\x0f\x76\x35\x50\xc3\x37\x7c\x11\x9f\x8f\xa7\x40\xd6\xd6\x10\x87\x41\x4f\xd8\xac\xea\x08\x9b\x57\x79\xcc\x7a\xdf\xd7\x8e\x17\xf6\x39\xe9\xe3\x55\xe5\x92\x37\x3f\x99\xf4\x9a\xdc\xb0\x62\x26\x5e\x38\x19\x45\xb9\x63\xf1\x87\x6b\x5c\x7f\x3c\x2c\xc4\x3b\xad\x7c\x90\x30\xb8\x11\x96\xcf\x08\x9f\xff\x90\x65\xf5\xa8\x0d\x49\xfb\xee\x5a\x2f\x8b\x6e\x5b\x8b\x2f\xd9\xe0\x99\xa3\x53\x11\x06\xc1\x68\xe8\xae\x14\xdc\x61\x1f\x9e\x9c\x03\xa6\x9a\xca\x0d\x67\xf7\x64\xf1\xe7\x46\x06\xf4\x15\x87\x41\xd0\x96\x80\xad\x0d\xbc\x1b\x48\xef\x21\x1b\x98\x59\xb9\x12\x76\xfd\x7f\x5c\x0f\x67\x6e\xd7\x41\x86\x41\xd9\xcc\x41\xa1\x8e\xbc\x31\xa6\x5b\xf8\x3f\x9c\xe5\xbb\x2f\xe1\x4a\xf3\x43\xc9\x99\xf6\xba\xdd\xbd\x92\xa9\x4e\x54\x2a\xe3\x4b\x71\xce\x17\x4e\xbb\xe7\x94\x21\x00\x65\x83\xae\x68\xbe\xa3\x83\xee\x1c\x53\x46\x76\xce\x74\x8f\xb2\x33\x0c\x71\x6e\x16\x4e\x61\x25\xae\x31\xea\x8b\x0e\xad\x38\x28\x3d\x54\xde\xc9\x51\xb1\xde\x44\x98\x8c\x4a\x7b\x7f\x25\xc3\x0f\xfc\xc1\x48\xe8\x82\x5e\xc3\xd4\xef\xd6\x0b\xfc\x67\x1a\x9a\x99\xd2\x2d\x2c\x96\x51\x26\x85\x42\x72\x7e\x54\xd7\xc3\xd7\x66\xd3\x1c\x8d\xb5\x32\x16\x5d\x37\xdc\x17\xdd\x49\xdb\x63\x30\x7b\x3e\xee\x8d\x50\x15\xbe\x12\x45\xc1\xdb\xa6\xa3\xd3\x17\x8b\x33\xa9\xb3\xd6\x34\x9a\x8c\xb7\xeb\x02\xc7\x37\xbb\x71\xd8\xc5\x0b\xba\x22\x38\x28\x5e\x5b\xd5\x8b\x53\xd1\x52\x65\xd1\xc5\x34\xb1\x63\x89\x81\x5a\x74\x7f\x1e\x4c\x8a\x48\xa1\x46\x40\x6e\xa3\x64\x98\x8d\xef\x0d\x38\x75\x7c\xc5\x62\x4e\xd4\x24\x2f\x75\x26\x2d\xa6\x2e\xea\x06\x7e\xa1\x19\xaf\xf3\xc8\x90\x4a\x6e\x84\xda\x2a\xc5\x6c\x2c\x9f\x5b\xb3\xea\xc0\xb3\xc3\xf6\x8a\xdc\x22\x26\xf6\x17\x9b\x47\x42\xbd\x91\xd4\x0e\x6d\x2e\x52\xac\x7d\x7b\xc1\x02\xdf\x49\xd3\x20\x85\xdd\xc2\x3e\xf8\xcc\xd9\xaf\x87\x1e\xf8\xf0\x3b\x95\xb9\xef\xf5\x9e\xe2\xbc\x5a\xbc\x32\x99\xaf\xbd\xf9\xca\x25\xcf\x0b\x2b\xb5\x53\x3a\xea\xed\x5c\x57\x6c\xe7\x8b\x75\x1d\xdf\x3d\x9b\xb2\xd3\x47\xbb\x63\x3f\x3b\xdd\xab\xef\xb2\x02\xaf\x0a\x6a\x94\x12\x3e\x3a\x6f\xcc\x6d\x34\x00\xe1\x63\x50\xe7\x9e\x5c\xa6\x82\x55\x46\x49\xe1\xa7\x3d\xb9\xe4\x9e\xe2\xab\x9e\xda\x50\x11\xb7\x65\xf7\xf1\xda\xf6\xfc\x1b\x6d\x4d\xa7\x50\x7e\x52\xc9\x33\x6b\x2f\xcc\x1b\x73\xeb\xeb\x71\x1b\x91\x44\x77\x7a\x0a\xdd\x99\xe7\x9e\x5f\x1f\xbb\x96\x78\x10\x7a\xed\x96\xf4\x38\xb8\x5d\xa2\x06\xb7\x44\x8b\xc7\x25\x35\xb4\xfe\x9c\xb7\xca\xec\x9b\xa3\xf1\x34\x5d\x75\xe7\x87\xf7\x47\x7d\xfb\x78\x96\x76\x93\xb2\xbf\xee\xee\x9c\x6c\xa7\xa0\x6f\x85\x47\x5b\x58\xaa\x15\xf4\x70\xa2\x57\x13\x5f\x73\xf7\xa9\x18\x47\xbd\x78\x86\x05\xff\x80\xf6\xa1\x6b\x50\xee\x9a\xcb\x0d\x09\x4c\xfd\x46\x0f\x73\xbd\x69\x4c\x82\x6f\xbc\x0e\x36\x3f\x76\x65\xe6\x49\xee\xd0\x7e\xd7\xcb\xa0\xed\xfd\x37\x6c\xb5\x4e\xb5\x54\xc3\x57\x41\x13\xfe\x1e\x00\x00\xff\xff\x84\xdd\x1d\x25\x2c\x15\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 5420, mode: os.FileMode(420), modTime: time.Unix(1527689483, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\xa9\x81\x1c\xa4\xad\xc2\x1c\xfa\xf2\x25\x07\x63\x91\x38\x4e\xba\xb8\xac\xe3\xb3\xdd\x1e\x8a\x6e\xdb\xa3\xad\x91\x42\x44\x22\x19\x92\x8a\xd7\x3d\xe4\xbf\x17\x43\x4a\xb6\x9c\xb5\x93\x6d\x8b\x05\xfa\xcd\x26\x9f\x19\xce\x33\xef\x7a\xe2\x06\x4c\xf9\x79\x7a\x73\xfd\x80\x1b\x18\x82\xc1\x12\x3f\x6b\xf6\xb1\xb1\x6e\xa4\x6a\x2d\x2a\x4c\x7e\x49\xde\xd7\xe9\x3f\x2e\x6e\x17\xe3\x19\x2c\x2e\x2e\x6f\xc7\x70\x37\xb9\xfd\x2b\xb0\x77\x9f\xe4\x27\xfb\xdb\x8b\xab\x2b\x18\xdd\x4d\xe6\x8b\xd9\xc5\x87\xc9\x02\xd8\xbb\xf7\x70\x7d\x37\x1b\x7f\xb8\x99\xc0\x8f\x63\x42\xbd\xff\xe1\x93\xfc\x25\x8d\x63\xb7\xd1\x08\xba\x5c\xa0\x75\x68\xc0\x3a\xd3\xac\x1c\xfc\x1a\x47\xf9\x72\xa4\xa4\x84\x77\xf6\xb1\x62\x57\x97\x31\x1d\x4c\x78\x8d\x40\x10\x21\xcb\x38\xba\x57\xd6\x01\xec\xfe\x37\x16\x4d\xff\xbf\xe6\xd6\xf6\xff\x5b\x5b\xd5\x2a\xc7\xdd\xbd\x32\x5e\x5e\x48\x17\xc7\x91\x2e\xa7\xdc\xda\x6b\x51\x6d\x01\x71\xe4\xd0\xba\xab\x4b\xff\x6a\x7b\xf6\x1c\xc7\x45\x23\x57\x20\xa4\x70\x49\x1a\xcc\xfc\xc8\x85\x84\x21\x7c\xd7\x71\xf8\xf5\x99\x60\x67\x67\x60\xd1\x35\x1a\xf2\xa6\xd6\x16\xdc\x3d\x42\xce\x1d\x5f\x72\x8b\x60\x57\xf7\x58\x73\xe0\x32\x07\x51\x93\x19\x16\x84\x23\x3b\x14\x70\x70\x48\x47\xdc\x6c\xc0\x70\x99\xab\xba\xda\x90\xae\x12\x25\x1a\xee\x30\x07\x32\xaa\xa7\x4a\x81\xbb\xe7\xce\x9f\x5a\x58\x71\x09\x4b\x04\xd3\x48\xe0\x25\x17\xd2\x3a\x52\xdc\x58\x21\x4b\xb2\x60\x5f\x91\x7d\xac\x96\x4a\x54\x68\xe0\x6e\xf6\x11\x34\x5f\x3d\xf0\x12\x59\xe0\x97\x68\x78\xd7\xf1\x49\x03\x91\x24\x05\x34\x46\x19\x22\x4d\xd9\x81\xc6\x84\x83\x38\x8e\x9e\x84\x46\xc3\xe6\xe8\xae\xb0\xe0\x4d\xe5\x92\x81\xa6\xb0\x05\x9e\x83\x0c\x06\xba\x59\x56\x62\x35\x48\x8f\x42\xc9\x0b\x83\x0c\xfe\xf8\x87\xdf\xff\xee\x38\xa8\x8d\x20\x29\x34\xf8\xd8\x08\x83\x83\x94\x42\xc7\xda\xd4\x18\x42\x10\xbc\x41\x37\xf7\xf1\x6a\xe5\xf2\xa5\xe4\x35\x61\x23\xcd\x7c\xd6\x1c\x03\xd2\x65\x80\xf9\x64\x3a\x06\xa3\xcb\x00\xf3\x39\x76\x0c\x46\x97\x2d\x8c\x52\xad\x07\xfb\x20\xf7\x78\x7b\x4c\x97\x9e\xc7\xb4\x75\xe4\x89\x31\xf9\x7e\x08\x4f\xbc\xe2\xec\x12\x4b\x21\xff\xc2\x2b\x91\x73\x27\x94\x4c\x52\xd6\xfe\xc1\x24\x8e\x22\x0f\x09\x6a\x26\xca\x8d\x6b\xed\x36\x49\x20\x47\x41\xd9\x71\xc9\x8e\x62\xc9\x25\x1d\x36\xb8\x67\x8b\x9d\x28\x97\xf8\x1f\xe3\xc7\x86\x57\x36\x09\x3c\x33\xf8\xbe\xc3\x07\x72\xaf\x28\x0f\x71\xeb\xe0\x5d\x98\x8e\xe3\x5b\x1f\x74\x02\x5b\x97\x64\x71\x94\xb2\xd1\x3d\xae\x1e\x12\x72\x8f\x28\x7c\x76\xfe\x66\x08\x52\x54\x94\xaf\x91\x41\xd7\x18\x49\xa7\x71\xf4\x1c\xc7\xd1\xd9\x19\x8c\x0c\x72\x87\xc0\xdb\x32\x13\xff\xc2\x1c\xf2\x25\x90\x09\x8c\xe2\xd1\x2b\xfe\xe1\x0e\xc3\xe6\x8e\x2f\x2b\x0c\x17\x5b\x06\xbd\x47\x87\xa0\x59\xcd\x1f\x70\x7a\xd3\xf5\x93\x24\xfd\xe1\x2d\x73\x7a\xb2\xb9\x51\x7a\xe1\x9f\x7e\x53\xae\x2f\xb6\xf2\x6c\xbe\x52\x30\x8e\xa8\x29\x8d\xea\x1c\xce\x87\x80\x9f\x71\xc5\x46\xaa\xae\xb9\xcc\x93\x81\x2e\xff\x49\x77\x54\x62\xa7\xa7\xa1\x7e\x4f\x95\xac\x36\x83\x0c\x76\x64\x3b\x71\x36\x96\x4f\x30\x04\xae\x35\xca\x3c\x51\x96\xfe\x0b\x43\x49\x48\x68\x5d\x8e\xe5\x53\x92\x32\xc6\xd2\x38\x0a\xf6\x1d\x7e\xd2\x3e\x56\x5e\xfd\xce\xe3\x7d\x81\xaf\x7f\x24\x8e\x4c\x06\x6b\x7a\x40\x28\x36\x15\x1a\x93\x9e\xa9\x73\x97\xab\x86\x8a\x70\xdd\xd7\x3d\x77\xb9\x6f\xde\x12\xd7\xd7\x3f\xe2\xe6\x0a\xad\x33\x6a\x83\x26\xd9\xce\xbe\x0c\xcc\x5e\x74\x77\xfa\xb8\x71\xaf\x7a\x5a\x19\xcb\x7e\x36\x5c\x27\x68\xa8\xda\x0a\x2e\x2a\x6a\xdf\x0a\x2c\x89\x42\xeb\x69\x58\x05\x3f\x50\x13\xe8\x87\xb4\x6f\xe3\xff\xfa\x92\x7d\xac\xf6\x9f\x39\xc0\xe7\x67\x2e\x0e\x3d\x52\xd4\x8e\x4d\x8d\x90\xae\x92\xa4\x3d\xfd\xba\x77\xd7\x5c\x38\x28\x94\x39\x4c\x32\x8e\xd6\x6c\x54\x29\x8b\x49\x0a\x67\x67\x70\x51\xd0\xe0\xef\x32\x52\x58\xc8\x95\xc4\x0c\x56\x84\xf0\x73\x73\x6d\x84\x43\x40\x99\x83\x2a\xfc\x81\x16\x1a\xe3\x83\xbe\xfa\x46\x2c\x0e\x38\xb0\x95\x97\xa2\xda\x2e\x05\xfb\x43\xd3\x34\x72\x54\xe7\x89\xa5\x0c\xcb\x3a\xe9\x76\x8f\xc8\x80\x9b\xd2\x02\x63\x2c\xfc\xef\x8d\xd6\xd5\x81\x12\x69\x85\x83\x54\x5b\x4f\xff\x59\x61\x88\x02\x2a\x94\xc1\x98\x94\x3c\xf3\xbd\xf7\xcb\xaa\x57\x02\xc1\x12\xcb\x26\xb8\x9e\x21\xcf\xd1\xb4\xe8\x40\xd7\x86\xf2\x39\x1f\xc2\x77\xcb\x8d\x43\xcb\x2e\x9b\xa2\xf0\xbb\x0e\x5d\x91\xbb\x0f\x5d\xad\xfa\x85\x17\x54\x6c\x0f\x43\xe8\x82\xf0\x36\x96\xe7\x43\xa0\xeb\x59\x23\xdf\x88\x62\x17\x26\xd3\x48\x29\x64\x79\x3e\xd8\xba\x38\x78\x29\x7d\x81\x0f\x8f\xb7\x13\x25\x49\x0f\x5c\xa3\x31\x7b\xd7\x2f\x5b\xe6\x9b\x01\x6f\x3d\x0e\x7f\xfb\x7b\x70\x25\xd9\xdc\x0a\x75\x47\x1d\x8b\xb9\xa6\x77\x8b\x64\x30\xbd\xf9\xd3\xdd\x7c\x31\x3c\xb1\xbe\x01\xd2\x7c\xf5\xd3\xef\x05\x66\x7a\x37\x5b\x0c\x4f\x72\x8f\xa1\x99\x7a\x08\xf3\xe7\xf9\x78\xd6\xe9\xa1\x99\x7e\x50\xcf\xc5\x7c\x7e\xfd\xe1\x76\xdc\xe1\x76\x3b\x2f\xa1\x9f\x8f\xf0\x7a\x39\xcd\x76\xb9\xea\x6a\x9d\x75\x61\x13\xaa\x71\xa2\x62\x0b\xac\xb5\x87\x0d\xfc\xda\x57\x76\x3b\xd0\x6b\x23\xf9\x68\x01\x86\xba\x06\xa5\x69\xb3\x81\x42\x54\xd8\x55\x1f\x11\xbb\x6e\x89\x79\x2b\x06\x27\xf6\xfc\x24\x3f\xd7\xca\xba\xd2\xa0\x3d\xef\x79\xb4\xf3\xda\xd6\x33\xdb\x72\x08\xfb\x5b\xaf\x1e\xbe\x54\xdb\x29\xf2\x40\xdf\xa1\x77\x98\x4a\x12\x28\x7d\xc5\x9c\x93\xa3\x86\x74\x9b\xcf\xff\x91\x49\xbb\xf1\xfb\x0d\xcd\xea\x27\x1d\x0c\xc1\xd5\x9a\xf9\x4d\x2a\xdd\xd6\x0a\x1d\xb5\xd3\xe1\x48\x42\xee\xef\x3a\xbb\x74\x6c\x15\x68\xd6\xb6\x5e\x9f\x82\x01\x9c\x2f\xbf\xd8\x30\x0e\xeb\xee\xaf\x5f\x6f\x68\x26\xa8\xd7\x3b\x38\x3d\x15\xc5\x29\x7e\x16\xd6\xd9\x43\xcf\x9c\x9d\x81\x43\x6e\x72\xb5\x96\xbe\xaf\x37\x0e\x2d\xac\x2a\xe4\xb2\xd1\xe0\xb8\x7d\xb0\xb0\xbe\x47\xe9\x47\x5b\xf8\x8e\x2b\x84\x14\xf6\xbe\x6b\x6e\x87\xec\xec\x14\x1e\xff\x2a\xdb\x5b\x2a\xfd\xb7\x74\xe7\xd6\xb7\xd6\xca\x0e\x0f\x1e\xf1\x5f\xaf\xa7\x5b\xb7\x29\xcb\x66\x58\xab\x27\xda\x97\x7b\x2d\xe7\x58\x74\x95\x24\x56\x49\xfb\xe1\x9f\x05\x3a\xfe\x5b\x5b\x14\x5b\x2e\x07\x9e\xed\xae\x32\x6f\xb5\x37\xe0\x85\x47\x76\x88\x76\xf8\x3c\x56\xec\x4e\xa3\x4c\x06\x5d\xdf\x18\x64\x90\x1b\xf1\x84\x86\x4d\xe7\x3f\xdd\x5e\x36\xa2\xca\x7f\x6a\xd0\x6c\xda\xc1\xd0\x7d\x3a\x85\x2c\xff\xb2\x68\x5e\x96\x54\xfb\x81\x92\xbe\xd6\x00\xa5\xa8\xb2\x2f\x5c\xb6\xcf\xe5\x39\xfe\x77\x00\x00\x00\xff\xff\x00\xef\xd0\xbf\x8f\x11\x00\x00")

func templates_testSingletonPsql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testTpl,
		"templates_test/singleton/psql_main_test.tpl",
	)
}

func templates_testSingletonPsql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.tpl", size: 4495, mode: os.FileMode(420), modTime: time.Unix(1527187113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x8e\x83\x30\x10\x44\x7b\xbe\x62\x84\x28\xe0\x04\xfe\x80\x93\xae\xba\xea\xae\x48\x11\x91\x0f\x70\xc2\x82\x2c\x39\x1b\x84\x17\x29\x92\xf1\xbf\x47\x18\x8b\x90\xce\xe3\x99\xb7\x3b\xdb\xcf\x7c\x43\x4b\x4e\x2e\xa3\xa3\x49\x4a\xc1\x97\x90\x13\xc3\x83\x6a\x2b\xf8\x0c\xf0\xbe\xc1\xa4\x79\x20\x14\x86\x3b\x7a\xd6\x28\x44\x5f\x2d\xe1\xfb\x07\xaa\x5d\x5f\x2e\x84\x94\x33\x7d\x32\xd5\x9f\xfb\x7f\x18\x8e\x36\x9a\xdd\x27\xeb\x8e\x72\xcb\x9e\xf4\x3d\x0e\x4b\x64\x94\x0b\x46\x3b\x4f\xda\x62\x81\x18\xb1\xf4\xab\x77\x50\xd4\x79\xe6\x32\xf7\xfe\x4d\x87\x90\xd7\x58\x6b\x7f\x7e\x6e\x27\x55\x71\x19\x71\x77\xec\x91\x54\xc8\x5e\x01\x00\x00\xff\xff\x2f\xea\xf2\xb5\x00\x01\x00\x00")

func templates_testSingletonPsql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testTpl,
		"templates_test/singleton/psql_suites_test.tpl",
	)
}

func templates_testSingletonPsql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.tpl", size: 256, mode: os.FileMode(420), modTime: time.Unix(1527187008, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templates_testSingletonPsql_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_upsertTpl,
		"templates_test/singleton/psql_upsert.tpl",
	)
}

func templates_testSingletonPsql_upsertTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_upsert.tpl", size: 1317, mode: os.FileMode(420), modTime: time.Unix(1527189206, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x41\x6f\xda\x4c\x10\x3d\x7b\x7f\xc5\x7c\xe8\x6b\x65\x57\x64\xa3\x5e\x53\x71\x48\x48\x0e\x55\x55\x84\x82\xf9\x01\x1b\x7b\x4c\x56\x2c\xbb\xd6\xee\x38\x81\x6e\xf6\xbf\x57\x6b\x43\x42\xc0\x48\x1c\x9a\x4a\x3d\x80\xe4\xf1\x9b\x79\x6f\xde\xcc\xd8\xfb\x0b\xf8\x9f\xc4\x83\xc2\x89\x58\xe1\x4c\xea\x45\xa3\x84\x85\xab\x11\xf0\x3c\x46\x79\x0c\xc3\x0b\xb8\xdd\x9b\x17\x20\x49\x0a\xc7\xc2\x21\x5c\x84\xc0\xde\x17\x98\xaa\xc6\x0a\x75\x9c\x5e\x77\xf1\xde\xe4\x27\x61\xcf\x4a\x2d\xc4\x0a\x55\x6f\xea\x59\xb2\xdf\xa7\x57\x8d\x2e\x80\xd0\x91\xf7\x87\xea\x43\x98\xd7\x0e\x2d\xa5\x04\x5f\x22\x42\xea\x05\xcf\x33\xf0\x2c\x21\x3e\x15\x56\x28\x85\x2a\xcd\x18\x4b\x64\x05\x0a\x75\xea\xfd\xa1\x8e\x10\xc6\x46\x35\x2b\xed\x32\x18\x8d\x4e\x62\xa6\x56\xae\x84\xdd\xfc\xc0\xcd\x2b\xda\xb3\x24\x21\x3e\x5b\xca\x3a\x1d\xc4\xff\x5a\xea\x05\xb4\xf2\xe0\x59\xd2\x23\x18\xad\x36\x50\x77\x79\xb0\xc4\x0d\x14\x5d\xe6\x20\x63\x49\x60\x2c\x71\x88\x65\x34\xc1\x0a\x5d\x9a\x95\xfc\x85\x7c\x82\xcf\x33\xc4\x32\xcd\x58\xf2\x24\x2c\xa0\x6d\x7f\xc6\xb2\xe4\xf2\x12\xae\x89\x70\x55\x13\xd0\x23\xc2\xf7\xc9\xec\xee\x3e\x07\x27\x4b\x04\x53\x81\xd0\x30\x9f\xc6\x08\x4b\xfa\xb4\x47\x92\x7d\xe7\xde\xde\xf8\xd0\x1a\x13\x89\xf6\x75\xcc\xc8\x36\x05\xa5\x51\xe0\x10\x3e\xf7\x95\x1c\x42\x5f\xf4\xf6\x26\xdf\xd4\xe8\x86\x40\xb6\xc1\xec\x5b\x5b\xf7\xbf\x11\x68\xa9\xb6\x66\xdd\xc5\x6e\xaa\x74\x30\xd7\xad\x4d\x64\xde\x48\x4f\x28\x04\xd7\x6a\xb9\x82\x4f\x6e\x30\x8c\x05\xb7\xe6\x79\x2f\x2b\xd0\x86\x80\x4f\xcc\xd8\x68\xc2\x35\x85\x50\xd0\x3a\xf6\x5a\x74\xcf\xfc\x46\x14\xcb\x85\x35\x8d\x2e\xd3\xcc\x7b\xd4\x65\x08\x2c\xe9\x20\x3f\x1b\x47\xf9\x3a\x6d\xab\xec\x57\x38\x0a\x3c\x18\xa9\xf8\x0d\x2e\xa4\x6e\x6b\x28\x87\xfb\xb1\x7c\x9d\x16\xb4\x1e\xc6\x0e\x77\x0c\x67\x81\x32\x96\x94\x58\xa1\x05\x5a\xf3\x7b\xa3\xd4\x83\x28\x96\x71\xea\xaf\xb3\xe8\x73\x97\x6f\x37\xfd\x54\xeb\x71\x26\xa8\xcb\x78\x31\x10\x9f\x2a\xa1\x1c\xb6\xb4\x1d\xf7\x59\xf3\x68\x5a\x8e\x13\xc3\x38\x9a\x42\x61\x1a\x4d\x6d\xe0\x70\xc7\x76\xd7\x99\x66\x7c\x1c\x41\x67\xaa\x7e\xf3\xe0\x58\x67\xba\xe3\x8d\x90\x96\x39\x82\xbe\xbe\x83\x0c\x9e\x85\x26\x30\x1a\xc1\x62\x61\x6c\x39\x84\x85\xa1\xab\xc1\xb0\xc3\x6f\x55\x1f\x1c\xd3\x7c\x7a\x7b\x9d\xdf\xf5\x1d\xd3\x47\x9c\xc6\x76\x2a\x67\x7d\x63\x38\xe7\x1f\x7b\x45\x7f\x74\xdf\xe2\xd1\xff\xb5\x75\xfb\x57\xb6\x2d\xb0\xdf\x01\x00\x00\xff\xff\x00\x2e\x64\x89\xb5\x07\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 1973, mode: os.FileMode(420), modTime: time.Unix(1527652631, 0)}
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
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates_test/singleton/psql_main_test.tpl": templates_testSingletonPsql_main_testTpl,
	"templates_test/singleton/psql_suites_test.tpl": templates_testSingletonPsql_suites_testTpl,
	"templates_test/singleton/psql_upsert.tpl": templates_testSingletonPsql_upsertTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
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
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.tpl": &bintree{templates_testSingletonPsql_main_testTpl, map[string]*bintree{}},
			"psql_suites_test.tpl": &bintree{templates_testSingletonPsql_suites_testTpl, map[string]*bintree{}},
			"psql_upsert.tpl": &bintree{templates_testSingletonPsql_upsertTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
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

