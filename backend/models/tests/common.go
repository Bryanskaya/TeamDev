package tests

import (
	"api/teamdev/controllers"
	"api/teamdev/objects"
	"api/teamdev/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"net/http/cookiejar"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func StubConnecton() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:?cache=shared")
	db.AutoMigrate(
		objects.Recipe{},
		objects.Account{},
		objects.Step{},
		objects.RecipeCategory{}, objects.Category{},
		objects.Ingredient{}, objects.RecipeIngredient{},
	)
	return db, err
}

func StubServer() (port uint16) {
	port = 9000
	path := "9000.txt"

	//utils.InitConfig(strconv.Itoa(int(port)))
	utils.InitLogger(path)

	db, err := StubConnecton()
	if err != nil {
		panic(err)
	}

	r := controllers.InitRouter(db)

	go controllers.RunRouter(r, port)
	return port
}

func CompareRecipe(t *testing.T, objA, objB objects.Recipe, msgAndArgs ...interface{}) (ok bool) {
	objA.CreatedAt = time.Time{}
	objB.CreatedAt = time.Time{}
	return assert.Equal(t, objA, objB, msgAndArgs)
}
func CompareRecipes(t *testing.T, listA, listB []objects.Recipe, msgAndArgs ...interface{}) (ok bool) {
	for i := 0; i < len(listA); i++ {
		listA[i].CreatedAt = time.Time{}
	}
	for i := 0; i < len(listB); i++ {
		listB[i].CreatedAt = time.Time{}
	}
	return assert.ElementsMatch(t, listA, listB, msgAndArgs)
}

type ClientE2E struct {
	client *http.Client
}

func NewClient() *ClientE2E {
	client := new(ClientE2E)
	jar, _ := cookiejar.New(nil)
	client.client = &http.Client{
		Jar: jar,
	}
	return client
}

func (this *ClientE2E) PostQuery(url string, body string) error {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	res, err := this.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Unexpected status %v", res.StatusCode))
	}
	return err
}

func (this *ClientE2E) GetQuery(url string, obj interface{}) error {
	res, err := this.client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Unexpected status %v", res.StatusCode))
		return err
	}

	return json.NewDecoder(res.Body).Decode(obj)
}

func (this *ClientE2E) GetCookie(ur string) []*http.Cookie {
	return nil //this.client.Jar.Cookies(&url.URL{});
}
