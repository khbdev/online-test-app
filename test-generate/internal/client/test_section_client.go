package client

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"time"

	testpb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/grpc"
)

type SectionClient struct {
    client testpb.TestServiceClient
    addr   string
}

// NewSectionClient – GRPC client yaratadi
func NewSectionClient(addr string) *SectionClient {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    c := testpb.NewTestServiceClient(conn)
    return &SectionClient{client: c, addr: addr}
}

func (s *SectionClient) GetSection(sectionID uint64) (*testpb.Section, error) {
    var lastErr error
    for i := 0; i < 3; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
        resp, err := s.client.GetFullSectionStructure(ctx, &testpb.GetFullSectionRequest{SectionId: sectionID})
        cancel() // shu yerda darhol cancel qilish

        if err == nil {
            section := resp.Section

            // random 10 ta savol
            if len(section.Questions) > 10 {
                rand.Seed(time.Now().UnixNano())
                rand.Shuffle(len(section.Questions), func(i, j int) {
                    section.Questions[i], section.Questions[j] = section.Questions[j], section.Questions[i]
                })
                section.Questions = section.Questions[:10]
            }

            return section, nil
        }

        lastErr = err
        log.Printf("[Attempt %d] Failed to get section: %v", i+1, err)
        time.Sleep(1 * time.Second)
    }

    return nil, fmt.Errorf("all retries failed: %w", lastErr)
}

