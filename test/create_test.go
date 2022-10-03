package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	pb "intermediate_server/internal/models/pb"

	resty "github.com/go-resty/resty/v2"
	"google.golang.org/protobuf/proto"
)

func TestCreateWithMainServer(t *testing.T) {
	archive, err := os.Open("Backup-Pizzeria Maria Express-22_08_2022_20_46-v3.2.2.0.zip")
	if err != nil {
		t.Error(err)
	}

	archiveStat, err := archive.Stat()
	if err != nil {
		t.Error(err)
	}

	contentBytes := make([]byte, archiveStat.Size())
	_, err = archive.Read(contentBytes)
	if err != nil {
		t.Error(err)
	}

	message := pb.BackupCreate{Id: "test id", File: &pb.File{Title: "Backup-Pizzeria Maria Express-22_08_2022_20_46-v3.2.2.0.zip", Content: contentBytes}}
	protoBytes, err := proto.Marshal(&message)

	fmt.Printf("backup size - %.2f Mb\n", float64(len(protoBytes))/(1<<20))

	client := resty.New()
	resp, err := client.R().SetBody(protoBytes).Post("http://localhost:8080/api/backup/save")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode() != http.StatusCreated {
		t.Error(fmt.Sprintf("expected 201 status, found %d, error - %s", resp.StatusCode(), resp.Body()))
	}
}
