
---

Kullanıcı oluştururken önce bir **entity (model)** instance’ı oluşturuyoruz:

```go
user := &model.User{
	Username:     username,
	Email:        email,
	PasswordHash: password,
}
```

* Burada aslında veritabanına yazacağımız **user objesini hazırlıyoruz**

Sonrasında kayıt işlemi:

```go
r.db.WithContext(ctx).Create(user)
```

Bu satırın yaptığı şey:

* GORM bunu arka tarafta şu SQL’e çevirir:
  `INSERT INTO users (...) VALUES (...)`
* `WithContext(ctx)` kullanımı önemli:

  * Request iptal edilirse veya timeout olursa
  * **DB işlemi de otomatik olarak iptal edilir**

Yani context, handler’dan başlayıp repository’ye kadar taşınır ve
**request lifecycle ile DB işlemi senkron kalır.**

---

Şimdi kullanıcıyı email ile bulduğumuz fonksiyona bakalım:

```go
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
```

Burada:

* İlk olarak boş bir `user` struct’ı oluşturuyoruz
* Ardından GORM ile sorgu atıyoruz:

Bu kodun SQL karşılığı:

```sql
SELECT * FROM users WHERE email = ? LIMIT 1;
```

Detaylar:

* `Where("email = ?", email)` → parametreli query (SQL injection’a karşı güvenli)
* `First(&user)` → ilk kaydı getirir ve `user` içine map eder
* Kayıt bulunamazsa veya başka bir hata olursa `err` döner

---

Özetle burada kurduğun akış şu:

* Handler → Service → Repository
* Context yukarıdan aşağı taşınıyor
* Repository katmanı:

  * `Create` ile veri yazıyor
  * `FindByEmail` ile veri okuyor


  if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
- burda gormun kendinden kaynaklı user bulamadığı için hata dönüyor ondan kaynaklı bunu ekledik.