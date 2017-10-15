package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Link struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
	Link string `json:"link"`
}

type Location struct {
	Id  int     `json:"id"`
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Sight struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Location    Location `json:"location"`
	Links       []Link   `json:"links"`
}

type Achievement struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type User struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Last_name    string        `json:"last_name"`
	Rating       int           `json:"rating"`
	Image        string        `json:"image"`
	Achievements []Achievement `json:"achievements"`
}

type Prize struct {
	Price       int         `json:"price"`
	Achievement Achievement `json:"achievement"`
}

var user User
var achievesList []Achievement
var achievementCounter int = 0

func sendMapItems(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("base.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, str)
}

func sendUserInfo(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func sendPrize(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	achieveIdStr := r.FormValue("id")
	achieveId, err := strconv.Atoi(achieveIdStr)
	if err != nil {
		fmt.Println("error:", err)
	}

	price := 100

	for _, userAchieveId := range user.Achievements {
		if userAchieveId.Id == achieveId {
			price = 0
		}
	}

	newAchievement := achievesList[achievementCounter]
	newAchievement.Id = achieveId

	if price != 0 {
		achievementCounter++
		achievementCounter = achievementCounter % 5
		user.Achievements = append(user.Achievements, newAchievement)
	}

	user.Rating += price

	prize := Prize{
		Price:       price,
		Achievement: newAchievement,
	}

	b, err := json.Marshal(prize)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func main() {

	user = User{
		Id:           1,
		Name:         "Curiosity",
		Last_name:    "Curiosity",
		Rating:       45,
		Image:        "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxMTEhUSEhMWFRUXFxgVFxUVFxcVFxgXFRUXFxUVFxUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGhAQGi0lHyYtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLf/AABEIAOEA4QMBIgACEQEDEQH/xAAcAAAABwEBAAAAAAAAAAAAAAAAAQIDBAUGBwj/xAA9EAABAwIEAwYDBwIGAgMAAAABAAIRAwQFEiExQVFhBhMicYGRobHRFDJCUsHh8AfxFSNicpKiM8JTgtL/xAAZAQADAQEBAAAAAAAAAAAAAAAAAQIDBAX/xAAlEQEBAAICAgEEAgMAAAAAAAAAAQIRAyESMQQTMkFhFFEiQnH/2gAMAwEAAhEDEQA/AJ4CUAkhGtnnA5qwFYxUcOq6BKwWN0orPnnOi5fkz07vhXuxIo1tITVzTO433BTFo8bZnDzCniOfuFxPQujVhjEHxaPbxHHmFsLG/a9odwPwPVYHELUjxD3GoUjAsRcw5T906EHZXhneO7jHl45yTVdF0KIxzWapYwWktEneJ3BEyPgqW+7SVGnpP7rtnNjXBfj5NvcOEbqougFn24+agB4jdNtxF7tZmTCPqw/49WNRqaLVAbfmY3IB+Kcur8U4B6lL6s3or8fJJa3VBjOfn7KPb3RcC7ht/PZShVaYExO55AfvPsj6uP8AafoZEU6JcT8TyQrjMIBy028eJPlxKRc3wPgZo0bnny8yUm2aXgT90bCdPVcnJyXO/p28PD4T9pNlSa78OVo57nqVcMho0A+ar6NWNA1JrXJ1lZXLTomO0xvia4aakDRRcYfswKVhR/yy7qqu6dmcSuz4/eO3D8q6ukXu0Xdp7IhkXQ4zGRJyKRlREJhHLEnIpKBQe0Xu0FJjoggNwEYSQlAK0DWR7VUYqTrqJ2WuVT2gs87QRw0Kx+RP8dun4uWuRi2PHDXz0T5zDUhw67hMvY6mSNj+Y6p2wDi6O8GvX5hcetvT2T3pOnupP2YAA8Tw5K57lgac7RPBzdiqi9fA6bCOSizafJIbUaMruJgHzH7Kixvwy2JIfH0Uh2bK/kIIPUftKYD+8qAfmyO/6kO+IKrGa7RajWYymHcY6RIlO0CWuLD5hN45Zlha8bEo7+9BeRBHhBHzn4rT33C9J1o8Q9x2HH5D3VbdvzGZ0PH02Q+0ltHcHPIjlHHz1U2xwwmjmcNAZ/SPeEt6u6PaXg9HwtHAkn22+alXNPMRl2M+Wh38lFw+oWsOusBoHnqT7ED1U6gY8BEkideAOw8+Psssva4juszMR4eHXmSptvT018I5lSLarwOpAjp+5USvbZiZcpm1xIqOAEgknlsoVS9zaRCJ9LSAZ8yAfZIsqOd+XjxKVi/S+aCKDQOKrxSV1cNhgbyUHIvQ4ZrGPJ+Rd5ondIjTU3Ikli1YIRpojTU3IkFiYQ+7RGmpZYk5EtGi930QUnIgjRdNQEoJMowVaCk1duhhJ2hLzKPe1hlIOk81ny/bWvD98YbELpr5U7CKVINl7J9fkoGLWYa/M3nqFZW1zmZlEDnA0Xny6j1srPwav6gGjCY3GsghQaNUhwAE5iJB2+KvbbDA/cEDmPkkX9Sm0hri2YjoeXkl0keOWQZQJYNodHHr6LIdnyX1Wt3ifbf9SthfXDalIgH6j24LN4NbGlUJ2PA8ITmcmNha7XHaa0DqQHHh6cfaVmb20MiZDoA9I+gHutQ+4zA5uGg91CrW2bKTvGkb8pjoICeGWhYg2+H/APjbGgIdz4nQrXX1kO5JbsATHoqih4fESOnXaTHop1bEjkIHLTl0UZ3vtWLNYbVms0bhswOZO/sI9lrcRt20mh5J135+h4DqstgFoRVzmIHP91r7+q2o0NcYbvqTr5BPks30UlUtJzpkOBbyG37q0u7IGnLRMcjqT1hNMw3KJkRv5eRhP0qYEDMSIkctVMq2WLHZ4Oh5LVYDZCM51Kz+Muyv0GvstF2TzEeIQIR/tIvLfjsuq4knzSYTNa+YHubOxITzHgr1I8TKWXsIRZUuERTSbISS1OkJJCAaISS1OkIoQDMdUacyoIC/RhJlHKtI1WY7SmmY/hVmFWYy7QBRn1i04pvKM8LEu8TuWyfsbeJBEefJLfdtaJdtyUX/ABbvninSaSenDqV5uUtepEvtDiPc0obBMbkfGVzup3lR0wXE/wA0Wsx+wfLA8yCdugEqpu6/dAFojXfy109lpxTRZq+4tbm3g1KdWlO2drmg84zAStR2ZuRXZB+83RVVHFqt1LK1Rz2jUNJJDTtICe7IO7u6gnwuEH9FfJjLNJwtXt1QLXZeZn36pt51g7zHGPfn9VaYxTiqDP4QfmqarGbWdeZnZc2+mmi2b6tI1J166aexViyxAZLuX9/kmbemS0naNuesb+in9rX93Zkt3IDffRL7rpXphMVxw5y2noG8eZVdTxWsDmznzU3Dq32cisAC5u2YBwBPGCNVLbixua2eqGExoAxoBA5taIK65hjMd6Y3K7aLshjnfDu6gE7SBw91e3VsAdIHILD0rYULhhpatfw00PELSuvW94O9OU9dlyZ++m2Po1d4UT4xqQdk/RvzTpu/NENHVW2h1aRBHBQLigMwPJVxyXLSM8rMdudXVao15LiZJlaPs5ipd4TuofbWgGuBCidl6c1AZXf6cnWeLegoFLazRDItXGbKJOFiS5qNA2USVkRZUgEIIsqCejXSMIkpo5KiGFV41UAbqrUMPI+yqMcoFzTooz+2r4rrOMDitcl2UHcwAtf2aw5lNmYtBfxOsx6rFXFImoBtr7K+ZeVBAIlcOfrT0p7XPaWgC0OA5Fp4giQQfMfILJYna5mTrHDmD15LTWuLU/uVJA2PFqcdg1tVMsq5CRtJBI8kplo9bY7CrQsafiTojs3ZaoqbQQRO5Erav7PUKbS59QvA2DnSJWMxiqHOLWtgTHonM5lbD8dTbYXl42oC8cYbpxga6qCyyLtQNuHNJwQjI1o6b8+qvWgCJ3XBnbL06ccZpFsWQCDv15hV3ay971gY06gfd46f2VxeVwNSsneuHflwIMn+6rgtuRckkitq2+ZhHD5FN4Ja5HEu9Ft8KwmlUb4jDjoCDE+6S/su1p1reHlIB14Lrx5fw57x/lAwy17yo1+mmjRxJ/E7oAPmtVeYays0teBIGh4j2VM26o0G/wCWA93MEHbgmKeOucYHhHFTnd+lYw3h926m40nEQ0/DyVxQIc6eCx19Wca2Y6ytjgNI93mPHmteHHecrH5GWsKx3bqp4wFU9n3HvGgevkpXagl9d3IcTsEXZ6nmqta3aZJ5+a6sr2ww6wdIoM8ISyxO06cAJWVbOK1GLEhzFLyJDmI0NopppDqalFiSWI0No2VGn8qNLQSgUouPNNpbUyGEmtqCN/NGCjQcrnmMWjG1t8hniC5vuNW+xTl3buDc5+7zBkH1Wj7QYWHtJA1WKfcvpGMzm+seh5ri5cNV6PDn5QGOn7pPkfpxUq1eXEASOgPufJCnSdXgi3LuBfT8BHmAMp9gequqtH7PSl0mRu9uQ6/6j4Z9Vjl66bxEuPDvvtvPzCoXWZc7QGOani4JdLvCBuXAzHlzWgZhge2JInjOuyyuVwaY4zJT2d1TpkB5M/zZWrMUpOObNtsDpt6LE43Qq0Hw7WNncxwQo30gcytLxY2bgmV3puHXlGoIBk9P5oqO5sXZ8wHpxj04qksrqq+rDBxXSMLwAil43lzj4iDpBPLkseSfTu8V4Xy9qKzPDpx/kpFRzmyIBHMH4kxOnNKxhhZUIa4kiJEGfORuELemagh8ac9D/wARLj7Ix3ZsstekFxzaDXludehJlSbGzdm1Bjckn4JRw5zHeGlWd1DSxnu4n5BSaFa5+4GU2cz3jc3qcxIV6t9I9GMQtml4IbGsTsPjoVrbSnkpRyCoAym14dVqNc8f7yB/11KtLjEqXdOh/DgPrC7fjySbri+Tbeo5vjtUvquPCdBwV/2JY3NPFUtdlIuJmoZP5Wj/ANitH2Po0w4lof65f0Cr3TvWLaBBKEcj7/shI/h/ZdDzyEkpZI/n9kklMEwkwlFEUAmEEaCBs5KUCmwlhIFAowkpxrT/AH0QCXtELMdpMGYW548/qtUAOft9ShVtg4ZXAAH82vw/ZRnhMpprxZ3DLcc6w+0pOMB2eP8A46b3R6uDQr9zsg0a0kD7zozf8QSQrOrg7RJEujg4wweg+qq68unKcoboXEENn8rRu74dYXncvHY9Pj5Jl+VRc3twXaMG/EaD1Ktba6cYziDxAiP55BVTW1S4gkRqfP8AZMl/i0cT8B8Fhe28W2M2bazYPL1XO7y0fTqZIMzp15Ldh5gEnqqu6ALwd4JK34bYjPHa17JYP3bc51cdT0WnqVDwcAdod9eiorC4bz9JS7upPEx5/qp5Md9nOkh7rrNMtjX8IPqrOg95EPLepILD/wAhMepCp6TiCP8AMdBHmR+/Q7qdbtqN17wP0nQbt5xzHEbjdY2f0ravxm0LSXCmanKTp6NmXeYKpWX+aDtGwAgA9ANFc3t+WNJ8TT0MscOMER7EKkfsTA11003W+OPTO1FddOLiSUh9weMIy0GOabuGFdE6Z0TWtctV2WtiATwWQomFtezN2C3LxWmHdYc32tAAgjlFK6nniKIpUpJKCJSSlkopTMhBHKCAcA6pQI800CnGjmpBYdy+CVl5n6/sk5+WnzRsE/zQIBYfy0+funKbQPvanl9TwTfeAfd/5cfTkgxCpR1Xz5chsqnEcMa8bkcoPHmrOU1eGGOPSB+v86qcsZZ2eGVl6Zu1tWkkB33iWDnlBE+5+STUweBpzVVUuCyowtP3QCR1cS/9QtNYXzXy089PI7fBcGXHj+HqzOqR2HlMHDzyWsNsCiFjqosq5mzlHDVZWmHOMDr+qt6dmBupDYbsnorkg0MGGs7TI6Tt/Oqdu8jabgHQ9oJ6kDUkdQNfTolXF6GAk7SB81lO0GJTWYWGDIcCOv8ACiYwrahUKzi4nMC0nUHUH0U65ptLDk4bt4jqOY+KqcWytDX0xAfJLRs1w+80dNQR0IUdmIkkRoQrkGzkBNOfJUs1RVB2a/2DvPkeu3PmoWQ7HQjcFWkw52qv+ztx4gs1VOquezrwHhXj7jPk+2uhsOiUkUjolyut5lEURQQKCJKIBAokGCCKUEHs+LOoNS0pGR/5St2aY5JH2dvILPzdN+P+2Kaw7kED4ny+qJ9QngQOX84rbG0byRfYmflCPMv4/wC2ID0feLZHDaZ4BJOE0+Sfmn+Pf7ZAVEm7MtI6fz+dFrDgdM8ETsCpwjygnDlHEb+tkruEbZR7NCVUrubFSm6I0IHL8P09ArH+o+EGhXzD7r/mBCoLKqG6kgg6EdFx5+3dj6bXDcRzgaqeLjXdZGwuWtd4TIOoKvKVcO1CytXIt2101c14ChtqfBM3V0ls1L2jxQ5cnHc/p9fVVbao7tjjq4SOexn/ANkvEiCS7c8VXvf4QdtT+n0VzuEdtX5y6i4zn1Z0qMktj/cMzfUclFZoZTVCpDs4MEEEHkQZB91JxNwFQ5dGuh7fJ4Do9Jj0WkJIpP1lSQ8VBlcYI0a75Nd068PLaDbOSXOIlGyNVRDiDuN1bYC6HiVUsdnIa77w+67n/oP6H08p9rUykaRBV7Re46XbnQJwlV+GVszARqpmvIrrl28yyynESRJ5FJzeaZFEpJKSXpJnkUgXmRJvN/IQQHTC1ANTgajAXO9Q1COE6iAQDYajypzKjQDZCBCWhCAzvazAW3NJzSJMaea4Zd4SG1TTf4XAxC9KuYuX/wBR8CHeCqG77kc1nnOtnGApWZp7at/mql217lnXyCm27GxllMUMF/zMx1HBc1/bWJ9G80k8VVYniJEBvEK2dZHZVFbCfFmSl1exVDWqPcZjVR6zX6aaAfutCcOOpVdcU3B0RK0meysVL3jZTKjczaPOHMP/ANHkj4PCjXNMztCcbIog8qp/7Mb/APlazuITqrRTgDdR3El2iZGZ5Gq17MCNNge4biZU78YftV4fYjdymFgIL3DUaDqPqhUumt0OyiiuamgMNUzduw2/9P7trvC6Oi6F9jZyXKOxLqf2gZnRG3VdipCQIOi6MbWdxiH9hZyRHDmclPyIsirdLxiB/hzPyhGcPZyU3KhCex4xB/w2n+VBTYQRseESoRoSjAUqFCEJUIIAkUJUIkASNBEgDWe7b22a3d01WglQ8Wo56T2niClZuBwO+qBuoKVa40QROwVRjNQ06jmO4OI+Kq614eCx8VbdTwy6bX8LBJ5KVcYO8/hKr/6LUMxqVHa6wF17um8k/pSjycvp9l6zxtA6qj7QYFUow4jQbldtJAVF2rpNdQfI4I+nINuFVwwnaSpNHB89PXQZ5/6n6p+nVph2o6KVdXQyQ08z+n6FZW2elRmqzAyqGDYET7ruuE2tOvbt2Iyx8Fw+zcHVeZJhdi7D0XUqMErWX8JsYXtpgH2dxP4Sss0OcYYu3dp8MFxTLSFyi9pvtamXL6wi9ehC7TDXUgKjjruuv9kMR72k2TJC4nVxCrUcGwYXSv6cVIcW9EY3vsV0eEISgEcLYiMqLIEsBGgEZAgloIBuEcJSAQBI5QKEIIAgURCNAEhCIuSS5AKKZuNWkdEbnpmtU0QHAu2FgPtNQH8xVJSwsFart61zLpxjQ6yqC0qSuXK2VpJK6P8A0wpNpUnRz1W4bf8AMrnnZS7DKTh1Uuvijltjekem3qXw5qpxi9zU3DoVmBiTuaTUxOQQlae3MsRc4VHD/Vp7pupVc8w2THhHkOKvsSY1zzprzTdvSY3YLK8ipiT2ew8iqJ3XX8L8LAFheyNHO91TgNAtux6c3vdGlzSdKo+1+CMqUi+BI1U6lcQhiNfPSc3mtNylpyPM4aNAlaz+nVw4Vy1+5Cl4Z2VJcC4LSYX2XFKsKg4KMfYrYNRwjahC6EihCEcI0AnKgjRoBkopRZkguQDhcizpslBBFF6LvERREoAFyQXJRcmnlABzlFrOS31FVYle5QUtmwH9RC11QDislQphuil9qsRL65LhCrqFSTK5OWVpjWrwZ8AqXV12VPhFfcK2a5VhekZTtEqVMu6R3uimvpg7hRbqgGtJCWdEjMXryHyiD88ADVOXVPM6RonrVmoEKJfy0bbs5RyUmwFb947gE5hNpFNojgrihZ9F04Y9ds7VIys/8qssNYXnUKwqWgaJKdtnCJCdkgltTrakApYCj0naKRTdKUNIaUJSg1EtCBCEaKUxoIQQQQSAUEnMEoEIASiJQLgia4IIcIijLwilBkJuoU456ae9ARKztJWbxCtJK0d0dFk79haTpIWWd6PTnna0l1TRpgcY0VJSqxsuiXlNr2uaQFj7XBnPqER4Qd1j7NY4BRcRPNaJluU3ZWOQQFLBIT0ZLbfmqztC7I0DmplxXI1WUxi7c93OErCQHVtYVlhcFwMyqfJJVrh9ItghFkkVHZsDYHU2nor6jRCzXZKk9tMZuK1dELfjvSahYpaFzYCzle8dTIDmugcYW2ASX27TuAU7NiM63FmloIKusOdmgwk3VKlSYX9052oGWm3M4lzg0QPMjXhvsmKOONkNFC4bq1smmIlxgfi2mRPCCTpBKmFVq2dLpEVWtxlpzeCoIJbBaJMZ4IE6g5AB/vZzSKONNdk/yqwzFo8TIyhxcA52uglrgeWk6ELQvGrSUJVU7HGAkGnW0bmJySI000Orv9O8gjcLFXl9ffbzTbmzfagB4qmXuRVEAMDgzL3UzLTxPVKr4+K57/5t0uUSVlCCbHv+lSzZAIIIBDkpqCCQFzSxsgggzD0iqgggItwqe+2KCCzyDJ4qiwnZBBYxSzKQ5BBMIV7sqGj+JBBMop/xnzV9b/8AjHmPmjQUZe1x1bAfuN8h8loqaCC6MPSKcCU1GgrhDCccggnCoD6/qjCCCZCai/CiQSqsSUEEEkv/2Q==",
		Achievements: []Achievement{},
	}

	achievesList = []Achievement{
		{
			Id:          0,
			Text:        "Вы дошли!",
			Description: "Вы дошли до этого места. Возьмите пряник!",
			Image:       "https://i.pinimg.com/236x/a5/76/99/a57699849fb0d8f69c8e4016457b5c66--job-well-done-quotes-congratulations-quotes.jpg",
		},
		{
			Id:          0,
			Text:        "Открыта новая локация!",
			Description: "Вы открыли ранее неизведанную территорию!",
			Image:       "https://dncache-mauganscorp.netdna-ssl.com/thumbseg/59/59246-bigthumbnail.jpg",
		},
		{
			Id:          0,
			Text:        "Тише едешь - дальше будешь!",
			Description: "Вы были самым медленным путником на данном маршруте",
			Image:       "https://www.google.ru/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=0ahUKEwj6gr_u6PLWAhWBa5oKHY5hAIUQjRwIBw&url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DD2yY1F1-HqI&psig=AOvVaw225Fq6P5yE6yDN9t81MYlv&ust=1508163447914390",
		},
		{
			Id:          0,
			Text:        "Родина не забудет",
			Description: "Посетив это место, вы поддержали отечественного производителя!",
			Image:       "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRRAQ1BiIu42_JwJ-38MRudrZgnjIgszwQ-cc1ty541N-TRUWFk",
		},
		{
			Id:          0,
			Text:        "Романтическая особа!",
			Description: "Теперь ты знаешь толк в соболазнении",
			Image:       "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSX_sb6LXk2F1Fce1aXvSsV4X6BPvGGnsNmiEsNYoZMJJd1LROD",
		},
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/api/getMapItems", sendMapItems)
	http.HandleFunc("/api/getUserInfo", sendUserInfo)
	http.HandleFunc("/api/sendPosition", sendPrize)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
