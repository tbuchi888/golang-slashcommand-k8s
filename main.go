package main

import (
        "bytes"
        "encoding/json"
        "net/http"
        "os"

        "github.com/favclip/ucon"
)

func main() {
        ucon.Orthodox()

        ucon.HandleFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
                if err := r.ParseForm(); err != nil {
                        w.Write([]byte(err.Error()))
                        return
                }

                w.Header().Set("Content-Type", "application/json")

                // Check method and token
                if r.Method != "POST" {
                    w.WriteHeader(http.StatusBadRequest)
                    return
                }
                if r.PostFormValue("token") != os.Getenv("GOBOTTOKEN") {
                    w.WriteHeader(http.StatusUnauthorized)
                    return
                }
                w.WriteHeader(http.StatusOK)

                data := map[string]string{
                        "response_type": "in_channel",
                        "username": "gobot",
                        "icon_url": "https://golang.org/doc/gopher/gophercolor.png",
                }

                // Switch by text
                t := r.PostFormValue("text")
                switch t {
/*--------------------------------------------------------------------------------
                        ここへ追記
                        case "条件となる文字列を指定" :
　　　　　　　　　　　　　　　　やらせたい処理などを記載しMD記法で応答メッセージをセット。
                                data["text"] = "## 応答メッセージ"
--------------------------------------------------------------------------------*/
                        case "" :
                                data["text"] = "## メッセージが設定されていません。"
                        case "今日の天気は？" :
                                data["text"] = "## わかりません。。。"
                        default:
                                data["text"] = "## " + t + "gobotより自動返信しています。"
                }

                var buf bytes.Buffer
                json.NewEncoder(&buf).Encode(data)
                w.Write(buf.Bytes())
        })

        ucon.ListenAndServe(":3000")
        ucon.DefaultMux.Prepare()
        http.Handle("/", ucon.DefaultMux)
}
