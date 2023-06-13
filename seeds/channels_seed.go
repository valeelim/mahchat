package seeds

import (
	"log"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/valeelim/mahchat/internal/database"
	"github.com/valeelim/mahchat/pkg/dao"
	"github.com/valeelim/mahchat/pkg/dto"
)

func seedChannel(db *database.Conn, count int) ([]*dto.GetChannelResponse, error) {
	var wg sync.WaitGroup
	wg.Add(count)
	var result []*dto.GetChannelResponse

	for i := 0; i < count; i++ {
		time.Sleep(10 * time.Millisecond)
		go func() {
			defer wg.Done()

			channel, err := dao.NewChannel(faker.Word())
			if err != nil {
				log.Println("new channel", err)
				return
			}
			if err := db.CreateChannel(channel); err != nil {
				log.Println("create channel error huaa", err)
				return
			}
			result = append(result, &dto.GetChannelResponse{
				ID: channel.ID,
				Name: channel.Name,
			})
		}()
	}
	wg.Wait()
	log.Println("channel seed generated")
	return result, nil
}
