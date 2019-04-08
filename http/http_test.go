package http

import (
	"fmt"
	"testing"
)

func TestUtf8ToGbk(t *testing.T) {
	data := `content=<?xml version="1.0" encoding="GBK"?><DATA><REQUEST><NCCUPNTF><MSGTYP>NCCUPNTF</MSGTYP><MSGNBR>201904041513980411</MSGNBR><PAYTYP>1</PAYTYP><REQNBR></REQNBR><MCHNBR>M000003906</MCHNBR><SEQNBR>00000000000000001258</SEQNBR><SUBSEQ>0000000001</SUBSEQ><CCYNBR>10</CCYNBR><TRSAMT>000000000000001</TRSAMT><ENDAMT></ENDAMT><BBKNBR></BBKNBR><PAYACC></PAYACC><ACCNAM></ACCNAM><RCVACC>120914731110888</RCVACC><RCVNAM>佛山众塑联供应链服务有限公司</RCVNAM><REFMCH>M000003906</REFMCH><REFORD>20190404094258199426</REFORD><SUBORD></SUBORD><PAYNBR>20190404094258199426</PAYNBR><YURREF></YURREF><ENDDAT></ENDDAT><RTNFLG>P</RTNFLG><RTNDSP>订单请求已发送</RTNDSP><RSV30Z>B000000004</RSV30Z></NCCUPNTF></REQUEST></DATA>&signature=Hd4RjkMXFUAZFY9I/2EWfRE/bFWQ8rXaEs7t9V1FAKNfmhJHX1MeDQ+m4NO+R1YrHnrSdPi0q0zpzCbo/gyLzKPhoBqz6Wm1AAxnxvC+VIs8peOrhpRB7hMfWWNaqf2IkD7Ugj4pyL4gdKeB4qtGYZca8QBFcYO++idcH1taXmU=`
	by1, err := Utf8ToGbk([]byte(data))
	if err != nil {
		t.Fatalf("转gbk失败:%s", err.Error())
	}
	by2, err := GbkToUtf8(by1)
	if err != nil {
		t.Fatalf("gbk转utf8失败:%s", err.Error())
	}
	fmt.Printf("转换后:\n%s\n", string(by2))
	if data != string(by2) {
		t.Fatalf("测试gbk转换函数失败")
	}
}
