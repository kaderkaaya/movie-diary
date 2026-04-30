şimdi burda ilk olarak routerde:
movie.GET("/list-movies/:movie_type", movieHandler.ListMovies)
bu sekilde istek atıyoruz çünkü burda movie type discovered, popular, toprated olabilir.

