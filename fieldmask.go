package fieldmask

import (
	"errors"
	"reflect"
	"strings"

	fieldmask "google.golang.org/genproto/protobuf/field_mask"
)

var (
	ErrorNilSource    = errors.New("source object is nil")
	ErrorNilDest      = errors.New("destination object is nil")
	ErrorTypeNotMatch = errors.New("types of source and destination objects do not match")
)

// Merge will take the fields of `source` that are included as
// paths in `mask` and write them to the corresponding fields of `dest`
func Merge(source, dest interface{}, mask *fieldmask.FieldMask) error {

	if source == nil {
		return ErrorNilSource
	}

	if dest == nil {
		return ErrorNilDest
	}

	if reflect.TypeOf(source) != reflect.TypeOf(dest) {
		return ErrorTypeNotMatch
	}

pathsloop:
	for _, fullpath := range mask.GetPaths() {
		srcVal := reflect.ValueOf(source).Elem()
		dstVal := reflect.ValueOf(dest).Elem()

		subpaths := strings.Split(fullpath, ".")
		for _, path := range subpaths {
			// return all fields when path equals to *
			if path == "*" {
				break
			}
			for dstVal.Kind() == reflect.Ptr {
				if dstVal.IsNil() {
					dstVal.Set(reflect.New(dstVal.Type().Elem()))
				}
				dstVal = dstVal.Elem()
				srcVal = srcVal.Elem()
			}

			// avoid panic when dstVal.Kind() is not struct
			if dstVal.Kind() != reflect.Struct {
				continue pathsloop
			}
			srcVal = srcVal.FieldByName(path)
			dstVal = dstVal.FieldByName(path)
		}

		dstVal.Set(srcVal)
	}

	return nil
}
