package seeds

import (
	"log"
	"sync"

	"github.com/go-faker/faker/v4"
	"github.com/valeelim/mahchat/internal/database"
	"github.com/valeelim/mahchat/pkg/dao"
	"github.com/valeelim/mahchat/pkg/dto"
)

func seedMessage(db *database.Conn, channels []*dto.GetChannelResponse, count int) {
	var wg sync.WaitGroup

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch *dto.GetChannelResponse) {
			defer wg.Done()
			for i := 0; i < count; i++ {
				groupMsg, err := dao.NewGroupMessage(ch.ID, 1, faker.Sentence())
				if err != nil {
					return
				}
				if err := db.CreateGroupMessage(groupMsg); err != nil {
					return
				}
			}
		}(ch)
	}
	wg.Wait()
	log.Println("done seeding message")
}
