package engine

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path"
// 	"strconv"
// 	"time"

// 	"github.com/adnsio/vmkit/assets"
// 	"github.com/adnsio/vmkit/util"
// )

// type ImagesJSON struct {
// 	Images []*Image
// }

// type Image struct {
// 	Name        string
// 	Description string
// 	Sources     []*ImageSource
// 	SSH         *ImageSSH
// 	Features    *ImageFeatures

// 	dir    string
// 	engine *Engine
// }

// type ImageFeatures struct {
// 	CloudInit bool
// }

// type ImageSSH struct {
// 	Username string
// }

// type ImageSource struct {
// 	URL      string
// 	Checksum string
// 	Arch     string
// }

// func (i *Image) IsDownloaded() bool {
// 	if _, err := os.Stat(i.Path()); err != nil {
// 		return false
// 	}

// 	return true
// }

// func (i *Image) Download(progressFunc func(totalBytes int64, downloadedBytes int64)) error {
// 	var src *ImageSource
// 	for _, imgSrc := range i.Sources {
// 		if imgSrc.Arch == i.engine.arch {
// 			src = imgSrc
// 			break
// 		}
// 	}

// 	if src == nil {
// 		return ErrImageSourceNotFound
// 	}

// 	headRes, err := http.Head(src.URL)
// 	if err != nil {
// 		return err
// 	}
// 	defer headRes.Body.Close()

// 	contentLength, err := strconv.ParseInt(headRes.Header.Get("content-length"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	getRes, err := http.Get(src.URL)
// 	if err != nil {
// 		return err
// 	}
// 	defer getRes.Body.Close()

// 	if err := util.MkdirAllIfNotExist(i.dir); err != nil {
// 		return err
// 	}

// 	tmpPath := fmt.Sprintf("%s.tmp", i.Path())
// 	imgFile, err := os.Create(tmpPath)
// 	if err != nil {
// 		if err := os.Remove(tmpPath); err != nil {
// 			return err
// 		}

// 		return err
// 	}
// 	defer imgFile.Close()

// 	progressReader := util.NewProgressReader(getRes.Body)

// 	go func() {
// 		for {
// 			if progressReader.Progress >= contentLength {
// 				return
// 			}

// 			progressFunc(contentLength, progressReader.Progress)

// 			time.Sleep(1)
// 		}
// 	}()

// 	if _, err := io.Copy(imgFile, progressReader); err != nil {
// 		if err := os.Remove(tmpPath); err != nil {
// 			return err
// 		}

// 		return err
// 	}

// 	if err := os.Rename(tmpPath, i.Path()); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (i *Image) Path() string {
// 	return path.Join(i.dir, fmt.Sprintf("%s.img", i.engine.arch))
// }

// func (e *Engine) reloadImages() error {
// 	e.Images = map[string]*Image{}

// 	var imgsJSON ImagesJSON
// 	if err := json.Unmarshal(assets.ImagesJSON, &imgsJSON); err != nil {
// 		return err
// 	}

// 	for _, img := range imgsJSON.Images {
// 		e.Images[img.Name] = img
// 		e.Images[img.Name].engine = e
// 		e.Images[img.Name].dir = path.Join(e.diskDir, img.Name)
// 	}

// 	return nil
// }
