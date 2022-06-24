package function

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/gocolly/colly/queue"
	"nhentaiAgnet/structType"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const webSiteUrl = "https://nhentai.net"

var comicName string
var comicPictureCount = 0
var nowDownloadIndex = 0

func Download(comicDownload structType.ComicDownload) (err error) {
	downloadUrl := fmt.Sprintf("%s/g/%d/", webSiteUrl, comicDownload.ComicId)
	localComicDir := comicDownload.OutputDir
	if localComicDir == "" {
		localComicDir = "./"
	}

	htmlCollector := colly.NewCollector()
	imageCollector := colly.NewCollector(colly.AllowURLRevisit())
	imageCollector.SetRequestTimeout(time.Second * 15)
	imageDownloadQueue, _ := queue.New(comicDownload.ThreadCount, &queue.InMemoryQueueStorage{MaxSize: 10000})
	retryImageDownloadList := make([]string, 0)
	if comicDownload.ProxyUrl != "" {
		//rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
		rp, err := proxy.RoundRobinProxySwitcher(comicDownload.ProxyUrl)
		if err != nil {
			return errors.New("proxy url is not correct")
		}
		htmlCollector.SetProxyFunc(rp)
		imageCollector.SetProxyFunc(rp)
	}

	htmlCollector.OnHTML("#info>h1.title", func(e *colly.HTMLElement) {
		comicName = e.Text
	})
	htmlCollector.OnHTML("#thumbnail-container .thumb-container img", func(e *colly.HTMLElement) {
		img := e.Attr("data-src")
		img = strings.Replace(img, "https://t", "https://i", 1)
		img = strings.Replace(img, "t.", ".", 1)
		imageDownloadQueue.AddURL(img)
		comicPictureCount++
	})

	err = htmlCollector.Visit(downloadUrl)
	if err != nil || comicName == "" || imageDownloadQueue.IsEmpty() {
		return errors.New("try download comic failed,please check the comic id is correct")
	}

	if comicDownload.DirNameUseComicId {
		localComicDir = path.Join(localComicDir, strconv.Itoa(comicDownload.ComicId))
	} else {
		localComicDir = path.Join(localComicDir, comicName)
	}
	err = os.MkdirAll(localComicDir, 0755)
	if err != nil {
		return errors.New("create comic dir failed")
	}
	imageCollector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("error downloading image %s\n", r.Request.URL)
		retryImageDownloadList = append(retryImageDownloadList, r.Request.URL.String())
	})
	imageCollector.OnResponse(func(r *colly.Response) {
		nowDownloadIndex++
		fmt.Printf("[%d/%d]%s\n", nowDownloadIndex, comicPictureCount, localComicDir+"/"+path.Base(r.Request.URL.Path))
		err = r.Save(localComicDir + "/" + path.Base(r.Request.URL.Path))
		if err != nil {
			fmt.Println("save image failed")
		}
	})
	retryCount := 0
	for {
		imageDownloadQueue.Run(imageCollector)
		if len(retryImageDownloadList) == 0 || !comicDownload.Retry || retryCount >= 3 {
			break
		}
		imageDownloadQueue.Threads = 1
		for _, url := range retryImageDownloadList {
			imageDownloadQueue.AddURL(url)
			fmt.Println("retry download image:", url)
		}
		retryImageDownloadList = make([]string, 0)
		retryCount++
	}
	if len(retryImageDownloadList) == 0 {
		if comicDownload.Zip {
			var zipFileName string
			if comicDownload.DirNameUseComicId {
				zipFileName = strconv.Itoa(comicDownload.ComicId)
			} else {
				zipFileName = comicName
			}
			fmt.Println("Zipping files....")
			Zip(localComicDir, zipFileName+".zip")
			os.RemoveAll(localComicDir)
		}
		fmt.Println("Have a nice time :)")
	} else {
		fmt.Println("Download completed,But some comic pictures is missing...:(")
	}
	return nil
}
