package cache

import (
	"log"
	"sync"
	"time"

	"github.com/terc1997/go-cms/internal/db"
)

type Cache struct {
	Data        interface{}
	Expiration  time.Time
	RefreshRate time.Duration
	Lock        sync.Mutex
}

var (
	instance *Cache
	once     sync.Once
)

// NewCache creates a new cache instance if it doesn't exist, otherwise returns the existing instance.
func NewCache(data interface{}, refreshRate time.Duration) *Cache {
	once.Do(func() {
		instance = &Cache{
			Data:        data,
			Expiration:  time.Now().Add(refreshRate),
			RefreshRate: refreshRate,
		}
		instance.scheduler() // Start cache scheduler
	})
	return instance
}

// scheduler updates the cache data periodically.
func (c *Cache) scheduler() {
	go func() {
		for {
			time.Sleep(c.RefreshRate)
			c.update()
		}
	}()
}

// update updates the cache data.
func (c *Cache) update() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	log.Printf("Updating cache...\n")
	dbc := db.NewDBConfig()
	articles, err := dbc.GetArticles()
	if err != nil {
		log.Fatalln("Failed to fetch articles")
	}
	// Perform cache update logic here (e.g., fetch data from database)
	var articleTitles []string
	for _, article := range articles {
		articleTitles = append(articleTitles, article.Title)
	}
	c.Data = articleTitles
	log.Println("Cache Data: ", c.Data)
	c.Expiration = time.Now().Add(c.RefreshRate)
}

// getData returns the cached data if it's still valid.
func (c *Cache) getData() (interface{}, error) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if time.Now().After(c.Expiration) {
		// Cache expired, trigger update
		c.update()
	}
	return c.Data, nil
}
