package upload

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gabriel-vasile/mimetype"
	"github.com/litsoftware/litmedia/internal/common/types"
	"github.com/litsoftware/litmedia/internal/ent/schema"
	"github.com/litsoftware/litmedia/pkg/random"
	"github.com/uniplaces/carbon"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strings"
	"time"
)

type UploadService struct {
	f    *os.File
	mime *mimetype.MIME

	allowedMime []string
	fileName    string
	savePath    string
	fullPath    string
	ext         string
	size        schema.MediaSize

	fileCheckSum string
	IsExists     bool
}

func (s *UploadService) Init() {
	s.allowedMime = []string{
		"image/gif",
		"image/png",
		"image/jpeg",
		"image/bmp",
		"image/webp",
		"image/x-icon",
		"image/vnd.microsoft.icon",
		"audio/midi",
		"audio/x-midi",
		"audio/aac",
		"audio/wave",
		"audio/wav",
		"audio/x-wav",
		"audio/x-pn-wav",
		"audio/webm",
		"audio/ogg",
		"audio/mpeg",
		"audio/3gpp",
		"audio/3gpp2",
		"video/3gpp2",
		"video/3gpp",
		"video/mpeg",
		"video/x-msvideo",
		"video/webm",
		"video/ogg",
		"text/plain",
		"text/html",
		"text/css",
		"text/javascript",
		"application/octet-stream",
		"application/pkcs12",
		"application/vnd.mspowerpoint",
		"application/xhtml+xml",
		"application/xml",
		"application/pdf",
	}
}

func (s *UploadService) getSize() int64 {
	sf, err := s.f.Stat()
	if err != nil {
		return 0
	}

	return sf.Size()
}

func (s *UploadService) isImage() bool {
	spew.Dump("isImage", s.mime)
	return strings.HasPrefix(s.mime.String(), "image")
}

func (s *UploadService) isVideo() bool {
	return strings.HasPrefix(s.mime.String(), "video")
}

func (s *UploadService) isAudio() bool {
	return strings.HasPrefix(s.mime.String(), "audio")
}

func (s *UploadService) isPdf() bool {
	return strings.HasSuffix(s.mime.String(), "pdf")
}

func (s *UploadService) GetMediaType() types.MediaType {
	if s.isImage() {
		return types.MediaType_Image
	}

	if s.isAudio() {
		return types.MediaType_Audio
	}

	if s.isVideo() {
		return types.MediaType_Video
	}

	if s.isPdf() {
		return types.MediaType_PDF
	}

	return types.MediaType_MediaType_UNUSED
}

func (s *UploadService) initMime() error {
	var err error
	_, err = s.f.Seek(io.SeekStart, 0)
	if err != nil {
		return err
	}

	s.mime, err = mimetype.DetectReader(s.f)
	if err != nil {
		return err
	}

	return nil
}

func (s *UploadService) generateSavePath(level types.MediaLevel) {
	t := carbon.NewCarbon(time.Now())
	var savePath string
	s.ext = s.mime.Extension()
	if level == types.MediaLevel_Sensitive {
		savePath = fmt.Sprintf("media/sensitive/%d/%s/%s", t.Year(), t.Month(), random.String(10))
	} else {
		savePath = fmt.Sprintf("media/general/%d/%s/%s", t.Year(), t.Month(), random.String(10))
	}

	s.fileName = fmt.Sprintf("%s%s", random.String(16), s.ext)
	s.savePath = savePath
	s.fullPath = s.savePath + "/" + s.fileName
}

func (s *UploadService) checkIsAllowedMime() error {
	if mimetype.EqualsAny(s.mime.String(), s.allowedMime...) {
		return nil
	} else {
		return errors.New("file type not allowed")
	}
}

func (s *UploadService) generateChecksum() error {
	h := sha256.New()
	if _, err := io.Copy(h, s.f); err != nil {
		return err
	}

	s.fileCheckSum = hex.EncodeToString(h.Sum(nil))
	return nil
}

func (s *UploadService) Parse(mf *os.File, level types.MediaLevel) error {
	s.f = mf
	return s.parse(level)
}

func (s *UploadService) ParseByPath(filePath string, level types.MediaLevel) error {
	var err error
	if filePath == "" {
		return errors.New("filePath is empty")
	}

	s.f, err = os.Open(filePath)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(s.f)

	return s.parse(level)
}

func (s *UploadService) parse(level types.MediaLevel) (err error) {
	size := schema.MediaSize{}
	err = s.initMime()
	if err != nil {
		return err
	}

	s.generateSavePath(level)

	err = s.generateChecksum()
	if err != nil {
		return err
	}

	if s.isImage() {
		_, err := s.f.Seek(io.SeekStart, 0)
		if err != nil {
			return err
		}

		if s.mime.String() == "image/jpeg" || s.mime.String() == "image/png" || s.mime.String() == "image/gif" {
			img, _, err := image.DecodeConfig(s.f)

			if err != nil {
				fmt.Println("image.DecodeConfig", err)
				return err
			}
			size.Width = img.Width
			size.Height = img.Height
			size.WhRate = float64(img.Width) / float64(img.Height)
		}
	}

	s.size = size
	return nil
}

func (s *UploadService) GetMime() *mimetype.MIME {
	return s.mime
}

func (s *UploadService) GetSavePath() string {
	return s.savePath
}

func (s *UploadService) GetFileName() string {
	return s.fileName
}

func (s *UploadService) GetFullPath() string {
	return s.fullPath
}

func (s *UploadService) GetSize() schema.MediaSize {
	return s.size
}

func (s *UploadService) GetExt() string {
	return s.ext
}

func (s *UploadService) GetCheckSum() string {
	return s.fileCheckSum
}

func NewUploadService() *UploadService {
	s := new(UploadService)
	s.Init()
	return s
}
