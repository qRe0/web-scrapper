package img_processing

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	dbs "web-scrapper/internal/db-structure"
)

func ProcessSaving(data dbs.ScrappedData) {
	filename := fmt.Sprintf("img/%s.jpg", data.Name)

	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("File %s already exists. Skipping...", filename)
		return
	} else if !os.IsNotExist(err) {
		log.Fatal(err)
	}

	response, err := http.Get(data.Img)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot close response body: %v\n", err)
	}
	// defer response.Body.Close()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Cannot close response body: %v\n", err)
		}
	}(response.Body)

	imgDir := "img"
	_, err = os.Stat(imgDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(imgDir, 0755)
		if err != nil {
			log.Fatalf("Cannot create directory: %v\n", err)
		}
	}

	imgFile, err := os.Create(fmt.Sprintf("img/%s.jpg", data.Name))
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot close response body: %v\n", err)
	}

	// defer imgFile.Close()
	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {
			log.Fatalf("Cannot close imgFile: %v\n", err)
		}
	}(imgFile)

	_, err = io.Copy(imgFile, response.Body)
	if err != nil {
		log.Fatalf("Cannot copy data: %v\n", err)
	}

	fmt.Printf("Saved %s.jpg\n", data.Name)
}
