package config

import (
	"fmt"
	"reflect"
	"strings"
)

type ifaceProcesser interface {
	process(string) error
}

// traverseIface traverses interface and executes processIface on given non-structural interface.
func traverseIface(processor ifaceProcesser, iface interface{}, parts ...string) error {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	// Get pointer value
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
		ift = ift.Elem()
	}

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			traverseIface(processor, v.Interface(), append(parts, tv)...)
		default:
			mapKey := strings.Join(append(parts, tv), ".")

			if err := processor.process(mapKey); err != nil {
				return fmt.Errorf("failed to proccess iface: %v", err)
			}
		}
	}

	return nil
}
