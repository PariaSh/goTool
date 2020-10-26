package evaluation

import "testing"

// _volume_dataset_ImageMagick-7.0.1_10_arm-gcc-O0_libMagickCore-7.Q16HDRI.so.0.0.0
// _volume_dataset_ImageMagick-7.0.1_10_x86-gcc-O0_libMagick++-7.Q16HDRI.so.0.0.0

func TestGetReg(t *testing.T)  {

	ss := rePrevious.FindStringSubmatch("_volume_dataset_ImageMagick-7.0.1_10_arm-gcc-O0_libMagickCore-7.Q16HDRI.so.0.0.0")
	t.Logf("%+v", ss)
}