# vkcloud-vision
Golang implementation Vk Cloud Vision API v1.0 - https://mcs.mail.ru/docs/ml/vision/about-vision

[![GoDoc](https://godoc.org/github.com/TheLastochka/vkcloud-vision?status.svg)](https://pkg.go.dev/github.com/TheLastochka/vkcloud-vision)

Installation
-----------
Last release:
```
go get github.com/TheLastochka/vkcloud-vision
```

Last commit:
```
go get github.com/TheLastochka/vkcloud-vision@main
```

Usage
-----------
### Face recognition
Setting fixed person id for person in image:
```go
client := &http.Client{
  Timeout: 15 * time.Second,
}
vis := vision.NewVisionClient(client, "YOUR_APIKEY")

respOk, respErr := vis.PersonsSet(vision.MetaSet{
  Space: "0",
  Images: []vision.ImageMeta{
    {
      Name:     "test_imgs/rach.jpeg",
      PersonId: 1,
    },
  },
})
if respErr != nil {
  fmt.Printf("error: %s\n", respErr.Body)
} else {
  fmt.Println("Ok")
  fmt.Println(respOk)
}
```

Face recognition in the image with the creation of new faces, if there are none:
```go
client := &http.Client{
  Timeout: 15 * time.Second,
}
vis := vision.NewVisionClient(client, "YOUR_APIKEY")

respOk, respErr := vis.PersonsRecognize(vision.MetaRecognize{
  Space: "0",
  Images: []vision.ImageMeta{
    {
      Name: "test_imgs/friends.jpg",
    },
  },
  CreateNew:       true,
  UpdateEmbedding: false,
})
if respErr != nil {
  fmt.Printf("error: %s\n", respErr.Body)
} else {
  fmt.Println("Ok")
  fmt.Printf("status object 0: %d\n", respOk.Body.Objects[0].Status)
  fmt.Printf("persons recognized: %d\n", len(respOk.Body.Objects[0].Persons))
}
```
More examples can be found in the examples directory.
### Not implemented yet
* Recognition of types of documents
* Object recognition
* Text recognition in documents
* Image processing methods
* Document recognition
* License plate recognition
* Text recognition in photos
* Content Recognition 18+

Features
--------

* Without external dependencies

Requirements
-----------

* Need at least `go1.19` or newer.

Documentation
-----------

You can read package documentation [here](https://pkg.go.dev/github.com/TheLastochka/vkcloud-vision).
Official docs: [here](https://mcs.mail.ru/docs/ml/vision/about-vision)