package service

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/valeelim/mahchat/pkg/chat"
	"github.com/valeelim/mahchat/pkg/dao"
	"github.com/valeelim/mahchat/pkg/dto"
	"github.com/valeelim/mahchat/pkg/repository"
)

type CreateChannel func(c *gin.Context, req dto.CreateChannelRequest) (*dto.CreateChannelResponse, error)

type GetChannel func(c *gin.Context) (*dto.GetChannelResponse, error)

type ServeChannel func(c *gin.Context, wg *sync.WaitGroup) func()

type CreateGroupMessage func(c *gin.Context, req dto.CreateGroupMessageRequest) error

type GetGroupMessages func(c *gin.Context) (*dto.GetGroupMessages, error)

func CreateChannelService(repo repository.Channel) CreateChannel {
	return func(c *gin.Context, req dto.CreateChannelRequest) (*dto.CreateChannelResponse, error) {
		channel, err := dao.NewChannel(req.Name)
		if err != nil {
			return nil, err
		}
		if err := repo.CreateChannel(channel); err != nil {
			return nil, err
		}
		return &dto.CreateChannelResponse{
			ID:   channel.ID,
			Name: channel.Name,
		}, nil
	}
}

func GetChannelByIDService(repo repository.Channel) GetChannel {
	return func(c *gin.Context) (*dto.GetChannelResponse, error) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return nil, err
		}
		channel, err := repo.GetChannelByID(id)
		if err != nil {
			return nil, err
		}
		return &dto.GetChannelResponse{
			ID:        channel.ID,
			Name:      channel.Name,
			CreatedAt: channel.CreatedAt,
		}, nil
	}
}

func ServeChannelService(repo repository.Channel, client *redis.Client) ServeChannel {
	return func(c *gin.Context, wg *sync.WaitGroup) func() {
		chat, close := chat.New(c.Param("id"), client)
		defer func(){
			wg.Done()
		}()
		
		chat.ServeWS(c)
		return close
	}
}

func CreateGroupMessageService(repo repository.GroupMessage) CreateGroupMessage {
	return func(c *gin.Context, req dto.CreateGroupMessageRequest) error {
		groupMessage, err := dao.NewGroupMessage(req.ChannelID, req.UserID, req.Content)
		if err != nil {
			return err
		}
		if err := repo.CreateGroupMessage(groupMessage); err != nil {
			return err
		}
		return nil
	}
}

func GetGroupMessageByChannelIDService(
	msgRepo repository.GroupMessage,
	channelRepo repository.Channel) GetGroupMessages {
	return func(c *gin.Context) (*dto.GetGroupMessages, error) {
		idInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return nil, err
		}
		if _, err := channelRepo.GetChannelByID(idInt); err != nil {
			return nil, err
		}
		result, err := msgRepo.GetGroupMessageByChannelID(c.Param("id"))
		if err != nil {
			return nil, err
		}
		return &dto.GetGroupMessages{Data: result}, nil
	}
}
