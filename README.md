Bitaksi API Gateway
Bu proje, Bitaksi microservice mimarisi altında **Gateway-Service** olarak çalışır ve driver servisleri ile iletişimi yönetir. Tek container içinde çalışacak şekilde tasarlanmıştır ve dış dünyaya yalnızca Gateway portu açılır.
İçerik
- Ön Koşullar
- Kurulum
- Çalıştırma
- API Dökümantasyonu
- Swagger UI
- Servisler
Ön Koşullar
- Docker & Docker Compose
- Go >= 1.25
- Postman (opsiyonel)
Kurulum
```bash
git clone https://github.com/kullanici_adiniz/bitaksi-gateway.git
cd bitaksi-gateway
```
Çalıştırma
### Docker ile
1. Container’ı build et ve çalıştır:

```bash
docker-compose up --build
```
2. Container ayağa kalktığında:
- MongoDB: 27017 portunda (localhost:27017)
- Gateway-Service: 9090 portunda (localhost:9090)

> Not: Driver-Service (8080) iç container’da çalışır ve dışa açılmaz. Gateway container içinden ona `http://localhost:8080` ile erişir.
### Go ile
```bash
go run main.go
```
Server 9090 portunda çalışacaktır.
API Dökümantasyonu
Tüm API endpointleri ve kullanım örnekleri için Postman dökümantasyonu:
[https://documenter.getpostman.com/view/25744155/2sB3dQtoQU](https://documenter.getpostman.com/view/25744155/2sB3dQtoQU)
Swagger UI
Uygulama çalışırken Swagger UI’ye tarayıcı üzerinden erişebilirsin:
http://localhost:9090/swagger/index.html
Servisler
- **Gateway-Service** → `:9090`
  - API çağrıları dış dünyaya açılır.
  - Driver servislerine istekleri yönlendirir.

- **Driver-Service** → `:8080` (iç container)
  - Sadece container içi Gateway üzerinden erişilebilir.
  - CRUD işlemleri, sürücü listeleme ve nearby sorgularını yönetir.
Örnek Curl İstekleri
### Sürücü Ekleme
```bash
Authorization: Bearer BITAKSI-TEST-TOKEN-12345
curl -X POST http://localhost:9090/drivers \
-H "Content-Type: application/json" \
-d {
  "firstName": "Can",
  "lastName": "Heyal",
  "plate": "34BCH123",
  "taksiType": "sarı",
  "carBrand": "Toyota",
  "carModel": "Corolla",
  "lat": 41.0082,
  "lon": 28.9784
}

```
### Sürücü Listeleme
Authorization: Bearer BITAKSI-TEST-TOKEN-12345
```bash
curl http://localhost:9090/drivers
```
### Sürücü Güncelleme
```bash
Authorization: Bearer BITAKSI-TEST-TOKEN-12345
curl -X PUT http://localhost:9090/drivers \
-H "Content-Type: application/json" \
-d {
  "id": "6935eb0b25c92d2a97825be1",
  "firstName": "sare",
  "lastName": "heyal",
  "plate": "34ABC1234",
  "taksiType": "siyah",
  "carBrand": "Kia",
  "carModel": "bbb",
  "lat": 41.02,
  "lon": 28.97
}

### Yakınındaki Sürücüler
Authorization: Bearer BITAKSI-TEST-TOKEN-12345
```bash
curl -x POST [http://localhost:9090/drivers](http://localhost:9090/drivers/nearby)\
-H "Content-Type: application/json" \
-d {
  "lat": 41.02,
  "lon": 28.97,
  "taksiType": "sarı"
}
```
###Logları Görüntülemek
Authorization: Bearer BITAKSI-TEST-TOKEN-12345
curl -X GET [http://localhost:9090/drivers](http://localhost:9090/internal/logs) \
-H "Content-Type: application/json"

Notlar
- Gateway-Service dış dünyaya açık olduğundan, tüm istekler buradan yönlendirilir.
- Driver-Service sadece container içi `localhost:8080` üzerinden çalışır ve dışa açılmaz.
- MongoDB bağlantısı `mongodb://mongo:27017` veya `localhost:27017` üzerinden sağlanabilir.
