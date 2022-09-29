package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"project/models"
	"project/pkg/dbhelpers"
	"fmt"
)

func (h handler) AddBook(c *gin.Context) {

	var book=[]models.Book{
		{ Title:"Think and Grow Rich",Author:"Napoleon Hill",Desc:"It inspire more people to become millionaires and billionaires" },
		{ Title:"The Alchemist",Author:"Paulo Coelho",Desc:"Paulo Coelho's masterpiece tells the mystical story of Santiago, an Andalusian shepherd boy who yearns to travel in search of a worldly treasure." },
		{ Title:"Rich Dad, Poor Dad",Author:"Robert T. Kiyosaki",Desc:"Rich Dad Poor Dad is Robert's story of growing up with two dads" },
	}
	h.DB.Create(&book)

	var addbook=models.Book{ Title:"How to Win Friends & Influence People",Author:"Dale Carnegie ",Desc:"Dale Carnegie’s rock-solid, time-tested advice has carried countless people up the ladder of success in their business and personal lives." }
	k := dbhelpers.InsertQuery(h.DB, &addbook)
	fmt.Println(k.RowsAffected)

	var data_payload []interface{}
	book = append(book,addbook)
	data_payload = append(data_payload, book)

    c.JSON(http.StatusOK, gin.H{"data":data_payload,"message": "Added a Book successfully",})
}