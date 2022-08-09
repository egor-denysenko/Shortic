package utility

import (
	"net/http"
	"testing"
)

func TestApiVersionHeaderCheck(t *testing.T) {
  type apiVersionTests struct {
        testName string
        requestHeader string
        currentApiVersion   string
        want bool 
  }

  tests := []apiVersionTests{
    {testName: "Correct Api version requested and correct client request header",requestHeader: "1.0.0", currentApiVersion: "1.0.0", want: true},
    {testName: "Mismatched client request header",requestHeader: "1.0.0", currentApiVersion: "1.0.1", want: false},
    {testName: "Mismatched server api version header",requestHeader: "0.9.0", currentApiVersion: "1.0.0", want: false},
  }

  for _, testCase := range tests {
    t.Run(testCase.testName, func(t *testing.T){
      req, _ := http.NewRequest("GET", "http://localhost:7878/health/", nil)
      req.Header.Add("Accept-Version",testCase.requestHeader)
      outcome := ApiVersionHeaderCheck(req,testCase.currentApiVersion)
      if testCase.want != outcome{
        t.Errorf("Api version non correct")
      }
    })
  } 

}
