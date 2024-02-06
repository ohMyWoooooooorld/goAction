package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//读取并取消封送源数据文件，获取信息源。
//一个 Feed 类型的指针切片 ([]*Feed)。这意味着这个函数可以返回一个 Feed 结构体的指针组成的切片。
//这种设计允许在切片中的每个元素（即每个 Feed）被修改时，这些更改会反映在原始数据结构中
func RetrieveFeeds() ([]*Feed ,error) {
	//打开文件
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	//使用defer关键字确保在函数返回之前文件被关闭。这是一个好的做法，因为它确保无论函数是否因为错误而提前返回，文件都会被正确关闭。
	defer file.Close()

	//Decoder 解码器
	//这部分创建了一个新的JSON解码器，该解码器从提供的file（一个已打开的文件）读取数据。
	//这部分使用解码器来解析JSON数据，并将解析的结果存储在feeds变量中。
	//这里，&feeds是一个指向feeds变量的指针，这意味着我们正在将解析的结果直接存储在feeds变量中，而不是存储一个副本
	//这个变量存储了解析过程中可能出现的任何错误。如果解析成功，err将为nil；如果出现错误，err将包含错误信息。
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds,err
}