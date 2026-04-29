# Auth Handler Notları (`internal/http/handlers/auth_handler.go`)

- **AuthHandler Struct ve Methodları**: `Register` fonksiyonu, `AuthHandler` struct'ına bağlı bir methoddur. `authHandler.service` sayesinde iş mantığına (business logic) erişir.

## gin.Context (`c`) Nedir?
- `c *gin.Context`: Request (istek) + Response (yanıt) + Context (taşıyıcı) bütünüdür.
- Request'i temsil eder.
- Response'u yazmayı sağlar.
- Middleware'lerle (ara katman) veri taşır.

**Sık Kullanılan gin.Context Metotları:**
- **JSON parse:** `c.ShouldBindJSON`
- **Response:** `c.JSON`
- **Parametreler:** `c.Param`, `c.Query`
- **Context:** `c.Request.Context()` (Bunu genellikle service katmanına geçiririz)

## ShouldBindJSON İşleyişi
- Gelen JSON body'yi alır ve belirttiğimiz struct'ın (örneğin `req model.RegisterRequest`) içine doldurur (map'ler). Gelen DTO'ları kontrol ederiz.
- Eğer JSON body, struct yapısına uymazsa hata (`err`) döner ve `utils.Fail` fonksiyonu çağırılarak hata durumu yönetilir (`http.StatusBadRequest`).

## Service'e İstek Gönderme
- `c.Request.Context()` servis katmanına geçirilerek context taşınır (`authHandler.service.Register(...)`).
- Eğer işlem sırasında bir hata olursa, `utils.Fail` ile `http.StatusInternalServerError` (500) hatası dönülür. İşlem başarılıysa `utils.Ok` ile HTTP 200 (OK) dönülerek başarılı yanıt verilir.
