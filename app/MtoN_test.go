package app

import "testing"

func TestB64ToDec(t *testing.T) {
	var b64 string = "BA"
	dec := B64ToDec(b64)
	if dec != 64 {
		t.Errorf("64进制转换10进制错误,转换结果为：%d", dec)
	}
}

func TestDecToB64(t *testing.T) {
	var dec int = 10
	b64 := DecToB64(dec)
	if b64 != "K" {
		t.Errorf("10进制转换64进制错误,转换结果为：%s", b64)
	}
}
