# Config Paketi Notları (`internal/config/config.go`)

- **godotenv.Load()**: `.env` dosyasındaki key-value'ları alıp process environment'a (çalışma ortamına) yükler.
- **Pointer Döndürme (`*Config`)**: Büyük struct'ları kopyalamak yerine pointer ile return ediyoruz. Bu sayede bellek kullanımı azalır ve performans artar. Dependency injection yapıyoruz ve böylece daha clean (temiz) bir yapı olur.
- **&Config**: Tüm config değerlerini tek bir struct içinde topluyoruz. `&Config` ile config'in adresini gönderiyoruz, kopya değil direkt config'i alıyoruz. (`*` ile de o adresteki değeri değiştirebilirsin.)
- **mustEnv Fonksiyonu**: Çevre değişkeni (env) dosyasındaki zorunlu (required) key'leri alır. Eğer key'in değeri boşsa, eksik yapılandırma ile devam etmemek için `log.Fatalf` çağırarak uygulamayı başlatmadan öldürür.
