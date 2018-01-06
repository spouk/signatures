package signatures

import (
	"os"
	"errors"
	"log"
)

const (
	randomSym = 0x3F
	prefix = "[signatures]"
)

var (
	errorWrongType = errors.New("undefined data type\n")


	rpm = Fileobj{
		Ext:       "RPM", Offset: 0, Hexsig: []byte{0xed, 0xab, 0xee, 0xdb},
		LengthSig: 4, Desc: "RedHat Package Manager (RPM) package",
	}
	bin = Fileobj{
		Ext:       "BIN", Offset: 0, Hexsig: []byte{0x53, 0x50, 0x30, 0x31},
		LengthSig: 4, Desc: "Amazon Kindle Update Package",
	}
	pic = Fileobj{
		Ext:       "PIC", Offset: 0, Hexsig: []byte{0x00},
		LengthSig: 1, Desc: "IBM Storyboard bitmap file",
	}
	ico = Fileobj{
		Ext:       "ICO", Offset: 0, Hexsig: []byte{0x00, 0x00, 0x01, 0x00},
		LengthSig: 4, Desc: "Computer icon encoded in ICO file format",
	}
	gif87 = Fileobj{
		Ext:       "GIF", Offset: 0, Hexsig: []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61},
		LengthSig: 6, Desc: "Image file encoded in the Graphics Interchange Format (GIF)",
	}
	gif89 = Fileobj{
		Ext:       "GIF", Offset: 0, Hexsig: []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61},
		LengthSig: 6, Desc: "Image file encoded in the Graphics Interchange Format (GIF)",
	}
	tif = Fileobj{
		Ext:       "TIF", Offset: 0, Hexsig: []byte{0x49, 0x49, 0x2A, 0x00},
		LengthSig: 4, Desc: "Tagged Image File Format",
	}
	tiff = Fileobj{
		Ext:       "TIFF", Offset: 0, Hexsig: []byte{0x4D, 0x4D, 0x00, 0x2A},
		LengthSig: 4, Desc: "Tagged Image File Format",
	}
	png = Fileobj{
		Ext:       "PNG", Offset: 0, Hexsig: []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
		LengthSig: 8, Desc: "Image encoded in the Portable Network Graphics format",
	}
	jpg = Fileobj{
		Ext:       "JPG", Offset: 0, Hexsig: []byte{0xFF, 0xD8, 0xFF, 0xDB},
		LengthSig: 4, Desc: "JPEG raw or in the JFIF or Exif file format",
	}
	jpg1 = Fileobj{
		Ext:       "JPG", Offset: 0, Hexsig: []byte{0xFF, 0xD8, 0xFF, 0xE0, randomSym, randomSym, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01},
		LengthSig: 12, Desc: "JPEG raw or in the JFIF or Exif file format",
	}
	jpg2 = Fileobj{
		Ext:       "JPG", Offset: 0, Hexsig: []byte{0xFF, 0xD8, 0xFF, 0xE1, randomSym, randomSym, 0x45, 0x78, 0x69, 0x66, 0x00, 0x00},
		LengthSig: 12, Desc: "JPEG raw or in the JFIF or Exif file format",
	}
	wav = Fileobj{
		Ext:       "WAV", Offset: 0, Hexsig: []byte{0x52, 0x49, 0x46, 0x46, randomSym, randomSym, randomSym, randomSym, 0x57, 0x41, 0x56, 0x45},
		LengthSig: 12, Desc: "Waveform Audio File Format",
	}
	avi = Fileobj{
		Ext:       "AVI", Offset: 0, Hexsig: []byte{0x52, 0x49, 0x46, 0x46, randomSym, randomSym, randomSym, randomSym, 0x41, 0x56, 0x49, 0x20},
		LengthSig: 12, Desc: "Audio Video Interleave video format",
	}
	mp3 = Fileobj{
		Ext:       "MP3", Offset: 0, Hexsig: []byte{0xFF, 0xFB},
		LengthSig: 2, Desc: "MPEG-1 Layer 3 file without an ID3 tag or with an ID3v1 tag (which's appended at the end of the file)",
	}
	mp32 = Fileobj{
		Ext:       "MP3", Offset: 0, Hexsig: []byte{0x49, 0x44, 0x33},
		LengthSig: 3, Desc: "MP3 file with an ID3v2 container",
	}
	bmp = Fileobj{
		Ext:       "BMP", Offset: 0, Hexsig: []byte{0x42, 0x4D},
		LengthSig: 2, Desc: "BMP file, a bitmap format used mostly in the Windows world",
	}
)

type Signature struct {
	StockSig []Fileobj
	log *log.Logger
}
type Fileobj struct {
	Ext       string
	Desc      string
	Offset    int
	Hexsig    []byte
	LengthSig int
}

func NewSignatureStock() *Signature {
	return  &Signature{
		log:log.New(os.Stdout, prefix, log.Ldate | log.Ltime),
		StockSig:[]Fileobj{png,jpg,jpg1,ico,rpm,bin,pic,gif87,gif89,tif,tiff,mp3,avi,mp32,bmp},
	}
}
func (s *Signature) FoundTypeFile(fin *os.File) (*Fileobj, error) {
	for _, x := range s.StockSig {
		resultCheck, err := x.checkType(fin)
		if err != nil {
			s.log.Printf(err.Error())
			return nil, err
		}
		if resultCheck {
			return &x, nil
		}
	}
	return nil, errorWrongType
}
func (f *Fileobj) checkType(fin *os.File) (bool, error) {
	_, err := fin.Seek(0, f.Offset)
	if err != nil {
		return false, err
	}
	b := make([]byte, f.LengthSig)
	_, err = fin.Read(b)
	if err != nil {
		return false, err
	}

	for index, x := range b{
		if f.Hexsig[index] == randomSym {
			continue
		} else {
			if x != f.Hexsig[index] {
				return false, nil
			}
		}
	}
	return true, nil
}
