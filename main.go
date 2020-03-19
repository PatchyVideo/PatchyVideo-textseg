package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-ego/gse"
)

type segRequest struct {
	Text string
}

type segResponse struct {
	Words []string
}

var (
	seger gse.Segmenter
)

func main() {
	fmt.Println("[+] TextSeg start")
	seger.LoadDict("touhou.txt,touhou2.txt,networds.txt,chs.txt,cht.txt,jpn.txt")
	fmt.Println("[+] Done loading dict")
	http.HandleFunc("/s/", segTextSearch)
	http.HandleFunc("/i/", segTextIndex)
	http.HandleFunc("/d/", segTextDisplay)
	http.HandleFunc("/t/", segTextTouhou)
	fmt.Println("[+] Serving ...")
	http.ListenAndServe("0.0.0.0:5005", nil)
}

func unique(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func segTextDisplay(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBodyStr := strings.ToLower(string(requestBody))

	//splitText := strings.FieldsFunc(requestBodyStr, split)
	//for i := 0; i < len(splitText); i++ {
	//tb := []byte(splitText[i])
	tb := []byte(requestBodyStr)
	segments := seger.Segment(tb)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	// The search mode is mainly used to provide search engines
	// with as many keywords as possible
	//fmt.Println(gse.ToString(segments, true))
	ret := [][]string{}
	for i := 0; i < len(segments); i++ {
		seg := segments[i]
		//wordType := seg.Token().Pos()
		//if wordType == "TH名詞" {
		ele := []string{seg.Token().Text(), seg.Token().Pos()}
		ret = append(ret, ele)
		//fmt.Printf("%s=>%s\n", )
		//}
	}
	//}

	js, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func segTextTouhou(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBodyStr := strings.ToLower(string(requestBody))

	//splitText := strings.FieldsFunc(requestBodyStr, split)
	//for i := 0; i < len(splitText); i++ {
	//tb := []byte(splitText[i])
	tb := []byte(requestBodyStr)
	segments := seger.Segment(tb)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	// The search mode is mainly used to provide search engines
	// with as many keywords as possible
	//fmt.Println(gse.ToString(segments, true))
	ret := []string{}
	for i := 0; i < len(segments); i++ {
		seg := segments[i]
		wordType := seg.Token().Pos()
		if wordType == "TH名詞" {
			ret = append(ret, seg.Token().Text())
		}
	}

	uniqueRet := unique(ret)
	//}

	js, err := json.Marshal(uniqueRet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func segTextSearch(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBodyStr := strings.ToLower(string(requestBody))
	segResult := make([]string, 0, len(requestBody))
	splitText := strings.FieldsFunc(requestBodyStr, split)
	for i := 0; i < len(splitText); i++ {
		tmpSegResult := seger.Cut(splitText[i], true)
		segResult = append(segResult, tmpSegResult...)
	}
	resp := segResponse{segResult}
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func segTextIndex(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBodyStr := strings.ToLower(string(requestBody))
	segResult := make([]string, 0, len(requestBody))
	splitText := strings.FieldsFunc(requestBodyStr, split)
	for i := 0; i < len(splitText); i++ {
		tmpSegResult := seger.CutSearch(splitText[i], true)
		segResult = append(segResult, tmpSegResult...)
	}

	allWords := unique(segResult)
	resp := segResponse{allWords}
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func split(r rune) bool {
	return r == ':' ||
		r == '.' ||
		r == '\n' ||
		r == '\r' ||
		r == '[' ||
		r == ']' ||
		r == ' ' ||
		r == '\t' ||
		r == '\v' ||
		r == '\f' ||
		r == '{' ||
		r == '}' ||
		r == '-' ||
		r == '_' ||
		r == '=' ||
		r == '+' ||
		r == '`' ||
		r == '~' ||
		r == '!' ||
		r == '@' ||
		r == '#' ||
		r == '$' ||
		r == '%' ||
		r == '^' ||
		r == '&' ||
		r == '*' ||
		r == '(' ||
		r == ')' ||
		r == ';' ||
		r == '\'' ||
		r == '"' ||
		r == ',' ||
		r == '<' ||
		r == '>' ||
		r == '/' ||
		r == '?' ||
		r == '\\' ||
		r == '|' ||
		r == '－' ||
		r == '＞' ||
		r == '＜' ||
		r == '。' ||
		r == '，' ||
		r == '《' ||
		r == '》' ||
		r == '【' ||
		r == '】' ||
		r == '　' ||
		r == '？' ||
		r == '！' ||
		r == '￥' ||
		r == '…' ||
		r == '（' ||
		r == '）' ||
		r == '、' ||
		r == '：' ||
		r == '；' ||
		r == '·' ||
		r == '「' ||
		r == '」' ||
		r == '『' ||
		r == '』' ||
		r == '〔' ||
		r == '〕' ||
		r == '［' ||
		r == '］' ||
		r == '｛' ||
		r == '｝' ||
		r == '｟' ||
		r == '｠' ||
		r == '〉' ||
		r == '〈' ||
		r == '〖' ||
		r == '〗' ||
		r == '〘' ||
		r == '〙' ||
		r == '〚' ||
		r == '〛' ||
		r == '゠' ||
		r == '＝' ||
		r == '‥' ||
		r == '※' ||
		r == '＊' ||
		r == '〽' ||
		r == '〓' ||
		r == '〇' ||
		r == '＂' ||
		r == '“' ||
		r == '”' ||
		r == '‘' ||
		r == '’' ||
		r == '＃' ||
		r == '＄' ||
		r == '％' ||
		r == '＆' ||
		r == '＇' ||
		r == '＋' ||
		r == '．' ||
		r == '／' ||
		r == '＠' ||
		r == '＼' ||
		r == '＾' ||
		r == '＿' ||
		r == '｀' ||
		r == '｜' ||
		r == '～' ||
		r == '｡' ||
		r == '｢' ||
		r == '｣' ||
		r == '､' ||
		r == '･' ||
		r == 'ｰ' ||
		r == 'ﾟ' ||
		r == '￠' ||
		r == '￡' ||
		r == '￢' ||
		r == '￣' ||
		r == '￤' ||
		r == '￨' ||
		r == '￩' ||
		r == '￪' ||
		r == '￫' ||
		r == '￬' ||
		r == '￭' ||
		r == '￮' ||
		r == '・' ||
		r == '◊' ||
		r == '→' ||
		r == '←' ||
		r == '↑' ||
		r == '↓' ||
		r == '↔'
}
