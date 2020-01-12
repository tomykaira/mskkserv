//go:generate mkdir -p ../gen
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/commands/candidates.proto mozc/commands/commands.proto mozc/commands/engine_builder.proto mozc/commands/renderer_command.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/config/config.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/converter/segmenter_data.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/protocol/state.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/proxy.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/renderer/renderer_style.proto
//go:generate protoc -I ../protos --go_out=plugins=grpc:$GOPATH/src mozc/user_dictionary/user_dictionary_storage.proto

package skkserv

import (
	"context"
	"log"
	"time"

	pb "github.com/tomykaira/mskkserv/mozc"
	commands "github.com/tomykaira/mskkserv/mozc/commands"
	"google.golang.org/grpc"
)

// GoogleJapaneseInput is a local google japanese input client.
type GoogleJapaneseInput struct {
	conn   *grpc.ClientConn
	client pb.ProxyClient
}

// NewGoogleJapaneseInput creates google japanese input based translation engine.
func NewGoogleJapaneseInput() (*GoogleJapaneseInput, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := pb.NewProxyClient(conn)
	return &GoogleJapaneseInput{conn: conn, client: client}, nil
}

// Search responds to translation queries.
func (g *GoogleJapaneseInput) Search(query string) (cands []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	activated := true

	r, err := g.client.SendKeyEvent(ctx, &pb.SendKeyEventRequest{
		KeyEvent: &commands.KeyEvent{
			KeyString: &query,
			Mode:      commands.CompositionMode_HIRAGANA.Enum(),
			Activated: &activated,
		},
		Context: &commands.Context{},
	})
	if err != nil {
		log.Printf("gRPC send key event error %v", err)
		return nil
	}

	log.Printf("Res %v", r)

	rc, err := g.client.SendCommand(ctx, &pb.SendCommandRequest{
		Command: &commands.SessionCommand{
			Type: commands.SessionCommand_REVERT.Enum(),
		},
		Context: &commands.Context{},
	})
	if err != nil {
		log.Printf("gRPC send revert command error %v", err)
		return nil
	}
	log.Printf("Res %v", rc)

	for _, v := range r.Output.Candidates.Candidate {
		cands = append(cands, *v.Value)
	}

	return
}

// Close cleanups resources used by the class.
func (g *GoogleJapaneseInput) Close() {
	g.conn.Close()
}
