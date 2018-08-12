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

	resp, err := requestRecognition("in.flac")
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Println(alt.Transcript)
		}
	}
}

func requestRecognition(filename string) (*speechpb.LongRunningRecognizeResponse, error) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalln("failed to create client:", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	op, err := client.LongRunningRecognize(ctx, &speechpb.LongRunningRecognizeRequest{
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
		return nil, err
	}

	return op.Wait(ctx)
}
