package service

import (
	"backend/api/dto"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms/ollama"
)

type UploadService interface {
	Upload(data dto.UploadRequest) error
	Load(data dto.LoadRequest) error
}

type uploadSvc struct {
	ctx context.Context
}

func NewUploadService() UploadService {
	ctx := context.Background()
	return &uploadSvc{ctx: ctx}
}

func (u *uploadSvc) Upload(data dto.UploadRequest) error {

	// if !validateFileType(data.FileType) {
	// 	return errors.New("invalid file type, only PDFs are accepted")
	// }

	return nil
}

func (u *uploadSvc) Load(data dto.LoadRequest) error {
	// Open the file
	file, err := os.Open(os.Getenv("TMP_FILE_DIR") + "test.pdf")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Obtain the file size
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error obtaining file info:", err)
		return err
	}
	size := fi.Size()

	// Create the PDF loader
	pdfLoader := documentloaders.NewPDF(file, size)

	// Now you can use pdfLoader to load the PDF and process it
	docs, err := pdfLoader.Load(context.Background())
	if err != nil {
		fmt.Println("Error loading PDF:", err)
		return err
	}

	llm, err := ollama.NewChat(ollama.WithLLMOptions(ollama.WithModel("llama2")))
	if err != nil {
		log.Fatal(err)
	}

	// We can use LoadStuffQA to create a chain that takes input documents and a question,
	// stuffs all the documents into the prompt of the llm and returns an answer to the
	// question. It is suitable for a small number of documents.
	stuffQAChain := chains.LoadStuffQA(llm)

	answer, err := chains.Call(u.ctx, stuffQAChain, map[string]any{
		"input_documents": docs,
		"question":        "What is this document about?",
	})
	if err != nil {
		return err
	}
	fmt.Println(answer)

	return nil
}
