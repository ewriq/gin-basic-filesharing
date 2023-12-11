package Handlers

import ( 
	"github.com/gin-gonic/gin" 
	"net/http"
)

func Home(c *gin.Context) {
	html := `
	<html>
	<body>
		<form method="POST" action="/upload" enctype="multipart/form-data">
			<input type="file" name="file" accept=".zip">
			<input type="submit" value="Yükle">
		</form>
	</body>
	</html>
`
c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
