# createCSVFile

## createCSV.goについて
読み込んだCSVファイル内の{keyword}文字列を、コマンドライン引数を反映した結果をもとに置換して、指定ファイルに保存します。
置換ルールは固定なので、下記以外の置換が必要な場合は、修正が必要です。
|キー|置換ルール|補足|
| --- | --- | --- |
|title|第2引数|イベントタイトル。空白が含まれている場合は、""で囲む。例）"Cybozu Frontend Monthly"|
|connpass|第3引数|イベントURL|
|eventdate|第4引数|開催日付。年4桁/月2桁/日2桁形式|
|preparelimit|eventdateの4日前（eventdateをもとに算出）||
|postproc|eventdateの5日後（eventdateをもとに算出）||

## 使い方
- コマンドライン引数で、置換用の文字列と、保存先のファイル名を指定します。
- プログラムと同じパスに、置換元ファイルrecords.csvが必要です。
- go run createCSV.go イベントタイトル（空白が含まれる場合は、""で囲む） エベントURL 開催日付（yyyy/MM/dd形式）置換結果の保存先ファイル名