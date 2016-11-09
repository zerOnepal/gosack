package main

import (
	b64 "encoding/base64"
	json "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type GithubContentResponse struct {
	Content     string `json:"content"`
	DownloadURL string `json:"download_url"`
	Encoding    string `json:"encoding"`
	Name        string `json:"name"`
}

type QA struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func main() {

	url := "https://api.github.com/repos/foss-np/np-l10n-glossary/contents/en2ne/fun.tra"
	qa := &QA{Q: "ram", A: "hari"}
	qa.Downloadfromurl(url)

	b, err := json.Marshal(qa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func (qa *QA) Downloadfromurl(url string) {

	tmp_path := "/tmp/"
	tokens := strings.Split(url, "/")
	fileName := tmp_path + tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// decode and print
	githubcontentresponse := GithubContentResponse{}
	json.Unmarshal(raw, &githubcontentresponse)
	contentsInside := githubcontentresponse.Content

	sDec, _ := b64.StdEncoding.DecodeString(contentsInside)

	words := strings.Split(string(sDec), "\n")
	// fmt.Println(words[10])
	// fmt.Println(words[11])
	// fmt.Println(words[0])
	// fmt.Println(len(words))

	// sellout random
	s1 := rand.NewSource(time.Now().UnixNano())
	randomNum := rand.New(s1).Intn(len(words))

	// return random word
	randWord := words[randomNum]
	questionAnswers := strings.Split(randWord, ";")
	qa.Q = questionAnswers[0]
	qa.A = questionAnswers[1]

	// // loop over the array
	// for _, word := range words {
	// 	questionAnswers := strings.Split(word, ";")
	// 	question := questionAnswers[0]
	// 	answer := questionAnswers[1]

	// 	if question == "angry birds" {
	// 		fmt.Println("answer", answer)
	// 		q = question
	// 		a = answer

	// 		fmt.Println(qa)
	// 		// return
	// 	}
	// }

}
