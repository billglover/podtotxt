package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalln("failed to create client:", err)
	}

	data, err := ioutil.ReadFile("in.flac")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: int32(48000),
			LanguageCode:    "en-GB",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: data,
			},
		},
	})
	if err != nil {
		log.Fatalln("failed to recognise:", err)
	}

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Println(alt.Transcript)
		}
	}
}
