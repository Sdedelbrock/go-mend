package mend

import (
	"errors"
	"reflect"
)

func Mend(dst, src interface{}) error {
	vdst := reflect.ValueOf(dst)
	vsrc := reflect.ValueOf(src)
	if src == nil || dst == nil {
		return errors.New("Nil value passed to mend")
	}
	if vsrc.Kind() == reflect.Ptr {
		vsrc = vsrc.Elem()
	}
	return mend(vdst, vsrc)
}

func mend(dst, src reflect.Value) error {
	switch dst.Kind() {
	case reflect.Struct:
		for i := 0; i < dst.NumField(); i++ {
			mend(dst.Field(i), src.Field(i))
		}
	case reflect.Slice:
		if !dst.CanAddr() {
			return nil
		}
		dst.Set(reflect.AppendSlice(dst, src))

	case reflect.Map:
		for _, key := range src.MapKeys() {
			srcIndex := src.MapIndex(key)
			dstIndex := dst.MapIndex(key)

			if !isZero(srcIndex) && isZero(dstIndex) {
				dst.SetMapIndex(key, srcIndex)
			} else {
				switch reflect.TypeOf(srcIndex.Interface()).Kind() {
				case reflect.Map:
					mend(dstIndex, srcIndex)
				}
			}
		}
	case reflect.Interface, reflect.Ptr:
		mend(dst.Elem(), src)

	default:
		if !isZero(src.Interface()) && isZero(dst.Interface()) {
			dst.Set(src)
		}

	}
	return nil
}

func isZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
