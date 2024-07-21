package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
	"log"
	"errors"
)

func main() {

	// get Args parameters
	title, connpass, eventdate, preparelimit, postproc, resultpath, err := getArgs()
	if (err != nil){
		fmt.Println(err)
		return
	}
	fmt.Printf("\n=== パラメータ ===\ntitle:%s\nconnpass:%s\ndate:%s\nprepare:%s\npostProc:%s\nresult:%s\n",
	 title,connpass,eventdate,preparelimit,postproc, resultpath)

	// get source csv data
	str := getSourceCSV()

	// replace parameters
	str = setParameters(str, title, connpass, eventdate, preparelimit, postproc)

	// csv書き出し
	createNewTaskCSV(str, resultpath)
}

func getArgs()(title, connpass, eventdate, preparelimit, postproc, resultpath string, err error){
	// get args
	if(len(os.Args) != 5){
		err = errors.New("パラメータに誤りがあります。\nUsage:\ngo run createCSV.go [event title] [connpass url] [event date] [result path]")
		return title, connpass, eventdate, preparelimit, postproc, resultpath, err
	}

	// title
	title = os.Args[1]
	// url
	connpass = os.Args[2]
	// eventdate
	date := strings.Split(os.Args[3], "/")
	year := date[0]
	month := date[1]
	day :=date[2]

	eventdate = os.Args[3]

	// result path(output)
	resultpath = os.Args[4]

	// limit date for prepare
	var preDate int
	preDate, _ = strconv.Atoi(day)
	preDate -=4
	preparelimit = fmt.Sprintf("%s/%s/%d", year, month, preDate)
	var postDate, _ = time.Parse("2006/01/02", os.Args[3])
	y, m, d := postDate.AddDate(0,0,5).Date()
	postproc = fmt.Sprintf("%d/%d/%d", y, m, d)

	return title, connpass, eventdate, preparelimit,postproc,resultpath,err
}

// 雛形のCSV読み込み
func getSourceCSV() (str string){
	// open file
	f, err := os.Open("./records.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read file
	buf := make([]byte, 1024)
	count, err := f.Read(buf) // reader to buf
	if err != nil {
		log.Fatal(err)
	}
	str = string(buf[:count])
	return str
}

func setParameters(str, title, connpass, eventdate, preparelimit, postproc string) (string){
	// replace
	// title置換
	str = strings.Replace(str, "{title}", title, -1)
	// connpass URL置換
	str = strings.Replace(str, "{connpass}", connpass, -1)
	// 開催日付
	str = strings.Replace(str, "{eventdate}", eventdate, -1)
	// 準備期限
	str = strings.Replace(str, "{preparelimit}", preparelimit, -1)
	// 開催後のCybozuTech動画掲載目安日
	str = strings.Replace(str, "{postproc}", postproc, -1)

	// 結果を返却
	return str
}

func createNewTaskCSV(str, resultpath string){
	outputpath := resultpath
	// 出力パスに拡張子があれば、そのまま。拡張子がなければ、".csv"を付加する
	if( !strings.HasSuffix(resultpath, ".csv") ){
		outputpath = fmt.Sprintf("./%s.csv", resultpath)
	}

	// 書き込み可としてファイルを開く
	f, err := os.Create(outputpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// CSV書き出し
	data := []byte(str)
	count, err := f.Write(data) // buf to writer
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("write %d bytes\n", count)
}
