package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIDispatchErr(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toIDispatchErr(v, nil)
}

func TestToInt64Err(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toInt64Err(v, nil)
}

func TestToInt32Err(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toInt32Err(v, nil)
}

func TestToFloat64Err(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toFloat64Err(v, nil)
}

func TestToFloat32Err(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toFloat32Err(v, nil)
}

func TestToStringErr(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toStringErr(v, nil)
}

func TestToBoolErr(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toBoolErr(v, nil)
}

func TestToTimeErr(t *testing.T) {
	v := &ole.VARIANT{}
	_, _ = toTimeErr(v, nil)
}

func TestVariantToIDispatch(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToIDispatch(v)
}

func TestVariantToInt64(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToInt64(v)
}

func TestVariantToInt32(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToInt32(v)
}

func TestVariantToFloat64(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToFloat64(v)
}

func TestVariantToFloat32(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToFloat32(v)
}

func TestVariantToString(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToString(v)
}

func TestVariantToBool(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToBool(v)
}

func TestVariantToTime(t *testing.T) {
	v := &ole.VARIANT{}
	_ = variantToTime(v)
}
