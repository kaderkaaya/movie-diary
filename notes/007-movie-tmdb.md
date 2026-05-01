
---

İlk olarak router tarafında şu şekilde bir endpoint tanımlıyoruz:

```go
movie.GET("/list-movies/:movie_type", movieHandler.ListMovies)
```

Burada `movie_type` parametresi dinamik olarak geliyor. Bu değer `trending`, `popular`, `top_rated` veya `discover` olabilir.

Daha sonra service katmanında, gelen `movieType` değerine göre The Movie Database (TMDB) API’sine istek atıyoruz.

Öncelikle bir HTTP client oluşturuyoruz ve ardından `switch-case` yapısı ile hangi endpoint’e istek atacağımızı belirliyoruz:

```go
switch movieType {
case "trending":
	url = fmt.Sprintf(
		"https://api.themoviedb.org/3/trending/movie/day?api_key=%s",
		apiKey,
	)

case "popular", "top_rated":
	url = fmt.Sprintf(
		"https://api.themoviedb.org/3/movie/%s?api_key=%s&language=en-US&page=%d",
		movieType, apiKey, page,
	)

default:
	url = fmt.Sprintf(
		"https://api.themoviedb.org/3/discover/movie?api_key=%s&language=en-US&page=%d&include_adult=false&include_video=false&sort_by=popularity.desc&with_genres=%d&primary_release_year=%d&vote_average.gte=%f",
		apiKey, page, genreID, year, rating,
	)
}
```

* Burada `movieType` değerine göre doğru URL oluşturuluyor.
* Ardından API’den gelen response okunuyor:

```go
body, err := io.ReadAll(response.Body)
```

Bu adımda response body byte dizisi olarak alınıyor.

Daha sonra gelen veriyi işlemek için bir loop kullanıyoruz:

```go
for _, tmdbMovie := range tmdbResp.Results
```

Bu loop içinde, her bir film verisini kendi modelimize mapliyoruz.

Örneğin, film yılını parse etmek için:

```go
fmt.Sscanf(tmdbMovie.ReleaseDate[:4], "%d", &movieYear)
```

Burada `"2014-11-05"` gibi bir tarih string’inden sadece yıl kısmını alıp integer’a çeviriyoruz:

```
"2014-11-05" → 2014
```

Son olarak service katmanında iki farklı mapping işlemi yapıyoruz:

* İlk mapping: TMDB’den gelen veriyi kendi **entity modelimize** dönüştürmek
* İkinci mapping: Client’a döneceğimiz **response modeline** dönüştürmek

Bu sayede hem dış API yapısından bağımsız kalıyoruz hem de client’a daha temiz ve kontrol edilebilir bir response sağlıyoruz.

---
