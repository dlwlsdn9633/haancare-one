package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DeliveryRecord struct {
	InvNo          string `json:"invNo"`          // 운송장 번호
	GdsNm          string `json:"gdsNm"`          // 상품명
	AcperNm        string `json:"acperNm"`        // 수령인 이름
	AcperTel       string `json:"acperTel"`       // 수령인 전화번호
	AcperRdnmBadr  string `json:"acperRdnmBadr"`  // 수령인 도로명 주소
	SnperNm        string `json:"snperNm"`        // 발송인 이름
	DlvYmd         string `json:"dlvYmd"`         // 배송일자
	DlvEmpNm       string `json:"dlvEmpNm"`       // 배송기사 이름
	DlvEmpScanCpno string `json:"dlvEmpScanCpno"` // 배송기사 연락처
	TotalCnt       int    `json:"totalCnt"`       // 전체 데이터 개수
	PickYmd        string `json:"pickYmd"`        // 집하일자
}

type SearchFilter struct {
	SrchPickYmd      string      `json:"srchPickYmd"`
	SrchPickYmdStrt  string      `json:"srchPickYmdStrt"`
	SrchPickYmdEnd   string      `json:"srchPickYmdEnd"`
	CboSrchCustSctCd string      `json:"cboSrchCustSctCd"`
	SrchCustCd       string      `json:"srchCustCd"`
	SrchCustNm       string      `json:"srchCustNm"`
	CboSrchWkSctCd   string      `json:"cboSrchWkSctCd"`
	JobCustCd        interface{} `json:"jobCustCd"`
	TabIdx           string      `json:"tabIdx"`
	RowCount         int         `json:"rowCount"`
	DispCount        int         `json:"dispCount"`
	PickYmd          string      `json:"pickYmd"`
	ColNm            string      `json:"colNm"`
	UstRtgSctCd      string      `json:"ustRtgSctCd"`
	FstmIstrYmd      string      `json:"fstmIstrYmd"`
	Status           string      `json:"_STATUS_"`
}

func main() {
	InitLogger()
	e := echo.New()
	e.HideBanner = true

	// 1. URL 설정 (제공해주신 필터 포함 전체 경로)
	// targetURL := "https://pid.alps.llogis.com:18210/pid/ftr/hdarvmgr/daily/dtls?filter=%7B%22srchPickYmd%22%3A%22%22%2C%22srchPickYmdStrt%22%3A%2220260201%22%2C%22srchPickYmdEnd%22%3A%2220260227%22%2C%22cboSrchCustSctCd%22%3A%2210%22%2C%22srchCustCd%22%3A%22982718%22%2C%22srchCustNm%22%3A%22%EC%A3%BC%EC%8B%9D%ED%9A%8C%EC%82%AC%20%ED%95%9C%EC%BC%80%EC%96%B4%22%2C%22cboSrchWkSctCd%22%3A%22%22%2C%22jobCustCd%22%3Anull%2C%22tabIdx%22%3A%22%22%2C%22rowCount%22%3A0%2C%22dispCount%22%3A1000%2C%22pickYmd%22%3A%22%22%2C%22colNm%22%3A%22totCnt%22%2C%22ustRtgSctCd%22%3A%22%22%2C%22fstmIstrYmd%22%3A%22%22%2C%22_STATUS_%22%3A%22U%22%7D"
	// filterData := SearchFilter{
	// 	SrchPickYmd:      "",
	// 	SrchPickYmdStrt:  "20260201",
	// 	SrchPickYmdEnd:   "20260227",
	// 	CboSrchCustSctCd: "10",
	// 	SrchCustCd:       "982718",
	// 	SrchCustNm:       "주식회사 한케어",
	// 	CboSrchWkSctCd:   "",
	// 	JobCustCd:        nil, // null 처리
	// 	TabIdx:           "",
	// 	RowCount:         0,
	// 	DispCount:        1000,
	// 	PickYmd:          "",
	// 	ColNm:            "totCnt",
	// 	UstRtgSctCd:      "",
	// 	FstmIstrYmd:      "",
	// 	Status:           "U",
	// }
	// filterJSON, err := json.Marshal(filterData)
	// if err != nil {
	// 	log.Fatal(fmt.Sprint("failed to marshal data"))
	// 	return
	// }

	// baseURL := "https://pid.alps.llogis.com:18210/pid/ftr/hdarvmgr/daily/dtls"
	// params := url.Values{}
	// params.Add("filter", string(filterJSON)) // JSON 문자열을 filter 파라미터에 담기
	// targetURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// // 2. HTTP 클라이언트 생성
	// client := &http.Client{}

	// // 3. GET 요청 객체 생성
	// req, err := http.NewRequest("GET", targetURL, nil)
	// if err != nil {
	// 	log.Fatalf("Error creating request: %v", err)
	// }

	// // 4. 헤더 설정
	// authToken := "eyJ0eXAiOiJKV1QiLCJyZWdEYXRlIjoxNzcyMzc4MzI4NjA4LCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3NzIzODkxMjgsImFjY2Vzc1VzZXIiOnsidXNlcklkIjoiOTgyNzE4IiwidXNlck5hbWUiOiLso7zsi53tmozsgqwg7ZWc7LyA7Ja0IiwiZGVwdElkIjpudWxsLCJkZXB0TmFtZSI6bnVsbCwicm9sZXMiOm51bGwsInBlcm1pc3Npb25zIjpudWxsLCJhdXRoZW50aWNhdGlvblRpbWUiOjE3NzIzNzY1ODAzNDMsImFjY2Vzc1RpbWUiOjE3NzIzNzgzMjg2MDgsInN5c3RtSWQiOiIzIiwibWFjQWRkcmVzcyI6Im5vcm1hbC1icm93c2VyIiwibGdpbklwIjoiMjIxLjE1My4xMzQuMTUwIn19.OFb4gBJxLH5eN9UqJo1inMsItRBEpOtFXnsPkKZsl74"

	// req.Header.Set("Authorization", authToken)
	// req.Header.Set("Referer", "https://partner.alps.llogis.com/")
	// // 브라우저 요청처럼 보이기 위해 User-Agent를 추가하는 것이 좋습니다.
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	// // 5. 요청 실행
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// defer resp.Body.Close()

	// // 6. 결과 출력
	// fmt.Println("Response Status:", resp.Status)

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response: %v", err)
	// }

	// fmt.Println("Response Body:")
	// fmt.Println(string(body))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	StartCronJobs()
	InitRouter(e)
	e.Logger.Fatal(e.Start(":8080"))
}
