package curl
import ("encoding/json"
        "io/ioutil"
        "net/http"
        //"bytes"
        "logs"
        "net/url"
        "strings"
       )

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}
func handleErr() {
        if x := recover(); x != nil {
                logs.Logger.Error("curl failed: %#v", x)
        }
}
func  CurlPost(addr string, buf interface{}) {
        client := &http.Client{}
        defer handleErr()
        // 将需要上传的JSON转为Byte
        v, _ := json.Marshal(buf)
        params:=string(v);
        values:=url.Values{};
        values.Set("params",params);
        // 上传JSON数据
        //postDataBytes := []byte(values.Encode())
        //req, e := http.NewRequest("POST", addr,bytes.NewReader(postDataBytes))
        req, e := http.NewRequest("POST", addr,strings.NewReader(values.Encode()))
        if e != nil {
            // 提交异常,返回错误
            logs.Logger.Infof("new request failed %#v",req);
        }
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded;param=value")
        // Body Type
        // 完成后断开连接
        req.Header.Set("Connection", "close")
        res,err := client.Do(req)
        checkErr(err);
        body,err := ioutil.ReadAll(res.Body)
        checkErr(err);
        bodystr := string(body);
        if res.StatusCode == 200 {
                logs.Logger.Infof("curl sucess %#v:%s",string(v),bodystr);
        }else{
                logs.Logger.Errorf("curl failed %#v:%s",string(v),bodystr);
        }
}
