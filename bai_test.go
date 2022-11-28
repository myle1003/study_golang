package main

import "testing"
import "os"

func TestCreateArrayBai(t *testing.T) {
	boBai := createArrayBai()

	if len(boBai) !=16 {
		t.Errorf("mong muon co 16 la bai, thuc te la %v ", len(boBai))
	}
}

func TestChiaBai(t *testing.T) {
	boBai := createArrayBai()

	bo1,bo2 := chiaBai(boBai,8)

	if len(bo1) !=8 {
		t.Errorf("bo1 mong muon co 8 la bai, thuc te la %v ", len(bo1))
	}

	if len(bo2) !=8 {
		t.Errorf("bo2 mong muon co 8 la bai, thuc te la %v ", len(bo2))
	}
}

func TestLuaVaTaoFile(t *testing.T) {
	// boBai := createArrayBai()
	boBai := arrayBai{"3 co","4 ro", "5 bich"}

	err := boBai.luufile("_testBai")

	if err != nil {
		t.Errorf("luu bai xuong file bi loi %v", err)
	}

	boBai, err = taoBaiTuFile("_testBai")
	if err != nil {
		t.Errorf("doc bai xuong file bi loi %v", err)
	}

	if len(boBai) !=3 {
		t.Errorf("bo1 mong muon co 3 la bai, thuc te la %v ", len(boBai))
	}
	os.Remove("_testBai")

}


