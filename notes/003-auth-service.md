
---

Kullanıcı kayıt olurken şifreyi düz (plain text) saklamak yerine hash’lemek gerekir. Bu yüzden önce bir helper fonksiyon yazıyoruz:

```go
func HashPassword(plain string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	return string(hashedPassword), err
}
```

Burada:

* Fonksiyon, kullanıcıdan gelen **plain password**’ü alır.
* Geriye iki değer döner:

  * `string`: Hash’lenmiş şifre
  * `error`: İşlem sırasında oluşabilecek hata

`bcrypt.GenerateFromPassword` fonksiyonu:

* String değil, **byte array (`[]byte`)** bekler → bu yüzden `[]byte(plain)` kullanıyoruz
* Şifreyi alır, **güvenli şekilde hash’ler** ve sonucu yine byte olarak döner

`bcryptCost` ise hash işleminin maliyetini belirler:

* Düşük cost → daha hızlı ama daha az güvenli
* Yüksek cost → daha yavaş ama daha güvenli

Son olarak:

* Dönen `[]byte` değeri `string`’e çevrilir
* Çünkü veritabanına genellikle **string olarak kaydederiz**

Kullanım tarafında ise:

```go
hashedPassword, err := utils.HashPassword(password)
if err != nil {
	return nil, apperror.ErrPasswordEmpty
}
```

* Eğer hashleme sırasında bir hata oluşursa, bu hatayı kontrol edip uygun bir **application error** döneriz.
* Hata yoksa, `hashedPassword` artık veritabanına kaydedilmeye hazırdır.

---

