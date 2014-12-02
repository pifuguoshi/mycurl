package curl
import ("encoding/json"
        "io/ioutil"
        "net/http"
        "bytes"
        "logs"
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
        // 上传JSON数据
        req, e := http.NewRequest("POST", addr, bytes.NewReader(v))
        if e != nil {
            // 提交异常,返回错误
            logs.Logger.Infof("new request failed %#v",req);
        }
        // Body Type
        req.Header.Set("Content-Type", "application/json")
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
