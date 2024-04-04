# Web Scrapper (Go + Colly) by qRe0

## Used technologies
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Colly](https://img.shields.io/badge/Colly-0C8B7C?style=for-the-badge)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

## Overview 
This Go language program allows you to retrieve elements from a browser page and save them to a database, then download the images in `.jpg` format

## Project structure
1. `cmd/main.go` - Program source code
2. `internal/database/scrapped_data.db` - Database to store scrapped data
3. `internal/db_processing/writeDataToDatabase` - .go file to act with database and write data in it
4. `internal/db_structure_struct.go` - .go file with structure to handle data from target website
5. `internal/img_processing/processSaving.go` - .go file to save image from link to `.jpg` file
6. `internal/img_processing/saveImgToFile.go` - .go file to act with DB field `img` and call `processSaving` function
```shell
web_scrapper
├───cmd
├───img
└───internal
    ├───database
    ├───db_processing
    ├───db_structure
    └───img_processing
```

## Additional notes
* [Colly framework](https://github.com/gocolly/colly)<br>
* [Target site](https://scrapeme.live/shop/)