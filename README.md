# signatures
Signatures of typical files with functionality for verification

check the file type by signature

install 
---
```
 go get github.com/spouk/signatures
```

example usage
---

```go

//open a file of the type you want to install
  f, err := os.Open("qrcode/noimage.jpg")
	checkerror(err)
	defer f.Close()
//make instance 
	sig:= signatures.NewSignatureStock()
	fo, err := sig.FoundTypeFile(f)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Found type: %v\n", fo)
	}

```

wbr//spouk
