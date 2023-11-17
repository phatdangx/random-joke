package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"random-joke/config"
	"random-joke/model"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/gommon/log"
)

type NameService struct {
	client *http.Client
	redis  *redis.Client
}

func NewNameService(client *http.Client, redisClient *redis.Client) *NameService {
	return &NameService{
		client: client,
		redis:  redisClient,
	}
}

func (ns *NameService) fetchNameFromCache() (*model.Name, error) {
	// Fetch the latest name from the cache
	keys, err := ns.redis.SMembers("randomName:index").Result()
	if err != nil || len(keys) == 0 {
		return nil, fmt.Errorf("no names found in cache")
	}

	// Retrieve the most recently added name
	latestKey := keys[len(keys)-1]
	cachedNameStr, err := ns.redis.Get(latestKey).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON string back into a Name struct
	var cachedName model.Name
	err = json.Unmarshal([]byte(cachedNameStr), &cachedName)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling name from cache: %v", err)
	}

	return &cachedName, nil
}

func (ns *NameService) cacheName(name *model.Name) error {
	// Generate a unique key for the name, e.g., using a timestamp
	key := "randomName:" + strconv.FormatInt(time.Now().UnixNano(), 10)

	// Marshal
	nameStr, err := json.Marshal(name)
	if err != nil {
		return err
	}

	// Cache the name with the unique key and set expiration for 5 minutes
	err = ns.redis.Set(key, nameStr, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	// Add the key to the index
	return ns.redis.SAdd("randomName:index", key).Err()
}

func (ns *NameService) GetRandomName() (name *model.Name, err error) {
	// Send GET req
	resp, err := ns.client.Get(config.Config.ExternalService.RandomName)
	if err != nil || resp.StatusCode != 200 {
		// API call failed, try fetching from cache
		cachedName, cacheErr := ns.fetchNameFromCache()
		if cacheErr != nil {
			log.Printf("fetch name from cache error: %v", cacheErr)
			return nil, cacheErr
		}
		return cachedName, nil
	}

	defer resp.Body.Close()

	// Read the resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read random name body error: %v", err)
		return nil, err
	}

	// Unmarshal the JSON data into struct
	err = json.Unmarshal(body, &name)
	if err != nil {
		log.Errorf("unmarshall random name error: %v", err)
		return nil, err
	}
	if ns.redis != nil {
		cacheErr := ns.cacheName(name)
		if cacheErr != nil {
			log.Printf("cache name error: %v", cacheErr)
			// Even if caching fails, return the fetched name
		}
	}
	return name, nil
}
