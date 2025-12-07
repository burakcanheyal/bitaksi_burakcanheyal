package application

import "errors"

// Bu dosya servis içinde dönebilecek tüm error mesajlarını
// front/gateway tarafında standardize etmek için kullanılır.

var CustomErrorMap = map[string]struct {
	Status  int
	Message string
}{
	// Generic
	"ERR_BAD_REQUEST": {400, "Invalid request"},
	"ERR_INTERNAL":    {500, "Internal server error"},
	"Kayıtlı Plaka":   {400, "Plaka zaten kayıtlı"},

	// Driver Specific
	"ERR_MISSING_ID":         {400, "Driver ID missing"},
	"Hatalı ID":              {400, "Hatalı Id"},
	"Sürücü Bulunamadı":      {404, "Sürücü Bulunamadı"},
	"ERR_VALIDATION":         {400, "Validation error"},
	"Koordinat Hatası : Lon": {400, "Lon koordinatı validasyon hatası"},
	"Koordinat Hatası: Lat":  {400, "Lat koordinatı validasyon hatası"},
	"Taksi Tipi Hatası":      {400, "Taksi Tipi hatası"},
}
var (
	ErrBadRequest      = errors.New("ERR_BAD_REQUEST")
	ErrInternal        = errors.New("ERR_INTERNAL")
	ErrMissingID       = errors.New("ERR_MISSING_ID")
	ErrInvalidObjectID = errors.New("ERR_INVALID_OBJECTID")
	ErrDriverNotFound  = errors.New("ERR_NOT_FOUND")
	ErrValidation      = errors.New("ERR_VALIDATION")
)

// Wrap: dışarıya error kodu döndermek için
func Wrap(code string) error {
	return errors.New(code)
}
