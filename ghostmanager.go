package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

var AUTH string

type GhostManager struct {
	Configs Config
	Cookie  []*http.Cookie
	BaseUrl string
}

func NewGhostManager() *GhostManager {
	homeDir, _ := os.UserHomeDir()
	AUTH = filepath.Join(homeDir, ".tgug")
	fmt.Println(AUTH)
	g := new(GhostManager)
	g.Configs = loadConfig()
	g.BaseUrl = g.Configs.Domain + "/ghost/api/v3/admin"
	g.handleAuth()

	return g
}

type Config struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func newConfig() Config {
	var config Config

	fmt.Println("Ghost login information is needed to be configured. Type your Ghost blog information.")

	fmt.Print("Ghost blog domain (https://example.com) : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	config.Domain = text

	fmt.Print("Username (example@exmple.com) : ")
	reader = bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	config.Username = text

	fmt.Print("Password : ")
	reader = bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	config.Password = text

	fmt.Println("\nReview your information :")
	fmt.Println(config)

	byteArray, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(filepath.Join(AUTH, "auth.json"), byteArray, 0755)

	return config
}

func loadConfig() Config {
	var config Config
	jsonFile, err := os.Open(filepath.Join(AUTH, "auth.json"))
	defer jsonFile.Close()
	if err != nil {
		config = newConfig()
	} else {
		byteArray, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal(byteArray, &config)
	}

	return config
}

func (g *GhostManager) handleAuth() {
	payload := map[string]string{
		"username": g.Configs.Username,
		"password": g.Configs.Password,
	}
	byteArray, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	authentication := bytes.NewBuffer(byteArray)

	client := http.Client{}
	req, _ := http.NewRequest("POST", g.BaseUrl+"/session/", authentication)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	g.Cookie = resp.Cookies()
}

func (g *GhostManager) createMultipartFormData(path string) (*multipart.Writer, io.Reader) {

	const paramName = "file"

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
	}
	if err := file.Close(); err != nil {
		log.Println(err)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "image/jpeg")
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		escapeQuotes(paramName), escapeQuotes(fi.Name())))
	part, err := writer.CreatePart(h)
	if err != nil {
		log.Println(err)
	}
	_, err = part.Write(fileContents)
	if err != nil {
		log.Println(err)
	}

	params := map[string]string{
		"purpose": "image",
		"ref":     path,
	}
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		log.Println(err)
	}

	return writer, body
}

func (g *GhostManager) UploadImage(path string) (imageURL string, err error) {
	writer, body := g.createMultipartFormData(path)
	var uri = g.BaseUrl + "/images/upload"

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("cookie", g.Cookie[0].String())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("cannot close body reader: %v\n", err)
		}
	}()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBody := string(content[:])
	if response.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("wrong http code: %d (%q)",
			response.StatusCode, responseBody)
	}

	var images ImageResponse
	if err := json.Unmarshal(content[:], &images); err != nil {
		return "", err
	}

	return images.Images[0].URL, nil
}

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
