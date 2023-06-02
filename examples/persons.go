package main

import (
	"fmt"
	"log"
	"os"

	vision "github.com/TheLastochka/vkcloud-vision"

	"github.com/joho/godotenv"
)

func main() {
	// get token from env
	godotenv.Load()
	token := os.Getenv("AI_VISION_TOKEN")
	if token == "" {
		log.Fatal("AI_VISION_TOKEN is not set")
	}

	vis := vision.NewVisionClient(token)

	// vis.PersonsSet(vision.MetaSet{
	// 	Space: "0",
	// 	Images: []vision.MetaImageSet{
	// 		{
	// 			Name:     "test_imgs/rach.jpeg",
	// 			PersonId: 1,
	// 		},
	// 	},
	// })

	respOk, respErr := vis.PersonsRecognize(vision.MetaRecognize{
		Space: "0",
		Images: []vision.MetaImageRecognize{
			{
				Name: "test_imgs/friends.jpg",
			},
		},
		CreateNew:       false,
		UpdateEmbedding: false,
	})

	if respErr != nil {
		fmt.Printf("error: %s\n", respErr.Body)
	} else {
		fmt.Println("Ok")
		fmt.Printf("status object 0: %d\n", respOk.Body.Objects[0].Status)
		fmt.Printf("persons recognized: %d\n", len(respOk.Body.Objects[0].Persons))
	}

	// vis.PersonsDelete(vision.MetaDelete{
	// 	Space: "0",
	// 	Images: []vision.MetaImageSet{
	// 		{
	// 			Name:     "test_imgs/rach.jpeg",
	// 			PersonId: 1,
	// 		},
	// 	},
	// })

	// respOk, respErr := vis.PersonsTruncate(vision.MetaTruncate{
	// 	Space: "0",
	// })
	// if respErr != nil {
	// 	fmt.Printf("error: %s\n", respErr.Body)
	// } else {
	// 	fmt.Println("Ok")
	// 	fmt.Printf("status: %d\n", respOk.Status)
	// }
}
