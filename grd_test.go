package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/softlandia/xlib"
)

func TestCollExclude(t *testing.T) {
	testFileName := "CollExclude1_test.txt"
	f, _ := os.Create(testFileName)
	f.WriteString("7397758.472745	5044558.321537	2263.000000	1\n\r")
	f.WriteString("7402003.114255	5038343.000243	9999999999.990000	3\n\r")
	/*fmt.Fprintf(f, "7397758.472745	5044558.321537	2263.000000	1\n")
	fmt.Fprintf(f, "7402003.114255	5038343.000243	9999999999.990000	3\n")*/
	f.Close()

	CollExclude1(testFileName)

	testFileName = xlib.ChangeFileExt(testFileName, ".xyz")
	var err error
	f, err = os.Open(testFileName)
	if err != nil {
		t.Errorf("<CollExclude1> file open error: %v", err)
		return
	}
	var (
		v1,
		v2,
		v3 float64
	)

	n, err := fmt.Fscanf(f, "%f %f %f", &v1, &v2, &v3)
	if err != nil {
		t.Errorf("<CollExclude1> on file^ '%s' read occure error: %v", testFileName, err)
		return
	}
	if n != 3 {
		t.Errorf("<CollExclude1> expected 3 collums, got %d", n)
		return
	}
	if v3 != 2263.0 {
		t.Errorf("<CollExclude1> on line 1 coll 3 expected 2263.0, got %f", v3)
		return
	}
	f.Close()
}
