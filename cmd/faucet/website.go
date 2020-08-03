// Code generated by go-bindata.
// sources:
// faucet.html
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x6d\x93\xdb\x36\x92\xfe\x3c\xfe\x15\x1d\x9e\xbd\x92\xce\x43\x52\x33\x63\x7b\x7d\x12\xa9\x94\xd7\x9b\xdd\xf3\xd5\x5d\x92\x4a\x9c\xba\xdb\xca\xa6\xae\x40\xb2\x25\xc2\x03\x02\x0c\x00\x4a\xa3\x4c\xe9\xbf\x5f\x35\x40\x52\xd4\xcb\x4c\xec\xb5\xaf\x6a\xfd\x61\x4c\x02\x8d\x46\xa3\xfb\x69\xf4\x0b\x95\x7c\xf5\xe7\xef\xde\xbe\xff\xdb\xf7\xdf\x40\x69\x2b\xb1\x78\x92\xd0\x7f\x20\x98\x5c\xa5\x01\xca\x60\xf1\xe4\x22\x29\x91\x15\x8b\x27\x17\x17\x49\x85\x96\x41\x5e\x32\x6d\xd0\xa6\x41\x63\x97\xe1\xeb\x60\x3f\x51\x5a\x5b\x87\xf8\x6b\xc3\xd7\x69\xf0\x3f\xe1\x4f\x6f\xc2\xb7\xaa\xaa\x99\xe5\x99\xc0\x00\x72\x25\x2d\x4a\x9b\x06\xef\xbe\x49\xb1\x58\xe1\x60\x9d\x64\x15\xa6\xc1\x9a\xe3\xa6\x56\xda\x0e\x48\x37\xbc\xb0\x65\x5a\xe0\x9a\xe7\x18\xba\x97\x4b\xe0\x92\x5b\xce\x44\x68\x72\x26\x30\xbd\x0a\x16\x4f\x88\x8f\xe5\x56\xe0\xe2\xfe\x3e\xfa\x16\xed\x46\xe9\xdb\xdd\x6e\x06\x6f\x1a\x5b\xa2\xb4\x3c\x67\x16\x0b\xf8\x0b\x6b\x72\xb4\x49\xec\x29\xdd\x22\xc1\xe5\x2d\x94\x1a\x97\x69\x40\xa2\x9b\x59\x1c\xe7\x85\xfc\x60\xa2\x5c\xa8\xa6\x58\x0a\xa6\x31\xca\x55\x15\xb3\x0f\xec\x2e\x16\x3c\x33\xb1\xdd\x70\x6b\x51\x87\x99\x52\xd6\x58\xcd\xea\xf8\x26\xba\x89\xfe\x18\xe7\xc6\xc4\xfd\x58\x54\x71\x19\xe5\xc6\x04\xa0\x51\xa4\x81\xb1\x5b\x81\xa6\x44\xb4\x01\xc4\x8b\x7f\x6c\xdf\xa5\x92\x36\x64\x1b\x34\xaa\xc2\xf8\x45\xf4\xc7\x68\xea\xb6\x1c\x0e\x3f\xbe\x2b\x6d\x6b\x72\xcd\x6b\x0b\x46\xe7\x1f\xbd\xef\x87\x5f\x1b\xd4\xdb\xf8\x26\xba\x8a\xae\xda\x17\xb7\xcf\x07\x13\x2c\x92\xd8\x33\x5c\x7c\x16\xef\x50\x2a\xbb\x8d\xaf\xa3\x17\xd1\x55\x5c\xb3\xfc\x96\xad\xb0\xe8\x76\xa2\xa9\xa8\x1b\xfc\x62\xfb\x3e\x64\xc3\x0f\xc7\x26\xfc\x12\x9b\x55\xaa\x42\x69\xa3\x0f\x26\xbe\x8e\xae\x5e\x47\xd3\x6e\xe0\x94\xbf\xdb\x80\x8c\x46\x5b\x5d\x44\x6b\xd4\x84\x5c\x11\xe6\x28\x2d\x6a\xb8\xa7\xd1\x8b\x8a\xcb\xb0\x44\xbe\x2a\xed\x0c\xae\xa6\xd3\x67\xf3\x73\xa3\xeb\xd2\x0f\x17\xdc\xd4\x82\x6d\x67\xb0\x14\x78\xe7\x87\x98\xe0\x2b\x19\x72\x8b\x95\x99\x81\xe7\xec\x26\x76\x6e\xcf\x5a\xab\x95\x46\x63\xda\xcd\x6a\x65\xb8\xe5\x4a\xce\x08\x51\xcc\xf2\x35\x9e\xa3\x35\x35\x93\x27\x0b\x58\x66\x94\x68\x2c\x1e\x09\x92\x09\x95\xdf\xfa\x31\xe7\xcd\xc3\x43\xe4\x4a\x28\x3d\x83\x4d\xc9\xdb\x65\xe0\x36\x82\x5a\x63\xcb\x1e\x6a\x56\x14\x5c\xae\x66\xf0\xaa\x6e\xcf\x03\x15\xd3\x2b\x2e\x67\x30\xdd\x2f\x49\xe2\x4e\x8d\x49\xec\x2f\xae\x27\x17\x49\xa6\x8a\xad\xb3\x61\xc1\xd7\x90\x0b\x66\x4c\x1a\x1c\xa9\xd8\x5d\x48\x07\x04\x74\x0f\x31\x2e\xbb\xa9\x83\x39\xad\x36\x01\xb8\x8d\xd2\xc0\x0b\x11\x66\xca\x5a\x55\xcd\xe0\x8a\xc4\x6b\x97\x1c\xf1\x13\xa1\x58\x85\x57\xd7\xdd\xe4\x45\x52\x5e\x75\x4c\x2c\xde\xd9\xd0\xd9\xa7\xb7\x4c\xb0\x48\x78\xb7\x76\xc9\x60\xc9\xc2\x8c\xd9\x32\x00\xa6\x39\x0b\x4b\x5e\x14\x28\xd3\xc0\xea\x06\x09\x47\x7c\x01\xc3\xeb\xef\x81\xdb\xaf\xbc\xea\xe4\x8a\x0b\xbe\x6e\x8f\x35\x78\x3c\x3a\xe1\xc3\x87\x78\x0d\xed\x83\x5a\x2e\x0d\xda\x70\x70\xa6\x01\x31\x97\x75\x63\xc3\x95\x56\x4d\xdd\xcf\x5f\x24\x6e\x14\x78\x91\x06\x8d\x16\x41\x7b\xfd\xbb\x47\xbb\xad\x5b\x55\x04\xfd\xc1\x95\xae\x42\xb2\x84\x56\x22\x80\x5a\xb0\x1c\x4b\x25\x0a\xd4\x69\xf0\xa3\xca\x39\x13\x20\xfd\x99\xe1\xa7\x1f\xfe\x13\x5a\x93\x71\xb9\x82\xad\x6a\x34\x7c\x63\x4b\xd4\xd8\x54\xc0\x8a\x82\xe0\x1a\x45\xd1\x40\x10\x87\xdd\x53\x51\xc3\xcc\xca\x3d\xd5\x45\x92\x35\xd6\xaa\x9e\x30\xb3\x12\x32\x2b\xc3\x02\x97\xac\x11\x16\x0a\xad\xea\x42\x6d\x64\x68\xd5\x6a\x45\x91\xce\x1f\xc2\x2f\x0a\xa0\x60\x96\xb5\x53\x69\xd0\xd1\x76\x36\x64\xa6\x56\x75\x53\xb7\x56\xf4\x83\x78\x57\x33\x59\x60\x41\x36\x17\x06\x83\xc5\x5f\xf9\x1a\xa1\x42\x7f\x96\x8b\x63\x48\xe4\x4c\xa3\x0d\x87\x4c\x4f\x80\x91\xc4\x5e\x18\x7f\x24\x68\xff\x25\x8d\xe8\x38\xf5\x47\xa8\x50\x36\x70\xf0\x16\x6a\xba\x57\x82\xc5\xfd\xbd\x66\x72\x85\xf0\x94\x17\x77\x97\xf0\x94\x55\xaa\x91\x16\x66\x29\x44\x6f\xdc\xa3\xd9\xed\x0e\xb8\x03\x24\x82\x2f\x12\xf6\x18\xbc\x41\xc9\x5c\xf0\xfc\x36\x0d\x2c\x47\x9d\xde\xdf\x13\xf3\xdd\x6e\x0e\xf7\xf7\x7c\x09\x4f\xa3\x1f\x30\x67\xb5\xcd\x4b\xb6\xdb\xad\x74\xf7\x1c\xe1\x1d\xe6\x8d\xc5\xf1\xe4\xfe\x1e\x85\xc1\xdd\xce\x34\x59\xc5\xed\xb8\x5b\x4e\xe3\xb2\xd8\xed\x48\xe6\x56\xce\xdd\x0e\x62\x62\x2a\x0b\xbc\x83\xa7\xd1\xf7\xa8\xb9\x2a\x0c\x78\xfa\x24\x66\x8b\x24\x16\x7c\xd1\xae\x3b\x54\x52\xdc\x88\x3d\x5e\x62\x02\x4c\x8f\x73\xe7\x36\x4e\xd4\xa1\xa4\x67\xbc\x60\x15\xf6\xd2\xb7\x78\x30\xdc\xe2\x2d\x6e\xd3\xe0\xfe\x7e\xb8\xb6\x9d\xcd\x99\x10\x19\x23\xbd\xf8\xa3\xf5\x8b\x7e\x43\xc2\xe9\x9a\x1b\x97\x52\x2d\x3a\x09\xf6\x62\x7f\xa4\x5b\x1f\x5d\x5c\x56\xd5\x33\xb8\xb9\x1e\xdc\x5a\xe7\x3c\xfe\xd5\x91\xc7\xdf\x9c\x25\xae\x99\x44\x01\xee\x6f\x68\x2a\x26\xba\xe7\xd6\x5b\x06\xce\x77\xbc\x28\xa4\x3b\xba\x17\xad\xbf\xeb\xa7\x73\x50\x6b\xd4\x4b\xa1\x36\x33\x60\x8d\x55\x73\xa8\xd8\x5d\x1f\xef\x6e\xa6\xd3\xa1\xdc\x94\x0a\xb2\x4c\xa0\xbb\x5d\x34\xfe\xda\xa0\xb1\xa6\xbf\x4b\xfc\x94\xfb\x4b\x57\x4a\x81\xd2\x60\x71\xa4\x0d\xda\x91\x54\xeb\xa8\x06\xa6\xef\x95\x79\x56\xf6\xa5\x52\x7d\x08\x19\x8a\xd1\xb2\x1e\x44\xbb\x60\x91\x58\xbd\xa7\xbb\x48\x6c\xf1\x49\x21\x40\x53\x8a\xf7\x50\x04\xf0\x37\x1a\x9d\xbd\x46\xd4\x3e\xbf\x20\xc8\x82\x7b\x4d\x62\x5b\x7c\xc6\xce\x04\xc2\x8c\x19\xfc\x98\xed\x5d\xa4\xdf\x6f\xef\x5e\x3f\x77\xff\x12\x99\xb6\x19\x32\xfb\x31\x02\x2c\x1b\x59\x0c\xce\xef\xee\xce\xcf\x15\xa0\x91\x7c\x8d\xda\x70\xbb\xfd\x58\x09\xb0\xd8\x8b\xe0\xdf\x0f\x45\x48\x62\xab\x1f\xc7\xda\xf0\xe5\x0b\x39\xf7\xef\xa5\x24\x37\x8b\x7f\x57\x1b\x28\x14\x1a\xb0\x25\x37\x40\xc1\xf5\xeb\x24\x2e\x6f\x7a\x92\x7a\xf1\x9e\x26\x9c\x52\x61\xe9\x52\x0b\xe0\x06\x74\x23\x5d\xe4\x55\x12\x6c\x89\x87\xe9\x48\x1b\xa4\x23\x78\xaf\x28\xa5\x5b\xa3\xb4\x50\x31\xc1\x73\xae\x1a\x03\x2c\xb7\x4a\x1b\x58\x6a\x55\x01\xde\x95\xac\x31\x96\x18\xd1\xf5\xc1\xd6\x8c\x0b\xe7\x4b\xce\xa4\xa0\x34\xb0\x3c\x6f\xaa\x86\x52\x52\xb9\x02\x94\xaa\x59\x95\xad\x2c\x56\x81\x0f\x4c\x42\xc9\x55\x2f\x8f\xa9\x59\x05\xcc\x5a\x96\xdf\x9a\x4b\xe8\x6e\x05\x60\x1a\xc1\x72\x2c\x68\x55\xae\xaa\x4a\x49\xb8\xd1\x05\xd4\x4c\xdb\x2d\x98\xc3\xdc\x82\xe5\xb9\x8b\x72\x11\xbc\x91\x5b\x25\x11\x4a\xb6\x76\x12\xc2\x7b\x5f\x4e\x90\x5c\x7f\x61\x39\x66\x4a\xf5\xd4\x50\xb1\x6d\xb7\x5d\x2b\xfd\x86\xdb\x92\x7b\xf5\xd4\xa8\x2b\x5a\x5a\x80\xe0\x15\xb7\x26\x4a\xe2\x7a\x7f\xa3\xee\x63\xb3\x08\x4b\xa5\xf9\x6f\x94\xd8\x88\xe1\xf5\x69\x8f\x2e\x97\xee\x6e\x74\x56\x17\xb8\xb4\x33\x78\xe1\xef\xc6\x63\x1c\xb7\x15\xd0\x39\x10\x77\x3c\x5d\x65\x49\x01\x67\x06\x37\x3e\x9d\xf5\x89\x44\x61\x07\x12\x14\x47\x50\xf3\x9b\xbe\x7e\x5d\xdf\xf5\x72\xf4\x39\xf1\xb4\x67\x42\x08\x38\x54\xca\x9a\xf7\x6a\xbc\x84\x8a\xdd\x22\x30\x48\xd8\x51\x85\xdc\x0a\xed\xea\x2b\xee\xfa\x03\xb1\xdd\x20\xda\xaf\xc9\x75\xd3\x1f\x3c\x43\x2e\x57\xcf\xae\xa7\x1e\x91\xf4\x40\xec\x9f\x5d\x4f\xb9\xb4\xea\xd9\xf5\x74\x7a\x37\xfd\xc8\x7f\xcf\xae\xa7\x4a\x3e\xbb\x9e\xda\x12\x9f\x5d\x4f\x9f\x5d\xdf\x0c\xb1\xec\x47\xba\xcc\x92\xa8\xd0\xd0\x6e\x1d\xc4\x03\xb0\x4c\xaf\xd0\xa6\xc1\xff\xb2\x4c\x35\x76\x96\x09\x26\x6f\x83\x85\x13\x97\xb2\x0d\x87\x82\xf3\xf9\x29\xd4\xcc\x10\x24\x48\x62\x87\x92\xb6\x17\x62\x60\x6c\x1a\xad\x55\x23\x29\x2a\x02\x9d\xd9\x79\xa8\x1c\x11\xca\x48\x31\x93\x28\xc9\x74\xbc\x78\xab\xea\x6d\xe8\x98\xb8\xe5\x27\x6a\x34\x4d\x5d\x2b\x6d\xa3\xa1\x3a\x19\xd5\x41\x02\x4d\xfc\x7a\xfa\xf2\xf5\xab\x47\xc5\x37\x94\x65\xbb\x33\xf4\x12\xb2\x4c\xad\x11\x7c\x4e\x9f\xa9\x3b\x60\xb2\x80\x25\xd7\x08\x6c\xc3\xb6\x5f\x25\x71\xe1\x2a\xb0\xcf\x47\xed\xb2\xf5\xae\x7f\x2a\xd8\x76\x2e\x7f\x09\x75\x93\x09\x6e\x4a\x60\x20\x71\x03\x89\xb1\x5a\xc9\xd5\xc2\x8d\xe6\x54\x92\xba\x57\xa8\x95\xb1\x8f\x99\x1f\xab\x0c\x8b\xe2\x0c\x00\xbe\x94\xfd\x37\x9b\x4d\xd4\x69\xd2\x19\xbf\x44\x51\xc7\x74\xfd\x35\x92\xdb\x6d\xec\xdd\x48\xc9\xf8\x6b\x5e\xa4\xd7\xaf\xaf\x5f\xbd\xba\x7e\xf1\x6f\xaf\x5f\xbe\xbc\x7e\xfd\xe2\xe5\x43\xc8\xa0\x43\x7d\x26\x30\x7c\x1a\xfd\xad\xa2\xaa\xb5\xcf\xa1\x3d\x5e\xba\xdc\x8d\x22\x74\x41\x35\x88\x0e\xfe\x61\x0c\x35\x92\x12\x91\x90\x89\xb3\x39\xc4\x27\xa0\xc8\xc1\xe8\x11\xc9\x3e\x13\x5a\x1d\x7c\x08\x29\xaa\xb1\x74\xc2\xae\x98\xe7\x4a\xf6\x70\xba\x04\xc3\xab\x5a\x6c\x21\xdf\x5b\xfd\x3c\xae\x1e\x34\xca\xef\xc2\xea\xd0\x6c\x1e\x64\x2e\xfa\x57\xaa\x40\x8a\xfa\xa6\x31\x39\xd6\xae\xcb\x4b\x91\xf4\x4f\xdb\xdf\x98\xb4\x5c\x62\x17\x71\x23\xf8\x4e\x8a\x2d\x34\x06\x61\xa9\x34\x14\x98\x35\xab\x95\x4b\x13\x34\xd4\x9a\xaf\x99\xc5\x2e\xcc\x9a\x16\x15\x3d\x28\x06\x95\x0d\xa5\x3c\x62\x90\x81\xfc\x4d\x35\x90\x33\x09\x56\xb3\xfc\xd6\x7b\x4a\xa3\x35\x79\x4a\x8d\xfe\x34\x7d\xa0\xcf\x50\xa8\x8d\x23\xf1\xe7\x5e\x72\x14\x2e\xea\x1b\x44\x28\xd5\x06\xaa\x26\x77\x0e\x49\x51\xdd\x1d\x62\xc3\xb8\x85\x46\x5a\x2e\xbc\x3e\x6d\xa3\x25\xe5\x08\x78\x10\xa5\x4f\x6a\xbf\x04\xab\xc5\xfb\x12\xcf\xa4\x44\x7d\xd5\x06\x1a\xdf\x7a\x72\xa8\xb5\xb2\x98\x93\x41\x81\xad\x18\x97\x86\x2c\xe2\xf2\x00\xac\x3e\xa2\xaa\xeb\x9f\xda\x87\x7d\x87\xd2\x4d\xc7\x31\xfc\x55\xa8\x8c\x09\x58\x13\xd2\x33\x41\xe9\x9c\x82\x52\xd1\xd1\x07\xda\x32\x96\xd9\xc6\x80\x5a\xba\x51\x2f\x39\xad\x5f\x33\x4d\x16\xc4\xaa\xb6\x90\xb6\xfd\x35\x1a\x33\xa8\xd7\x6d\xd7\x90\x5e\xa9\x72\x3f\x98\xef\xb5\x9e\xc2\xcf\xbf\xcc\x9f\xb4\xa2\xfc\x19\x97\x0e\x12\x84\x6f\x7f\x64\x5b\x32\x0b\xb9\x46\x66\xd1\x40\x2e\x94\x69\xb4\x97\xb0\xd0\xaa\x06\x92\xb2\xe3\xd4\x71\xa6\x89\xda\xed\xd6\x31\x19\x97\xcc\x94\x93\xb6\x3d\xa8\xd1\x59\xa9\x9f\xeb\xc6\x2f\x08\x75\x63\x62\xc0\xd3\xe9\x1c\x78\xd2\xf1\x8d\x04\xca\x95\x2d\xe7\xc0\x9f\x3f\xef\x89\x2f\xf8\x12\xc6\x1d\xc5\xcf\xfc\x97\xc8\xde\x45\xb4\x0b\xa4\x29\x0c\x77\x73\x1b\xb6\x7c\x4c\x2d\x78\x8e\x63\x7e\x09\x57\x93\x79\x37\x9b\x69\x64\xb7\xdd\x5b\x6b\x47\xff\x9f\xfb\xbb\x9b\x1f\x6a\xc6\x29\xff\x40\x37\xbe\xf6\x37\xc0\x60\xc5\x8d\x85\x46\x0b\x68\x7d\xd8\x9b\xa0\x37\x88\xa3\x1b\x6a\xe5\x04\x97\xed\x43\x8b\xa9\xee\x08\x9e\x4d\x64\x50\x16\xe3\xff\xf8\xf1\xbb\x6f\x23\x63\x35\x97\x2b\xbe\xdc\x8e\xef\x1b\x2d\x66\xf0\x74\x1c\xfc\x4b\xa3\x45\x30\xf9\x79\xfa\x4b\xb4\x66\xa2\xc1\x4b\x67\xef\x99\xfb\x7b\xb2\xcb\x25\xb4\x8f\x33\x38\xdc\x70\x37\x99\xcc\xcf\xf7\x49\x06\x6d\x1d\x8d\x06\xed\x98\x08\x7b\xe0\x1f\xeb\x88\x41\x85\xb6\x54\xce\x75\x35\xe6\x4a\x4a\xcc\x2d\x34\xb5\x92\xad\x4a\x40\x28\x63\xf6\x40\xec\x28\xd2\x53\x50\xb4\xf4\xa9\x0b\xd6\xff\x8d\xd9\x8f\x2a\xbf\x45\x3b\x1e\x8f\x37\x5c\x16\x6a\x13\x09\xe5\xaf\xda\x88\x9c\x54\xe5\x4a\x40\x9a\xa6\xd0\x46\xd1\x60\x02\x5f\x43\xb0\x31\x14\x4f\x03\x98\xd1\x23\x3d\x4d\xe0\x39\x1c\x2f\x2f\x29\xde\x3f\x87\x20\x66\x35\x0f\x26\xde\x1d\x3a\xc5\x2b\x59\xa1\x31\x6c\x85\x43\x01\x5d\x65\xd4\x83\x8c\xce\x51\x99\x15\xa4\xe0\x0c\x54\x33\x6d\xd0\x93\x44\x54\x8d\x77\x68\x23\xcc\x3a\xb2\x34\x05\xd9\x08\xb1\x07\xa9\x77\x8a\x79\x07\xbf\x03\xf2\xc8\xc7\x9a\xaf\xd2\x14\xa8\x34\x25\x15\x17\xfb\x95\x64\x7c\x5f\x44\x4f\x22\x8a\x0b\xfb\x15\x93\xf9\x10\xcd\x07\xdc\xb0\xf8\x3d\x76\x58\x1c\xf3\xc3\xe2\x01\x86\xae\x67\xf1\x18\x3f\xdf\xe3\x18\xb0\x73\x03\x0f\x70\x93\x4d\x95\xa1\x7e\x8c\x9d\xef\x59\xb4\xec\x9c\xaa\xdf\x49\x3b\x58\x7b\x09\x57\xaf\x26\x0f\x70\x47\xad\xd5\x83\xcc\xa5\xb2\xdb\xf1\xbd\x60\x5b\xca\x99\x60\x64\x55\xfd\xd6\xb5\x18\x46\x97\x2e\xe2\xce\xa0\xe7\x70\xe9\x9a\xc7\x33\x18\xb9\x37\x9a\xe7\x15\xba\x55\x2f\xa7\xd3\xe9\x25\x74\x5f\x5d\xfe\xc4\xc8\x09\x75\x83\xbb\x07\xe4\x31\x4d\x9e\x53\xdc\xff\x1c\x89\x5a\x1e\xbd\x4c\xed\xfb\x67\x48\xd5\xc7\x86\x03\xb1\xe0\x0f\x7f\x80\x93\xd9\x43\x18\xc7\x31\xfc\x17\xa3\x32\x5c\x08\xd7\x3d\x70\x4d\x83\x9e\xbe\xe2\xc6\xb8\x62\xdc\x40\xa1\x24\xb6\x6b\x3e\xed\xda\x3f\x91\xb1\x25\x83\x05\x4c\x8f\x05\xa4\xeb\x70\x10\x16\xce\x44\x8b\x01\xdf\xc3\x40\x70\xb1\x1b\xee\x77\xb0\x92\x57\x08\x5f\xa5\x10\x04\xc3\xc5\x27\x14\x44\xd0\x33\xbb\x30\x68\xdf\x7b\x5b\x8c\xdb\xe8\x78\x2e\x76\x4d\x2e\xe1\x66\x3a\x9d\x4e\x4e\x84\xd8\xed\xd5\xfb\xa6\xa6\xb4\x09\x98\xdc\xba\x2b\xb1\xd7\xad\x4b\x1c\x29\x05\xa2\x2b\x4d\x40\xae\x84\xf0\x39\x4b\xbb\x94\x14\xdc\x36\x4f\x52\x08\xaf\xe6\x67\xa2\xe8\x40\x93\x83\xa3\x1d\x9b\xe7\x8c\xee\x8f\x4d\x74\xa8\xb3\x23\xe2\xf0\xea\xc0\x28\x07\xf6\x3a\x6f\x98\x8b\x5e\x6e\xbe\xd7\xe8\x91\xb9\xf6\xf6\x3a\xd6\xd9\x40\x7e\xcf\xe7\xf9\xd5\x47\x1e\xa3\x9f\xae\x1b\x53\x8e\x8f\x04\x9d\xcc\x4f\x6d\xf3\xce\xa2\xa6\x2c\x59\x51\xc8\x22\x5b\x50\x29\xa0\xf1\xc4\x24\x2e\x55\xd7\x18\x6a\x94\x05\xea\x2e\xa5\xf0\x99\x3d\x25\x80\x07\x26\xf3\x55\xe5\x10\x4e\x9f\xe8\x30\x2e\x25\x53\x12\x01\x00\x8e\x9c\xc0\x01\xf5\x00\xa9\x44\x8c\x82\xd5\x06\x0b\x48\xc1\x7f\x04\x1f\x4f\xa2\x46\xf2\xbb\xf1\x24\x6c\xdf\x8f\x79\x74\xf3\xf3\xbe\x4c\xec\xc4\x7e\x9e\x42\x90\x58\x0d\xbc\x48\x47\x01\x3c\x3f\xe7\x82\x14\x75\x47\x8b\xbd\x04\xc3\xa5\x00\x89\x2d\x16\xae\x0f\xea\xeb\xb5\xbf\x07\x19\xcb\x6f\x57\xae\x10\x9a\x51\xaa\x35\x3e\x61\xcb\xd6\xcc\x32\xed\xb8\x4e\xe6\xb0\x27\x6f\x0b\xc5\x9c\x8c\x33\x07\x5f\x91\xba\x76\x2b\xf4\x9f\x28\xdc\x5b\xa6\x74\x81\x3a\xd4\xac\xe0\x8d\x99\xc1\x8b\xfa\x6e\xfe\xf7\xee\x13\x8e\x6b\x0a\x3f\x2a\x6a\xad\x71\x71\x22\x51\xdb\x65\x7c\x0e\x41\x12\x13\xc1\xef\xb1\xe9\x0f\x3b\xfc\xf8\x0e\x67\x5a\xdf\xd0\x7f\x1a\x6f\xc7\x2b\x5e\x14\x02\x49\xe0\x3d\x7b\x72\x46\xb2\xff\xd0\xa5\x0e\xb7\x84\xb6\xe7\xbd\x5f\xb3\x03\x14\x06\x1f\x59\xd0\xb7\xcf\x47\x04\x80\x90\x8e\xcc\x9d\xce\xdb\x62\xdb\x0d\xeb\x91\xd3\x45\xfb\x53\x8a\xa2\xd1\x2e\xd7\x1a\x87\x2d\xc0\x2e\x61\x64\x28\xf7\x2b\xcc\x68\x12\x95\x4d\xc5\x24\xff\x0d\xc7\x14\x97\x26\x5e\x57\xae\x1f\x1f\x9c\x5e\xc9\x27\xc2\xec\x1b\xe5\xa3\x2e\xc6\x8d\x5a\x25\x8e\x3a\xeb\xbe\xd8\xd7\xf6\x33\x98\xce\x47\x9f\xa8\xa1\xf3\xbb\x84\x19\xd3\x30\x7c\x09\xbb\xe0\x0b\x5a\xd1\xee\xdd\x5c\xc6\xf4\xc8\x77\x32\x5c\x7e\x2e\xd5\x26\x1d\xdd\x4c\x7b\x21\xbd\xa1\x9d\x9d\x47\x2d\xd6\x4e\x8c\x41\x52\x76\xae\xb9\x80\x9b\xe9\x97\x90\xd6\x77\x43\x8e\x4e\x60\x35\xaf\xb1\x00\x96\x5b\xbe\xc6\xff\x87\x83\x7c\x01\x25\x7f\xb2\x88\x84\xc3\x4e\x79\x0e\xa6\x07\xf2\xd2\x6c\xaf\xdb\x7f\x25\x7f\x83\xd8\x69\xf8\x39\x04\x67\x0f\xf2\x20\x12\x8f\x08\x8f\x5c\xfb\x61\xbf\x77\x1f\x98\x82\xe3\x98\x42\xd9\x6e\xff\x71\x74\x12\x95\xb6\x12\xe3\x20\xb1\xee\x47\x32\x24\x73\xcf\xc1\x31\xf0\xc3\x87\x29\xdd\xee\xb0\x90\xa1\xfa\x1d\x8f\xea\x2c\x18\x24\x27\x7d\x2d\xd6\x65\x22\xb0\xdb\xff\x96\x28\x8e\xe1\x47\xcb\xb4\x05\x06\x3f\xbd\x83\xa6\x2e\x98\xf5\x9f\x72\x28\x3e\xfa\x4f\x25\xdd\x8f\x8d\x32\xa6\x0d\x2c\x95\xde\x30\x5d\xb4\xfd\x19\x5b\xe2\xd6\x7d\xca\xe9\x52\x3f\x83\xf6\x1d\xdd\x62\x6b\x26\xc6\x27\x75\xdf\xd3\xf1\x28\x1a\x9a\x7c\x34\x89\x90\xe5\xe5\x29\xa1\x8b\x58\xfd\xbe\x29\x7c\xeb\x4a\x80\xf1\xd3\xb1\x2d\xb9\x99\x44\xcc\x5a\x3d\x1e\x1d\x80\x61\x34\x21\xbb\x5e\x0d\x4a\xb2\x7e\x79\x72\xe0\x56\x8f\xf1\xd8\x27\xd3\x7d\x22\xd0\x91\xe7\xc6\x8c\x3d\xae\x46\x97\x03\xde\x87\xb0\x1a\x3d\x1b\xf5\x86\xda\xbb\xf7\xfe\x1c\xe9\x59\x49\x0e\x58\x8f\xc8\xcb\x46\x27\xdb\xb3\xa2\x78\x4b\xfe\x33\x0e\xce\x78\xfa\x31\x3a\x26\xbd\xb2\xfd\x7d\xfd\xa8\x96\xfd\xcf\x32\x1e\x50\x31\x2f\x46\x93\xc8\x34\x99\xef\x4d\x8c\x5f\xf6\x05\x58\x47\xe6\xc0\x7b\x1c\x0a\x4e\x12\x0a\xda\xe2\x30\xa9\x08\x8f\x92\x90\x47\xa2\x46\xbb\xa5\x3f\xd5\xee\x92\x14\x3e\x9d\xf4\xad\xad\x6f\x0c\x25\x57\xbe\xf5\xbf\xc1\xcc\xb8\x4e\x02\xb4\x78\x77\xdd\x1c\xdf\xb5\x79\xf3\xfd\xbb\x41\xe7\xa6\xf7\x88\xb1\xe3\xde\xff\x0e\xf0\x5c\x9f\xe4\xec\x0f\x0f\x37\x9b\x4d\xb4\x52\x6a\x25\xfc\x4f\x0e\xfb\x46\x4a\xcc\x6a\x1e\x7d\x30\x01\x30\xb3\x95\x39\x14\xb8\x44\xbd\x18\xb0\x6f\xbb\x2b\x49\xec\x7f\x12\x97\xc4\xfe\x57\xbf\xff\x17\x00\x00\xff\xff\x31\x9f\x54\x5e\x06\x2c\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
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
